// Code generated by protoc-gen-go.
// source: proto2.proto
// DO NOT EDIT!

package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ToBeExtened struct {
	Foo              *int32 `protobuf:"varint,1,req,name=foo" json:"foo,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToBeExtened) Reset() {
	*m = ToBeExtened{}
}
func (m *ToBeExtened) String() string {
	return proto.CompactTextString(m)
}
func (*ToBeExtened) ProtoMessage() {}
func (*ToBeExtened) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0}
}

var extRange_ToBeExtened = []proto.ExtensionRange{
	{10, 20},
}

func (*ToBeExtened) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ToBeExtened
}

func (m *ToBeExtened) GetFoo() int32 {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return 0
}

func init() {
	proto.RegisterType((*ToBeExtened)(nil), "grpc.testing.ToBeExtened")
}

func init() {
	proto.RegisterFile("proto2.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 85 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd2, 0x03, 0x53, 0x42, 0x3c, 0xe9, 0x45, 0x05, 0xc9, 0x7a, 0x25, 0xa9, 0xc5, 0x25,
	0x99, 0x79, 0xe9, 0x4a, 0xaa, 0x5c, 0xdc, 0x21, 0xf9, 0x4e, 0xa9, 0xae, 0x15, 0x25, 0xa9, 0x79,
	0xa9, 0x29, 0x42, 0x02, 0x5c, 0xcc, 0x69, 0xf9, 0xf9, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0xac, 0x41,
	0x20, 0xa6, 0x16, 0x0b, 0x07, 0x97, 0x80, 0x28, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xed, 0xbc,
	0xc2, 0x43, 0x00, 0x00, 0x00,
}
