package tx

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/types/entities"
	error1 "github.com/consensys/orchestrate/pkg/types/error"
	"github.com/consensys/orchestrate/pkg/types/ethereum"
	"github.com/consensys/orchestrate/pkg/utils"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type Envelope struct {
	ID            string `validate:"uuid4,required"`
	Headers       map[string]string
	ContextLabels map[string]string
	Method
	JobType
	Tx             `mapstructure:",squash"`
	Chain          `mapstructure:",squash"`
	Contract       `mapstructure:",squash"`
	Private        `mapstructure:",squash"`
	Receipt        *ethereum.Receipt
	Errors         []*error1.Error
	InternalLabels map[string]string
}

func NewEnvelope() *Envelope {
	return &Envelope{
		Headers:        make(map[string]string),
		ContextLabels:  make(map[string]string),
		Errors:         make([]*error1.Error, 0),
		InternalLabels: make(map[string]string),
	}
}

func (e *Envelope) SafeEnvelope() *Envelope {
	if e.Headers == nil {
		e.Headers = make(map[string]string)
	}
	if e.ContextLabels == nil {
		e.ContextLabels = make(map[string]string)
	}
	if e.Errors == nil {
		e.Errors = make([]*error1.Error, 0)
	}
	if e.InternalLabels == nil {
		e.InternalLabels = make(map[string]string)
	}
	return e
}

func (e *Envelope) GetID() string {
	return e.ID
}

func (e *Envelope) SetID(id string) *Envelope {
	e.ID = id
	return e
}

func (e *Envelope) GetErrors() []*error1.Error {
	return e.Errors
}

// Error returns string representation of errors encountered by envelope
func (e *Envelope) Error() string {
	if len(e.GetErrors()) == 0 {
		return ""
	}
	return fmt.Sprintf("%q", e.GetErrors())
}

func (e *Envelope) AppendError(err *error1.Error) *Envelope {
	e.Errors = append(e.Errors, err)
	return e
}

func (e *Envelope) AppendErrors(errs []*error1.Error) *Envelope {
	e.Errors = append(e.Errors, errs...)
	return e
}
func (e *Envelope) SetReceipt(receipt *ethereum.Receipt) *Envelope {
	e.Receipt = receipt
	return e
}
func (e *Envelope) GetReceipt() *ethereum.Receipt {
	return e.Receipt
}

func (e *Envelope) IsEthSendTransaction() bool {
	return e.JobType == JobType_ETH_TX
}

// IsEthSendRawTransaction for a classic Ethereum transaction
func (e *Envelope) IsEthSendRawTransaction() bool {
	return e.JobType == JobType_ETH_RAW_TX
}

// IsEthSendTesseraMarkingTransaction for Quorum Constellation
func (e *Envelope) IsEthSendTesseraMarkingTransaction() bool {
	return e.JobType == JobType_ETH_TESSERA_MARKING_TX
}

// IsEthSendTesseraPrivateTransaction for Quorum Tessera
func (e *Envelope) IsEthSendTesseraPrivateTransaction() bool {
	return e.JobType == JobType_ETH_TESSERA_PRIVATE_TX
}

// IsEthSendRawTransaction for Besu/EEA
func (e *Envelope) IsEeaSendMarkingTransaction() bool {
	return e.JobType == JobType_ETH_EEA_MARKING_TX
}

func (e *Envelope) IsEeaSendPrivateTransaction() bool {
	return e.JobType == JobType_ETH_EEA_PRIVATE_TX
}

// IsEthSendRawTransaction for Besu/EEA with Privacy Group
func (e *Envelope) IsEeaSendPrivateTransactionPrivacyGroup() bool {
	return e.IsEeaSendPrivateTransaction() && e.PrivacyGroupID != ""
}

// IsEthSendRawTransaction for Besu/EEA with PrivateFor
func (e *Envelope) IsEeaSendPrivateTransactionPrivateFor() bool {
	return e.IsEeaSendPrivateTransaction() && len(e.PrivateFor) > 0
}

// IsResendingJobTx in case ParentJob and envelopeID are equal
func (e *Envelope) IsResendingJobTx() bool {
	return e.GetParentJobUUID() == e.GetJobUUID()
}

