// Code generated by protoc-gen-go. DO NOT EDIT.
// source: req.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	req.proto

It has these top-level messages:
	Req
*/
package core

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

// 请求
type Req struct {
	Uid    string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Req) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Req) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Req)(nil), "core.Req")
}

func init() { proto.RegisterFile("req.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4a, 0x2d, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xce, 0x2f, 0x4a, 0x55, 0x72, 0xe6, 0x62, 0x0e,
	0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0x31, 0x85, 0xc4, 0xb8, 0xd8, 0x72, 0x53, 0x4b, 0x32, 0xf2, 0x53, 0x24, 0x98, 0xc0, 0x82,
	0x50, 0x9e, 0x90, 0x10, 0x17, 0x4b, 0x52, 0x7e, 0x4a, 0xa5, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f,
	0x10, 0x98, 0x9d, 0xc4, 0x06, 0x36, 0xd1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xbe, 0xc8,
	0xd0, 0x5e, 0x00, 0x00, 0x00,
}