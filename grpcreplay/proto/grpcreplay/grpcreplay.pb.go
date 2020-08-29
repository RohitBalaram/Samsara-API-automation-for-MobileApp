// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcreplay.proto

package grpcreplay

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Entry_Kind int32

const (
	Entry_TYPE_UNSPECIFIED Entry_Kind = 0
	// A unary request.
	// method: the full name of the method
	// message: the request proto
	// is_error: false
	// ref_index: 0
	Entry_REQUEST Entry_Kind = 1
	// A unary response.
	// method: the full name of the method
	// message:
	//   if is_error: a google.rpc.Status proto
	//   else:        the response proto
	// ref_index: index in the sequence of Entries of matching request (1-based)
	Entry_RESPONSE Entry_Kind = 2
	// A method that creates a stream.
	// method: the full name of the method
	// message:
	//   if is_error: a google.rpc.Status proto
	//   else:        nil
	// ref_index: 0
	Entry_CREATE_STREAM Entry_Kind = 3
	// A call to Send on the client returned by a stream-creating method.
	// method: unset
	// message: the proto being sent
	// is_error: false
	// ref_index: index of matching CREATE_STREAM entry (1-based)
	Entry_SEND Entry_Kind = 4
	// A call to Recv on the client returned by a stream-creating method.
	// method: unset
	// message:
	//   if is_error: a google.rpc.Status proto, or nil on EOF
	//   else:        the received message
	// ref_index: index of matching CREATE_STREAM entry
	Entry_RECV Entry_Kind = 5
)

var Entry_Kind_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "REQUEST",
	2: "RESPONSE",
	3: "CREATE_STREAM",
	4: "SEND",
	5: "RECV",
}

var Entry_Kind_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"REQUEST":          1,
	"RESPONSE":         2,
	"CREATE_STREAM":    3,
	"SEND":             4,
	"RECV":             5,
}

func (x Entry_Kind) String() string {
	return proto.EnumName(Entry_Kind_name, int32(x))
}

func (Entry_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4aab7311a0847da, []int{0, 0}
}

// An Entry represents a single RPC activity, typically a request or response.
type Entry struct {
	Kind     Entry_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=grpcreplay.Entry_Kind" json:"kind,omitempty"`
	Method   string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Message  *any.Any   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	IsError  bool       `protobuf:"varint,4,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
	RefIndex int32      `protobuf:"varint,5,opt,name=ref_index,json=refIndex,proto3" json:"ref_index,omitempty"`
	// for SEND/RECV, index of CREATE_STREAM
	// since the start of this recording, in golang duration unit
	Delay                int64    `protobuf:"varint,6,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4aab7311a0847da, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetKind() Entry_Kind {
	if m != nil {
		return m.Kind
	}
	return Entry_TYPE_UNSPECIFIED
}

func (m *Entry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Entry) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Entry) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *Entry) GetRefIndex() int32 {
	if m != nil {
		return m.RefIndex
	}
	return 0
}

func (m *Entry) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func init() {
	proto.RegisterEnum("grpcreplay.Entry_Kind", Entry_Kind_name, Entry_Kind_value)
	proto.RegisterType((*Entry)(nil), "grpcreplay.Entry")
}

func init() { proto.RegisterFile("grpcreplay.proto", fileDescriptor_a4aab7311a0847da) }

var fileDescriptor_a4aab7311a0847da = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5d, 0x68, 0xa1, 0x0c, 0x6a, 0xd6, 0x09, 0x21, 0x45, 0x2f, 0x0d, 0xa7, 0xc6, 0xc3,
	0x92, 0xe0, 0x13, 0x10, 0x18, 0x13, 0x62, 0x44, 0xdc, 0x16, 0x13, 0x2f, 0x36, 0xc5, 0x2e, 0xb5,
	0x11, 0x5a, 0xb2, 0xad, 0x89, 0x7d, 0x24, 0xdf, 0xd2, 0xb4, 0x40, 0xf4, 0x36, 0xdf, 0xcc, 0x97,
	0xf9, 0x67, 0x80, 0xc7, 0x7a, 0xff, 0xae, 0xd5, 0x7e, 0x1b, 0x96, 0x62, 0xaf, 0xb3, 0x22, 0x43,
	0xf8, 0xeb, 0x5c, 0x0f, 0xe2, 0x2c, 0x8b, 0xb7, 0x6a, 0x54, 0x4f, 0xd6, 0x5f, 0x9b, 0x51, 0x98,
	0x1e, 0xb5, 0xe1, 0x4f, 0x03, 0x4c, 0x4a, 0x0b, 0x5d, 0xe2, 0x2d, 0x18, 0x9f, 0x49, 0x1a, 0xd9,
	0xcc, 0x61, 0xee, 0xe5, 0xb8, 0x2f, 0xfe, 0x6d, 0xac, 0x05, 0xf1, 0x90, 0xa4, 0x91, 0xac, 0x1d,
	0xec, 0x43, 0x6b, 0xa7, 0x8a, 0x8f, 0x2c, 0xb2, 0x1b, 0x0e, 0x73, 0x3b, 0xf2, 0x48, 0x28, 0xa0,
	0xbd, 0x53, 0x79, 0x1e, 0xc6, 0xca, 0x6e, 0x3a, 0xcc, 0xed, 0x8e, 0x7b, 0xe2, 0x10, 0x2d, 0x4e,
	0xd1, 0x62, 0x92, 0x96, 0xf2, 0x24, 0xe1, 0x00, 0xac, 0x24, 0x0f, 0x94, 0xd6, 0x99, 0xb6, 0x0d,
	0x87, 0xb9, 0x96, 0x6c, 0x27, 0x39, 0x55, 0x88, 0x37, 0xd0, 0xd1, 0x6a, 0x13, 0x24, 0x69, 0xa4,
	0xbe, 0x6d, 0xd3, 0x61, 0xae, 0x29, 0x2d, 0xad, 0x36, 0xf3, 0x8a, 0xb1, 0x07, 0x66, 0xa4, 0xb6,
	0x61, 0x69, 0xb7, 0x1c, 0xe6, 0x36, 0xe5, 0x01, 0x86, 0x6f, 0x60, 0x54, 0x37, 0x62, 0x0f, 0xb8,
	0xff, 0xba, 0xa4, 0x60, 0xb5, 0xf0, 0x96, 0x34, 0x9d, 0xdf, 0xcf, 0x69, 0xc6, 0xcf, 0xb0, 0x0b,
	0x6d, 0x49, 0xcf, 0x2b, 0xf2, 0x7c, 0xce, 0xf0, 0x1c, 0x2c, 0x49, 0xde, 0xf2, 0x69, 0xe1, 0x11,
	0x6f, 0xe0, 0x15, 0x5c, 0x4c, 0x25, 0x4d, 0x7c, 0x0a, 0x3c, 0x5f, 0xd2, 0xe4, 0x91, 0x37, 0xd1,
	0x02, 0xc3, 0xa3, 0xc5, 0x8c, 0x1b, 0x55, 0x25, 0x69, 0xfa, 0xc2, 0xcd, 0x75, 0xab, 0x7e, 0xe2,
	0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x2a, 0xe4, 0x30, 0x6d, 0x01, 0x00, 0x00,
}