func (e *Envelope) IsOneTimeKeySignature() bool {
	if v, ok := e.InternalLabels[TxFromLabel]; ok {
		return v == TxFromOneTimeKey
	}

	return false
}

func (e *Envelope) IsParentJob() bool {
	return e.GetParentJobUUID() == ""
}

func (e *Envelope) IsChildJob() bool {
	return e.GetParentJobUUID() != ""
}

func (e *Envelope) Carrier() opentracing.TextMapCarrier {
	return e.ContextLabels
}

func (e *Envelope) OnlyWarnings() bool {
	for _, err := range e.GetErrors() {
		if !errors.IsWarning(err) {
			return false
		}
	}
	return true
}

func (e *Envelope) GetHeaders() map[string]string {
	return e.Headers
}

func (e *Envelope) SetHeaders(headers map[string]string) *Envelope {
	if headers != nil {
		e.Headers = headers
	}
	return e
}

func (e *Envelope) GetHeadersValue(key string) string {
	return e.Headers[key]
}
func (e *Envelope) SetHeadersValue(key, value string) *Envelope {
	e.Headers[key] = value
	return e
}

func (e *Envelope) GetInternalLabels() map[string]string {
	return e.InternalLabels
}

func (e *Envelope) GetInternalLabelsValue(key string) string {
	return e.InternalLabels[key]
}

func (e *Envelope) SetInternalLabels(internalLabels map[string]string) *Envelope {
	if internalLabels != nil {
		e.InternalLabels = internalLabels
	}
	return e
}

func (e *Envelope) AppendInternalLabels(internalLabels map[string]string) *Envelope {
	if internalLabels == nil {
		return e
	}

	for k, v := range internalLabels {
		e.InternalLabels[k] = v
	}

	return e
}

func (e *Envelope) SetInternalLabelsValue(key, value string) *Envelope {
	e.InternalLabels[key] = value
	return e
}

func (e *Envelope) SetContextLabelsValue(key, value string) *Envelope {
	e.ContextLabels[key] = value
	return e
}

func (e *Envelope) SetContextLabels(ctxLabels map[string]string) *Envelope {
	if ctxLabels != nil {
		e.ContextLabels = ctxLabels
	}
	return e
}

func (e *Envelope) Validate() []error {
	err := utils.GetValidator().Struct(e)
	if err != nil {
		return utils.HandleValidatorError(err.(validator.ValidationErrors))
	}

	if e.IsEeaSendPrivateTransaction() && len(e.GetPrivateFor()) > 0 && e.GetPrivacyGroupID() != "" {
		return []error{errors.DataError("privacyGroupId and privateFor fields are mutually exclusive")}
	}

	if e.IsEeaSendPrivateTransaction() && len(e.GetPrivateFor()) == 0 && e.GetPrivacyGroupID() == "" {
		return []error{errors.DataError("privacyGroupId or privateFor is missing")}
	}
	return nil
}

func (e *Envelope) GetContextLabelsValue(key string) string {
	return e.ContextLabels[key]
}
func (e *Envelope) GetContextLabels() map[string]string {
	return e.ContextLabels
}

type Tx struct {
	From            *ethcommon.Address
	To              *ethcommon.Address
	Gas             *uint64
	GasPrice        *hexutil.Big
	GasFeeCap       *hexutil.Big
	GasTipCap       *hexutil.Big
	AccessList      []*ethereum.AccessTuple
	TransactionType string
	Value           *hexutil.Big
	Nonce           *uint64
	Data            hexutil.Bytes   `validate:"omitempty"`
	Raw             hexutil.Bytes   `validate:"omitempty,required_with_all=TxHash"`
	TxHash          *ethcommon.Hash `validate:"omitempty,required_with_all=Raw"`
}

