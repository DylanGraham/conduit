// Code generated by protoc-gen-go. DO NOT EDIT.
// source: public/api.proto

/*
Package conduit_public is a generated protocol buffer package.

It is generated from these files:
	public/api.proto

It has these top-level messages:
	HistogramValue
	Histogram
	MetricValue
	MetricDatapoint
	MetricSeries
	MetricMetadata
	MetricResponse
	MetricRequest
	Empty
	VersionInfo
	ListPodsResponse
	Pod
	TapRequest
*/
package conduit_public

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import conduit_common "github.com/runconduit/conduit/controller/gen/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetricName int32

const (
	MetricName_REQUEST_RATE MetricName = 0
	MetricName_LATENCY      MetricName = 1
	MetricName_SUCCESS_RATE MetricName = 2
)

var MetricName_name = map[int32]string{
	0: "REQUEST_RATE",
	1: "LATENCY",
	2: "SUCCESS_RATE",
}
var MetricName_value = map[string]int32{
	"REQUEST_RATE": 0,
	"LATENCY":      1,
	"SUCCESS_RATE": 2,
}

func (x MetricName) String() string {
	return proto.EnumName(MetricName_name, int32(x))
}
func (MetricName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TimeWindow int32

const (
	TimeWindow_TEN_SEC  TimeWindow = 0
	TimeWindow_ONE_MIN  TimeWindow = 1
	TimeWindow_TEN_MIN  TimeWindow = 2
	TimeWindow_ONE_HOUR TimeWindow = 3
)

var TimeWindow_name = map[int32]string{
	0: "TEN_SEC",
	1: "ONE_MIN",
	2: "TEN_MIN",
	3: "ONE_HOUR",
}
var TimeWindow_value = map[string]int32{
	"TEN_SEC":  0,
	"ONE_MIN":  1,
	"TEN_MIN":  2,
	"ONE_HOUR": 3,
}

func (x TimeWindow) String() string {
	return proto.EnumName(TimeWindow_name, int32(x))
}
func (TimeWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AggregationType int32

const (
	AggregationType_TARGET_POD    AggregationType = 0
	AggregationType_TARGET_DEPLOY AggregationType = 1
	AggregationType_SOURCE_POD    AggregationType = 2
	AggregationType_SOURCE_DEPLOY AggregationType = 3
	AggregationType_MESH          AggregationType = 4
)

var AggregationType_name = map[int32]string{
	0: "TARGET_POD",
	1: "TARGET_DEPLOY",
	2: "SOURCE_POD",
	3: "SOURCE_DEPLOY",
	4: "MESH",
}
var AggregationType_value = map[string]int32{
	"TARGET_POD":    0,
	"TARGET_DEPLOY": 1,
	"SOURCE_POD":    2,
	"SOURCE_DEPLOY": 3,
	"MESH":          4,
}

func (x AggregationType) String() string {
	return proto.EnumName(AggregationType_name, int32(x))
}
func (AggregationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type HistogramLabel int32

const (
	HistogramLabel_MIN HistogramLabel = 0
	HistogramLabel_P50 HistogramLabel = 1
	HistogramLabel_P95 HistogramLabel = 2
	HistogramLabel_P99 HistogramLabel = 3
	HistogramLabel_MAX HistogramLabel = 4
)

var HistogramLabel_name = map[int32]string{
	0: "MIN",
	1: "P50",
	2: "P95",
	3: "P99",
	4: "MAX",
}
var HistogramLabel_value = map[string]int32{
	"MIN": 0,
	"P50": 1,
	"P95": 2,
	"P99": 3,
	"MAX": 4,
}

func (x HistogramLabel) String() string {
	return proto.EnumName(HistogramLabel_name, int32(x))
}
func (HistogramLabel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type HistogramValue struct {
	Label HistogramLabel `protobuf:"varint,1,opt,name=label,enum=conduit.public.HistogramLabel" json:"label,omitempty"`
	Value int64          `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *HistogramValue) Reset()                    { *m = HistogramValue{} }
func (m *HistogramValue) String() string            { return proto.CompactTextString(m) }
func (*HistogramValue) ProtoMessage()               {}
func (*HistogramValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HistogramValue) GetLabel() HistogramLabel {
	if m != nil {
		return m.Label
	}
	return HistogramLabel_MIN
}

func (m *HistogramValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Histogram struct {
	Values []*HistogramValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *Histogram) Reset()                    { *m = Histogram{} }
func (m *Histogram) String() string            { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()               {}
func (*Histogram) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Histogram) GetValues() []*HistogramValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type MetricValue struct {
	// Types that are valid to be assigned to Value:
	//	*MetricValue_Counter
	//	*MetricValue_Gauge
	//	*MetricValue_Histogram
	Value isMetricValue_Value `protobuf_oneof:"value"`
}

func (m *MetricValue) Reset()                    { *m = MetricValue{} }
func (m *MetricValue) String() string            { return proto.CompactTextString(m) }
func (*MetricValue) ProtoMessage()               {}
func (*MetricValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isMetricValue_Value interface {
	isMetricValue_Value()
}

type MetricValue_Counter struct {
	Counter int64 `protobuf:"varint,1,opt,name=counter,oneof"`
}
type MetricValue_Gauge struct {
	Gauge float64 `protobuf:"fixed64,2,opt,name=gauge,oneof"`
}
type MetricValue_Histogram struct {
	Histogram *Histogram `protobuf:"bytes,3,opt,name=histogram,oneof"`
}

func (*MetricValue_Counter) isMetricValue_Value()   {}
func (*MetricValue_Gauge) isMetricValue_Value()     {}
func (*MetricValue_Histogram) isMetricValue_Value() {}

func (m *MetricValue) GetValue() isMetricValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricValue) GetCounter() int64 {
	if x, ok := m.GetValue().(*MetricValue_Counter); ok {
		return x.Counter
	}
	return 0
}

func (m *MetricValue) GetGauge() float64 {
	if x, ok := m.GetValue().(*MetricValue_Gauge); ok {
		return x.Gauge
	}
	return 0
}

func (m *MetricValue) GetHistogram() *Histogram {
	if x, ok := m.GetValue().(*MetricValue_Histogram); ok {
		return x.Histogram
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricValue_OneofMarshaler, _MetricValue_OneofUnmarshaler, _MetricValue_OneofSizer, []interface{}{
		(*MetricValue_Counter)(nil),
		(*MetricValue_Gauge)(nil),
		(*MetricValue_Histogram)(nil),
	}
}

func _MetricValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricValue)
	// value
	switch x := m.Value.(type) {
	case *MetricValue_Counter:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Counter))
	case *MetricValue_Gauge:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Gauge))
	case *MetricValue_Histogram:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Histogram); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MetricValue.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricValue)
	switch tag {
	case 1: // value.counter
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricValue_Counter{int64(x)}
		return true, err
	case 2: // value.gauge
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricValue_Gauge{math.Float64frombits(x)}
		return true, err
	case 3: // value.histogram
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Histogram)
		err := b.DecodeMessage(msg)
		m.Value = &MetricValue_Histogram{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MetricValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricValue)
	// value
	switch x := m.Value.(type) {
	case *MetricValue_Counter:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Counter))
	case *MetricValue_Gauge:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *MetricValue_Histogram:
		s := proto.Size(x.Histogram)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type MetricDatapoint struct {
	Value       *MetricValue `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64        `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *MetricDatapoint) Reset()                    { *m = MetricDatapoint{} }
func (m *MetricDatapoint) String() string            { return proto.CompactTextString(m) }
func (*MetricDatapoint) ProtoMessage()               {}
func (*MetricDatapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetricDatapoint) GetValue() *MetricValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricDatapoint) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

type MetricSeries struct {
	Name       MetricName         `protobuf:"varint,1,opt,name=name,enum=conduit.public.MetricName" json:"name,omitempty"`
	Metadata   *MetricMetadata    `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	Datapoints []*MetricDatapoint `protobuf:"bytes,3,rep,name=datapoints" json:"datapoints,omitempty"`
}

func (m *MetricSeries) Reset()                    { *m = MetricSeries{} }
func (m *MetricSeries) String() string            { return proto.CompactTextString(m) }
func (*MetricSeries) ProtoMessage()               {}
func (*MetricSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MetricSeries) GetName() MetricName {
	if m != nil {
		return m.Name
	}
	return MetricName_REQUEST_RATE
}

func (m *MetricSeries) GetMetadata() *MetricMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MetricSeries) GetDatapoints() []*MetricDatapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

type MetricMetadata struct {
	TargetPod    string `protobuf:"bytes,1,opt,name=targetPod" json:"targetPod,omitempty"`
	TargetDeploy string `protobuf:"bytes,2,opt,name=targetDeploy" json:"targetDeploy,omitempty"`
	SourcePod    string `protobuf:"bytes,3,opt,name=sourcePod" json:"sourcePod,omitempty"`
	SourceDeploy string `protobuf:"bytes,4,opt,name=sourceDeploy" json:"sourceDeploy,omitempty"`
	Component    string `protobuf:"bytes,5,opt,name=component" json:"component,omitempty"`
}

func (m *MetricMetadata) Reset()                    { *m = MetricMetadata{} }
func (m *MetricMetadata) String() string            { return proto.CompactTextString(m) }
func (*MetricMetadata) ProtoMessage()               {}
func (*MetricMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MetricMetadata) GetTargetPod() string {
	if m != nil {
		return m.TargetPod
	}
	return ""
}

func (m *MetricMetadata) GetTargetDeploy() string {
	if m != nil {
		return m.TargetDeploy
	}
	return ""
}

func (m *MetricMetadata) GetSourcePod() string {
	if m != nil {
		return m.SourcePod
	}
	return ""
}

func (m *MetricMetadata) GetSourceDeploy() string {
	if m != nil {
		return m.SourceDeploy
	}
	return ""
}

func (m *MetricMetadata) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

type MetricResponse struct {
	Metrics []*MetricSeries `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *MetricResponse) Reset()                    { *m = MetricResponse{} }
func (m *MetricResponse) String() string            { return proto.CompactTextString(m) }
func (*MetricResponse) ProtoMessage()               {}
func (*MetricResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MetricResponse) GetMetrics() []*MetricSeries {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type MetricRequest struct {
	Metrics   []MetricName    `protobuf:"varint,1,rep,packed,name=metrics,enum=conduit.public.MetricName" json:"metrics,omitempty"`
	Window    TimeWindow      `protobuf:"varint,2,opt,name=window,enum=conduit.public.TimeWindow" json:"window,omitempty"`
	GroupBy   AggregationType `protobuf:"varint,3,opt,name=groupBy,enum=conduit.public.AggregationType" json:"groupBy,omitempty"`
	FilterBy  *MetricMetadata `protobuf:"bytes,4,opt,name=filterBy" json:"filterBy,omitempty"`
	Summarize bool            `protobuf:"varint,5,opt,name=summarize" json:"summarize,omitempty"`
}

func (m *MetricRequest) Reset()                    { *m = MetricRequest{} }
func (m *MetricRequest) String() string            { return proto.CompactTextString(m) }
func (*MetricRequest) ProtoMessage()               {}
func (*MetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MetricRequest) GetMetrics() []MetricName {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *MetricRequest) GetWindow() TimeWindow {
	if m != nil {
		return m.Window
	}
	return TimeWindow_TEN_SEC
}

func (m *MetricRequest) GetGroupBy() AggregationType {
	if m != nil {
		return m.GroupBy
	}
	return AggregationType_TARGET_POD
}

func (m *MetricRequest) GetFilterBy() *MetricMetadata {
	if m != nil {
		return m.FilterBy
	}
	return nil
}

func (m *MetricRequest) GetSummarize() bool {
	if m != nil {
		return m.Summarize
	}
	return false
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type VersionInfo struct {
	GoVersion      string `protobuf:"bytes,1,opt,name=goVersion" json:"goVersion,omitempty"`
	BuildDate      string `protobuf:"bytes,2,opt,name=buildDate" json:"buildDate,omitempty"`
	ReleaseVersion string `protobuf:"bytes,3,opt,name=releaseVersion" json:"releaseVersion,omitempty"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *VersionInfo) GetGoVersion() string {
	if m != nil {
		return m.GoVersion
	}
	return ""
}

func (m *VersionInfo) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *VersionInfo) GetReleaseVersion() string {
	if m != nil {
		return m.ReleaseVersion
	}
	return ""
}

type ListPodsResponse struct {
	Pods []*Pod `protobuf:"bytes,1,rep,name=pods" json:"pods,omitempty"`
}

func (m *ListPodsResponse) Reset()                    { *m = ListPodsResponse{} }
func (m *ListPodsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPodsResponse) ProtoMessage()               {}
func (*ListPodsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListPodsResponse) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type Pod struct {
	Name                string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PodIP               string                    `protobuf:"bytes,2,opt,name=podIP" json:"podIP,omitempty"`
	Deployment          string                    `protobuf:"bytes,3,opt,name=deployment" json:"deployment,omitempty"`
	Status              string                    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Added               bool                      `protobuf:"varint,5,opt,name=added" json:"added,omitempty"`
	SinceLastReport     *google_protobuf.Duration `protobuf:"bytes,6,opt,name=sinceLastReport" json:"sinceLastReport,omitempty"`
	ControllerNamespace string                    `protobuf:"bytes,7,opt,name=controllerNamespace" json:"controllerNamespace,omitempty"`
	ControlPlane        bool                      `protobuf:"varint,8,opt,name=controlPlane" json:"controlPlane,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Pod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pod) GetPodIP() string {
	if m != nil {
		return m.PodIP
	}
	return ""
}

func (m *Pod) GetDeployment() string {
	if m != nil {
		return m.Deployment
	}
	return ""
}

func (m *Pod) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Pod) GetAdded() bool {
	if m != nil {
		return m.Added
	}
	return false
}

func (m *Pod) GetSinceLastReport() *google_protobuf.Duration {
	if m != nil {
		return m.SinceLastReport
	}
	return nil
}

func (m *Pod) GetControllerNamespace() string {
	if m != nil {
		return m.ControllerNamespace
	}
	return ""
}

func (m *Pod) GetControlPlane() bool {
	if m != nil {
		return m.ControlPlane
	}
	return false
}

type TapRequest struct {
	// Types that are valid to be assigned to Target:
	//	*TapRequest_Pod
	//	*TapRequest_Deployment
	Target isTapRequest_Target `protobuf_oneof:"target"`
	// validation of these fields happens on the server
	MaxRps    float32 `protobuf:"fixed32,3,opt,name=maxRps" json:"maxRps,omitempty"`
	ToPort    uint32  `protobuf:"varint,4,opt,name=toPort" json:"toPort,omitempty"`
	ToIP      string  `protobuf:"bytes,5,opt,name=toIP" json:"toIP,omitempty"`
	FromPort  uint32  `protobuf:"varint,6,opt,name=fromPort" json:"fromPort,omitempty"`
	FromIP    string  `protobuf:"bytes,7,opt,name=fromIP" json:"fromIP,omitempty"`
	Scheme    string  `protobuf:"bytes,8,opt,name=scheme" json:"scheme,omitempty"`
	Method    string  `protobuf:"bytes,9,opt,name=method" json:"method,omitempty"`
	Authority string  `protobuf:"bytes,10,opt,name=authority" json:"authority,omitempty"`
	Path      string  `protobuf:"bytes,11,opt,name=path" json:"path,omitempty"`
}

func (m *TapRequest) Reset()                    { *m = TapRequest{} }
func (m *TapRequest) String() string            { return proto.CompactTextString(m) }
func (*TapRequest) ProtoMessage()               {}
func (*TapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type isTapRequest_Target interface {
	isTapRequest_Target()
}

type TapRequest_Pod struct {
	Pod string `protobuf:"bytes,1,opt,name=pod,oneof"`
}
type TapRequest_Deployment struct {
	Deployment string `protobuf:"bytes,2,opt,name=deployment,oneof"`
}

func (*TapRequest_Pod) isTapRequest_Target()        {}
func (*TapRequest_Deployment) isTapRequest_Target() {}

func (m *TapRequest) GetTarget() isTapRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TapRequest) GetPod() string {
	if x, ok := m.GetTarget().(*TapRequest_Pod); ok {
		return x.Pod
	}
	return ""
}

func (m *TapRequest) GetDeployment() string {
	if x, ok := m.GetTarget().(*TapRequest_Deployment); ok {
		return x.Deployment
	}
	return ""
}

func (m *TapRequest) GetMaxRps() float32 {
	if m != nil {
		return m.MaxRps
	}
	return 0
}

func (m *TapRequest) GetToPort() uint32 {
	if m != nil {
		return m.ToPort
	}
	return 0
}

func (m *TapRequest) GetToIP() string {
	if m != nil {
		return m.ToIP
	}
	return ""
}

func (m *TapRequest) GetFromPort() uint32 {
	if m != nil {
		return m.FromPort
	}
	return 0
}

func (m *TapRequest) GetFromIP() string {
	if m != nil {
		return m.FromIP
	}
	return ""
}

func (m *TapRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *TapRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *TapRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *TapRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TapRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TapRequest_OneofMarshaler, _TapRequest_OneofUnmarshaler, _TapRequest_OneofSizer, []interface{}{
		(*TapRequest_Pod)(nil),
		(*TapRequest_Deployment)(nil),
	}
}

func _TapRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TapRequest)
	// target
	switch x := m.Target.(type) {
	case *TapRequest_Pod:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Pod)
	case *TapRequest_Deployment:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Deployment)
	case nil:
	default:
		return fmt.Errorf("TapRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _TapRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TapRequest)
	switch tag {
	case 1: // target.pod
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Target = &TapRequest_Pod{x}
		return true, err
	case 2: // target.deployment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Target = &TapRequest_Deployment{x}
		return true, err
	default:
		return false, nil
	}
}

func _TapRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TapRequest)
	// target
	switch x := m.Target.(type) {
	case *TapRequest_Pod:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Pod)))
		n += len(x.Pod)
	case *TapRequest_Deployment:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Deployment)))
		n += len(x.Deployment)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HistogramValue)(nil), "conduit.public.HistogramValue")
	proto.RegisterType((*Histogram)(nil), "conduit.public.Histogram")
	proto.RegisterType((*MetricValue)(nil), "conduit.public.MetricValue")
	proto.RegisterType((*MetricDatapoint)(nil), "conduit.public.MetricDatapoint")
	proto.RegisterType((*MetricSeries)(nil), "conduit.public.MetricSeries")
	proto.RegisterType((*MetricMetadata)(nil), "conduit.public.MetricMetadata")
	proto.RegisterType((*MetricResponse)(nil), "conduit.public.MetricResponse")
	proto.RegisterType((*MetricRequest)(nil), "conduit.public.MetricRequest")
	proto.RegisterType((*Empty)(nil), "conduit.public.Empty")
	proto.RegisterType((*VersionInfo)(nil), "conduit.public.VersionInfo")
	proto.RegisterType((*ListPodsResponse)(nil), "conduit.public.ListPodsResponse")
	proto.RegisterType((*Pod)(nil), "conduit.public.Pod")
	proto.RegisterType((*TapRequest)(nil), "conduit.public.TapRequest")
	proto.RegisterEnum("conduit.public.MetricName", MetricName_name, MetricName_value)
	proto.RegisterEnum("conduit.public.TimeWindow", TimeWindow_name, TimeWindow_value)
	proto.RegisterEnum("conduit.public.AggregationType", AggregationType_name, AggregationType_value)
	proto.RegisterEnum("conduit.public.HistogramLabel", HistogramLabel_name, HistogramLabel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	Stat(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*MetricResponse, error)
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionInfo, error)
	ListPods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPodsResponse, error)
	Tap(ctx context.Context, in *TapRequest, opts ...grpc.CallOption) (Api_TapClient, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Stat(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*MetricResponse, error) {
	out := new(MetricResponse)
	err := grpc.Invoke(ctx, "/conduit.public.Api/Stat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionInfo, error) {
	out := new(VersionInfo)
	err := grpc.Invoke(ctx, "/conduit.public.Api/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListPods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := grpc.Invoke(ctx, "/conduit.public.Api/ListPods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Tap(ctx context.Context, in *TapRequest, opts ...grpc.CallOption) (Api_TapClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Api_serviceDesc.Streams[0], c.cc, "/conduit.public.Api/Tap", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiTapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Api_TapClient interface {
	Recv() (*conduit_common.TapEvent, error)
	grpc.ClientStream
}

type apiTapClient struct {
	grpc.ClientStream
}

func (x *apiTapClient) Recv() (*conduit_common.TapEvent, error) {
	m := new(conduit_common.TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Api service

type ApiServer interface {
	Stat(context.Context, *MetricRequest) (*MetricResponse, error)
	Version(context.Context, *Empty) (*VersionInfo, error)
	ListPods(context.Context, *Empty) (*ListPodsResponse, error)
	Tap(*TapRequest, Api_TapServer) error
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.public.Api/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Stat(ctx, req.(*MetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.public.Api/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.public.Api/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListPods(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Tap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServer).Tap(m, &apiTapServer{stream})
}

type Api_TapServer interface {
	Send(*conduit_common.TapEvent) error
	grpc.ServerStream
}

type apiTapServer struct {
	grpc.ServerStream
}

func (x *apiTapServer) Send(m *conduit_common.TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conduit.public.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stat",
			Handler:    _Api_Stat_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Api_Version_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _Api_ListPods_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tap",
			Handler:       _Api_Tap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "public/api.proto",
}

func init() { proto.RegisterFile("public/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x16, 0x45, 0x59, 0x97, 0x23, 0x5b, 0xe1, 0x3f, 0xc9, 0x1f, 0xb0, 0x6a, 0xea, 0xaa, 0x5c,
	0xb4, 0x86, 0x17, 0x72, 0xea, 0x26, 0x01, 0xdc, 0x22, 0x08, 0x64, 0x89, 0x88, 0x0c, 0xf8, 0xa2,
	0x8e, 0xe4, 0xb4, 0x05, 0x0a, 0x18, 0x63, 0x71, 0x4c, 0x13, 0x20, 0x39, 0x0c, 0x39, 0x4c, 0xea,
	0xee, 0xbb, 0xed, 0xb6, 0x7d, 0x82, 0x3e, 0x40, 0xb7, 0x7d, 0xb9, 0x62, 0x2e, 0xa4, 0x2e, 0x91,
	0x83, 0xae, 0x3c, 0xe7, 0x3b, 0xdf, 0x7c, 0x3c, 0x73, 0x6e, 0x16, 0x58, 0x49, 0x7e, 0x1d, 0x06,
	0xf3, 0x03, 0x92, 0x04, 0xfd, 0x24, 0x65, 0x9c, 0xa1, 0xce, 0x9c, 0xc5, 0x5e, 0x1e, 0xf0, 0xbe,
	0xf2, 0x74, 0x77, 0x7d, 0xc6, 0xfc, 0x90, 0x1e, 0x48, 0xef, 0x75, 0x7e, 0x73, 0xe0, 0xe5, 0x29,
	0xe1, 0x01, 0x8b, 0x15, 0xbf, 0xfb, 0x70, 0xce, 0xa2, 0x88, 0xc5, 0x07, 0xea, 0x8f, 0x02, 0x9d,
	0x9f, 0xa1, 0x33, 0x0e, 0x32, 0xce, 0xfc, 0x94, 0x44, 0x6f, 0x48, 0x98, 0x53, 0xf4, 0x0c, 0xb6,
	0x42, 0x72, 0x4d, 0x43, 0xdb, 0xe8, 0x19, 0x7b, 0x9d, 0xc3, 0xdd, 0xfe, 0xea, 0x67, 0xfa, 0x25,
	0xfd, 0x54, 0xb0, 0xb0, 0x22, 0xa3, 0x47, 0xb0, 0xf5, 0x4e, 0x5c, 0xb7, 0xab, 0x3d, 0x63, 0xcf,
	0xc4, 0xca, 0x70, 0x86, 0xd0, 0x2a, 0xe9, 0xe8, 0x05, 0xd4, 0x25, 0x9a, 0xd9, 0x46, 0xcf, 0xdc,
	0x6b, 0x7f, 0x44, 0x59, 0x06, 0x82, 0x35, 0xdb, 0xf9, 0xcd, 0x80, 0xf6, 0x19, 0xe5, 0x69, 0x30,
	0x57, 0x01, 0x76, 0xa1, 0x31, 0x67, 0x79, 0xcc, 0x69, 0x2a, 0x43, 0x34, 0xc7, 0x15, 0x5c, 0x00,
	0xe8, 0x31, 0x6c, 0xf9, 0x24, 0xf7, 0x55, 0x18, 0xc6, 0xb8, 0x82, 0x95, 0x89, 0x8e, 0xa0, 0x75,
	0x5b, 0xa8, 0xdb, 0x66, 0xcf, 0xd8, 0x6b, 0x1f, 0x7e, 0x72, 0xef, 0xe7, 0xc7, 0x15, 0xbc, 0x60,
	0x1f, 0x37, 0xf4, 0xcb, 0x1c, 0x1f, 0x1e, 0xa8, 0x30, 0x46, 0x84, 0x93, 0x84, 0x05, 0x31, 0x47,
	0x5f, 0x17, 0xaf, 0x36, 0xa4, 0xe4, 0xa7, 0xeb, 0x92, 0x4b, 0x61, 0xeb, 0x94, 0xa0, 0x2f, 0x60,
	0x9b, 0x07, 0x11, 0xcd, 0x38, 0x89, 0x92, 0xab, 0x28, 0xd3, 0xf9, 0x6a, 0x97, 0xd8, 0x59, 0xe6,
	0xfc, 0x63, 0xc0, 0xb6, 0xba, 0x39, 0xa5, 0x69, 0x40, 0x33, 0xd4, 0x87, 0x5a, 0x4c, 0x22, 0xaa,
	0x2b, 0xd2, 0xdd, 0xfc, 0x95, 0x73, 0x12, 0x51, 0x2c, 0x79, 0xe8, 0x5b, 0x68, 0x46, 0x94, 0x13,
	0x8f, 0x70, 0x22, 0xf5, 0x37, 0xe4, 0x5a, 0xdd, 0x39, 0xd3, 0x2c, 0x5c, 0xf2, 0xd1, 0x2b, 0x00,
	0xaf, 0x78, 0x5f, 0x66, 0x9b, 0xb2, 0x52, 0x9f, 0x6f, 0xbe, 0x5d, 0xe6, 0x01, 0x2f, 0x5d, 0x71,
	0xfe, 0x36, 0xa0, 0xb3, 0xaa, 0x8e, 0x9e, 0x40, 0x8b, 0x93, 0xd4, 0xa7, 0x7c, 0xc2, 0x3c, 0xf9,
	0x88, 0x16, 0x5e, 0x00, 0xc8, 0x81, 0x6d, 0x65, 0x8c, 0x68, 0x12, 0xb2, 0x3b, 0x19, 0x71, 0x0b,
	0xaf, 0x60, 0x42, 0x21, 0x63, 0x79, 0x3a, 0xa7, 0x42, 0xc1, 0x54, 0x0a, 0x25, 0x20, 0x14, 0x94,
	0xa1, 0x15, 0x6a, 0x4a, 0x61, 0x19, 0x13, 0x0a, 0x73, 0x16, 0x25, 0x2c, 0xa6, 0x31, 0xb7, 0xb7,
	0x94, 0x42, 0x09, 0x38, 0xe3, 0x22, 0x66, 0x4c, 0xb3, 0x84, 0xc5, 0x19, 0x45, 0x2f, 0xa0, 0x11,
	0x49, 0xa4, 0x68, 0xd7, 0x27, 0x9b, 0x93, 0xa0, 0x4a, 0x84, 0x0b, 0xb2, 0xf3, 0x7b, 0x15, 0x76,
	0x0a, 0xa9, 0xb7, 0x39, 0xcd, 0x38, 0x7a, 0xb6, 0xaa, 0xf4, 0xf1, 0x02, 0x16, 0x54, 0x74, 0x08,
	0xf5, 0xf7, 0x41, 0xec, 0xb1, 0xf7, 0x32, 0x1f, 0x1b, 0x2e, 0xcd, 0x82, 0x88, 0xfe, 0x20, 0x19,
	0x58, 0x33, 0xd1, 0x11, 0x34, 0xfc, 0x94, 0xe5, 0xc9, 0xf1, 0x9d, 0xcc, 0x51, 0xe7, 0xc3, 0xc2,
	0x0d, 0x7c, 0x3f, 0xa5, 0xbe, 0xdc, 0x0a, 0xb3, 0xbb, 0x84, 0xe2, 0x82, 0x2f, 0x5a, 0xe6, 0x26,
	0x08, 0x39, 0x4d, 0x8f, 0x55, 0xfa, 0xfe, 0x43, 0xcb, 0x14, 0x7c, 0x59, 0x9c, 0x3c, 0x8a, 0x48,
	0x1a, 0xfc, 0x4a, 0x65, 0x6a, 0x9b, 0x78, 0x01, 0x38, 0x0d, 0xd8, 0x72, 0xa3, 0x84, 0xdf, 0x39,
	0x6f, 0xa1, 0xfd, 0x86, 0xa6, 0x59, 0xc0, 0xe2, 0x93, 0xf8, 0x86, 0x89, 0x5b, 0x3e, 0xd3, 0x40,
	0xd1, 0x14, 0x25, 0x20, 0xbc, 0xd7, 0x79, 0x10, 0x7a, 0x23, 0xc2, 0xa9, 0xee, 0x88, 0x05, 0x80,
	0xbe, 0x84, 0x4e, 0x4a, 0x43, 0x4a, 0x32, 0x5a, 0x08, 0xa8, 0x9e, 0x58, 0x43, 0x9d, 0xef, 0xc0,
	0x3a, 0x0d, 0x32, 0xd1, 0x65, 0x59, 0x59, 0xd8, 0xaf, 0xa0, 0x96, 0x30, 0xaf, 0xa8, 0xea, 0xc3,
	0xf5, 0x57, 0x4e, 0x98, 0x87, 0x25, 0xc1, 0xf9, 0xb3, 0x0a, 0xa6, 0xe8, 0x2e, 0xb4, 0x34, 0x7d,
	0x2d, 0x3d, 0x61, 0x8f, 0x60, 0x2b, 0x61, 0xde, 0xc9, 0x44, 0x87, 0xa6, 0x0c, 0xb4, 0x0b, 0xe0,
	0xc9, 0x6e, 0x8b, 0x44, 0x93, 0xa9, 0x90, 0x96, 0x10, 0xf4, 0x18, 0xea, 0x19, 0x27, 0x3c, 0xcf,
	0x74, 0x87, 0x6a, 0x4b, 0xa8, 0x11, 0xcf, 0xa3, 0x9e, 0x4e, 0x9e, 0x32, 0xd0, 0x10, 0x1e, 0x64,
	0x41, 0x3c, 0xa7, 0xa7, 0x24, 0xe3, 0x98, 0x26, 0x2c, 0xe5, 0x76, 0x5d, 0x6f, 0x2e, 0xb5, 0xe9,
	0xfb, 0xc5, 0xa6, 0xef, 0x8f, 0xf4, 0xa6, 0xc7, 0xeb, 0x37, 0xd0, 0x53, 0x78, 0x38, 0x67, 0x31,
	0x4f, 0x59, 0x18, 0xd2, 0x54, 0x74, 0x58, 0x96, 0x90, 0x39, 0xb5, 0x1b, 0xf2, 0xfb, 0x9b, 0x5c,
	0x62, 0x98, 0x34, 0x3c, 0x09, 0x49, 0x4c, 0xed, 0xa6, 0x8c, 0x69, 0x05, 0x73, 0xfe, 0xaa, 0x02,
	0xcc, 0x48, 0x52, 0x74, 0x38, 0x02, 0x33, 0x29, 0x26, 0x7b, 0x5c, 0xc1, 0xc2, 0x40, 0xbd, 0x95,
	0x5c, 0x54, 0xb5, 0x6b, 0x2d, 0x1b, 0x11, 0xf9, 0x05, 0x27, 0x99, 0xcc, 0x54, 0x15, 0x6b, 0x4b,
	0xe0, 0x9c, 0x4d, 0xc4, 0x73, 0x45, 0x96, 0x76, 0xb0, 0xb6, 0x44, 0x1d, 0x38, 0x3b, 0x99, 0xe8,
	0xe1, 0x95, 0x67, 0xd4, 0x85, 0xe6, 0x4d, 0xca, 0xa2, 0x49, 0x91, 0x9c, 0x1d, 0x5c, 0xda, 0x42,
	0x47, 0x9c, 0x4f, 0x26, 0xfa, 0xb5, 0xda, 0x92, 0x55, 0x98, 0xdf, 0xd2, 0x48, 0x3d, 0x4d, 0x54,
	0x41, 0x5a, 0x32, 0x1e, 0xca, 0x6f, 0x99, 0x67, 0xb7, 0x14, 0xae, 0x2c, 0xd1, 0x8a, 0x24, 0xe7,
	0xb7, 0x2c, 0x0d, 0xf8, 0x9d, 0x0d, 0xaa, 0x15, 0x4b, 0x40, 0x44, 0x95, 0x10, 0x7e, 0x6b, 0xb7,
	0x55, 0x54, 0xe2, 0x7c, 0xdc, 0x84, 0xba, 0xda, 0x5e, 0xfb, 0x2f, 0x01, 0x16, 0xc3, 0x8d, 0x2c,
	0xd8, 0xc6, 0xee, 0xf7, 0x97, 0xee, 0x74, 0x76, 0x85, 0x07, 0x33, 0xd7, 0xaa, 0xa0, 0x36, 0x34,
	0x4e, 0x07, 0x33, 0xf7, 0x7c, 0xf8, 0x93, 0x65, 0x08, 0xf7, 0xf4, 0x72, 0x38, 0x74, 0xa7, 0x53,
	0xe5, 0xae, 0xee, 0x0f, 0x00, 0x16, 0x63, 0x2e, 0xc8, 0x33, 0xf7, 0xfc, 0x6a, 0xea, 0x0e, 0xd5,
	0xcd, 0x8b, 0x73, 0xf7, 0xea, 0xec, 0xe4, 0xdc, 0x32, 0x0a, 0x8f, 0x30, 0xaa, 0x68, 0x1b, 0x9a,
	0xc2, 0x33, 0xbe, 0xb8, 0xc4, 0x96, 0xb9, 0x4f, 0xe0, 0xc1, 0xda, 0xd0, 0xa3, 0x0e, 0xc0, 0x6c,
	0x80, 0x5f, 0xbb, 0xb3, 0xab, 0xc9, 0xc5, 0xc8, 0xaa, 0xa0, 0xff, 0xc1, 0x8e, 0xb6, 0x47, 0xee,
	0xe4, 0xf4, 0x42, 0x84, 0xd2, 0x01, 0x98, 0x5e, 0x5c, 0xe2, 0xa1, 0x2b, 0x29, 0x55, 0x41, 0xd1,
	0xb6, 0xa6, 0x98, 0xa8, 0x09, 0xb5, 0x33, 0x77, 0x3a, 0xb6, 0x6a, 0xfb, 0x2f, 0x97, 0x7e, 0x43,
	0xc8, 0x1f, 0x05, 0xa8, 0x01, 0xa6, 0x88, 0xa5, 0x22, 0x0e, 0x93, 0xe7, 0x4f, 0x2d, 0x43, 0x1e,
	0x8e, 0x9e, 0x5b, 0x55, 0x75, 0x38, 0xb2, 0x4c, 0xc9, 0x19, 0xfc, 0x68, 0xd5, 0x0e, 0xff, 0xa8,
	0x82, 0x39, 0x48, 0x02, 0xf4, 0x1a, 0x6a, 0x53, 0x4e, 0x38, 0xfa, 0x6c, 0xf3, 0xe2, 0xd1, 0xcd,
	0xd6, 0xdd, 0xbd, 0xcf, 0xad, 0xe6, 0xdb, 0xa9, 0xa0, 0x57, 0xd0, 0x28, 0xd6, 0xc8, 0xff, 0xd7,
	0xc9, 0x72, 0x15, 0x75, 0x3f, 0xf8, 0x47, 0xbd, 0xb4, 0x98, 0x9c, 0x0a, 0x72, 0xa1, 0x59, 0xac,
	0x8d, 0xfb, 0x14, 0x7a, 0xeb, 0xf0, 0xfa, 0x9e, 0x91, 0x71, 0x98, 0x33, 0x92, 0xa0, 0x0f, 0x37,
	0x77, 0x39, 0x39, 0x5d, 0xbb, 0xf4, 0xe9, 0x5f, 0x65, 0x33, 0x92, 0xb8, 0xef, 0xc4, 0x7f, 0xa4,
	0xca, 0x53, 0xe3, 0xba, 0x2e, 0x07, 0xfc, 0x9b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xae, 0xce,
	0x2d, 0x4b, 0xfc, 0x09, 0x00, 0x00,
}
