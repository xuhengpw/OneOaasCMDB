// Code generated by protoc-gen-go.
// source: control.proto
// DO NOT EDIT!

/*
Package grpc_testing is a generated protocol buffer package.

It is generated from these files:
	control.proto
	messages.proto
	payloads.proto
	services.proto
	stats.proto

It has these top-level messages:
	PoissonParams
	UniformParams
	DeterministicParams
	ParetoParams
	ClosedLoopParams
	LoadParams
	SecurityParams
	ClientConfig
	ClientStatus
	Mark
	ClientArgs
	ServerConfig
	ServerArgs
	ServerStatus
	CoreRequest
	CoreResponse
	Void
	Scenario
	Scenarios
	Payload
	EchoStatus
	SimpleRequest
	SimpleResponse
	StreamingInputCallRequest
	StreamingInputCallResponse
	ResponseParameters
	StreamingOutputCallRequest
	StreamingOutputCallResponse
	ReconnectParams
	ReconnectInfo
	ByteBufferParams
	SimpleProtoParams
	ComplexProtoParams
	PayloadConfig
	ServerStats
	HistogramParams
	HistogramData
	ClientStats
*/
package grpc_testing

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

type ClientType int32

const (
	ClientType_SYNC_CLIENT ClientType = 0
	ClientType_ASYNC_CLIENT ClientType = 1
)