func (e *Envelope) GetTransaction() (*ethtypes.Transaction, error) {
	// TODO: Use custom validation with https://godoc.org/gopkg.in/go-playground/validator.v10#Validate.StructFiltered

	data := e.MustGetDataBytes()
	nonce, err := e.GetNonceUint64()
	if err != nil {
		return nil, err
	}
	value, err := e.GetValueBig()
	if value == nil || err != nil {
		_ = e.SetValue((*hexutil.Big)(big.NewInt(0)))
	}

	gas, err := e.GetGasUint64()
	if err != nil {
		if e.IsEeaSendPrivateTransaction() {
			gas = 0
		} else {
			return nil, err
		}
	}

	gasPrice, err := e.GetGasPriceBig()
	if err != nil {
		return nil, err
	}

	if e.IsContractCreation() {
		// Create contract deployment transaction
		return ethtypes.NewContractCreation(
			nonce,
			value.ToInt(),
			gas,
			gasPrice.ToInt(),
			data,
		), nil
	}

	to, err := e.GetToAddress()
	if err != nil {
		return nil, err
	}

	// Create transaction
	return ethtypes.NewTransaction(
		nonce,
		to,
		value.ToInt(),
		gas,
		gasPrice.ToInt(),
		data,
	), nil
}

// FROM

func (e *Envelope) GetFrom() *ethcommon.Address {
	return e.From
}

func (e *Envelope) GetFromAddress() (ethcommon.Address, error) {
	if e.From == nil {
		return ethcommon.Address{}, errors.DataError("no from is filled")
	}
	return *e.From, nil
}

func (e *Envelope) MustGetFromAddress() ethcommon.Address {
	if e.From == nil {
		return ethcommon.Address{}
	}
	return *e.From
}

func (e *Envelope) GetFromString() string {
	if e.From == nil {
		return ""
	}
	return e.From.Hex()
}

func (e *Envelope) SetFromString(from string) error {
	if from != "" {
		if !ethcommon.IsHexAddress(from) {
			return errors.DataError("invalid from - got %s", from)
		}
		_ = e.SetFrom(ethcommon.HexToAddress(from))
	}
	return nil
}

func (e *Envelope) MustSetFromString(from string) *Envelope {
	_ = e.SetFrom(ethcommon.HexToAddress(from))
	return e
}

func (e *Envelope) SetFrom(from ethcommon.Address) *Envelope {
	e.From = &from
	return e
}

// TO

func (e *Envelope) GetTo() *ethcommon.Address {
	return e.To
}

func (e *Envelope) GetToAddress() (ethcommon.Address, error) {
	if e.To == nil {
		return ethcommon.Address{}, errors.DataError("no to is filled")
	}
	return *e.To, nil
}

func (e *Envelope) MustGetToAddress() ethcommon.Address {
	if e.To == nil {
		return ethcommon.Address{}
	}
	return *e.To
}

func (e *Envelope) GetToString() string {
	if e.To == nil {
		return ""
	}
	return e.To.Hex()
}

func (e *Envelope) MustSetToString(to string) *Envelope {
	_ = e.SetTo(ethcommon.HexToAddress(to))
	return e
}

func (e *Envelope) SetToString(to string) error {
	if to != "" {
		if !ethcommon.IsHexAddress(to) {
			return errors.DataError("invalid to - got %s", to)
		}
		_ = e.SetTo(ethcommon.HexToAddress(to))
	}
	return nil
}

func (e *Envelope) SetTo(to ethcommon.Address) *Envelope {
	e.To = &to
	return e
}

// GAS

func (e *Envelope) GetGas() *uint64 {
	return e.Gas
}
func (e *Envelope) GetGasUint64() (uint64, error) {
	if e.Gas == nil {
		return 0, errors.DataError("no gas is filled")
	}
	return *e.Gas, nil
}
func (e *Envelope) MustGetGasUint64() uint64 {
	if e.Gas == nil {
		return 0
	}
	return *e.Gas
}
func (e *Envelope) GetGasString() string {
	if e.Gas == nil {
		return ""
	}
	return strconv.FormatUint(*e.Gas, 10)
}
func (e *Envelope) SetGasString(gas string) error {
	if gas != "" {
		g, err := strconv.ParseUint(gas, 10, 32)
		if err != nil {
			return errors.DataError("invalid gas - got %s", gas)
		}
		_ = e.SetGas(g)
	}
	return nil
}
func (e *Envelope) SetGas(gas uint64) *Envelope {
	e.Gas = &(&struct{ x uint64 }{gas}).x
	return e
}

// NONCE

