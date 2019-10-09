// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ss_gateway.proto

package ss

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

type EMsgID int32

const (
	EMsgID_Invalid     EMsgID = 0
	EMsgID_ClientNew   EMsgID = 1
	EMsgID_ClientAuth  EMsgID = 2
	EMsgID_ClientClose EMsgID = 3
	EMsgID_Other       EMsgID = 4096
)

var EMsgID_name = map[int32]string{
	0:    "Invalid",
	1:    "ClientNew",
	2:    "ClientAuth",
	3:    "ClientClose",
	4096: "Other",
}

var EMsgID_value = map[string]int32{
	"Invalid":     0,
	"ClientNew":   1,
	"ClientAuth":  2,
	"ClientClose": 3,
	"Other":       4096,
}

func (x EMsgID) String() string {
	return proto.EnumName(EMsgID_name, int32(x))
}

func (EMsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{0}
}

type Head struct {
	Msgid                EMsgID   `protobuf:"varint,1,opt,name=msgid,proto3,enum=ss.EMsgID" json:"msgid,omitempty"`
	ClientId             uint32   `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Uid                  uint64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head) Reset()         { *m = Head{} }
func (m *Head) String() string { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()    {}
func (*Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{0}
}

func (m *Head) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Head.Unmarshal(m, b)
}
func (m *Head) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Head.Marshal(b, m, deterministic)
}
func (m *Head) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head.Merge(m, src)
}
func (m *Head) XXX_Size() int {
	return xxx_messageInfo_Head.Size(m)
}
func (m *Head) XXX_DiscardUnknown() {
	xxx_messageInfo_Head.DiscardUnknown(m)
}

var xxx_messageInfo_Head proto.InternalMessageInfo

func (m *Head) GetMsgid() EMsgID {
	if m != nil {
		return m.Msgid
	}
	return EMsgID_Invalid
}

func (m *Head) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Head) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type ClientNewReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientNewReq) Reset()         { *m = ClientNewReq{} }
func (m *ClientNewReq) String() string { return proto.CompactTextString(m) }
func (*ClientNewReq) ProtoMessage()    {}
func (*ClientNewReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{1}
}

func (m *ClientNewReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientNewReq.Unmarshal(m, b)
}
func (m *ClientNewReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientNewReq.Marshal(b, m, deterministic)
}
func (m *ClientNewReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientNewReq.Merge(m, src)
}
func (m *ClientNewReq) XXX_Size() int {
	return xxx_messageInfo_ClientNewReq.Size(m)
}
func (m *ClientNewReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientNewReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClientNewReq proto.InternalMessageInfo

type ClientNewRes struct {
	Close                bool     `protobuf:"varint,1,opt,name=close,proto3" json:"close,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientNewRes) Reset()         { *m = ClientNewRes{} }
func (m *ClientNewRes) String() string { return proto.CompactTextString(m) }
func (*ClientNewRes) ProtoMessage()    {}
func (*ClientNewRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{2}
}

func (m *ClientNewRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientNewRes.Unmarshal(m, b)
}
func (m *ClientNewRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientNewRes.Marshal(b, m, deterministic)
}
func (m *ClientNewRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientNewRes.Merge(m, src)
}
func (m *ClientNewRes) XXX_Size() int {
	return xxx_messageInfo_ClientNewRes.Size(m)
}
func (m *ClientNewRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientNewRes.DiscardUnknown(m)
}

var xxx_messageInfo_ClientNewRes proto.InternalMessageInfo

func (m *ClientNewRes) GetClose() bool {
	if m != nil {
		return m.Close
	}
	return false
}

type ClientAuthReq struct {
	Openid               string   `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid,omitempty"`
	Openkey              string   `protobuf:"bytes,2,opt,name=openkey,proto3" json:"openkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientAuthReq) Reset()         { *m = ClientAuthReq{} }
func (m *ClientAuthReq) String() string { return proto.CompactTextString(m) }
func (*ClientAuthReq) ProtoMessage()    {}
func (*ClientAuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{3}
}

func (m *ClientAuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientAuthReq.Unmarshal(m, b)
}
func (m *ClientAuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientAuthReq.Marshal(b, m, deterministic)
}
func (m *ClientAuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientAuthReq.Merge(m, src)
}
func (m *ClientAuthReq) XXX_Size() int {
	return xxx_messageInfo_ClientAuthReq.Size(m)
}
func (m *ClientAuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientAuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClientAuthReq proto.InternalMessageInfo

func (m *ClientAuthReq) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *ClientAuthReq) GetOpenkey() string {
	if m != nil {
		return m.Openkey
	}
	return ""
}

type ClientAuthRes struct {
	Close                bool     `protobuf:"varint,1,opt,name=close,proto3" json:"close,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientAuthRes) Reset()         { *m = ClientAuthRes{} }
func (m *ClientAuthRes) String() string { return proto.CompactTextString(m) }
func (*ClientAuthRes) ProtoMessage()    {}
func (*ClientAuthRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{4}
}

func (m *ClientAuthRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientAuthRes.Unmarshal(m, b)
}
func (m *ClientAuthRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientAuthRes.Marshal(b, m, deterministic)
}
func (m *ClientAuthRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientAuthRes.Merge(m, src)
}
func (m *ClientAuthRes) XXX_Size() int {
	return xxx_messageInfo_ClientAuthRes.Size(m)
}
func (m *ClientAuthRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientAuthRes.DiscardUnknown(m)
}

var xxx_messageInfo_ClientAuthRes proto.InternalMessageInfo

func (m *ClientAuthRes) GetClose() bool {
	if m != nil {
		return m.Close
	}
	return false
}

type ClientCloseReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCloseReq) Reset()         { *m = ClientCloseReq{} }
func (m *ClientCloseReq) String() string { return proto.CompactTextString(m) }
func (*ClientCloseReq) ProtoMessage()    {}
func (*ClientCloseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{5}
}

func (m *ClientCloseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCloseReq.Unmarshal(m, b)
}
func (m *ClientCloseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCloseReq.Marshal(b, m, deterministic)
}
func (m *ClientCloseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCloseReq.Merge(m, src)
}
func (m *ClientCloseReq) XXX_Size() int {
	return xxx_messageInfo_ClientCloseReq.Size(m)
}
func (m *ClientCloseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCloseReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCloseReq proto.InternalMessageInfo

type ClientCloseRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCloseRes) Reset()         { *m = ClientCloseRes{} }
func (m *ClientCloseRes) String() string { return proto.CompactTextString(m) }
func (*ClientCloseRes) ProtoMessage()    {}
func (*ClientCloseRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2059a492e69a85d, []int{6}
}

func (m *ClientCloseRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCloseRes.Unmarshal(m, b)
}
func (m *ClientCloseRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCloseRes.Marshal(b, m, deterministic)
}
func (m *ClientCloseRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCloseRes.Merge(m, src)
}
func (m *ClientCloseRes) XXX_Size() int {
	return xxx_messageInfo_ClientCloseRes.Size(m)
}
func (m *ClientCloseRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCloseRes.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCloseRes proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ss.EMsgID", EMsgID_name, EMsgID_value)
	proto.RegisterType((*Head)(nil), "ss.Head")
	proto.RegisterType((*ClientNewReq)(nil), "ss.ClientNewReq")
	proto.RegisterType((*ClientNewRes)(nil), "ss.ClientNewRes")
	proto.RegisterType((*ClientAuthReq)(nil), "ss.ClientAuthReq")
	proto.RegisterType((*ClientAuthRes)(nil), "ss.ClientAuthRes")
	proto.RegisterType((*ClientCloseReq)(nil), "ss.ClientCloseReq")
	proto.RegisterType((*ClientCloseRes)(nil), "ss.ClientCloseRes")
}

func init() { proto.RegisterFile("ss_gateway.proto", fileDescriptor_b2059a492e69a85d) }

var fileDescriptor_b2059a492e69a85d = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x24, 0x4d, 0xda, 0x4c, 0xdf, 0xc4, 0x65, 0x10, 0x09, 0x78, 0x09, 0x41, 0x21,
	0x78, 0xc8, 0x41, 0x3f, 0x41, 0xa9, 0x82, 0x39, 0xa8, 0xb8, 0x17, 0x8f, 0x25, 0x76, 0x87, 0x34,
	0x18, 0x93, 0xb6, 0xb3, 0xb5, 0xf4, 0xe6, 0x47, 0x97, 0x64, 0xeb, 0x9f, 0x0a, 0xde, 0xe6, 0x37,
	0x33, 0xcc, 0xf3, 0xcc, 0x03, 0x82, 0x79, 0x56, 0x16, 0x9a, 0xb6, 0xc5, 0x2e, 0x5b, 0xae, 0x5b,
	0xdd, 0xa2, 0xcd, 0x9c, 0x3c, 0xc1, 0xe0, 0x96, 0x0a, 0x85, 0x31, 0xb8, 0xaf, 0x5c, 0x56, 0x2a,
	0xb2, 0x62, 0x2b, 0x0d, 0x2f, 0x21, 0x63, 0xce, 0x6e, 0xee, 0xb8, 0xcc, 0xaf, 0xa5, 0x19, 0xe0,
	0x29, 0xf8, 0xf3, 0xba, 0xa2, 0x46, 0xcf, 0x2a, 0x15, 0xd9, 0xb1, 0x95, 0x06, 0x72, 0x64, 0x1a,
	0xb9, 0x42, 0x01, 0xce, 0xa6, 0x52, 0x91, 0x13, 0x5b, 0xe9, 0x40, 0x76, 0x65, 0x12, 0xc2, 0xff,
	0x69, 0x3f, 0xbd, 0xa7, 0xad, 0xa4, 0x55, 0x72, 0x76, 0xc0, 0x8c, 0xc7, 0xe0, 0xce, 0xeb, 0x96,
	0xa9, 0x17, 0x1c, 0x49, 0x03, 0xc9, 0x04, 0x02, 0xb3, 0x35, 0xd9, 0xe8, 0x85, 0xa4, 0x15, 0x9e,
	0x80, 0xd7, 0x2e, 0xa9, 0xd9, 0x1b, 0xf3, 0xe5, 0x9e, 0x30, 0x82, 0x61, 0x57, 0xbd, 0xd0, 0xae,
	0xf7, 0xe2, 0xcb, 0x4f, 0x4c, 0xce, 0x0f, 0x4f, 0xfc, 0xa5, 0x24, 0x20, 0x34, 0x6b, 0xd3, 0x0e,
	0x3b, 0x87, 0xbf, 0x3b, 0x7c, 0xf1, 0x08, 0x9e, 0xc9, 0x00, 0xc7, 0x30, 0xcc, 0x9b, 0xb7, 0xa2,
	0xae, 0x94, 0xf8, 0x87, 0x01, 0xf8, 0x5f, 0xaf, 0x08, 0x0b, 0x43, 0x80, 0x6f, 0x41, 0x61, 0xe3,
	0x11, 0x8c, 0x7f, 0xdc, 0x11, 0x0e, 0x02, 0xb8, 0x0f, 0x7a, 0x41, 0x6b, 0xf1, 0x1e, 0x3f, 0x7b,
	0x7d, 0xf4, 0x57, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x42, 0x48, 0x28, 0x8e, 0x01, 0x00,
	0x00,
}