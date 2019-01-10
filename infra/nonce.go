package infra

import (
	"context"
	"fmt"
	"hash"
	"hash/fnv"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/core/services"
)

// StripeMutex is an object that allows fine grained locking based on keys
//
// It ensures that if `key1 == key2` then lock associated with `key1` is the same as the one associated with `key2`
// It holds a stable number of locks in memory that use can control
//
// It is inspired from Java lib Guava: https://github.com/google/guava/wiki/StripedExplained
type StripeMutex struct {
	stripes []*sync.Mutex
	pool    *sync.Pool
}

// Lock acquire lock for a given key
func (m *StripeMutex) Lock(key string) {
	l, _ := m.getLock(key)
	l.Lock()
}

// Unlock release lock for a given key
func (m *StripeMutex) Unlock(key string) {
	l, _ := m.getLock(key)
	l.Unlock()
}

func (m *StripeMutex) getLock(key string) (*sync.Mutex, error) {
	h := m.pool.Get().(hash.Hash64)
	defer m.pool.Put(h)
	h.Reset()
	_, err := h.Write([]byte(key))
	return m.stripes[h.Sum64()%uint64(len(m.stripes))], err
}

// NewStripeMutex creates a stripe mutex
func NewStripeMutex(stripes uint) *StripeMutex {
	m := &StripeMutex{
		make([]*sync.Mutex, stripes),
		&sync.Pool{New: func() interface{} { return fnv.New64() }},
	}
	for i := 0; i < len(m.stripes); i++ {
		m.stripes[i] = &sync.Mutex{}
	}

	return m
}

// SafeNonce allow to manipulate nonce in a concurrently safe manner
type SafeNonce struct {
	mux   *sync.Mutex
	value uint64
}

// Lock acquire lock
func (n *SafeNonce) Lock() error {
	n.mux.Lock()
	return nil
}

// Unlock release lock
func (n *SafeNonce) Unlock() error {
	n.mux.Unlock()
	return nil
}

// Get retrieve nonce value
// Warning: it does not acquire the lock
func (n *SafeNonce) Get() (uint64, error) {
	return n.value, nil
}

// Set sets nonce value
// Warning: it does not acquire the lock
func (n *SafeNonce) Set(v uint64) error {
	n.value = v
	return nil
}

// CalibrateNonceFunc allows to calibrate managed nonce
type CalibrateNonceFunc func(chainID *big.Int, a common.Address) (uint64, error)

// CacheNonceManager allows to manage nonce
type CacheNonceManager struct {
	mux    *StripeMutex
	nonces *sync.Map

	calibrate CalibrateNonceFunc
}

// NewCacheNonceManager creates a new cache nonce
func NewCacheNonceManager(calibrate CalibrateNonceFunc, stripes uint) *CacheNonceManager {
	return &CacheNonceManager{
		mux:       NewStripeMutex(stripes),
		nonces:    &sync.Map{},
		calibrate: calibrate,
	}
}

func computeKey(chainID *big.Int, a common.Address) string {
	return fmt.Sprintf("%v-%v", chainID.Text(16), a.Hex())
}

// Obtain return a locked SafeNonce for given chain and address
func (c *CacheNonceManager) Obtain(chainID *big.Int, a common.Address) (services.NonceLocker, error) {
	key := computeKey(chainID, a)
	mux, err := c.mux.getLock(key)
	if err != nil {
		return nil, err
	}
	// Lock key
	mux.Lock()
	defer mux.Unlock()

	// Retrieve nonce from cache
	n, ok := c.nonces.LoadOrStore(key, &SafeNonce{mux: mux, value: 0})
	rv := n.(*SafeNonce)
	if !ok {
		// If nonce has just been created we compute its initial value
		rv.value, err = c.calibrate(chainID, a)
		if err != nil {
			return rv, err
		}
	}
	return rv, nil
}

// NewEthClientNonceCalibrate returns a function to get nonce initial values from an Eth client
func NewEthClientNonceCalibrate(ec *EthClient) CalibrateNonceFunc {
	return func(chainID *big.Int, a common.Address) (uint64, error) {
		return ec.PendingNonceAt(context.Background(), a)
	}
}
