// Code generated by protoc-gen-go.
// source: stats.proto
// DO NOT EDIT!

package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerStats struct {
	// wall clock time change in seconds since last reset
	TimeElapsed float64 `protobuf:"fixed64,1,opt,name=time_elapsed,json=timeElapsed" json:"time_elapsed,omitempty"`
	// change in user time (in seconds) used by the server since last reset
	TimeUser    float64 `protobuf:"fixed64,2,opt,name=time_user,json=timeUser" json:"time_user,omitempty"`
	// change in server time (in seconds) used by the server process and all
	// threads since last reset
	TimeSystem  float64 `protobuf:"fixed64,3,opt,name=time_system,json=timeSystem" json:"time_system,omitempty"`
}

func (m *ServerStats) Reset() {
	*m = ServerStats{}
}
func (m *ServerStats) String() string {
	return proto.CompactTextString(m)
}
func (*ServerStats) ProtoMessage() {}
func (*ServerStats) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0}
}

// Histogram params based on grpc/support/histogram.c
type HistogramParams struct {
	Resolution  float64 `protobuf:"fixed64,1,opt,name=resolution" json:"resolution,omitempty"`
	MaxPossible float64 `protobuf:"fixed64,2,opt,name=max_possible,json=maxPossible" json:"max_possible,omitempty"`
}

func (m *HistogramParams) Reset() {
	*m = HistogramParams{}
}
func (m *HistogramParams) String() string {
	return proto.CompactTextString(m)
}
func (*HistogramParams) ProtoMessage() {}
func (*HistogramParams) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{1}
}

// Histogram data based on grpc/support/histogram.c
type HistogramData struct {
	Bucket       []uint32 `protobuf:"varint,1,rep,name=bucket" json:"bucket,omitempty"`
	MinSeen      float64  `protobuf:"fixed64,2,opt,name=min_seen,json=minSeen" json:"min_seen,omitempty"`
	MaxSeen      float64  `protobuf:"fixed64,3,opt,name=max_seen,json=maxSeen" json:"max_seen,omitempty"`
	Sum          float64  `protobuf:"fixed64,4,opt,name=sum" json:"sum,omitempty"`
	SumOfSquares float64  `protobuf:"fixed64,5,opt,name=sum_of_squares,json=sumOfSquares" json:"sum_of_squares,omitempty"`
	Count        float64  `protobuf:"fixed64,6,opt,name=count" json:"count,omitempty"`
}

func (m *HistogramData) Reset() {
	*m = HistogramData{}
}
func (m *HistogramData) String() string {
	return proto.CompactTextString(m)
}
func (*HistogramData) ProtoMessage() {}
func (*HistogramData) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{2}
}

type ClientStats struct {
	// Latency histogram. Data points are in nanoseconds.
	Latencies   *HistogramData `protobuf:"bytes,1,opt,name=latencies" json:"latencies,omitempty"`
	// See ServerStats for details.
	TimeElapsed float64 `protobuf:"fixed64,2,opt,name=time_elapsed,json=timeElapsed" json:"time_elapsed,omitempty"`
	TimeUser    float64 `protobuf:"fixed64,3,opt,name=time_user,json=timeUser" json:"time_user,omitempty"`
	TimeSystem  float64 `protobuf:"fixed64,4,opt,name=time_system,json=timeSystem" json:"time_system,omitempty"`
}

func (m *ClientStats) Reset() {
	*m = ClientStats{}
}
func (m *ClientStats) String() string {
	return proto.CompactTextString(m)
}
func (*ClientStats) ProtoMessage() {}
func (*ClientStats) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{3}
}

func (m *ClientStats) GetLatencies() *HistogramData {
	if m != nil {
		return m.Latencies
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerStats)(nil), "grpc.testing.ServerStats")
	proto.RegisterType((*HistogramParams)(nil), "grpc.testing.HistogramParams")
	proto.RegisterType((*HistogramData)(nil), "grpc.testing.HistogramData")
	proto.RegisterType((*ClientStats)(nil), "grpc.testing.ClientStats")
}

