package client

import (
	"context"

	"github.com/ConsenSys/orchestrate/pkg/errors"
	"github.com/ConsenSys/orchestrate/pkg/types/keymanager"
	"github.com/ConsenSys/orchestrate/pkg/types/keymanager/ethereum"
	zksnarks "github.com/ConsenSys/orchestrate/pkg/types/keymanager/zk-snarks"
	healthz "github.com/heptiolabs/healthcheck"
)

var _ KeyManagerClient = &NonHTTPClient{}

type NonHTTPClient struct {
}

func NewNonHTTPClient() KeyManagerClient {
	return &NonHTTPClient{}
}

func (c *NonHTTPClient) Checker() healthz.Check {
	return func() error {
		return nil
	}
}

func (c *NonHTTPClient) ETHCreateAccount(ctx context.Context, request *ethereum.CreateETHAccountRequest) (*ethereum.ETHAccountResponse, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHImportAccount(ctx context.Context, request *ethereum.ImportETHAccountRequest) (*ethereum.ETHAccountResponse, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHSign(ctx context.Context, address string, request *keymanager.SignPayloadRequest) (string, error) {
	return "", errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHSignTypedData(ctx context.Context, address string, request *ethereum.SignTypedDataRequest) (string, error) {
	return "", errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHSignTransaction(ctx context.Context, address string, request *ethereum.SignETHTransactionRequest) (string, error) {
	return "", errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHSignQuorumPrivateTransaction(ctx context.Context, address string, request *ethereum.SignQuorumPrivateTransactionRequest) (string, error) {
	return "", errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHSignEEATransaction(ctx context.Context, address string, request *ethereum.SignEEATransactionRequest) (string, error) {
	return "", errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHListAccounts(ctx context.Context, namespace string) ([]string, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHListNamespaces(ctx context.Context) ([]string, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHGetAccount(ctx context.Context, address, namespace string) (*ethereum.ETHAccountResponse, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHVerifySignature(ctx context.Context, request *ethereum.VerifyPayloadRequest) error {
	return errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ETHVerifyTypedDataSignature(ctx context.Context, request *ethereum.VerifyTypedDataRequest) error {
	return errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ZKSCreateAccount(ctx context.Context, request *zksnarks.CreateZKSAccountRequest) (*zksnarks.ZKSAccountResponse, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ZKSSign(ctx context.Context, address string, request *keymanager.SignPayloadRequest) (string, error) {
	return "", errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ZKSListAccounts(ctx context.Context, namespace string) ([]string, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ZKSListNamespaces(ctx context.Context) ([]string, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ZKSGetAccount(ctx context.Context, address, namespace string) (*zksnarks.ZKSAccountResponse, error) {
	return nil, errors.DependencyFailureError("KeyManager is disabled")
}

func (c *NonHTTPClient) ZKSVerifySignature(ctx context.Context, request *zksnarks.VerifyPayloadRequest) error {
	return errors.DependencyFailureError("KeyManager is disabled")
}
