// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/types/tx/tx.proto

package tx

import (
	error1 "github.com/consensys/orchestrate/pkg/types/error"
	ethereum "github.com/consensys/orchestrate/pkg/types/ethereum"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Method int32

const (
	// Classic Ethereum Transaction
	Method_ETH_SENDRAWTRANSACTION Method = 0
	// Quorum Constellation
	Method_ETH_SENDPRIVATETRANSACTION Method = 1
	// Quorum Tessera
	Method_ETH_SENDRAWPRIVATETRANSACTION Method = 2
	// Besu Orion
	Method_EEA_SENDPRIVATETRANSACTION Method = 3
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "ETH_SENDRAWTRANSACTION",
		1: "ETH_SENDPRIVATETRANSACTION",
		2: "ETH_SENDRAWPRIVATETRANSACTION",
		3: "EEA_SENDPRIVATETRANSACTION",
	}
	Method_value = map[string]int32{
		"ETH_SENDRAWTRANSACTION":        0,
		"ETH_SENDPRIVATETRANSACTION":    1,
		"ETH_SENDRAWPRIVATETRANSACTION": 2,
		"EEA_SENDPRIVATETRANSACTION":    3,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_types_tx_tx_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_pkg_types_tx_tx_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_pkg_types_tx_tx_proto_rawDescGZIP(), []int{0}
}

type JobType int32

const (
	JobType_ETH_TX                 JobType = 0
	JobType_ETH_RAW_TX             JobType = 1
	JobType_ETH_ORION_MARKING_TX   JobType = 2
	JobType_ETH_ORION_EEA_TX       JobType = 3
	JobType_ETH_TESSERA_MARKING_TX JobType = 4
	JobType_ETH_TESSERA_PRIVATE_TX JobType = 5
)

// Enum value maps for JobType.
var (
	JobType_name = map[int32]string{
		0: "ETH_TX",
		1: "ETH_RAW_TX",
		2: "ETH_ORION_MARKING_TX",
		3: "ETH_ORION_EEA_TX",
		4: "ETH_TESSERA_MARKING_TX",
		5: "ETH_TESSERA_PRIVATE_TX",
	}
	JobType_value = map[string]int32{
		"ETH_TX":                 0,
		"ETH_RAW_TX":             1,
		"ETH_ORION_MARKING_TX":   2,
		"ETH_ORION_EEA_TX":       3,
		"ETH_TESSERA_MARKING_TX": 4,
		"ETH_TESSERA_PRIVATE_TX": 5,
	}
)

func (x JobType) Enum() *JobType {
	p := new(JobType)
	*p = x
	return p
}

func (x JobType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_types_tx_tx_proto_enumTypes[1].Descriptor()
}

func (JobType) Type() protoreflect.EnumType {
	return &file_pkg_types_tx_tx_proto_enumTypes[1]
}

func (x JobType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobType.Descriptor instead.
func (JobType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_types_tx_tx_proto_rawDescGZIP(), []int{1}
}

type TxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Technical header (optional)
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Name of the Chain as registered on the chain registry
	// e.g. 1 for mainnet, 3 for Ropsten
	Chain string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	// Type of a protocol that should be used
	Method Method `protobuf:"varint,3,opt,name=method,proto3,enum=tx.Method" json:"method,omitempty"`
	// Params for the transaction
	Params *Params `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	// ID of the Request in UUID RFC 4122, ISO/IEC 9834-8:2005 format
	// e.g a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	// [option]
	ContextLabels map[string]string `protobuf:"bytes,6,rep,name=context_labels,json=contextLabels,proto3" json:"context_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Type of the job
	JobType JobType `protobuf:"varint,7,opt,name=jobType,proto3,enum=tx.JobType" json:"jobType,omitempty"`
}

func (x *TxRequest) Reset() {
	*x = TxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_types_tx_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxRequest) ProtoMessage() {}

func (x *TxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_types_tx_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxRequest.ProtoReflect.Descriptor instead.
func (*TxRequest) Descriptor() ([]byte, []int) {
	return file_pkg_types_tx_tx_proto_rawDescGZIP(), []int{0}
}

func (x *TxRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *TxRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *TxRequest) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_ETH_SENDRAWTRANSACTION
}

func (x *TxRequest) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *TxRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TxRequest) GetContextLabels() map[string]string {
	if x != nil {
		return x.ContextLabels
	}
	return nil
}

