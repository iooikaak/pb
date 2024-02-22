// Code generated by WARHORSE protoc-gen-go. DO NOT EDIT EXCEPET SERVER VERSION.
// source: message.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:

	message.proto
	services.proto
	user.proto

It has these top-level messages:

	InteralMsg
	UserOnline
	Omit
*/
package user

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

type MessageType int32

const (
	MessageType_MT_OMIT                      MessageType = 0
	MessageType_MT_USER_LOGIN                MessageType = 1
	MessageType_MT_USER_LOGOUT               MessageType = 2
	MessageType_MT_USER_MODIFY_INFO          MessageType = 3
	MessageType_MT_USER_UPLOAD_LOCATION      MessageType = 4
	MessageType_MT_USER_BIND_OR_UNBIND_GAMES MessageType = 5
)

var MessageType_name = map[int32]string{
	0: "MT_OMIT",
	1: "MT_USER_LOGIN",
	2: "MT_USER_LOGOUT",
	3: "MT_USER_MODIFY_INFO",
	4: "MT_USER_UPLOAD_LOCATION",
	5: "MT_USER_BIND_OR_UNBIND_GAMES",
}
var MessageType_value = map[string]int32{
	"MT_OMIT":                      0,
	"MT_USER_LOGIN":                1,
	"MT_USER_LOGOUT":               2,
	"MT_USER_MODIFY_INFO":          3,
	"MT_USER_UPLOAD_LOCATION":      4,
	"MT_USER_BIND_OR_UNBIND_GAMES": 5,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InteralMsg struct {
	Mt  MessageType `protobuf:"varint,1,opt,name=mt,enum=user.MessageType" json:"mt" xml:"mt,omitempty"`
	Uid int64       `protobuf:"varint,2,opt,name=uid" json:"uid" xml:"uid,omitempty"`
	T   int64       `protobuf:"varint,3,opt,name=t" json:"t" xml:"t,omitempty"`
	M   string      `protobuf:"bytes,4,opt,name=m" json:"m,omitempty" xml:"m,omitempty"`
}

func (m *InteralMsg) Reset()                    { *m = InteralMsg{} }
func (m *InteralMsg) String() string            { return proto.CompactTextString(m) }
func (*InteralMsg) ProtoMessage()               {}
func (*InteralMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InteralMsg) GetMt() MessageType {
	if m != nil {
		return m.Mt
	}
	return MessageType_MT_OMIT
}

func (m *InteralMsg) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *InteralMsg) GetT() int64 {
	if m != nil {
		return m.T
	}
	return 0
}

func (m *InteralMsg) GetM() string {
	if m != nil {
		return m.M
	}
	return ""
}

type UserOnline struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id" xml:"user_id,omitempty"`
}

func (m *UserOnline) Reset()                    { *m = UserOnline{} }
func (m *UserOnline) String() string            { return proto.CompactTextString(m) }
func (*UserOnline) ProtoMessage()               {}
func (*UserOnline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserOnline) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterEnum("user.MessageType", MessageType_name, MessageType_value)
}

func init() {}

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0xa4, 0xb6, 0x38, 0xb5, 0x65, 0x3b, 0x1e, 0x12, 0xd0, 0x43, 0x2c, 0x08, 0xc1,
	0x43, 0x0e, 0xfa, 0x0b, 0xa2, 0x69, 0xcb, 0x42, 0x36, 0x2b, 0xdb, 0xcd, 0xc1, 0x83, 0x2c, 0x95,
	0x2c, 0x25, 0xd0, 0xa4, 0x25, 0xd9, 0x1e, 0xfc, 0x29, 0xfe, 0x5b, 0x49, 0x24, 0xd0, 0xdb, 0xfb,
	0xde, 0x9b, 0x99, 0x07, 0x03, 0xb3, 0xca, 0xb4, 0xed, 0x6e, 0x6f, 0xa2, 0x53, 0x73, 0xb4, 0x47,
	0x1c, 0x9d, 0x5b, 0xd3, 0x2c, 0xbf, 0x00, 0x58, 0x6d, 0x4d, 0xb3, 0x3b, 0xf0, 0x76, 0x8f, 0x8f,
	0xe0, 0x54, 0xd6, 0x27, 0x01, 0x09, 0xe7, 0x2f, 0x8b, 0xa8, 0x1b, 0x88, 0xf8, 0xff, 0x92, 0xfa,
	0x39, 0x19, 0xe9, 0x54, 0x16, 0x29, 0xb8, 0xe7, 0xb2, 0xf0, 0x9d, 0x80, 0x84, 0xae, 0xec, 0x24,
	0xde, 0x02, 0xb1, 0xbe, 0xdb, 0x33, 0xb1, 0x1d, 0x55, 0xfe, 0x28, 0x20, 0xe1, 0x8d, 0x24, 0xd5,
	0xf2, 0x09, 0x20, 0x6f, 0x4d, 0x23, 0xea, 0x43, 0x59, 0x1b, 0xf4, 0x60, 0xd2, 0xdd, 0xd4, 0x65,
	0xd1, 0x77, 0xb8, 0x72, 0xdc, 0x21, 0x2b, 0x9e, 0x7f, 0x09, 0x4c, 0x2f, 0x8a, 0x70, 0x0a, 0x13,
	0xae, 0xb4, 0xe0, 0x4c, 0xd1, 0x2b, 0x5c, 0xc0, 0x8c, 0x2b, 0x9d, 0x6f, 0x57, 0x52, 0xa7, 0x62,
	0xc3, 0x32, 0x4a, 0x10, 0x61, 0x7e, 0x61, 0x89, 0x5c, 0x51, 0x07, 0x3d, 0xb8, 0x1b, 0x3c, 0x2e,
	0x12, 0xb6, 0xfe, 0xd4, 0x2c, 0x5b, 0x0b, 0xea, 0xe2, 0x3d, 0x78, 0x43, 0x90, 0x7f, 0xa4, 0x22,
	0x4e, 0x74, 0x2a, 0xde, 0x63, 0xc5, 0x44, 0x46, 0x47, 0x18, 0xc0, 0xc3, 0x10, 0xbe, 0xb1, 0x2c,
	0xd1, 0x42, 0xea, 0x3c, 0xeb, 0xd5, 0x26, 0xe6, 0xab, 0x2d, 0xbd, 0xfe, 0x1e, 0xf7, 0xef, 0x7a,
	0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xa9, 0xd0, 0x03, 0x3f, 0x01, 0x00, 0x00,
}