// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/grpc_web/grpc_web.proto

package grpc_web

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/transformation"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GrpcWeb support is enabled be default. Use this extension to disable it.
type GrpcWeb struct {
	// Disable grpc web support.
	Disable              bool     `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcWeb) Reset()         { *m = GrpcWeb{} }
func (m *GrpcWeb) String() string { return proto.CompactTextString(m) }
func (*GrpcWeb) ProtoMessage()    {}
func (*GrpcWeb) Descriptor() ([]byte, []int) {
	return fileDescriptor_0849cead00dee4e7, []int{0}
}
func (m *GrpcWeb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcWeb.Unmarshal(m, b)
}
func (m *GrpcWeb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcWeb.Marshal(b, m, deterministic)
}
func (m *GrpcWeb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcWeb.Merge(m, src)
}
func (m *GrpcWeb) XXX_Size() int {
	return xxx_messageInfo_GrpcWeb.Size(m)
}
func (m *GrpcWeb) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcWeb.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcWeb proto.InternalMessageInfo

func (m *GrpcWeb) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcWeb)(nil), "grpc_web.options.gloo.solo.io.GrpcWeb")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/grpc_web/grpc_web.proto", fileDescriptor_0849cead00dee4e7)
}

var fileDescriptor_0849cead00dee4e7 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0x87, 0xd0,
	0x89, 0x05, 0x99, 0xc5, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xe9, 0x39, 0xf9, 0xf9, 0x10, 0xa2, 0xcc,
	0x50, 0x3f, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x3f, 0xbd, 0xa8, 0x20, 0x39, 0xbe, 0x3c,
	0x35, 0x09, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x85, 0xf3, 0xa1, 0x2a, 0xf5,
	0x40, 0x3a, 0xf5, 0x40, 0x26, 0xea, 0x65, 0xe6, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x55,
	0xea, 0x83, 0x58, 0x10, 0x4d, 0x52, 0x42, 0xa9, 0x15, 0x25, 0x10, 0xc1, 0xd4, 0x8a, 0x12, 0xa8,
	0x98, 0x13, 0x11, 0x8e, 0x28, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0xcb, 0x2f, 0xca, 0x4d, 0x04, 0xf1,
	0xf5, 0x0b, 0x12, 0x8b, 0x12, 0x73, 0x53, 0x4b, 0x52, 0x8b, 0x8a, 0x21, 0x66, 0x28, 0x29, 0x73,
	0xb1, 0xbb, 0x17, 0x15, 0x24, 0x87, 0xa7, 0x26, 0x09, 0x49, 0x70, 0xb1, 0xa7, 0x64, 0x16, 0x27,
	0x26, 0xe5, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xb8, 0x4e, 0xfe, 0x3b, 0xbe,
	0xb2, 0x30, 0xae, 0x78, 0x24, 0xc7, 0x18, 0xe5, 0x8a, 0x37, 0x28, 0x0a, 0xb2, 0xd3, 0xe1, 0x2e,
	0x81, 0xf9, 0x07, 0x5b, 0x88, 0x24, 0xb1, 0x81, 0x2d, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xc7, 0x90, 0xa8, 0xa0, 0x5a, 0x01, 0x00, 0x00,
}

func (this *GrpcWeb) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcWeb)
	if !ok {
		that2, ok := that.(GrpcWeb)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disable != that1.Disable {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}