func (e *Envelope) GetNonce() *uint64 {
	return e.Nonce
}
func (e *Envelope) GetNonceUint64() (uint64, error) {
	if e.Nonce == nil {
		return 0, errors.DataError("no nonce is filled")
	}
	return *e.Nonce, nil
}
func (e *Envelope) MustGetNonceUint64() uint64 {
	if e.Nonce == nil {
		return 0
	}
	return *e.Nonce
}
func (e *Envelope) GetNonceString() string {
	return utils.ValueToString(e.Nonce)
}

func (e *Envelope) SetNonceString(nonce string) error {
	if nonce != "" {
		g, err := strconv.ParseUint(nonce, 10, 32)
		if err != nil {
			return errors.DataError("invalid nonce - got %s", nonce)
		}
		_ = e.SetNonce(g)
	}
	return nil
}
func (e *Envelope) SetNonce(nonce uint64) *Envelope {
	e.Nonce = &(&struct{ x uint64 }{nonce}).x
	return e
}

// GASPRICE

func (e *Envelope) GetGasPrice() *hexutil.Big {
	return e.GasPrice
}

func (e *Envelope) GetGasPriceBig() (*hexutil.Big, error) {
	if e.GasPrice == nil {
		return nil, errors.DataError("no gasPrice is filled")
	}
	return e.GasPrice, nil
}

func (e *Envelope) GetGasPriceString() string {
	if e.GasPrice == nil {
		return ""
	}
	return e.GasPrice.String()
}

func (e *Envelope) SetGasPriceString(gasPrice string) error {
	if gasPrice != "" {
		v, err := hexutil.DecodeBig(gasPrice)
		if err != nil {
			return errors.DataError("invalid gasPrice - got %s", gasPrice)
		}
		_ = e.SetGasPrice((*hexutil.Big)(v))
	}

	return nil
}

func (e *Envelope) SetGasPrice(gasPrice *hexutil.Big) *Envelope {
	e.GasPrice = gasPrice
	return e
}

// GasFeeCap
func (e *Envelope) SetGasFeeCapString(gasFeeCap string) error {
	if gasFeeCap != "" {
		v, err := hexutil.DecodeBig(gasFeeCap)
		if err != nil {
			return errors.DataError("invalid gasFeeCap - got %s", gasFeeCap)
		}
		_ = e.SetFeeCap((*hexutil.Big)(v))
	}
	return nil
}

func (e *Envelope) SetFeeCap(gasFeeCap *hexutil.Big) *Envelope {
	e.GasFeeCap = gasFeeCap
	return e
}

func (e *Envelope) GetGasFeeCapString() string {
	if e.GasFeeCap == nil {
		return ""
	}
	return e.GasFeeCap.String()
}

func (e *Envelope) GetGasFeeCap() *hexutil.Big {
	return e.GasFeeCap
}

// GasTipCap
func (e *Envelope) SetGasTipCapString(gasTipCap string) error {

	if gasTipCap != "" {
		v, err := hexutil.DecodeBig(gasTipCap)
		if err != nil {
			return errors.DataError("invalid gasTipCap - got %s", gasTipCap)
		}
		_ = e.SetTipCap((*hexutil.Big)(v))
	}
	return nil
}

func (e *Envelope) SetTipCap(gasTipCap *hexutil.Big) *Envelope {
	e.GasTipCap = gasTipCap
	return e
}

func (e *Envelope) GetGasTipCap() *hexutil.Big {
	return e.GasTipCap
}

// @TODO: Remove ***String() func helpers since they might not be needed anymore
func (e *Envelope) GetGasTipCapString() string {
	if e.GasTipCap == nil {
		return ""
	}
	return e.GasTipCap.String()
}

// AccessList

func (e *Envelope) SetAccessList(accessList []*ethereum.AccessTuple) *Envelope {
	e.AccessList = accessList
	return e
}

func (e *Envelope) GetAccessList() []*ethereum.AccessTuple {
	return e.AccessList
}

// Transaction Type

func (e *Envelope) SetTransactionType(txType string) *Envelope {
	e.TransactionType = txType
	return e
}

func (e *Envelope) GetTransactionType() string {
	return e.TransactionType
}

// VALUE

