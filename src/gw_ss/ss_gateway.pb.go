// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ss_gateway.proto

package gw_ss

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
	return fileDescriptor_d4ce98bec259b67b, []int{0}
}

type Head struct {
	Msgid                EMsgID   `protobuf:"varint,1,opt,name=msgid,proto3,enum=gw.ss.EMsgID" json:"msgid,omitempty"`
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
	return fileDescriptor_d4ce98bec259b67b, []int{0}
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
	return fileDescriptor_d4ce98bec259b67b, []int{1}
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
	return fileDescriptor_d4ce98bec259b67b, []int{2}
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
	return fileDescriptor_d4ce98bec259b67b, []int{3}
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
	return fileDescriptor_d4ce98bec259b67b, []int{4}
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
	return fileDescriptor_d4ce98bec259b67b, []int{5}
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
	return fileDescriptor_d4ce98bec259b67b, []int{6}
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
	proto.RegisterEnum("gw.ss.EMsgID", EMsgID_name, EMsgID_value)
	proto.RegisterType((*Head)(nil), "gw.ss.Head")
	proto.RegisterType((*ClientNewReq)(nil), "gw.ss.ClientNewReq")
	proto.RegisterType((*ClientNewRes)(nil), "gw.ss.ClientNewRes")
	proto.RegisterType((*ClientAuthReq)(nil), "gw.ss.ClientAuthReq")
	proto.RegisterType((*ClientAuthRes)(nil), "gw.ss.ClientAuthRes")
	proto.RegisterType((*ClientCloseReq)(nil), "gw.ss.ClientCloseReq")
	proto.RegisterType((*ClientCloseRes)(nil), "gw.ss.ClientCloseRes")
}

func init() { proto.RegisterFile("proto/ss_gateway.proto", fileDescriptor_d4ce98bec259b67b) }

var fileDescriptor_d4ce98bec259b67b = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0xd3, 0xa4, 0xcd, 0xd4, 0xc4, 0x65, 0x90, 0x12, 0xf0, 0x12, 0xa2, 0x42, 0xf0,
	0x10, 0x41, 0x3f, 0x41, 0xa9, 0x82, 0x39, 0xa8, 0xb8, 0x47, 0x2f, 0x25, 0x76, 0x87, 0x34, 0x18,
	0x93, 0xb6, 0xbb, 0x35, 0xf4, 0xe6, 0x47, 0x97, 0xec, 0xd6, 0x3f, 0x15, 0xbc, 0xcd, 0xef, 0xbd,
	0x61, 0xde, 0x4b, 0x16, 0xc6, 0xcb, 0x75, 0xa3, 0x9a, 0x4b, 0x29, 0x67, 0x45, 0xae, 0xa8, 0xcd,
	0xb7, 0xa9, 0x16, 0xd0, 0x29, 0xda, 0x54, 0xca, 0xf8, 0x19, 0xfa, 0x77, 0x94, 0x0b, 0x3c, 0x05,
	0xe7, 0x4d, 0x16, 0xa5, 0x08, 0xad, 0xc8, 0x4a, 0x82, 0x2b, 0x3f, 0xd5, 0x76, 0x7a, 0x7b, 0x2f,
	0x8b, 0xec, 0x86, 0x1b, 0x0f, 0x4f, 0xc0, 0x9b, 0x57, 0x25, 0xd5, 0x6a, 0x56, 0x8a, 0xb0, 0x17,
	0x59, 0x89, 0xcf, 0x87, 0x46, 0xc8, 0x04, 0x32, 0xb0, 0x37, 0xa5, 0x08, 0xed, 0xc8, 0x4a, 0xfa,
	0xbc, 0x1b, 0xe3, 0x00, 0x0e, 0xa7, 0xda, 0x7d, 0xa0, 0x96, 0xd3, 0x2a, 0x3e, 0xdb, 0x63, 0x89,
	0xc7, 0xe0, 0xcc, 0xab, 0x46, 0x92, 0xce, 0x1c, 0x72, 0x03, 0xf1, 0x04, 0x7c, 0xb3, 0x35, 0xd9,
	0xa8, 0x05, 0xa7, 0x15, 0x8e, 0xc1, 0x6d, 0x96, 0x54, 0xef, 0xba, 0x79, 0x7c, 0x47, 0x18, 0xc2,
	0xa0, 0x9b, 0x5e, 0x69, 0xab, 0xbb, 0x78, 0xfc, 0x0b, 0xe3, 0xf3, 0xfd, 0x13, 0xff, 0x25, 0x31,
	0x08, 0xcc, 0xda, 0xb4, 0xc3, 0xae, 0xe1, 0x5f, 0x45, 0x5e, 0x3c, 0x81, 0x6b, 0xfe, 0x01, 0x8e,
	0x60, 0x90, 0xd5, 0xef, 0x79, 0x55, 0x0a, 0x76, 0x80, 0x3e, 0x78, 0xdf, 0x9f, 0xc2, 0x2c, 0x0c,
	0x00, 0x7e, 0x02, 0x59, 0x0f, 0x8f, 0x60, 0xf4, 0xeb, 0x0e, 0xb3, 0x11, 0xc0, 0x79, 0x54, 0x0b,
	0x5a, 0xb3, 0x8f, 0xe8, 0xc5, 0xd5, 0x0f, 0x70, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x3f,
	0x6d, 0xec, 0x9a, 0x01, 0x00, 0x00,
}