func init() {
	proto.RegisterFile("stats.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xe3, 0x30,
	0x10, 0xc5, 0x95, 0xa6, 0xed, 0xb6, 0x93, 0x76, 0x77, 0x65, 0xad, 0x56, 0x41, 0x95, 0xf8, 0x13,
	0x71, 0xe8, 0x29, 0x07, 0x38, 0x71, 0x06, 0x24, 0x6e, 0x54, 0x0d, 0x9c, 0x23, 0x37, 0x4c, 0x2b,
	0x8b, 0xc4, 0x0e, 0x99, 0x09, 0x2a, 0x1f, 0x09, 0xf1, 0x25, 0x71, 0x9c, 0x08, 0x0a, 0x48, 0x70,
	0x49, 0xf2, 0x7e, 0x6f, 0x34, 0xe3, 0xc9, 0x33, 0x04, 0xc4, 0x92, 0x29, 0x2e, 0x2b, 0xc3, 0x46,
	0x4c, 0x36, 0x55, 0x99, 0xc5, 0x8c, 0xc4, 0x4a, 0x6f, 0x22, 0x0d, 0x41, 0x82, 0xd5, 0x23, 0x56,
	0x49, 0x53, 0x22, 0x8e, 0x60, 0xc2, 0xaa, 0xc0, 0x14, 0x73, 0x59, 0x12, 0xde, 0x85, 0xde, 0xa1,
	0x37, 0xf7, 0x96, 0x41, 0xc3, 0x2e, 0x5b, 0x24, 0x66, 0x30, 0x76, 0x25, 0x35, 0x61, 0x15, 0xf6,
	0x9c, 0x3f, 0x6a, 0xc0, 0xad, 0xd5, 0xe2, 0x00, 0x5c, 0x6d, 0x4a, 0x4f, 0xc4, 0x58, 0x84, 0xbe,
	0xb3, 0xa1, 0x41, 0x89, 0x23, 0xd1, 0x0d, 0xfc, 0xb9, 0x52, 0xc4, 0x66, 0x53, 0xc9, 0x62, 0x21,
	0xed, 0x83, 0xc4, 0x3e, 0x40, 0x85, 0x64, 0xf2, 0x9a, 0x95, 0xd1, 0xdd, 0xc4, 0x1d, 0xd2, 0x9c,
	0xa9, 0x90, 0xdb, 0xb4, 0x34, 0x44, 0x6a, 0x95, 0x63, 0x37, 0x33, 0xb0, 0x6c, 0xd1, 0xa1, 0xe8,
	0xc5, 0x83, 0xe9, 0x5b, 0xdb, 0x0b, 0xc9, 0x52, 0xfc, 0x87, 0xe1, 0xaa, 0xce, 0xee, 0x91, 0x6d,
	0x43, 0x7f, 0x3e, 0x5d, 0x76, 0x4a, 0xec, 0xc1, 0xa8, 0x50, 0x3a, 0x25, 0x44, 0xdd, 0x35, 0xfa,
	0x65, 0x75, 0x62, 0xa5, 0xb3, 0xec, 0x1c, 0x67, 0xf9, 0x9d, 0x25, 0xb7, 0xce, 0xfa, 0x0b, 0x3e,
	0xd5, 0x45, 0xd8, 0x77, 0xb4, 0xf9, 0x14, 0xc7, 0xf0, 0xdb, 0xbe, 0x52, 0xb3, 0x4e, 0xe9, 0xa1,
	0x96, 0xf6, 0xb4, 0xe1, 0xc0, 0x99, 0x13, 0x4b, 0xaf, 0xd7, 0x49, 0xcb, 0xc4, 0x3f, 0x18, 0x64,
	0xa6, 0xd6, 0x1c, 0x0e, 0x9d, 0xd9, 0x8a, 0xe8, 0xd9, 0x83, 0xe0, 0x3c, 0x57, 0xa8, 0xb9, 0xfd,
	0xe9, 0x67, 0x30, 0xce, 0x25, 0xa3, 0xce, 0x94, 0x6d, 0xd3, 0xec, 0x1f, 0x9c, 0xcc, 0xe2, 0xdd,
	0x94, 0xe2, 0x0f, 0xbb, 0x2d, 0xdf, 0xab, 0xbf, 0xe4, 0xd5, 0xfb, 0x21, 0x2f, 0xff, 0xfb, 0xbc,
	0xfa, 0x9f, 0xf3, 0x5a, 0x0d, 0xdd, 0xa5, 0x39, 0x7d, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xea, 0x75,
	0x34, 0x90, 0x43, 0x02, 0x00, 0x00,
}