func (x *TxRequest) GetJobType() JobType {
	if x != nil {
		return x.JobType
	}
	return JobType_ETH_TX
}

type TxEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*TxEnvelope_TxRequest
	//	*TxEnvelope_TxResponse
	Msg            isTxEnvelope_Msg  `protobuf_oneof:"msg"`
	InternalLabels map[string]string `protobuf:"bytes,1,rep,name=internal_labels,json=internalLabels,proto3" json:"internal_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TxEnvelope) Reset() {
	*x = TxEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_types_tx_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxEnvelope) ProtoMessage() {}

func (x *TxEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_types_tx_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxEnvelope.ProtoReflect.Descriptor instead.
func (*TxEnvelope) Descriptor() ([]byte, []int) {
	return file_pkg_types_tx_tx_proto_rawDescGZIP(), []int{1}
}

func (m *TxEnvelope) GetMsg() isTxEnvelope_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *TxEnvelope) GetTxRequest() *TxRequest {
	if x, ok := x.GetMsg().(*TxEnvelope_TxRequest); ok {
		return x.TxRequest
	}
	return nil
}

func (x *TxEnvelope) GetTxResponse() *TxResponse {
	if x, ok := x.GetMsg().(*TxEnvelope_TxResponse); ok {
		return x.TxResponse
	}
	return nil
}

func (x *TxEnvelope) GetInternalLabels() map[string]string {
	if x != nil {
		return x.InternalLabels
	}
	return nil
}

type isTxEnvelope_Msg interface {
	isTxEnvelope_Msg()
}

type TxEnvelope_TxRequest struct {
	TxRequest *TxRequest `protobuf:"bytes,2,opt,name=tx_request,json=txRequest,proto3,oneof"`
}

type TxEnvelope_TxResponse struct {
	TxResponse *TxResponse `protobuf:"bytes,3,opt,name=tx_response,json=txResponse,proto3,oneof"`
}

func (*TxEnvelope_TxRequest) isTxEnvelope_Msg() {}

func (*TxEnvelope_TxResponse) isTxEnvelope_Msg() {}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sender of the transaction - Ethereum Account Address
	// e.g 0xAf84242d70aE9D268E2bE3616ED497BA28A7b62C
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// The address of the receiver. null when it’s a contract creation transaction.
	// e.g. 0xAf84242d70aE9D268E2bE3616ED497BA28A7b62C
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// Integer of the gas provided for the transaction execution.
	Gas string `protobuf:"bytes,4,opt,name=gas,proto3" json:"gas,omitempty"`
	// Integer of the gas price used for each paid gas.
	// e.g 0xaf23b
	GasPrice  string `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasFeeCap string `protobuf:"bytes,6,opt,name=gas_fee_cap,json=gasFeeCap,proto3" json:"gas_fee_cap,omitempty"`
	GasTipCap string `protobuf:"bytes,7,opt,name=gas_tip_cap,json=gasTipCap,proto3" json:"gas_tip_cap,omitempty"`
	// Integer of the value sent with this transaction.
	// e.g 0xaf23
	Value string `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	// Integer of a nonce.
	Nonce string `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Hash of the method signature (4 bytes) followed by encoded parameters.
	// e.g 0xa9059cbb000000000000000000000000ff778b716fc07d98839f48ddb88d8be583beb684000000000000000000000000000000000000000000000000002386f26fc10000
	Data            string                  `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	TransactionType string                  `protobuf:"bytes,11,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
	AccessList      []*ethereum.AccessTuple `protobuf:"bytes,12,rep,name=access_list,json=accessList,proto3" json:"access_list,omitempty"`
	// The signed, RLP encoded transaction
	Raw string `protobuf:"bytes,13,opt,name=raw,proto3" json:"raw,omitempty"`
	//***********
	// ADDITIONAL CONTRACT FIELD DATA
	//***********
	// Contract identifier
	// e.g. "ERC20[v1.0.0]"
	Contract string `protobuf:"bytes,14,opt,name=contract,proto3" json:"contract,omitempty"`
	// Signature of the method to call on contract
	// e.g "transfer(address,uint256)"
	MethodSignature string `protobuf:"bytes,15,opt,name=method_signature,json=methodSignature,proto3" json:"method_signature,omitempty"`
	// Arguments to feed on transaction call
	Args []string `protobuf:"bytes,16,rep,name=args,proto3" json:"args,omitempty"`
	//***********
	// PRIVATE ONLY FIELDS
	//***********
	PrivateFor     []string `protobuf:"bytes,17,rep,name=private_for,json=privateFor,proto3" json:"private_for,omitempty"`
	PrivateFrom    string   `protobuf:"bytes,18,opt,name=private_from,json=privateFrom,proto3" json:"private_from,omitempty"`
	PrivateTxType  string   `protobuf:"bytes,19,opt,name=private_tx_type,json=privateTxType,proto3" json:"private_tx_type,omitempty"`
	PrivacyGroupId string   `protobuf:"bytes,20,opt,name=privacy_group_id,json=privacyGroupId,proto3" json:"privacy_group_id,omitempty"`
	MandatoryFor   []string `protobuf:"bytes,21,rep,name=mandatory_for,json=mandatoryFor,proto3" json:"mandatory_for,omitempty"`
	PrivacyFlag    int32    `protobuf:"varint,22,opt,name=privacy_flag,json=privacyFlag,proto3" json:"privacy_flag,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_types_tx_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_types_tx_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_pkg_types_tx_tx_proto_rawDescGZIP(), []int{2}
}

func (x *Params) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Params) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Params) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *Params) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *Params) GetGasFeeCap() string {
	if x != nil {
		return x.GasFeeCap
	}
	return ""
}

func (x *Params) GetGasTipCap() string {
	if x != nil {
		return x.GasTipCap
	}
	return ""
}

func (x *Params) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Params) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Params) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Params) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *Params) GetAccessList() []*ethereum.AccessTuple {
	if x != nil {
		return x.AccessList
	}
	return nil
}

func (x *Params) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *Params) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *Params) GetMethodSignature() string {
	if x != nil {
		return x.MethodSignature
	}
	return ""
}

func (x *Params) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Params) GetPrivateFor() []string {
	if x != nil {
		return x.PrivateFor
	}
	return nil
}

func (x *Params) GetPrivateFrom() string {
	if x != nil {
		return x.PrivateFrom
	}
	return ""
}

func (x *Params) GetPrivateTxType() string {
	if x != nil {
		return x.PrivateTxType
	}
	return ""
}

func (x *Params) GetPrivacyGroupId() string {
	if x != nil {
		return x.PrivacyGroupId
	}
	return ""
}

func (x *Params) GetMandatoryFor() []string {
	if x != nil {
		return x.MandatoryFor
	}
	return nil
}

func (x *Params) GetPrivacyFlag() int32 {
	if x != nil {
		return x.PrivacyFlag
	}
	return 0
}

type TxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Extra information (optional)
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the Response in UUID RFC 4122, ISO/IEC 9834-8:2005 format
	// e.g a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the Job created as part of transaction request
	// e.g 15276759-bbc6-4ead-ad51-ddfecf79cf09
	JobUUID string `protobuf:"bytes,8,opt,name=jobUUID,proto3" json:"jobUUID,omitempty"`
	// [option]
	ContextLabels map[string]string     `protobuf:"bytes,3,rep,name=context_labels,json=contextLabels,proto3" json:"context_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Transaction   *ethereum.Transaction `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Receipt       *ethereum.Receipt     `protobuf:"bytes,5,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// Name of the Chain as registered on the chain registry
	// e.g. 1 for mainnet, 3 for Ropsten
	Chain  string          `protobuf:"bytes,7,opt,name=chain,proto3" json:"chain,omitempty"`
	Errors []*error1.Error `protobuf:"bytes,6,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *TxResponse) Reset() {
	*x = TxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_types_tx_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResponse) ProtoMessage() {}

func (x *TxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_types_tx_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxResponse.ProtoReflect.Descriptor instead.
func (*TxResponse) Descriptor() ([]byte, []int) {
	return file_pkg_types_tx_tx_proto_rawDescGZIP(), []int{3}
}

func (x *TxResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *TxResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TxResponse) GetJobUUID() string {
	if x != nil {
		return x.JobUUID
	}
	return ""
}

func (x *TxResponse) GetContextLabels() map[string]string {
	if x != nil {
		return x.ContextLabels
	}
	return nil
}

func (x *TxResponse) GetTransaction() *ethereum.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TxResponse) GetReceipt() *ethereum.Receipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *TxResponse) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *TxResponse) GetErrors() []*error1.Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_pkg_types_tx_tx_proto protoreflect.FileDescriptor

var file_pkg_types_tx_tx_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x74,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x74, 0x78, 0x1a, 0x1b, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9d, 0x03, 0x0a, 0x09, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x78, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x74, 0x78, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x22,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x74, 0x78, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x78, 0x2e,
	0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x6a,
	0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x74,
	0x78, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x86, 0x02, 0x0a, 0x0a, 0x54, 0x78, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12,
	0x2e, 0x0a, 0x0a, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x78, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x74, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x0b, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x78, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x78,
	0x2e, 0x54, 0x78, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a,
	0x41, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x89, 0x05, 0x0a, 0x06, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x67, 0x61, 0x73, 0x5f, 0x66,
	0x65, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61,
	0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x12, 0x1e, 0x0a, 0x0b, 0x67, 0x61, 0x73, 0x5f, 0x74,
	0x69, 0x70, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61,
	0x73, 0x54, 0x69, 0x70, 0x43, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61,
	0x77, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x15,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x46,
	0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x66, 0x6c,
	0x61, 0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x46, 0x6c, 0x61, 0x67, 0x22, 0xd7, 0x03, 0x0a, 0x0a, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x78, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a,
	0x6f, 0x62, 0x55, 0x55, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f,
	0x62, 0x55, 0x55, 0x49, 0x44, 0x12, 0x48, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x78, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x37, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a,
	0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x87, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x54,
	0x48, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x52, 0x41, 0x57, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x54, 0x48, 0x5f, 0x53, 0x45,
	0x4e, 0x44, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x54, 0x48, 0x5f, 0x53, 0x45,
	0x4e, 0x44, 0x52, 0x41, 0x57, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x45, 0x41,
	0x5f, 0x53, 0x45, 0x4e, 0x44, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x2a, 0x8d, 0x01, 0x0a, 0x07, 0x4a, 0x6f,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x54, 0x48, 0x5f, 0x54, 0x58, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x54, 0x48, 0x5f, 0x52, 0x41, 0x57, 0x5f, 0x54, 0x58, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x54, 0x48, 0x5f, 0x4f, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x41, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x58, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x54, 0x48, 0x5f, 0x4f, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x45, 0x41, 0x5f, 0x54, 0x58, 0x10,
	0x03, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x54, 0x48, 0x5f, 0x54, 0x45, 0x53, 0x53, 0x45, 0x52, 0x41,
	0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x58, 0x10, 0x04, 0x12, 0x1a, 0x0a,
	0x16, 0x45, 0x54, 0x48, 0x5f, 0x54, 0x45, 0x53, 0x53, 0x45, 0x52, 0x41, 0x5f, 0x50, 0x52, 0x49,
	0x56, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x58, 0x10, 0x05, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x79,
	0x73, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_types_tx_tx_proto_rawDescOnce sync.Once
	file_pkg_types_tx_tx_proto_rawDescData = file_pkg_types_tx_tx_proto_rawDesc
)

func file_pkg_types_tx_tx_proto_rawDescGZIP() []byte {
	file_pkg_types_tx_tx_proto_rawDescOnce.Do(func() {
		file_pkg_types_tx_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_types_tx_tx_proto_rawDescData)
	})
	return file_pkg_types_tx_tx_proto_rawDescData
}

var file_pkg_types_tx_tx_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_types_tx_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_types_tx_tx_proto_goTypes = []interface{}{
	(Method)(0),                  // 0: tx.Method
	(JobType)(0),                 // 1: tx.JobType
	(*TxRequest)(nil),            // 2: tx.TxRequest
	(*TxEnvelope)(nil),           // 3: tx.TxEnvelope
	(*Params)(nil),               // 4: tx.Params
	(*TxResponse)(nil),           // 5: tx.TxResponse
	nil,                          // 6: tx.TxRequest.HeadersEntry
	nil,                          // 7: tx.TxRequest.ContextLabelsEntry
	nil,                          // 8: tx.TxEnvelope.InternalLabelsEntry
	nil,                          // 9: tx.TxResponse.HeadersEntry
	nil,                          // 10: tx.TxResponse.ContextLabelsEntry
	(*ethereum.AccessTuple)(nil), // 11: ethereum.AccessTuple
	(*ethereum.Transaction)(nil), // 12: ethereum.Transaction
	(*ethereum.Receipt)(nil),     // 13: ethereum.Receipt
	(*error1.Error)(nil),         // 14: error.Error
}
var file_pkg_types_tx_tx_proto_depIdxs = []int32{
	6,  // 0: tx.TxRequest.headers:type_name -> tx.TxRequest.HeadersEntry
	0,  // 1: tx.TxRequest.method:type_name -> tx.Method
	4,  // 2: tx.TxRequest.params:type_name -> tx.Params
	7,  // 3: tx.TxRequest.context_labels:type_name -> tx.TxRequest.ContextLabelsEntry
	1,  // 4: tx.TxRequest.jobType:type_name -> tx.JobType
	2,  // 5: tx.TxEnvelope.tx_request:type_name -> tx.TxRequest
	5,  // 6: tx.TxEnvelope.tx_response:type_name -> tx.TxResponse
	8,  // 7: tx.TxEnvelope.internal_labels:type_name -> tx.TxEnvelope.InternalLabelsEntry
	11, // 8: tx.Params.access_list:type_name -> ethereum.AccessTuple
	9,  // 9: tx.TxResponse.headers:type_name -> tx.TxResponse.HeadersEntry
	10, // 10: tx.TxResponse.context_labels:type_name -> tx.TxResponse.ContextLabelsEntry
	12, // 11: tx.TxResponse.transaction:type_name -> ethereum.Transaction
	13, // 12: tx.TxResponse.receipt:type_name -> ethereum.Receipt
	14, // 13: tx.TxResponse.errors:type_name -> error.Error
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_pkg_types_tx_tx_proto_init() }
func file_pkg_types_tx_tx_proto_init() {
	if File_pkg_types_tx_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_types_tx_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_types_tx_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxEnvelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_types_tx_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_types_tx_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_types_tx_tx_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TxEnvelope_TxRequest)(nil),
		(*TxEnvelope_TxResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_types_tx_tx_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_types_tx_tx_proto_goTypes,
		DependencyIndexes: file_pkg_types_tx_tx_proto_depIdxs,
		EnumInfos:         file_pkg_types_tx_tx_proto_enumTypes,
		MessageInfos:      file_pkg_types_tx_tx_proto_msgTypes,
	}.Build()
	File_pkg_types_tx_tx_proto = out.File
	file_pkg_types_tx_tx_proto_rawDesc = nil
	file_pkg_types_tx_tx_proto_goTypes = nil
	file_pkg_types_tx_tx_proto_depIdxs = nil
}