func (e *Envelope) GetValue() *hexutil.Big {
	return e.Value
}
func (e *Envelope) GetValueBig() (*hexutil.Big, error) {
	if e.Value == nil {
		return nil, errors.DataError("no value is filled")
	}
	return e.Value, nil
}

func (e *Envelope) GetValueString() string {
	if e.Value == nil {
		return ""
	}
	return e.Value.String()
}

func (e *Envelope) SetValueString(value string) error {
	if value != "" {
		v, err := hexutil.DecodeBig(value)
		if err != nil {
			return errors.DataError("invalid value - got %s", value)
		}
		_ = e.SetValue((*hexutil.Big)(v))
	}
	return nil
}

func (e *Envelope) SetValue(value *hexutil.Big) *Envelope {
	e.Value = value
	return e
}

// DATA

func (e *Envelope) GetDataString() string {
	return e.Data.String()
}

func (e *Envelope) GetData() hexutil.Bytes {
	return e.Data
}

func (e *Envelope) MustGetDataBytes() []byte {
	return e.Data
}

func (e *Envelope) SetData(data []byte) *Envelope {
	e.Data = data
	return e
}

func (e *Envelope) SetDataString(data string) error {
	var err error
	e.Data, err = hexutil.Decode(data)
	if err != nil {
		return errors.DataError("invalid data")
	}
	return nil
}

func (e *Envelope) MustSetDataString(data string) *Envelope {
	e.Data = utils.StringToHexBytes(data)
	return e
}

// RAW

func (e *Envelope) GetShortRaw() string {
	return utils.ShortString(e.Raw.String(), 30)
}

func (e *Envelope) GetRaw() hexutil.Bytes {
	return e.Raw
}

func (e *Envelope) GetRawString() string {
	return e.Raw.String()
}

func (e *Envelope) MustGetRawBytes() []byte {
	return e.Raw
}

func (e *Envelope) SetRaw(raw []byte) *Envelope {
	e.Raw = raw
	return e
}

func (e *Envelope) SetRawString(raw string) error {
	var err error
	e.Raw, err = hexutil.Decode(raw)
	if err != nil {
		return errors.DataError("invalid raw")
	}
	return nil
}

func (e *Envelope) MustSetRawString(raw string) *Envelope {
	if raw != "" {
		e.Raw = hexutil.MustDecode(raw)
	}

	return e
}

// TXHASH

func (e *Envelope) GetTxHash() *ethcommon.Hash {
	return e.TxHash
}

func (e *Envelope) GetTxHashValue() (ethcommon.Hash, error) {
	if e.TxHash == nil {
		return ethcommon.Hash{}, errors.DataError("no tx hash is filled")
	}
	return *e.TxHash, nil
}

func (e *Envelope) MustGetTxHashValue() ethcommon.Hash {
	if e.TxHash == nil {
		return ethcommon.Hash{}
	}
	return *e.TxHash
}

func (e *Envelope) GetTxHashString() string {
	if e.TxHash == nil {
		return ""
	}
	return e.TxHash.Hex()
}

func (e *Envelope) SetTxHash(hash ethcommon.Hash) *Envelope {
	e.TxHash = &hash
	return e
}

func (e *Envelope) SetTxHashString(txHash string) error {
	if txHash != "" {
		h, err := hexutil.Decode(txHash)
		if err != nil || len(h) != ethcommon.HashLength {
			return errors.DataError("invalid txHash - got %s", txHash)
		}
		_ = e.SetTxHash(ethcommon.BytesToHash(h))
	}
	return nil
}

func (e *Envelope) MustSetTxHashString(txHash string) *Envelope {
	_ = e.SetTxHash(ethcommon.HexToHash(txHash))
	return e
}

type Chain struct {
	ChainID   *big.Int
	ChainName string
	ChainUUID string `validate:"omitempty,uuid4"`
}

func (e *Envelope) GetChainID() *big.Int {
	return e.ChainID
}

func (e *Envelope) GetChainIDString() string {
	if e.ChainID == nil {
		return ""
	}
	return e.ChainID.String()
}

func (e *Envelope) SetChainID(chainID *big.Int) *Envelope {
	e.ChainID = chainID
	return e
}

func (e *Envelope) SetChainIDUint64(chainID uint64) *Envelope {
	e.ChainID = big.NewInt(int64(chainID))
	return e
}

