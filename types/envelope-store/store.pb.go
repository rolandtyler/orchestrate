// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/envelope-store/store.proto

package envelope_store

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	common "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/common"
	envelope "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/envelope"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StoreRequest struct {
	Envelope             *envelope.Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{0}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetEnvelope() *envelope.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

type StoreResponse struct {
	// Status of Envelope
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Last update date of envelope stored
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// Envelope
	Envelope *envelope.Envelope `protobuf:"bytes,3,opt,name=envelope,proto3" json:"envelope,omitempty"`
	// Error
	Err                  *common.Error `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{1}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StoreResponse) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *StoreResponse) GetEnvelope() *envelope.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

func (m *StoreResponse) GetErr() *common.Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type TxHashRequest struct {
	// Chain ID the transaction has been sent to
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Hash of the transaction
	TxHash               string   `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxHashRequest) Reset()         { *m = TxHashRequest{} }
func (m *TxHashRequest) String() string { return proto.CompactTextString(m) }
func (*TxHashRequest) ProtoMessage()    {}
func (*TxHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{2}
}

func (m *TxHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxHashRequest.Unmarshal(m, b)
}
func (m *TxHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxHashRequest.Marshal(b, m, deterministic)
}
func (m *TxHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxHashRequest.Merge(m, src)
}
func (m *TxHashRequest) XXX_Size() int {
	return xxx_messageInfo_TxHashRequest.Size(m)
}
func (m *TxHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxHashRequest proto.InternalMessageInfo

func (m *TxHashRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *TxHashRequest) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type IDRequest struct {
	// Envelope identifier
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{3}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SetStatusRequest struct {
	// Envelope identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStatusRequest) Reset()         { *m = SetStatusRequest{} }
func (m *SetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetStatusRequest) ProtoMessage()    {}
func (*SetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{4}
}

func (m *SetStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStatusRequest.Unmarshal(m, b)
}
func (m *SetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStatusRequest.Marshal(b, m, deterministic)
}
func (m *SetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStatusRequest.Merge(m, src)
}
func (m *SetStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SetStatusRequest.Size(m)
}
func (m *SetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetStatusRequest proto.InternalMessageInfo

func (m *SetStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetStatusRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type LoadPendingRequest struct {
	// Pending duration in nanoseconds
	Duration             int64    `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadPendingRequest) Reset()         { *m = LoadPendingRequest{} }
func (m *LoadPendingRequest) String() string { return proto.CompactTextString(m) }
func (*LoadPendingRequest) ProtoMessage()    {}
func (*LoadPendingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{5}
}

func (m *LoadPendingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPendingRequest.Unmarshal(m, b)
}
func (m *LoadPendingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPendingRequest.Marshal(b, m, deterministic)
}
func (m *LoadPendingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPendingRequest.Merge(m, src)
}
func (m *LoadPendingRequest) XXX_Size() int {
	return xxx_messageInfo_LoadPendingRequest.Size(m)
}
func (m *LoadPendingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadPendingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadPendingRequest proto.InternalMessageInfo

func (m *LoadPendingRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type LoadPendingResponse struct {
	// Pending envelopes
	Envelopes []*envelope.Envelope `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
	// Error
	Err                  *common.Error `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoadPendingResponse) Reset()         { *m = LoadPendingResponse{} }
func (m *LoadPendingResponse) String() string { return proto.CompactTextString(m) }
func (*LoadPendingResponse) ProtoMessage()    {}
func (*LoadPendingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8943bd8e2a955ec2, []int{6}
}

func (m *LoadPendingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPendingResponse.Unmarshal(m, b)
}
func (m *LoadPendingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPendingResponse.Marshal(b, m, deterministic)
}
func (m *LoadPendingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPendingResponse.Merge(m, src)
}
func (m *LoadPendingResponse) XXX_Size() int {
	return xxx_messageInfo_LoadPendingResponse.Size(m)
}
func (m *LoadPendingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadPendingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadPendingResponse proto.InternalMessageInfo

func (m *LoadPendingResponse) GetEnvelopes() []*envelope.Envelope {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func (m *LoadPendingResponse) GetErr() *common.Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreRequest)(nil), "envelopestore.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "envelopestore.StoreResponse")
	proto.RegisterType((*TxHashRequest)(nil), "envelopestore.TxHashRequest")
	proto.RegisterType((*IDRequest)(nil), "envelopestore.IDRequest")
	proto.RegisterType((*SetStatusRequest)(nil), "envelopestore.SetStatusRequest")
	proto.RegisterType((*LoadPendingRequest)(nil), "envelopestore.LoadPendingRequest")
	proto.RegisterType((*LoadPendingResponse)(nil), "envelopestore.LoadPendingResponse")
}

func init() { proto.RegisterFile("types/envelope-store/store.proto", fileDescriptor_8943bd8e2a955ec2) }

var fileDescriptor_8943bd8e2a955ec2 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x5b, 0xe8, 0x36, 0xd3, 0x16, 0x21, 0x23, 0x41, 0xc8, 0x82, 0x5a, 0x72, 0xe2, 0x82,
	0xbd, 0x5a, 0x6e, 0x20, 0x38, 0xb4, 0xbb, 0xb0, 0x45, 0x1c, 0x50, 0x5a, 0x2e, 0x5c, 0x2a, 0x37,
	0xf1, 0x26, 0xd1, 0xa6, 0x71, 0xb0, 0x1d, 0xb4, 0xfd, 0x2e, 0x0e, 0xfc, 0x1e, 0x8a, 0x13, 0x87,
	0x26, 0x6c, 0xa9, 0xc4, 0xa5, 0xf2, 0x64, 0xde, 0x3c, 0xbf, 0xbc, 0x37, 0x0d, 0x4c, 0xd5, 0x2e,
	0x63, 0x92, 0xb0, 0xf4, 0x07, 0x4b, 0x78, 0xc6, 0x5e, 0x49, 0xc5, 0x05, 0x23, 0xfa, 0x17, 0x67,
	0x82, 0x2b, 0x8e, 0xc6, 0xa6, 0xa7, 0x1f, 0x3a, 0xcf, 0x9b, 0x03, 0xf5, 0xa1, 0x44, 0x3b, 0x76,
	0xd9, 0xf6, 0xf9, 0x76, 0xcb, 0x53, 0xc2, 0x84, 0xe0, 0xa2, 0xea, 0x4c, 0x42, 0xce, 0xc3, 0x84,
	0x11, 0x5d, 0x6d, 0xf2, 0x6b, 0xa2, 0xe2, 0x2d, 0x93, 0x8a, 0x6e, 0xb3, 0x12, 0xe0, 0xbe, 0x87,
	0xd1, 0xb2, 0xb8, 0xc2, 0x63, 0xdf, 0x73, 0x26, 0x15, 0xc2, 0x30, 0x30, 0xe4, 0x76, 0x67, 0xda,
	0x79, 0x39, 0x3c, 0x47, 0xb8, 0xbe, 0xed, 0xb2, 0x3a, 0x78, 0x35, 0xc6, 0xfd, 0xd5, 0x81, 0x71,
	0x45, 0x20, 0x33, 0x9e, 0x4a, 0x86, 0x1e, 0x43, 0x5f, 0x2a, 0xaa, 0x72, 0xa9, 0xe7, 0x2d, 0xaf,
	0xaa, 0xd0, 0x3b, 0x18, 0x25, 0x54, 0xaa, 0x75, 0x9e, 0x05, 0x54, 0xb1, 0xc0, 0xee, 0x6a, 0x76,
	0x07, 0x97, 0x0a, 0xb1, 0x51, 0x88, 0x57, 0x46, 0xa1, 0x37, 0x2c, 0xf0, 0x5f, 0x4b, 0x78, 0x43,
	0x58, 0xef, 0xb8, 0x30, 0x34, 0x81, 0x1e, 0x13, 0xc2, 0xbe, 0xa7, 0xa1, 0x63, 0x5c, 0x7a, 0x83,
	0x2f, 0x0b, 0x6f, 0xbc, 0xa2, 0xe3, 0xce, 0x61, 0xbc, 0xba, 0xbd, 0xa2, 0x32, 0x32, 0xaf, 0xfe,
	0x14, 0x06, 0x7e, 0x44, 0xe3, 0x74, 0x1d, 0x07, 0x95, 0xf4, 0x13, 0x5d, 0x2f, 0x02, 0xf4, 0x04,
	0x4e, 0xd4, 0xed, 0x3a, 0xa2, 0x32, 0xd2, 0xb2, 0x2d, 0xaf, 0xaf, 0xf4, 0xa8, 0x7b, 0x0a, 0xd6,
	0xe2, 0xc2, 0x10, 0x3c, 0x80, 0x6e, 0x3d, 0xda, 0x8d, 0x03, 0xf7, 0x0d, 0x3c, 0x5c, 0x32, 0xb5,
	0xd4, 0xaf, 0x7f, 0x00, 0xb3, 0xe7, 0x56, 0x77, 0xdf, 0x2d, 0xf7, 0x0c, 0xd0, 0x67, 0x4e, 0x83,
	0x2f, 0x2c, 0x0d, 0xe2, 0x34, 0x34, 0xd3, 0x0e, 0x0c, 0x82, 0x5c, 0x50, 0x15, 0xf3, 0x54, 0x73,
	0xf4, 0xbc, 0xba, 0x76, 0x23, 0x78, 0xd4, 0x98, 0xa8, 0xe2, 0x38, 0x03, 0xab, 0xde, 0x25, 0xbb,
	0x33, 0xed, 0x1d, 0x30, 0xee, 0x0f, 0xc8, 0x38, 0xd7, 0x3d, 0xe4, 0xdc, 0xf9, 0xcf, 0x1e, 0xdc,
	0xd7, 0x99, 0xa3, 0x99, 0x39, 0x9c, 0xe2, 0xc6, 0xc2, 0xe2, 0xfd, 0x9d, 0x72, 0x9e, 0xdd, 0xdd,
	0xac, 0x04, 0x7e, 0x82, 0x51, 0xa1, 0x7b, 0xb6, 0x2b, 0xd3, 0x40, 0x6d, 0x74, 0x23, 0xa4, 0x23,
	0x5c, 0x33, 0x18, 0x94, 0x5c, 0x8b, 0x0b, 0x64, 0xb7, 0x90, 0x75, 0x4e, 0x47, 0x38, 0xe6, 0x60,
	0x7d, 0x34, 0xa9, 0xfd, 0x37, 0xc9, 0x5b, 0xb0, 0xea, 0xe8, 0xd1, 0xa4, 0x0d, 0x6d, 0x2d, 0x85,
	0xd3, 0x34, 0x19, 0xad, 0x60, 0xb8, 0x97, 0x24, 0x7a, 0xd1, 0x1a, 0xff, 0x7b, 0x2f, 0x1c, 0xf7,
	0x5f, 0x90, 0x52, 0xd2, 0xec, 0xea, 0xdb, 0x87, 0x30, 0x56, 0x09, 0xdd, 0x14, 0x97, 0x91, 0x79,
	0xf1, 0x2c, 0x5d, 0xee, 0x24, 0xf1, 0x93, 0x98, 0xa5, 0x8a, 0x5c, 0x0b, 0xe2, 0x73, 0x51, 0x7c,
	0x8b, 0xa8, 0x7f, 0x43, 0xb2, 0x9b, 0x10, 0x87, 0xb1, 0x22, 0x77, 0x7d, 0xa8, 0x36, 0x7d, 0xfd,
	0x5f, 0x7d, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xf5, 0x1a, 0xe3, 0xc7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	// Store an envelope
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// LoadByTxHash load an envelope by transaction hash
	LoadByTxHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// LoadByID load an envelope by identifier
	LoadByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// GetStatus returns envelope status
	GetStatus(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// SetStatus set an envelope status
	SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*common.Error, error)
	// LoadPending load envelopes of pending transactions
	LoadPending(ctx context.Context, in *LoadPendingRequest, opts ...grpc.CallOption) (*LoadPendingResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) LoadByTxHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/LoadByTxHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) LoadByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/LoadByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetStatus(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) LoadPending(ctx context.Context, in *LoadPendingRequest, opts ...grpc.CallOption) (*LoadPendingResponse, error) {
	out := new(LoadPendingResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/LoadPending", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	// Store an envelope
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	// LoadByTxHash load an envelope by transaction hash
	LoadByTxHash(context.Context, *TxHashRequest) (*StoreResponse, error)
	// LoadByID load an envelope by identifier
	LoadByID(context.Context, *IDRequest) (*StoreResponse, error)
	// GetStatus returns envelope status
	GetStatus(context.Context, *IDRequest) (*StoreResponse, error)
	// SetStatus set an envelope status
	SetStatus(context.Context, *SetStatusRequest) (*common.Error, error)
	// LoadPending load envelopes of pending transactions
	LoadPending(context.Context, *LoadPendingRequest) (*LoadPendingResponse, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) Store(ctx context.Context, req *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedStoreServer) LoadByTxHash(ctx context.Context, req *TxHashRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadByTxHash not implemented")
}
func (*UnimplementedStoreServer) LoadByID(ctx context.Context, req *IDRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadByID not implemented")
}
func (*UnimplementedStoreServer) GetStatus(ctx context.Context, req *IDRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedStoreServer) SetStatus(ctx context.Context, req *SetStatusRequest) (*common.Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (*UnimplementedStoreServer) LoadPending(ctx context.Context, req *LoadPendingRequest) (*LoadPendingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadPending not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_LoadByTxHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).LoadByTxHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/LoadByTxHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).LoadByTxHash(ctx, req.(*TxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_LoadByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).LoadByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/LoadByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).LoadByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetStatus(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).SetStatus(ctx, req.(*SetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_LoadPending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadPendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).LoadPending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/LoadPending",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).LoadPending(ctx, req.(*LoadPendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envelopestore.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Store_Store_Handler,
		},
		{
			MethodName: "LoadByTxHash",
			Handler:    _Store_LoadByTxHash_Handler,
		},
		{
			MethodName: "LoadByID",
			Handler:    _Store_LoadByID_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Store_GetStatus_Handler,
		},
		{
			MethodName: "SetStatus",
			Handler:    _Store_SetStatus_Handler,
		},
		{
			MethodName: "LoadPending",
			Handler:    _Store_LoadPending_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types/envelope-store/store.proto",
}