var ClientType_name = map[int32]string{
	0: "SYNC_CLIENT",
	1: "ASYNC_CLIENT",
}
var ClientType_value = map[string]int32{
	"SYNC_CLIENT":  0,
	"ASYNC_CLIENT": 1,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

type ServerType int32

const (
	ServerType_SYNC_SERVER ServerType = 0
	ServerType_ASYNC_SERVER ServerType = 1
	ServerType_ASYNC_GENERIC_SERVER ServerType = 2
)

var ServerType_name = map[int32]string{
	0: "SYNC_SERVER",
	1: "ASYNC_SERVER",
	2: "ASYNC_GENERIC_SERVER",
}
var ServerType_value = map[string]int32{
	"SYNC_SERVER":          0,
	"ASYNC_SERVER":         1,
	"ASYNC_GENERIC_SERVER": 2,
}

func (x ServerType) String() string {
	return proto.EnumName(ServerType_name, int32(x))
}
func (ServerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

type RpcType int32

const (
	RpcType_UNARY RpcType = 0
	RpcType_STREAMING RpcType = 1
)

var RpcType_name = map[int32]string{
	0: "UNARY",
	1: "STREAMING",
}
var RpcType_value = map[string]int32{
	"UNARY":     0,
	"STREAMING": 1,
}

func (x RpcType) String() string {
	return proto.EnumName(RpcType_name, int32(x))
}
func (RpcType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

// Parameters of poisson process distribution, which is a good representation
// of activity coming in from independent identical stationary sources.
type PoissonParams struct {
	// The rate of arrivals (a.k.a. lambda parameter of the exp distribution).
	OfferedLoad float64 `protobuf:"fixed64,1,opt,name=offered_load,json=offeredLoad" json:"offered_load,omitempty"`
}

func (m *PoissonParams) Reset() {
	*m = PoissonParams{}
}
func (m *PoissonParams) String() string {
	return proto.CompactTextString(m)
}
func (*PoissonParams) ProtoMessage() {}
func (*PoissonParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

type UniformParams struct {
	InterarrivalLo float64 `protobuf:"fixed64,1,opt,name=interarrival_lo,json=interarrivalLo" json:"interarrival_lo,omitempty"`
	InterarrivalHi float64 `protobuf:"fixed64,2,opt,name=interarrival_hi,json=interarrivalHi" json:"interarrival_hi,omitempty"`
}

func (m *UniformParams) Reset() {
	*m = UniformParams{}
}
func (m *UniformParams) String() string {
	return proto.CompactTextString(m)
}
func (*UniformParams) ProtoMessage() {}
func (*UniformParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

type DeterministicParams struct {
	OfferedLoad float64 `protobuf:"fixed64,1,opt,name=offered_load,json=offeredLoad" json:"offered_load,omitempty"`
}

func (m *DeterministicParams) Reset() {
	*m = DeterministicParams{}
}
func (m *DeterministicParams) String() string {
	return proto.CompactTextString(m)
}
func (*DeterministicParams) ProtoMessage() {}
func (*DeterministicParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

type ParetoParams struct {
	InterarrivalBase float64 `protobuf:"fixed64,1,opt,name=interarrival_base,json=interarrivalBase" json:"interarrival_base,omitempty"`
	Alpha            float64 `protobuf:"fixed64,2,opt,name=alpha" json:"alpha,omitempty"`
}

func (m *ParetoParams) Reset() {
	*m = ParetoParams{}
}
func (m *ParetoParams) String() string {
	return proto.CompactTextString(m)
}
func (*ParetoParams) ProtoMessage() {}
func (*ParetoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

// Once an RPC finishes, immediately start a new one.
// No configuration parameters needed.
type ClosedLoopParams struct {
}

func (m *ClosedLoopParams) Reset() {
	*m = ClosedLoopParams{}
}
func (m *ClosedLoopParams) String() string {
	return proto.CompactTextString(m)
}
func (*ClosedLoopParams) ProtoMessage() {}
func (*ClosedLoopParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

type LoadParams struct {
	// Types that are valid to be assigned to Load:
	//	*LoadParams_ClosedLoop
	//	*LoadParams_Poisson
	//	*LoadParams_Uniform
	//	*LoadParams_Determ
	//	*LoadParams_Pareto
	Load isLoadParams_Load `protobuf_oneof:"load"`
}

func (m *LoadParams) Reset() {
	*m = LoadParams{}
}
func (m *LoadParams) String() string {
	return proto.CompactTextString(m)
}
func (*LoadParams) ProtoMessage() {}
func (*LoadParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

type isLoadParams_Load interface {
	isLoadParams_Load()
}

type LoadParams_ClosedLoop struct {
	ClosedLoop *ClosedLoopParams `protobuf:"bytes,1,opt,name=closed_loop,json=closedLoop,oneof"`
}
type LoadParams_Poisson struct {
	Poisson *PoissonParams `protobuf:"bytes,2,opt,name=poisson,oneof"`
}
type LoadParams_Uniform struct {
	Uniform *UniformParams `protobuf:"bytes,3,opt,name=uniform,oneof"`
}
type LoadParams_Determ struct {
	Determ *DeterministicParams `protobuf:"bytes,4,opt,name=determ,oneof"`
}
type LoadParams_Pareto struct {
	Pareto *ParetoParams `protobuf:"bytes,5,opt,name=pareto,oneof"`
}

func (*LoadParams_ClosedLoop) isLoadParams_Load() {}
func (*LoadParams_Poisson) isLoadParams_Load() {}
func (*LoadParams_Uniform) isLoadParams_Load() {}
func (*LoadParams_Determ) isLoadParams_Load() {}
func (*LoadParams_Pareto) isLoadParams_Load() {}

func (m *LoadParams) GetLoad() isLoadParams_Load {
	if m != nil {
		return m.Load
	}
	return nil
}

func (m *LoadParams) GetClosedLoop() *ClosedLoopParams {
	if x, ok := m.GetLoad().(*LoadParams_ClosedLoop); ok {
		return x.ClosedLoop
	}
	return nil
}

func (m *LoadParams) GetPoisson() *PoissonParams {
	if x, ok := m.GetLoad().(*LoadParams_Poisson); ok {
		return x.Poisson
	}
	return nil
}

func (m *LoadParams) GetUniform() *UniformParams {
	if x, ok := m.GetLoad().(*LoadParams_Uniform); ok {
		return x.Uniform
	}
	return nil
}

func (m *LoadParams) GetDeterm() *DeterministicParams {
	if x, ok := m.GetLoad().(*LoadParams_Determ); ok {
		return x.Determ
	}
	return nil
}

func (m *LoadParams) GetPareto() *ParetoParams {
	if x, ok := m.GetLoad().(*LoadParams_Pareto); ok {
		return x.Pareto
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadParams) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadParams_OneofMarshaler, _LoadParams_OneofUnmarshaler, _LoadParams_OneofSizer, []interface{}{
		(*LoadParams_ClosedLoop)(nil),
		(*LoadParams_Poisson)(nil),
		(*LoadParams_Uniform)(nil),
		(*LoadParams_Determ)(nil),
		(*LoadParams_Pareto)(nil),
	}
}

func _LoadParams_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadParams)
	// load
	switch x := m.Load.(type) {
	case *LoadParams_ClosedLoop:
		b.EncodeVarint(1 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClosedLoop); err != nil {
			return err
		}
	case *LoadParams_Poisson:
		b.EncodeVarint(2 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Poisson); err != nil {
			return err
		}
	case *LoadParams_Uniform:
		b.EncodeVarint(3 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Uniform); err != nil {
			return err
		}
	case *LoadParams_Determ:
		b.EncodeVarint(4 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Determ); err != nil {
			return err
		}
	case *LoadParams_Pareto:
		b.EncodeVarint(5 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pareto); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadParams.Load has unexpected type %T", x)
	}
	return nil
}

func _LoadParams_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadParams)
	switch tag {
	case 1: // load.closed_loop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClosedLoopParams)
		err := b.DecodeMessage(msg)
		m.Load = &LoadParams_ClosedLoop{msg}
		return true, err
	case 2: // load.poisson
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PoissonParams)
		err := b.DecodeMessage(msg)
		m.Load = &LoadParams_Poisson{msg}
		return true, err
	case 3: // load.uniform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UniformParams)
		err := b.DecodeMessage(msg)
		m.Load = &LoadParams_Uniform{msg}
		return true, err
	case 4: // load.determ
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeterministicParams)
		err := b.DecodeMessage(msg)
		m.Load = &LoadParams_Determ{msg}
		return true, err
	case 5: // load.pareto
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ParetoParams)
		err := b.DecodeMessage(msg)
		m.Load = &LoadParams_Pareto{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadParams_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadParams)
	// load
	switch x := m.Load.(type) {
	case *LoadParams_ClosedLoop:
		s := proto.Size(x.ClosedLoop)
		n += proto.SizeVarint(1 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadParams_Poisson:
		s := proto.Size(x.Poisson)
		n += proto.SizeVarint(2 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadParams_Uniform:
		s := proto.Size(x.Uniform)
		n += proto.SizeVarint(3 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadParams_Determ:
		s := proto.Size(x.Determ)
		n += proto.SizeVarint(4 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LoadParams_Pareto:
		s := proto.Size(x.Pareto)
		n += proto.SizeVarint(5 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// presence of SecurityParams implies use of TLS
type SecurityParams struct {
	UseTestCa          bool   `protobuf:"varint,1,opt,name=use_test_ca,json=useTestCa" json:"use_test_ca,omitempty"`
	ServerHostOverride string `protobuf:"bytes,2,opt,name=server_host_override,json=serverHostOverride" json:"server_host_override,omitempty"`
}

func (m *SecurityParams) Reset() {
	*m = SecurityParams{}
}
func (m *SecurityParams) String() string {
	return proto.CompactTextString(m)
}
func (*SecurityParams) ProtoMessage() {}
func (*SecurityParams) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

type ClientConfig struct {
	// List of targets to connect to. At least one target needs to be specified.
	ServerTargets             []string        `protobuf:"bytes,1,rep,name=server_targets,json=serverTargets" json:"server_targets,omitempty"`
	ClientType                ClientType      `protobuf:"varint,2,opt,name=client_type,json=clientType,enum=grpc.testing.ClientType" json:"client_type,omitempty"`
	SecurityParams            *SecurityParams `protobuf:"bytes,3,opt,name=security_params,json=securityParams" json:"security_params,omitempty"`
	// How many concurrent RPCs to start for each channel.
	// For synchronous client, use a separate thread for each outstanding RPC.
	OutstandingRpcsPerChannel int32 `protobuf:"varint,4,opt,name=outstanding_rpcs_per_channel,json=outstandingRpcsPerChannel" json:"outstanding_rpcs_per_channel,omitempty"`
	// Number of independent client channels to create.
	// i-th channel will connect to server_target[i % server_targets.size()]
	ClientChannels            int32 `protobuf:"varint,5,opt,name=client_channels,json=clientChannels" json:"client_channels,omitempty"`
	// Only for async client. Number of threads to use to start/manage RPCs.
	AsyncClientThreads        int32   `protobuf:"varint,7,opt,name=async_client_threads,json=asyncClientThreads" json:"async_client_threads,omitempty"`
	RpcType                   RpcType `protobuf:"varint,8,opt,name=rpc_type,json=rpcType,enum=grpc.testing.RpcType" json:"rpc_type,omitempty"`
	// The requested load for the entire client (aggregated over all the threads).
	LoadParams                *LoadParams      `protobuf:"bytes,10,opt,name=load_params,json=loadParams" json:"load_params,omitempty"`
	PayloadConfig             *PayloadConfig   `protobuf:"bytes,11,opt,name=payload_config,json=payloadConfig" json:"payload_config,omitempty"`
	HistogramParams           *HistogramParams `protobuf:"bytes,12,opt,name=histogram_params,json=histogramParams" json:"histogram_params,omitempty"`
	// Specify the cores we should run the client on, if desired
	CoreList                  []int32 `protobuf:"varint,13,rep,name=core_list,json=coreList" json:"core_list,omitempty"`
	CoreLimit                 int32   `protobuf:"varint,14,opt,name=core_limit,json=coreLimit" json:"core_limit,omitempty"`
}

func (m *ClientConfig) Reset() {
	*m = ClientConfig{}
}
func (m *ClientConfig) String() string {
	return proto.CompactTextString(m)
}
func (*ClientConfig) ProtoMessage() {}
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7}
}

func (m *ClientConfig) GetSecurityParams() *SecurityParams {
	if m != nil {
		return m.SecurityParams
	}
	return nil
}

func (m *ClientConfig) GetLoadParams() *LoadParams {
	if m != nil {
		return m.LoadParams
	}
	return nil
}

func (m *ClientConfig) GetPayloadConfig() *PayloadConfig {
	if m != nil {
		return m.PayloadConfig
	}
	return nil
}

func (m *ClientConfig) GetHistogramParams() *HistogramParams {
	if m != nil {
		return m.HistogramParams
	}
	return nil
}

type ClientStatus struct {
	Stats *ClientStats `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
}

func (m *ClientStatus) Reset() {
	*m = ClientStatus{}
}
func (m *ClientStatus) String() string {
	return proto.CompactTextString(m)
}
func (*ClientStatus) ProtoMessage() {}
func (*ClientStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8}
}

func (m *ClientStatus) GetStats() *ClientStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Request current stats
type Mark struct {
	// if true, the stats will be reset after taking their snapshot.
	Reset_ bool `protobuf:"varint,1,opt,name=reset" json:"reset,omitempty"`
}

func (m *Mark) Reset() {
	*m = Mark{}
}
func (m *Mark) String() string {
	return proto.CompactTextString(m)
}
func (*Mark) ProtoMessage() {}
func (*Mark) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9}
}

type ClientArgs struct {
	// Types that are valid to be assigned to Argtype:
	//	*ClientArgs_Setup
	//	*ClientArgs_Mark
	Argtype isClientArgs_Argtype `protobuf_oneof:"argtype"`
}

func (m *ClientArgs) Reset() {
	*m = ClientArgs{}
}
func (m *ClientArgs) String() string {
	return proto.CompactTextString(m)
}
func (*ClientArgs) ProtoMessage() {}
func (*ClientArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10}
}

type isClientArgs_Argtype interface {
	isClientArgs_Argtype()
}

type ClientArgs_Setup struct {
	Setup *ClientConfig `protobuf:"bytes,1,opt,name=setup,oneof"`
}
type ClientArgs_Mark struct {
	Mark *Mark `protobuf:"bytes,2,opt,name=mark,oneof"`
}

func (*ClientArgs_Setup) isClientArgs_Argtype() {}
func (*ClientArgs_Mark) isClientArgs_Argtype() {}

func (m *ClientArgs) GetArgtype() isClientArgs_Argtype {
	if m != nil {
		return m.Argtype
	}
	return nil
}

func (m *ClientArgs) GetSetup() *ClientConfig {
	if x, ok := m.GetArgtype().(*ClientArgs_Setup); ok {
		return x.Setup
	}
	return nil
}

func (m *ClientArgs) GetMark() *Mark {
	if x, ok := m.GetArgtype().(*ClientArgs_Mark); ok {
		return x.Mark
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientArgs) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientArgs_OneofMarshaler, _ClientArgs_OneofUnmarshaler, _ClientArgs_OneofSizer, []interface{}{
		(*ClientArgs_Setup)(nil),
		(*ClientArgs_Mark)(nil),
	}
}

func _ClientArgs_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ClientArgs_Setup:
		b.EncodeVarint(1 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Setup); err != nil {
			return err
		}
	case *ClientArgs_Mark:
		b.EncodeVarint(2 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mark); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientArgs.Argtype has unexpected type %T", x)
	}
	return nil
}

func _ClientArgs_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientArgs)
	switch tag {
	case 1: // argtype.setup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientConfig)
		err := b.DecodeMessage(msg)
		m.Argtype = &ClientArgs_Setup{msg}
		return true, err
	case 2: // argtype.mark
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mark)
		err := b.DecodeMessage(msg)
		m.Argtype = &ClientArgs_Mark{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientArgs_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ClientArgs_Setup:
		s := proto.Size(x.Setup)
		n += proto.SizeVarint(1 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientArgs_Mark:
		s := proto.Size(x.Mark)
		n += proto.SizeVarint(2 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ServerConfig struct {
	ServerType         ServerType      `protobuf:"varint,1,opt,name=server_type,json=serverType,enum=grpc.testing.ServerType" json:"server_type,omitempty"`
	SecurityParams     *SecurityParams `protobuf:"bytes,2,opt,name=security_params,json=securityParams" json:"security_params,omitempty"`
	// Port on which to listen. Zero means pick unused port.
	Port               int32 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	// Only for async server. Number of threads used to serve the requests.
	AsyncServerThreads int32 `protobuf:"varint,7,opt,name=async_server_threads,json=asyncServerThreads" json:"async_server_threads,omitempty"`
	// Specify the number of cores to limit server to, if desired
	CoreLimit          int32 `protobuf:"varint,8,opt,name=core_limit,json=coreLimit" json:"core_limit,omitempty"`
	// payload config, used in generic server
	PayloadConfig      *PayloadConfig `protobuf:"bytes,9,opt,name=payload_config,json=payloadConfig" json:"payload_config,omitempty"`
	// Specify the cores we should run the server on, if desired
	CoreList           []int32 `protobuf:"varint,10,rep,name=core_list,json=coreList" json:"core_list,omitempty"`
}

func (m *ServerConfig) Reset() {
	*m = ServerConfig{}
}
func (m *ServerConfig) String() string {
	return proto.CompactTextString(m)
}
func (*ServerConfig) ProtoMessage() {}
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11}
}

func (m *ServerConfig) GetSecurityParams() *SecurityParams {
	if m != nil {
		return m.SecurityParams
	}
	return nil
}

func (m *ServerConfig) GetPayloadConfig() *PayloadConfig {
	if m != nil {
		return m.PayloadConfig
	}
	return nil
}

type ServerArgs struct {
	// Types that are valid to be assigned to Argtype:
	//	*ServerArgs_Setup
	//	*ServerArgs_Mark
	Argtype isServerArgs_Argtype `protobuf_oneof:"argtype"`
}

func (m *ServerArgs) Reset() {
	*m = ServerArgs{}
}
func (m *ServerArgs) String() string {
	return proto.CompactTextString(m)
}
func (*ServerArgs) ProtoMessage() {}
func (*ServerArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{12}
}

type isServerArgs_Argtype interface {
	isServerArgs_Argtype()
}

type ServerArgs_Setup struct {
	Setup *ServerConfig `protobuf:"bytes,1,opt,name=setup,oneof"`
}
type ServerArgs_Mark struct {
	Mark *Mark `protobuf:"bytes,2,opt,name=mark,oneof"`
}

func (*ServerArgs_Setup) isServerArgs_Argtype() {}
func (*ServerArgs_Mark) isServerArgs_Argtype() {}

func (m *ServerArgs) GetArgtype() isServerArgs_Argtype {
	if m != nil {
		return m.Argtype
	}
	return nil
}

func (m *ServerArgs) GetSetup() *ServerConfig {
	if x, ok := m.GetArgtype().(*ServerArgs_Setup); ok {
		return x.Setup
	}
	return nil
}

func (m *ServerArgs) GetMark() *Mark {
	if x, ok := m.GetArgtype().(*ServerArgs_Mark); ok {
		return x.Mark
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerArgs) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerArgs_OneofMarshaler, _ServerArgs_OneofUnmarshaler, _ServerArgs_OneofSizer, []interface{}{
		(*ServerArgs_Setup)(nil),
		(*ServerArgs_Mark)(nil),
	}
}

func _ServerArgs_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ServerArgs_Setup:
		b.EncodeVarint(1 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Setup); err != nil {
			return err
		}
	case *ServerArgs_Mark:
		b.EncodeVarint(2 << 3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mark); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerArgs.Argtype has unexpected type %T", x)
	}
	return nil
}

func _ServerArgs_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerArgs)
	switch tag {
	case 1: // argtype.setup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerConfig)
		err := b.DecodeMessage(msg)
		m.Argtype = &ServerArgs_Setup{msg}
		return true, err
	case 2: // argtype.mark
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mark)
		err := b.DecodeMessage(msg)
		m.Argtype = &ServerArgs_Mark{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerArgs_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerArgs)
	// argtype
	switch x := m.Argtype.(type) {
	case *ServerArgs_Setup:
		s := proto.Size(x.Setup)
		n += proto.SizeVarint(1 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerArgs_Mark:
		s := proto.Size(x.Mark)
		n += proto.SizeVarint(2 << 3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ServerStatus struct {
	Stats *ServerStats `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
	// the port bound by the server
	Port  int32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	// Number of cores available to the server
	Cores int32 `protobuf:"varint,3,opt,name=cores" json:"cores,omitempty"`
}

func (m *ServerStatus) Reset() {
	*m = ServerStatus{}
}
func (m *ServerStatus) String() string {
	return proto.CompactTextString(m)
}
func (*ServerStatus) ProtoMessage() {}
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13}
}

func (m *ServerStatus) GetStats() *ServerStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type CoreRequest struct {
}

func (m *CoreRequest) Reset() {
	*m = CoreRequest{}
}
func (m *CoreRequest) String() string {
	return proto.CompactTextString(m)
}
func (*CoreRequest) ProtoMessage() {}
func (*CoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{14}
}

type CoreResponse struct {
	// Number of cores available on the server
	Cores int32 `protobuf:"varint,1,opt,name=cores" json:"cores,omitempty"`
}

func (m *CoreResponse) Reset() {
	*m = CoreResponse{}
}
func (m *CoreResponse) String() string {
	return proto.CompactTextString(m)
}
func (*CoreResponse) ProtoMessage() {}
func (*CoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{15}
}

type Void struct {
}

func (m *Void) Reset() {
	*m = Void{}
}
func (m *Void) String() string {
	return proto.CompactTextString(m)
}
func (*Void) ProtoMessage() {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{16}
}

// A single performance scenario: input to qps_json_driver
type Scenario struct {
	// Human readable name for this scenario
	Name                  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Client configuration
	ClientConfig          *ClientConfig `protobuf:"bytes,2,opt,name=client_config,json=clientConfig" json:"client_config,omitempty"`
	// Number of clients to start for the test
	NumClients            int32 `protobuf:"varint,3,opt,name=num_clients,json=numClients" json:"num_clients,omitempty"`
	// Server configuration
	ServerConfig          *ServerConfig `protobuf:"bytes,4,opt,name=server_config,json=serverConfig" json:"server_config,omitempty"`
	// Number of servers to start for the test
	NumServers            int32 `protobuf:"varint,5,opt,name=num_servers,json=numServers" json:"num_servers,omitempty"`
	// Warmup period, in seconds
	WarmupSeconds         int32 `protobuf:"varint,6,opt,name=warmup_seconds,json=warmupSeconds" json:"warmup_seconds,omitempty"`
	// Benchmark time, in seconds
	BenchmarkSeconds      int32 `protobuf:"varint,7,opt,name=benchmark_seconds,json=benchmarkSeconds" json:"benchmark_seconds,omitempty"`
	// Number of workers to spawn locally (usually zero)
	SpawnLocalWorkerCount int32 `protobuf:"varint,8,opt,name=spawn_local_worker_count,json=spawnLocalWorkerCount" json:"spawn_local_worker_count,omitempty"`
}

func (m *Scenario) Reset() {
	*m = Scenario{}
}
func (m *Scenario) String() string {
	return proto.CompactTextString(m)
}
func (*Scenario) ProtoMessage() {}
func (*Scenario) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{17}
}

func (m *Scenario) GetClientConfig() *ClientConfig {
	if m != nil {
		return m.ClientConfig
	}
	return nil
}

func (m *Scenario) GetServerConfig() *ServerConfig {
	if m != nil {
		return m.ServerConfig
	}
	return nil
}

// A set of scenarios to be run with qps_json_driver
type Scenarios struct {
	Scenarios []*Scenario `protobuf:"bytes,1,rep,name=scenarios" json:"scenarios,omitempty"`
}

func (m *Scenarios) Reset() {
	*m = Scenarios{}
}
func (m *Scenarios) String() string {
	return proto.CompactTextString(m)
}
func (*Scenarios) ProtoMessage() {}
func (*Scenarios) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{18}
}

func (m *Scenarios) GetScenarios() []*Scenario {
	if m != nil {
		return m.Scenarios
	}
	return nil
}

func init() {
	proto.RegisterType((*PoissonParams)(nil), "grpc.testing.PoissonParams")
	proto.RegisterType((*UniformParams)(nil), "grpc.testing.UniformParams")
	proto.RegisterType((*DeterministicParams)(nil), "grpc.testing.DeterministicParams")
	proto.RegisterType((*ParetoParams)(nil), "grpc.testing.ParetoParams")
	proto.RegisterType((*ClosedLoopParams)(nil), "grpc.testing.ClosedLoopParams")
	proto.RegisterType((*LoadParams)(nil), "grpc.testing.LoadParams")
	proto.RegisterType((*SecurityParams)(nil), "grpc.testing.SecurityParams")
	proto.RegisterType((*ClientConfig)(nil), "grpc.testing.ClientConfig")
	proto.RegisterType((*ClientStatus)(nil), "grpc.testing.ClientStatus")
	proto.RegisterType((*Mark)(nil), "grpc.testing.Mark")
	proto.RegisterType((*ClientArgs)(nil), "grpc.testing.ClientArgs")
	proto.RegisterType((*ServerConfig)(nil), "grpc.testing.ServerConfig")
	proto.RegisterType((*ServerArgs)(nil), "grpc.testing.ServerArgs")
	proto.RegisterType((*ServerStatus)(nil), "grpc.testing.ServerStatus")
	proto.RegisterType((*CoreRequest)(nil), "grpc.testing.CoreRequest")
	proto.RegisterType((*CoreResponse)(nil), "grpc.testing.CoreResponse")
	proto.RegisterType((*Void)(nil), "grpc.testing.Void")
	proto.RegisterType((*Scenario)(nil), "grpc.testing.Scenario")
	proto.RegisterType((*Scenarios)(nil), "grpc.testing.Scenarios")
	proto.RegisterEnum("grpc.testing.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("grpc.testing.ServerType", ServerType_name, ServerType_value)
	proto.RegisterEnum("grpc.testing.RpcType", RpcType_name, RpcType_value)
}

func init() {
	proto.RegisterFile("control.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0x8d, 0x14, 0xc9, 0x96, 0x86, 0x92, 0xac, 0x6f, 0xbf, 0xa4, 0x60, 0x1c, 0x27, 0x6d, 0xd8,
	0x16, 0x0d, 0x5c, 0xc0, 0x29, 0xd4, 0x02, 0x69, 0xd1, 0x8b, 0x40, 0x56, 0x85, 0xd8, 0x80, 0xe3,
	0xba, 0x2b, 0x27, 0x45, 0xae, 0x08, 0x9a, 0x5a, 0x4b, 0x44, 0x24, 0x2e, 0xbb, 0x4b, 0xc6, 0xf0,
	0x2b, 0xf4, 0x99, 0xfa, 0x1c, 0x7d, 0x8d, 0xbe, 0x42, 0x67, 0xff, 0x64, 0x52, 0x11, 0x10, 0xb7,
	0xbd, 0xe3, 0xce, 0x9c, 0xb3, 0x3b, 0x3b, 0x67, 0x66, 0x96, 0xd0, 0x8d, 0x79, 0x9a, 0x0b, 0xbe,
	0x38, 0xc8, 0x04, 0xcf, 0x39, 0xe9, 0xcc, 0x44, 0x16, 0x1f, 0xe4, 0x4c, 0xe6, 0x49, 0x3a, 0xdb,
	0xed, 0x65, 0xd1, 0xf5, 0x82, 0x47, 0x53, 0x69, 0xbc, 0xbb, 0x9e, 0xcc, 0xa3, 0xdc, 0x2e, 0x82,
	0x01, 0x74, 0xcf, 0x78, 0x22, 0x25, 0x4f, 0xcf, 0x22, 0x11, 0x2d, 0x25, 0x79, 0x02, 0x1d, 0x7e,
	0x79, 0xc9, 0x04, 0x9b, 0x86, 0x8a, 0xe4, 0xd7, 0x3e, 0xab, 0x3d, 0xad, 0x51, 0xcf, 0xda, 0x4e,
	0xd0, 0x14, 0x44, 0xd0, 0x7d, 0x9d, 0x26, 0x97, 0x5c, 0x2c, 0x2d, 0xe7, 0x2b, 0xd8, 0x49, 0xd2,
	0x9c, 0x89, 0x48, 0x88, 0xe4, 0x7d, 0xb4, 0x40, 0xa2, 0xa5, 0xf5, 0xca, 0xe6, 0x13, 0xfe, 0x01,
	0x70, 0x9e, 0xf8, 0xf5, 0x0f, 0x81, 0x47, 0x49, 0xf0, 0x3d, 0xfc, 0xff, 0x27, 0x86, 0x96, 0x65,
	0x92, 0x26, 0x78, 0x8b, 0xf8, 0xf6, 0xc1, 0xfd, 0x02, 0x1d, 0x04, 0xb3, 0x9c, 0x5b, 0xca, 0xd7,
	0xf0, 0xbf, 0xca, 0x91, 0x17, 0x91, 0x64, 0x96, 0xd7, 0x2f, 0x3b, 0x0e, 0xd1, 0x4e, 0xee, 0x41,
	0x33, 0x5a, 0x64, 0xf3, 0xc8, 0x46, 0x65, 0x16, 0x01, 0x81, 0xfe, 0x68, 0xc1, 0xa5, 0x3a, 0x80,
	0x67, 0x66, 0xdb, 0xe0, 0x8f, 0x3a, 0x80, 0x3a, 0xcf, 0x9e, 0x32, 0x04, 0x2f, 0xd6, 0x10, 0x8c,
	0x8b, 0x67, 0x7a, 0x7f, 0x6f, 0xf0, 0xf8, 0xa0, 0xac, 0xc3, 0xc1, 0xfa, 0x1e, 0x47, 0x77, 0x28,
	0xc4, 0x2b, 0x1b, 0x79, 0x0e, 0xdb, 0x99, 0x51, 0x42, 0x9f, 0xee, 0x0d, 0x1e, 0x56, 0xe9, 0x15,
	0x99, 0x90, 0xeb, 0xd0, 0x8a, 0x58, 0x18, 0x39, 0xfc, 0xbb, 0x9b, 0x88, 0x15, 0xad, 0x14, 0xd1,
	0xa2, 0xc9, 0x8f, 0xb0, 0x35, 0xd5, 0x49, 0xf6, 0x1b, 0x9a, 0xf7, 0xa4, 0xca, 0xdb, 0x20, 0x00,
	0xb2, 0x2d, 0x85, 0x7c, 0x07, 0x5b, 0x99, 0xce, 0xb3, 0xdf, 0xd4, 0xe4, 0xdd, 0xb5, 0x68, 0x4b,
	0x1a, 0x28, 0x96, 0xc1, 0x1e, 0x6e, 0x41, 0x43, 0x09, 0x17, 0x5c, 0x40, 0x6f, 0xc2, 0xe2, 0x42,
	0x24, 0xf9, 0xb5, 0xcd, 0xe0, 0x63, 0xf0, 0x0a, 0xc9, 0x42, 0xc5, 0x0f, 0xe3, 0x48, 0x67, 0xb0,
	0x45, 0xdb, 0x68, 0x3a, 0x47, 0xcb, 0x28, 0x22, 0xdf, 0xc0, 0x3d, 0xc9, 0xc4, 0x7b, 0x26, 0xc2,
	0x39, 0x47, 0x08, 0xc7, 0x2f, 0x91, 0x4c, 0x99, 0xce, 0x55, 0x9b, 0x12, 0xe3, 0x3b, 0x42, 0xd7,
	0xcf, 0xd6, 0x13, 0xfc, 0xde, 0x84, 0xce, 0x68, 0x91, 0xb0, 0x34, 0x1f, 0xf1, 0xf4, 0x32, 0x99,
	0x91, 0x2f, 0xa1, 0x67, 0xb7, 0xc8, 0x23, 0x31, 0x63, 0xb9, 0xc4, 0x53, 0xee, 0x22, 0xb9, 0x6b,
	0xac, 0xe7, 0xc6, 0x48, 0x7e, 0x50, 0x5a, 0x2a, 0x5a, 0x98, 0x5f, 0x67, 0xe6, 0x80, 0xde, 0xc0,
	0x5f, 0xd7, 0x52, 0x01, 0xce, 0xd1, 0xaf, 0x34, 0x74, 0xdf, 0x64, 0x0c, 0x3b, 0xd2, 0x5e, 0x2b,
	0xcc, 0xf4, 0xbd, 0xac, 0x24, 0x7b, 0x55, 0x7a, 0xf5, 0xee, 0xb4, 0x27, 0xab, 0xb9, 0x78, 0x01,
	0x7b, 0xbc, 0xc8, 0xb1, 0x4d, 0xd3, 0x29, 0xa2, 0x43, 0x64, 0xca, 0x30, 0xc3, 0xb0, 0xe3, 0x79,
	0x94, 0xa6, 0x6c, 0xa1, 0xe5, 0x6a, 0xd2, 0x07, 0x25, 0x0c, 0x45, 0xc8, 0x19, 0x13, 0x23, 0x03,
	0x50, 0x7d, 0x66, 0xaf, 0x60, 0x29, 0x52, 0xab, 0xd4, 0xa4, 0x3d, 0x63, 0xb6, 0x38, 0xa9, 0xb2,
	0x1a, 0xc9, 0xeb, 0x34, 0x0e, 0xdd, 0x8d, 0xe7, 0x82, 0xe1, 0xa4, 0xf0, 0xb7, 0x35, 0x9a, 0x68,
	0x9f, 0xbd, 0xab, 0xf1, 0x20, 0xa3, 0x85, 0xf1, 0x98, 0xd4, 0xb4, 0x74, 0x6a, 0xee, 0x57, 0xef,
	0x86, 0xa1, 0xe8, 0xbc, 0x6c, 0x0b, 0xf3, 0xa1, 0xf2, 0xa9, 0x34, 0x77, 0x09, 0x01, 0x9d, 0x90,
	0xb5, 0x7c, 0xde, 0xb4, 0x12, 0x85, 0xc5, 0x4d, 0x5b, 0x1d, 0x82, 0x1b, 0x5e, 0x61, 0xac, 0x35,
	0xf4, 0xbd, 0x8d, 0xad, 0x61, 0x30, 0x46, 0x66, 0xda, 0xcd, 0xca, 0x4b, 0x72, 0x04, 0xfd, 0x39,
	0x96, 0x30, 0x9f, 0xe1, 0x8e, 0x2e, 0x86, 0x8e, 0xde, 0xe5, 0x51, 0x75, 0x97, 0x23, 0x87, 0xb2,
	0x81, 0xec, 0xcc, 0xab, 0x06, 0xf2, 0x10, 0xda, 0x31, 0x17, 0x2c, 0x5c, 0xa0, 0xdd, 0xef, 0x62,
	0xe9, 0x34, 0x69, 0x4b, 0x19, 0x4e, 0x70, 0x4d, 0x1e, 0x01, 0x58, 0xe7, 0x32, 0xc9, 0xfd, 0x9e,
	0xce, 0x5f, 0xdb, 0x78, 0xd1, 0x10, 0xbc, 0x70, 0xb5, 0x38, 0xc1, 0xe1, 0x5b, 0x48, 0xf2, 0x0c,
	0x9a, 0x7a, 0x0c, 0xdb, 0x51, 0xf1, 0x60, 0x53, 0x79, 0x29, 0xa8, 0xa4, 0x06, 0x17, 0xec, 0x41,
	0xe3, 0x55, 0x24, 0xde, 0xa9, 0x11, 0x25, 0x98, 0x64, 0xb9, 0xed, 0x10, 0xb3, 0x08, 0x0a, 0x00,
	0xc3, 0x19, 0x8a, 0x99, 0x24, 0x03, 0xdc, 0x9c, 0xe5, 0x85, 0x9b, 0x43, 0xbb, 0x9b, 0x36, 0x37,
	0xd9, 0xc1, 0xd6, 0x34, 0x50, 0xf2, 0x14, 0x1a, 0x4b, 0xdc, 0xdf, 0xce, 0x1e, 0x52, 0xa5, 0xa8,
	0x93, 0x11, 0xaa, 0x11, 0x87, 0x6d, 0xd8, 0xc6, 0x4e, 0x51, 0x05, 0x10, 0xfc, 0x59, 0x87, 0xce,
	0x44, 0x37, 0x8f, 0x4d, 0x36, 0x6a, 0xed, 0x5a, 0x4c, 0x15, 0x48, 0x6d, 0x53, 0xef, 0x18, 0x82,
	0xe9, 0x1d, 0xb9, 0xfa, 0xde, 0xd4, 0x3b, 0xf5, 0x7f, 0xd1, 0x3b, 0x04, 0x1a, 0x19, 0x17, 0xb9,
	0xed, 0x11, 0xfd, 0x7d, 0x53, 0xe5, 0x2e, 0xb6, 0x0d, 0x55, 0x6e, 0xa3, 0xb2, 0x55, 0x5e, 0x55,
	0xb3, 0xb5, 0xa6, 0xe6, 0x86, 0xba, 0x6c, 0xff, 0xe3, 0xba, 0xac, 0x54, 0x13, 0x54, 0xab, 0x49,
	0xe9, 0x69, 0x02, 0xba, 0x85, 0x9e, 0x65, 0x01, 0xfe, 0xa3, 0x9e, 0x89, 0x93, 0xf3, 0x56, 0x55,
	0x7a, 0x03, 0x75, 0x55, 0xba, 0xca, 0x7e, 0xbd, 0x94, 0x7d, 0xac, 0x58, 0x75, 0x2f, 0x33, 0x0a,
	0x9b, 0xd4, 0x2c, 0x82, 0x2e, 0x78, 0x23, 0xfc, 0xa0, 0xec, 0xb7, 0x02, 0xb7, 0x0b, 0xbe, 0xc0,
	0xfe, 0xd0, 0x4b, 0x99, 0xf1, 0xd4, 0xbc, 0xc4, 0x86, 0x54, 0x2b, 0x93, 0xf0, 0xf9, 0x78, 0xc3,
	0x93, 0x69, 0xf0, 0x57, 0x1d, 0x5a, 0x93, 0x98, 0xa5, 0x91, 0x48, 0xb8, 0x3a, 0x33, 0x8d, 0x96,
	0xa6, 0xd8, 0xda, 0x54, 0x7f, 0xe3, 0x04, 0xed, 0xba, 0x01, 0x68, 0xf4, 0xa9, 0x7f, 0xac, 0x13,
	0x68, 0x27, 0x2e, 0xbf, 0x15, 0x9f, 0x82, 0x97, 0x16, 0x4b, 0x3b, 0x16, 0x5d, 0xe8, 0x80, 0x26,
	0xc3, 0x51, 0x33, 0xda, 0x3e, 0x1b, 0xee, 0x84, 0xc6, 0xc7, 0xb4, 0xa1, 0x1d, 0x59, 0x6e, 0x15,
	0x7b, 0x82, 0xb1, 0xb9, 0xf9, 0xac, 0x4e, 0x30, 0x1c, 0xa9, 0x9e, 0xab, 0xab, 0x48, 0x2c, 0x8b,
	0x0c, 0x31, 0x78, 0x06, 0xd6, 0xeb, 0x96, 0xc6, 0x74, 0x8d, 0x75, 0x62, 0x8c, 0xea, 0x07, 0xe7,
	0x82, 0xa5, 0xf1, 0x5c, 0x69, 0xb9, 0x42, 0x9a, 0xca, 0xee, 0xaf, 0x1c, 0x0e, 0xfc, 0x1c, 0x7c,
	0x99, 0x45, 0x57, 0x29, 0xfe, 0xa6, 0xc4, 0xf8, 0x33, 0x74, 0xc5, 0xc5, 0x3b, 0x7d, 0x83, 0x22,
	0x75, 0x55, 0x7e, 0x5f, 0xfb, 0x4f, 0x94, 0xfb, 0x57, 0xed, 0x1d, 0x29, 0x67, 0x30, 0x84, 0xb6,
	0x4b, 0xb8, 0xc4, 0xb7, 0xbf, 0x2d, 0xdd, 0x42, 0xbf, 0xa1, 0xde, 0xe0, 0x93, 0xb5, 0x7b, 0x5b,
	0x37, 0xbd, 0x01, 0xee, 0x3f, 0x73, 0x33, 0x4a, 0xb7, 0xfb, 0x0e, 0x78, 0x93, 0xb7, 0xa7, 0xa3,
	0x70, 0x74, 0x72, 0x3c, 0x3e, 0x3d, 0xef, 0xdf, 0x21, 0x7d, 0xe8, 0x0c, 0xcb, 0x96, 0xda, 0xfe,
	0xb1, 0x6b, 0x82, 0x0a, 0x61, 0x32, 0xa6, 0x6f, 0xc6, 0xb4, 0x4c, 0xb0, 0x96, 0x1a, 0xf1, 0xe1,
	0x9e, 0xb1, 0xbc, 0x1c, 0x9f, 0x8e, 0xe9, 0xf1, 0xca, 0x53, 0xdf, 0xff, 0x1c, 0xb6, 0xed, 0xbb,
	0x44, 0xda, 0xd0, 0x7c, 0x7d, 0x3a, 0xa4, 0x6f, 0x71, 0x87, 0x2e, 0x5e, 0xea, 0x9c, 0x8e, 0x87,
	0xaf, 0x8e, 0x4f, 0x5f, 0xf6, 0x6b, 0x17, 0x5b, 0xfa, 0x97, 0xf8, 0xdb, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x75, 0x59, 0xf4, 0x03, 0x4e, 0x0b, 0x00, 0x00,
}