func (e *Envelope) SetChainIDString(chainID string) error {
	if chainID != "" {
		v, ok := new(big.Int).SetString(chainID, 10)
		if !ok {
			return errors.DataError("invalid chainID - got %s", chainID)
		}
		_ = e.SetChainID(v)
	}
	return nil
}

func (e *Envelope) GetChainName() string {
	return e.ChainName
}

func (e *Envelope) SetChainName(chainName string) *Envelope {
	e.ChainName = chainName
	return e
}

func (e *Envelope) GetChainUUID() string {
	return e.ChainUUID
}

func (e *Envelope) SetChainUUID(chainUUID string) *Envelope {
	e.ChainUUID = chainUUID
	return e
}

type Contract struct {
	ContractName    string `validate:"omitempty,required_with_all=ContractTag"`
	ContractTag     string `validate:"omitempty"`
	MethodSignature string `validate:"omitempty,isValidMethodSig"`
	Args            []string
}

// IsContractCreation indicate whether the transaction is a contract deployment
func (e *Envelope) IsContractCreation() bool {
	return e.GetTo() == nil
}

// Short returns a short string representation of contract information
func (e *Envelope) MustGetMethodName() string {
	return strings.Split(e.MethodSignature, "(")[0]
}

func (e *Envelope) GetMethodSignature() string {
	return e.MethodSignature
}

func (e *Envelope) GetArgs() []string {
	return e.Args
}

func (e *Envelope) SetContractName(contractName string) *Envelope {
	e.ContractName = contractName
	return e
}

func (e *Envelope) SetMethodSignature(methodSignature string) *Envelope {
	e.MethodSignature = methodSignature
	return e
}
func (e *Envelope) SetArgs(args []string) *Envelope {
	e.Args = args
	return e
}

func (e *Envelope) SetContractTag(contractTag string) *Envelope {
	e.ContractTag = contractTag
	return e
}

func (e *Envelope) ShortContract() string {
	if e.ContractName == "" {
		return ""
	}

	if e.ContractTag == "" {
		return e.ContractName
	}

	return fmt.Sprintf("%v[%v]", e.ContractName, e.ContractTag)
}

type Private struct {
	PrivateFor     []string             `validate:"dive,base64"`
	MandatoryFor   []string             `validate:"dive,base64"`
	PrivateFrom    string               `validate:"omitempty,base64"`
	PrivateTxType  string               `validate:"omitempty,oneof=restricted unrestricted"`
	PrivacyGroupID string               `validate:"omitempty,base64"`
	PrivacyFlag    entities.PrivacyFlag `validate:"omitempty,isPrivacyFlag"`
}

func (e *Envelope) GetPrivateFor() []string {
	return e.PrivateFor
}
func (e *Envelope) SetPrivateFor(privateFor []string) *Envelope {
	e.PrivateFor = privateFor
	return e
}

func (e *Envelope) GetMandatoryFor() []string {
	return e.MandatoryFor
}

func (e *Envelope) SetMandatoryFor(mandatoryFor []string) *Envelope {
	e.MandatoryFor = mandatoryFor
	return e
}

func (e *Envelope) GetPrivacyFlag() entities.PrivacyFlag {
	return e.PrivacyFlag
}

func (e *Envelope) SetPrivacyFlag(privacyFlag int32) *Envelope {
	e.PrivacyFlag = entities.PrivacyFlag(privacyFlag)
	return e
}

func (e *Envelope) SetPrivateFrom(privateFrom string) *Envelope {
	e.PrivateFrom = privateFrom
	return e
}
func (e *Envelope) GetPrivateFrom() string {
	return e.PrivateFrom
}

func (e *Envelope) SetPrivateTxType(privateTxType string) *Envelope {
	e.PrivateTxType = privateTxType
	return e
}

func (e *Envelope) GetPrivateTxType() string {
	return e.PrivateTxType
}

func (e *Envelope) SetPrivacyGroupID(privacyGroupID string) *Envelope {
	e.PrivacyGroupID = privacyGroupID
	return e
}

func (e *Envelope) GetPrivacyGroupID() string {
	return e.PrivacyGroupID
}

