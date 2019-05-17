// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/abi/abi.proto

package abi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Contract holds contract information
type Contract struct {
	// Registry of contract
	// e.g. "registry.consensys.net/corestack"
	Registry string `protobuf:"bytes,1,opt,name=registry,proto3" json:"registry,omitempty"`
	// Name of contract
	// e.g. "ERC20"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Contract tag
	// e.g "v2.1.3"
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	// Contract ABI (optional)
	Abi []byte `protobuf:"bytes,4,opt,name=abi,proto3" json:"abi,omitempty"`
	// Contract deployment bytecode (optional)
	Bytecode []byte `protobuf:"bytes,5,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	// Contract deployed bytecode (optional)
	DeployedBytecode     []byte   `protobuf:"bytes,6,opt,name=deployedBytecode,proto3" json:"deployedBytecode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad872b616a735d7, []int{0}
}

func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return xxx_messageInfo_Contract.Size(m)
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *Contract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contract) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Contract) GetAbi() []byte {
	if m != nil {
		return m.Abi
	}
	return nil
}

func (m *Contract) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

func (m *Contract) GetDeployedBytecode() []byte {
	if m != nil {
		return m.DeployedBytecode
	}
	return nil
}

// Method
type Method struct {
	// Signature of the method to call on contract
	// e.g "transfer(address,uint256)"
	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// ABI of the method (Optional)
	Abi                  []byte   `protobuf:"bytes,2,opt,name=abi,proto3" json:"abi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method) Reset()         { *m = Method{} }
func (m *Method) String() string { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()    {}
func (*Method) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad872b616a735d7, []int{1}
}

func (m *Method) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method.Unmarshal(m, b)
}
func (m *Method) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method.Marshal(b, m, deterministic)
}
func (m *Method) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method.Merge(m, src)
}
func (m *Method) XXX_Size() int {
	return xxx_messageInfo_Method.Size(m)
}
func (m *Method) XXX_DiscardUnknown() {
	xxx_messageInfo_Method.DiscardUnknown(m)
}

var xxx_messageInfo_Method proto.InternalMessageInfo

func (m *Method) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Method) GetAbi() []byte {
	if m != nil {
		return m.Abi
	}
	return nil
}

func init() {
	proto.RegisterType((*Contract)(nil), "abi.Contract")
	proto.RegisterType((*Method)(nil), "abi.Method")
}

func init() { proto.RegisterFile("types/abi/abi.proto", fileDescriptor_dad872b616a735d7) }

var fileDescriptor_dad872b616a735d7 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4b, 0x43, 0x31,
	0x14, 0x85, 0x79, 0x6d, 0x7d, 0xb4, 0xc1, 0xa1, 0xc4, 0x25, 0x88, 0x43, 0xe9, 0x54, 0x04, 0x5f,
	0x06, 0x17, 0x41, 0x5c, 0xda, 0xd9, 0xa5, 0x6e, 0x6e, 0x37, 0x79, 0xd7, 0x18, 0xfa, 0x9a, 0xfb,
	0x48, 0xae, 0x43, 0xfe, 0x90, 0xbf, 0x53, 0x12, 0xec, 0x73, 0x70, 0xb8, 0x70, 0xce, 0xc7, 0xe1,
	0x72, 0x38, 0xe2, 0x86, 0xf3, 0x88, 0x49, 0x83, 0xf1, 0xe5, 0xba, 0x31, 0x12, 0x93, 0x9c, 0x83,
	0xf1, 0xdb, 0xef, 0x46, 0x2c, 0x0f, 0x14, 0x38, 0x82, 0x65, 0x79, 0x2b, 0x96, 0x11, 0x9d, 0x4f,
	0x1c, 0xb3, 0x6a, 0x36, 0xcd, 0x6e, 0x75, 0x9c, 0xbc, 0x94, 0x62, 0x11, 0xe0, 0x8c, 0x6a, 0x56,
	0x79, 0xd5, 0x72, 0x2d, 0xe6, 0x0c, 0x4e, 0xcd, 0x2b, 0x2a, 0xb2, 0x10, 0x30, 0x5e, 0x2d, 0x36,
	0xcd, 0xee, 0xfa, 0x58, 0x64, 0xf9, 0x69, 0x32, 0xa3, 0xa5, 0x1e, 0xd5, 0x55, 0xc5, 0x93, 0x97,
	0xf7, 0x62, 0xdd, 0xe3, 0x38, 0x50, 0xc6, 0x7e, 0x7f, 0xc9, 0xb4, 0x35, 0xf3, 0x8f, 0x6f, 0x9f,
	0x44, 0xfb, 0x8a, 0xfc, 0x49, 0xbd, 0xbc, 0x13, 0xab, 0xe4, 0x5d, 0x00, 0xfe, 0x8a, 0xf8, 0x5b,
	0xf3, 0x0f, 0x5c, 0x1a, 0xcc, 0xa6, 0x06, 0xfb, 0x97, 0xf7, 0x67, 0xe7, 0x79, 0x00, 0xd3, 0x59,
	0x3a, 0xeb, 0x03, 0x85, 0x84, 0xe1, 0x2d, 0x27, 0x6d, 0x07, 0x8f, 0x81, 0xf5, 0x47, 0xd4, 0x96,
	0x22, 0x3e, 0x24, 0x06, 0x7b, 0xd2, 0xe3, 0xc9, 0x75, 0xce, 0xb3, 0x9e, 0x06, 0x33, 0x6d, 0x5d,
	0xeb, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x40, 0x39, 0x4a, 0x44, 0x01, 0x00, 0x00,
}
