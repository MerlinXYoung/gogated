// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cs_gateway.proto

package gw_cs

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
	EMsgID_Invalid EMsgID = 0
	EMsgID_Auth    EMsgID = 1
	EMsgID_Other   EMsgID = 4096
)

var EMsgID_name = map[int32]string{
	0:    "Invalid",
	1:    "Auth",
	4096: "Other",
}

var EMsgID_value = map[string]int32{
	"Invalid": 0,
	"Auth":    1,
	"Other":   4096,
}

func (x EMsgID) String() string {
	return proto.EnumName(EMsgID_name, int32(x))
}

func (EMsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce1d9d98da9c443b, []int{0}
}

type Head struct {
	Msgid                EMsgID   `protobuf:"varint,1,opt,name=msgid,proto3,enum=gw.cs.EMsgID" json:"msgid,omitempty"`
	Seq                  uint32   `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	DownSeq              uint32   `protobuf:"varint,3,opt,name=down_seq,json=downSeq,proto3" json:"down_seq,omitempty"`
	Result               int32    `protobuf:"varint,4,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head) Reset()         { *m = Head{} }
func (m *Head) String() string { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()    {}
func (*Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d9d98da9c443b, []int{0}
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

func (m *Head) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Head) GetDownSeq() uint32 {
	if m != nil {
		return m.DownSeq
	}
	return 0
}

func (m *Head) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AuthReq struct {
	Openid               string   `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid,omitempty"`
	Openkey              string   `protobuf:"bytes,2,opt,name=openkey,proto3" json:"openkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d9d98da9c443b, []int{1}
}

func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReq.Unmarshal(m, b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
	return xxx_messageInfo_AuthReq.Size(m)
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

func (m *AuthReq) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *AuthReq) GetOpenkey() string {
	if m != nil {
		return m.Openkey
	}
	return ""
}

type AuthRes struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRes) Reset()         { *m = AuthRes{} }
func (m *AuthRes) String() string { return proto.CompactTextString(m) }
func (*AuthRes) ProtoMessage()    {}
func (*AuthRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d9d98da9c443b, []int{2}
}

func (m *AuthRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRes.Unmarshal(m, b)
}
func (m *AuthRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRes.Marshal(b, m, deterministic)
}
func (m *AuthRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRes.Merge(m, src)
}
func (m *AuthRes) XXX_Size() int {
	return xxx_messageInfo_AuthRes.Size(m)
}
func (m *AuthRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRes.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRes proto.InternalMessageInfo

func (m *AuthRes) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterEnum("gw.cs.EMsgID", EMsgID_name, EMsgID_value)
	proto.RegisterType((*Head)(nil), "gw.cs.Head")
	proto.RegisterType((*AuthReq)(nil), "gw.cs.AuthReq")
	proto.RegisterType((*AuthRes)(nil), "gw.cs.AuthRes")
}

func init() { proto.RegisterFile("proto/cs_gateway.proto", fileDescriptor_ce1d9d98da9c443b) }

var fileDescriptor_ce1d9d98da9c443b = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x8d, 0xeb, 0x9f, 0xed, 0xca, 0xa4, 0xdc, 0x87, 0x51, 0xf1, 0xa5, 0xd4, 0x97, 0xa2,
	0x50, 0x41, 0x1f, 0x7d, 0x12, 0x14, 0xdc, 0x83, 0x08, 0xf1, 0x03, 0x8c, 0xb8, 0x5c, 0xb2, 0xe1,
	0x6c, 0xba, 0x26, 0xb5, 0xec, 0xcd, 0x8f, 0x2e, 0x49, 0xd3, 0xb7, 0xf3, 0x3b, 0x37, 0x87, 0x73,
	0x08, 0xac, 0xda, 0x4e, 0x5b, 0x7d, 0xbf, 0x35, 0x1b, 0x25, 0x2c, 0x0d, 0xe2, 0x54, 0x7b, 0x03,
	0x63, 0x35, 0xd4, 0x5b, 0x53, 0xb6, 0x10, 0xbd, 0x91, 0x90, 0x78, 0x03, 0xf1, 0x8f, 0x51, 0x7b,
	0x99, 0xb3, 0x82, 0x55, 0x97, 0x0f, 0xcb, 0xda, 0x9f, 0xeb, 0xd7, 0x77, 0xa3, 0xd6, 0x2f, 0x7c,
	0xbc, 0x61, 0x06, 0x33, 0x43, 0xc7, 0xfc, 0xbc, 0x60, 0xd5, 0x92, 0x3b, 0x89, 0x57, 0x30, 0x97,
	0x7a, 0x68, 0x36, 0xce, 0x9e, 0x79, 0x3b, 0x75, 0xfc, 0x49, 0x47, 0x5c, 0x41, 0xd2, 0x91, 0xe9,
	0x0f, 0x36, 0x8f, 0x0a, 0x56, 0xc5, 0x3c, 0x50, 0xf9, 0x04, 0xe9, 0x73, 0x6f, 0x77, 0x7c, 0x7c,
	0xa2, 0x5b, 0x6a, 0x42, 0xeb, 0x82, 0x07, 0xc2, 0x1c, 0x52, 0xa7, 0xbe, 0xe9, 0xe4, 0xbb, 0x16,
	0x7c, 0xc2, 0xf2, 0x7a, 0x0a, 0x1b, 0x37, 0xa6, 0x0f, 0xc9, 0x88, 0x3b, 0x79, 0x7b, 0x07, 0xc9,
	0xb8, 0x17, 0x2f, 0x20, 0x5d, 0x37, 0xbf, 0xe2, 0xb0, 0x97, 0xd9, 0x19, 0xce, 0x21, 0x72, 0x99,
	0x8c, 0x21, 0x40, 0xfc, 0x61, 0x77, 0xd4, 0x65, 0x7f, 0xc5, 0x57, 0xe2, 0xbf, 0xe1, 0xf1, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0xb1, 0xab, 0x50, 0x20, 0x01, 0x00, 0x00,
}