func (e *Envelope) GetEnclaveKey() hexutil.Bytes {
	return utils.StringToHexBytes(e.InternalLabels[EnclaveKeyLabel])
}

func (e *Envelope) SetEnclaveKey(enclaveKey string) *Envelope {
	e.InternalLabels[EnclaveKeyLabel] = enclaveKey
	return e
}

func (e *Envelope) GetPriority() string {
	return e.InternalLabels[PriorityLabel]
}

func (e *Envelope) SetPriority(priority string) *Envelope {
	e.InternalLabels[PriorityLabel] = priority
	return e
}

func (e *Envelope) SetJobType(jobType JobType) *Envelope {
	e.JobType = jobType
	return e
}

func (e *Envelope) GetJobTypeString() string {
	return e.JobType.String()
}

func (e *Envelope) SetScheduleUUID(uuid string) *Envelope {
	e.InternalLabels[ScheduleUUIDLabel] = uuid
	return e
}

func (e *Envelope) GetScheduleUUID() string {
	return e.InternalLabels[ScheduleUUIDLabel]
}

func (e *Envelope) SetJobUUID(uuid string) *Envelope {
	e.InternalLabels[JobUUIDLabel] = uuid
	return e
}

func (e *Envelope) GetJobUUID() string {
	return e.InternalLabels[JobUUIDLabel]
}

func (e *Envelope) GetNextJobUUID() string {
	return e.ContextLabels[NextJobUUIDLabel]
}

func (e *Envelope) GetParentJobUUID() string {
	return e.ContextLabels[ParentJobUUIDLabel]
}

func (e *Envelope) GetExpectedNonce() string {
	return e.ContextLabels[ExpectedNonceLabel]
}

func (e *Envelope) SetNextJobUUID(uuid string) *Envelope {
	e.ContextLabels[NextJobUUIDLabel] = uuid
	return e
}

func (e *Envelope) TxRequest() *TxRequest {
	req := &TxRequest{
		Id:      e.ID,
		Headers: e.Headers,
		Chain:   e.GetChainName(),
		Method:  e.Method,
		JobType: e.JobType,
		Params: &Params{
			From:            e.GetFromString(),
			To:              e.GetToString(),
			Gas:             e.GetGasString(),
			GasPrice:        e.GetGasPriceString(),
			Value:           e.GetValueString(),
			Nonce:           e.GetNonceString(),
			Data:            e.GetDataString(),
			Contract:        e.ShortContract(),
			TransactionType: e.GetTransactionType(),
			MethodSignature: e.GetMethodSignature(),
			Args:            e.GetArgs(),
			Raw:             e.GetRawString(),
			PrivateFor:      e.GetPrivateFor(),
			PrivateFrom:     e.GetPrivateFrom(),
			PrivateTxType:   e.GetPrivateTxType(),
			PrivacyGroupId:  e.GetPrivacyGroupID(),
		},
		ContextLabels: e.ContextLabels,
	}

	return req
}

func (e *Envelope) fieldsToInternal() {
	if e.InternalLabels == nil {
		e.InternalLabels = make(map[string]string)
	}

	if e.GetChainID() != nil {
		e.InternalLabels[ChainIDLabel] = e.GetChainIDString()
	}
	if e.GetTxHash() != nil {
		e.InternalLabels[TxHashLabel] = e.GetTxHashString()
	}
	if e.GetChainUUID() != "" {
		e.InternalLabels[ChainUUIDLabel] = e.GetChainUUID()
	}
	if e.GetJobUUID() != "" {
		e.InternalLabels[JobUUIDLabel] = e.GetJobUUID()
	}
	if e.GetPriority() != "" {
		e.InternalLabels[PriorityLabel] = e.GetPriority()
	}
	if e.GetScheduleUUID() != "" {
		e.InternalLabels[ScheduleUUIDLabel] = e.GetScheduleUUID()
	}

}

