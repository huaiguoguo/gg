// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/template/template_msg.proto

package template

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TemplateMsgTest struct {
	I                    int32    `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	S                    string   `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateMsgTest) Reset()         { *m = TemplateMsgTest{} }
func (m *TemplateMsgTest) String() string { return proto.CompactTextString(m) }
func (*TemplateMsgTest) ProtoMessage()    {}
func (*TemplateMsgTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_msg_ddd9db27993ceaa4, []int{0}
}
func (m *TemplateMsgTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMsgTest.Unmarshal(m, b)
}
func (m *TemplateMsgTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMsgTest.Marshal(b, m, deterministic)
}
func (dst *TemplateMsgTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMsgTest.Merge(dst, src)
}
func (m *TemplateMsgTest) XXX_Size() int {
	return xxx_messageInfo_TemplateMsgTest.Size(m)
}
func (m *TemplateMsgTest) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMsgTest.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMsgTest proto.InternalMessageInfo

func (m *TemplateMsgTest) GetI() int32 {
	if m != nil {
		return m.I
	}
	return 0
}

func (m *TemplateMsgTest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type TemplateMsgTestAck struct {
	IntList              []int32  `protobuf:"varint,1,rep,packed,name=intList,proto3" json:"intList,omitempty"`
	StrList              []string `protobuf:"bytes,2,rep,name=strList,proto3" json:"strList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateMsgTestAck) Reset()         { *m = TemplateMsgTestAck{} }
func (m *TemplateMsgTestAck) String() string { return proto.CompactTextString(m) }
func (*TemplateMsgTestAck) ProtoMessage()    {}
func (*TemplateMsgTestAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_msg_ddd9db27993ceaa4, []int{1}
}
func (m *TemplateMsgTestAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMsgTestAck.Unmarshal(m, b)
}
func (m *TemplateMsgTestAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMsgTestAck.Marshal(b, m, deterministic)
}
func (dst *TemplateMsgTestAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMsgTestAck.Merge(dst, src)
}
func (m *TemplateMsgTestAck) XXX_Size() int {
	return xxx_messageInfo_TemplateMsgTestAck.Size(m)
}
func (m *TemplateMsgTestAck) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMsgTestAck.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMsgTestAck proto.InternalMessageInfo

func (m *TemplateMsgTestAck) GetIntList() []int32 {
	if m != nil {
		return m.IntList
	}
	return nil
}

func (m *TemplateMsgTestAck) GetStrList() []string {
	if m != nil {
		return m.StrList
	}
	return nil
}

type TemplateMsgTestNtf struct {
	IntList              []int32  `protobuf:"varint,1,rep,packed,name=intList,proto3" json:"intList,omitempty"`
	StrList              []string `protobuf:"bytes,2,rep,name=strList,proto3" json:"strList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateMsgTestNtf) Reset()         { *m = TemplateMsgTestNtf{} }
func (m *TemplateMsgTestNtf) String() string { return proto.CompactTextString(m) }
func (*TemplateMsgTestNtf) ProtoMessage()    {}
func (*TemplateMsgTestNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_msg_ddd9db27993ceaa4, []int{2}
}
func (m *TemplateMsgTestNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMsgTestNtf.Unmarshal(m, b)
}
func (m *TemplateMsgTestNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMsgTestNtf.Marshal(b, m, deterministic)
}
func (dst *TemplateMsgTestNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMsgTestNtf.Merge(dst, src)
}
func (m *TemplateMsgTestNtf) XXX_Size() int {
	return xxx_messageInfo_TemplateMsgTestNtf.Size(m)
}
func (m *TemplateMsgTestNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMsgTestNtf.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMsgTestNtf proto.InternalMessageInfo

func (m *TemplateMsgTestNtf) GetIntList() []int32 {
	if m != nil {
		return m.IntList
	}
	return nil
}

func (m *TemplateMsgTestNtf) GetStrList() []string {
	if m != nil {
		return m.StrList
	}
	return nil
}

func init() {
	proto.RegisterType((*TemplateMsgTest)(nil), "template.TemplateMsgTest")
	proto.RegisterType((*TemplateMsgTestAck)(nil), "template.TemplateMsgTestAck")
	proto.RegisterType((*TemplateMsgTestNtf)(nil), "template.TemplateMsgTestNtf")
}

func init() {
	proto.RegisterFile("pb/template/template_msg.proto", fileDescriptor_template_msg_ddd9db27993ceaa4)
}

var fileDescriptor_template_msg_ddd9db27993ceaa4 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x48, 0xd2, 0x2f,
	0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x85, 0x33, 0xe2, 0x73, 0x8b, 0xd3, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60, 0x62, 0x4a, 0xba, 0x5c, 0xfc, 0x21, 0x50, 0xb6, 0x6f, 0x71,
	0x7a, 0x48, 0x6a, 0x71, 0x89, 0x10, 0x0f, 0x17, 0x63, 0xa6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b,
	0x10, 0x63, 0x26, 0x88, 0x57, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x58, 0xac, 0xe4,
	0xc1, 0x25, 0x84, 0xa6, 0xdc, 0x31, 0x39, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x33, 0xaf, 0xc4, 0x27,
	0xb3, 0xb8, 0x44, 0x82, 0x51, 0x81, 0x59, 0x83, 0x35, 0x08, 0xc6, 0x05, 0xc9, 0x14, 0x97, 0x14,
	0x81, 0x65, 0x98, 0x14, 0x98, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x2c, 0x26, 0xf9, 0x95, 0xa4, 0x91,
	0x63, 0x52, 0x12, 0x1b, 0xd8, 0x4f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x76, 0x96,
	0xf7, 0xf5, 0x00, 0x00, 0x00,
}