func (e *Envelope) internalToFields() error {
	hash, ok := e.InternalLabels[TxHashLabel]
	if err := e.SetTxHashString(hash); err != nil && ok {
		return err
	}
	if err := e.SetChainIDString(e.InternalLabels[ChainIDLabel]); err != nil {
		return err
	}
	_ = e.SetChainUUID(e.InternalLabels[ChainUUIDLabel])
	_ = e.SetJobUUID(e.InternalLabels[JobUUIDLabel])
	if priority, ok := e.InternalLabels[PriorityLabel]; ok && priority != "" {
		_ = e.SetPriority(priority)
	}
	if txHash, ok := e.InternalLabels[TxHashLabel]; ok && txHash != "" {
		_ = e.SetTxHashString(txHash)
	}
	_ = e.SetScheduleUUID(e.InternalLabels[ScheduleUUIDLabel])
	return nil
}

func (e *Envelope) TxEnvelopeAsRequest() *TxEnvelope {
	e.fieldsToInternal()
	return &TxEnvelope{
		InternalLabels: e.InternalLabels,
		Msg:            &TxEnvelope_TxRequest{e.TxRequest()},
	}
}

func (e *Envelope) TxEnvelopeAsResponse() *TxEnvelope {
	e.fieldsToInternal()
	return &TxEnvelope{
		InternalLabels: e.InternalLabels,
		Msg:            &TxEnvelope_TxResponse{e.TxResponse()},
	}
}

func (e *Envelope) TxResponse() *TxResponse {
	return &TxResponse{
		Headers:       e.Headers,
		Id:            e.ID,
		JobUUID:       e.GetJobUUID(),
		ContextLabels: e.ContextLabels,
		Transaction: &ethereum.Transaction{
			From:       e.GetFromString(),
			Nonce:      e.GetNonceString(),
			To:         e.GetToString(),
			Value:      e.GetValueString(),
			Gas:        e.GetGasString(),
			GasPrice:   e.GetGasPriceString(),
			GasFeeCap:  e.GetGasFeeCapString(),
			GasTipCap:  e.GetGasTipCapString(),
			AccessList: e.GetAccessList(),
			TxType:     e.GetTransactionType(),
			Data:       e.GetDataString(),
			Raw:        e.GetRawString(),
			TxHash:     e.GetTxHashString(),
		},
		Chain:   e.GetChainName(),
		Receipt: e.Receipt,
		Errors:  e.Errors,
	}
}

func (e *Envelope) loadPtrFields(gas, nonce, gasPrice, gasFeeCap, gasTipCap, value, from, to string) []*error1.Error {
	errs := make([]*error1.Error, 0)
	if err := e.SetGasString(gas); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetNonceString(nonce); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetGasPriceString(gasPrice); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetGasFeeCapString(gasFeeCap); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetGasTipCapString(gasTipCap); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetValueString(value); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetFromString(from); err != nil {
		errs = append(errs, errors.FromError(err))
	}
	if err := e.SetToString(to); err != nil {
		errs = append(errs, errors.FromError(err))
	}

	return errs
}

// Attribute kafka partition and redis keys to well attribute nonce
// For a classic eth_sendRawTransaction transaction - <from>@<chainID>
// For a eea_sendRawTransaction with a privacyGroupID - <from>@eea-<privacyGroupID>@<chainID>
// For a eea_sendRawTransaction with a privateFor - <from>@eea-<hash(privateFor-privateFrom)>@<chainID>
func (e *Envelope) PartitionKey() string {

	// Return empty partition key for raw tx and one time key tx
	// Not able to format a correct partition key if From or ChainID are not set. In that case return empty partition key
	if e.IsEthSendRawTransaction() || e.IsOneTimeKeySignature() || e.GetFrom() == nil || e.GetChainID() == nil {
		return ""
	}

	switch {
	case e.IsEeaSendPrivateTransactionPrivacyGroup():
		return fmt.Sprintf("%v@eea-%v@%v", e.GetFromString(), e.GetPrivacyGroupID(), e.GetChainID().String())
	case e.IsEeaSendPrivateTransactionPrivateFor():
		l := append(e.GetPrivateFor(), e.GetPrivateFrom())
		sort.Strings(l)
		h := md5.New()
		_, _ = h.Write([]byte(strings.Join(l, "-")))
		return fmt.Sprintf("%v@eea-%v@%v", e.GetFromString(), fmt.Sprintf("%x", h.Sum(nil)), e.GetChainID().String())
	default:
		return fmt.Sprintf("%v@%v", e.GetFromString(), e.GetChainID().String())
	}
}
