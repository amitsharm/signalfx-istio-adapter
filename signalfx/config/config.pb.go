// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/signalfx/config/config.proto

/*
	Package config is a generated protocol buffer package.

	The `signalfx` adapter collects Istio metrics and trace spans and sends them
	to [SignalFx](https://signalfx.com).

	This adapter supports the [metric template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/metric/)
	and the [tracespan template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/tracespan/).

	If sending trace spans, this adapter can make use of certain conventions in
	the tracespan format that is configured to send to this adapter.  Here is an
	example tracespan spec that will work well:

	```yaml
	apiVersion: config.istio.io/v1alpha2
	kind: tracespan
	metadata:
	  name: signalfx
	spec:
	  traceId: request.headers["x-b3-traceid"] | ""
	  spanId: request.headers["x-b3-spanid"] | ""
	  parentSpanId: request.headers["x-b3-parentspanid"] | ""
	  # If the path contains query parameters, they will be split off and put into
	  # tags such that the span name sent to SignalFx will consist only of the path
	  # itself.
	  spanName: request.path | "/"
	  startTime: request.time
	  endTime: response.time
	  # If this is >=500, the span will get an 'error' tag
	  httpStatusCode: response.code | 0
	  clientSpan: context.reporter.kind == "outbound"
	  # Span tags below that do not have comments are useful but optional and will
	  # be passed to SignalFx unmodified. The tags that have comments are interpreted
	  # in a special manner, but are still optional.
	  spanTags:
	    # This gets put into the remoteEndpoint.ipv4 field
	    destination.ip: destination.ip | ip("0.0.0.0")
	    # This gets put into the remoteEndpoint.name field
	    destination.name: destination.name | "unknown"
	    destination.namespace: destination.namespace | "unknown"
	    request.host: request.host | ""
	    request.method: request.method | ""
	    request.path: request.path | ""
	    request.size: request.size | 0
	    request.useragent: request.useragent | ""
	    response.size: response.size | 0
	    # This gets put into the localEndpoint.name field
	    source.name: source.name | "unknown"
	    # This gets put into the localEndpoint.ipv4 field
	    source.ip: source.ip | ip("0.0.0.0")
	    source.namespace: source.namespace | "unknown"
	    source.version: source.labels["version"] | "unknown"
	 ```

	It is generated from these files:
		mixer/adapter/signalfx/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import strconv "strconv"

import types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Describes what kind of metric this is.
type Params_MetricConfig_Type int32

const (
	// None is the default and is invalid
	NONE Params_MetricConfig_Type = 0
	// Values with the same set of dimensions will be added together
	// as a continuously incrementing value.
	COUNTER Params_MetricConfig_Type = 1
	// A histogram distribution.  This will result in several metrics
	// emitted for each unique set of dimensions.
	HISTOGRAM Params_MetricConfig_Type = 2
)

var Params_MetricConfig_Type_name = map[int32]string{
	0: "NONE",
	1: "COUNTER",
	2: "HISTOGRAM",
}
var Params_MetricConfig_Type_value = map[string]int32{
	"NONE":      0,
	"COUNTER":   1,
	"HISTOGRAM": 2,
}

func (Params_MetricConfig_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{0, 0, 0}
}

// Configuration format for the `signalfx` adapter.
type Params struct {
	// Required. The set of metrics to send to SignalFx. If an Istio metric is
	// configured to be sent to this adapter, it must have a corresponding
	// description here.
	Metrics []*Params_MetricConfig `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// Optional. The URL of the SignalFx ingest server to use.  Will default to
	// the global ingest server if not specified.
	IngestUrl string `protobuf:"bytes,2,opt,name=ingest_url,json=ingestUrl,proto3" json:"ingest_url,omitempty"`
	// Required. The access token for the SignalFx organization that should
	// receive the metrics.  This can also be configured via an environment
	// variable `SIGNALFX_ACCESS_TOKEN` set on the adapter process which makes
	// it possible to use Kubernetes secrets to provide the token.  This field,
	// if specified, will take priority over the environment variable.
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Optional. Specifies how frequently to send metrics to SignalFx.  Metrics
	// reported to this adapter are collected and reported as a timeseries.
	// This will be rounded to the nearest second and rounded values less than
	// one second are not valid. Defaults to 10 seconds if not specified.
	DatapointInterval time.Duration `protobuf:"bytes,4,opt,name=datapoint_interval,json=datapointInterval,stdduration" json:"datapoint_interval"`
	// Optional.  If set to false, metrics won't be sent (but trace spans will
	// be sent, unless otherwise disabled).
	EnableMetrics bool `protobuf:"varint,5,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`
	// Optional.  If set to false, trace spans won't be sent (but metrics will
	// be sent, unless otherwise disabled).
	EnableTracing bool `protobuf:"varint,6,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`
	// Configuration for the Trace Span handler
	Tracing *Params_TracingConfig `protobuf:"bytes,7,opt,name=tracing" json:"tracing,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

// Describes what metrics should be sent to SignalFx and in what form.
type Params_MetricConfig struct {
	// Required.  The name of the metric as it is sent to the adapter.  In
	// Kubernetes this is of the form "<name>.metric.<namespace>" where
	// "<name>" is the name field of the metric resource, and "<namespace>"
	// is the namespace of the metric resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The metric type of the metric
	Type Params_MetricConfig_Type `protobuf:"varint,4,opt,name=type,proto3,enum=signalfx.config.Params_MetricConfig_Type" json:"type,omitempty"`
}

func (m *Params_MetricConfig) Reset()                    { *m = Params_MetricConfig{} }
func (*Params_MetricConfig) ProtoMessage()               {}
func (*Params_MetricConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

// Holds all of the tracing-specific configuration
type Params_TracingConfig struct {
	// Optional.  The number of trace spans that the adapter will buffer before
	// dropping them.  This defaults to 1000 spans but can be configured higher
	// if needed.  An error message will be logged if spans are dropped.
	BufferSize uint32 `protobuf:"varint,1,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	// The span tag that is used as the value of the localEndpoint.serviceName
	// field of the span sent to SignalFx
	LocalEndpointNameTagKey string `protobuf:"bytes,3,opt,name=local_endpoint_name_tag_key,json=localEndpointNameTagKey,proto3" json:"local_endpoint_name_tag_key,omitempty"`
	// The span tag that is used as the value of the localEndpoint.ipv4
	// field of the span sent to SignalFx
	LocalEndpointIpTagKey string `protobuf:"bytes,4,opt,name=local_endpoint_ip_tag_key,json=localEndpointIpTagKey,proto3" json:"local_endpoint_ip_tag_key,omitempty"`
	// The span tag that is used as the value of the remoteEndpoint.serviceName
	// field of the span sent to SignalFx
	RemoteEndpointNameTagKey string `protobuf:"bytes,5,opt,name=remote_endpoint_name_tag_key,json=remoteEndpointNameTagKey,proto3" json:"remote_endpoint_name_tag_key,omitempty"`
	// The span tag that is used as the value of the remoteEndpoint.ipv4
	// field of the span sent to SignalFx
	RemoteEndpointIpTagKey string `protobuf:"bytes,6,opt,name=remote_endpoint_ip_tag_key,json=remoteEndpointIpTagKey,proto3" json:"remote_endpoint_ip_tag_key,omitempty"`
	// If true, the local and remote endpoints will be swapped for
	// non-client spans.  This means that the above config options for
	// [local/remote]_endpoint_[name/ip]_tag_key with have a reversed
	// meaning for server spans.  The `clientSpan` field in the `tracespan`
	// instance that is used with this adapter determines what is a
	// "client" vs a "server" span.
	SwapLocalRemoteEndpoints bool `protobuf:"varint,7,opt,name=swap_local_remote_endpoints,json=swapLocalRemoteEndpoints,proto3" json:"swap_local_remote_endpoints,omitempty"`
}

func (m *Params_TracingConfig) Reset()                    { *m = Params_TracingConfig{} }
func (*Params_TracingConfig) ProtoMessage()               {}
func (*Params_TracingConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

func init() {
	proto.RegisterType((*Params)(nil), "signalfx.config.Params")
	proto.RegisterType((*Params_MetricConfig)(nil), "signalfx.config.Params.MetricConfig")
	proto.RegisterType((*Params_TracingConfig)(nil), "signalfx.config.Params.TracingConfig")
	proto.RegisterEnum("signalfx.config.Params_MetricConfig_Type", Params_MetricConfig_Type_name, Params_MetricConfig_Type_value)
}
func (x Params_MetricConfig_Type) String() string {
	s, ok := Params_MetricConfig_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, msg := range m.Metrics {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.IngestUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IngestUrl)))
		i += copy(dAtA[i:], m.IngestUrl)
	}
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintConfig(dAtA, i, uint64(types.SizeOfStdDuration(m.DatapointInterval)))
	n1, err := types.StdDurationMarshalTo(m.DatapointInterval, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.EnableMetrics {
		dAtA[i] = 0x28
		i++
		if m.EnableMetrics {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.EnableTracing {
		dAtA[i] = 0x30
		i++
		if m.EnableTracing {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Tracing != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Tracing.Size()))
		n2, err := m.Tracing.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Params_MetricConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_MetricConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Type != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func (m *Params_TracingConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_TracingConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BufferSize != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.BufferSize))
	}
	if len(m.LocalEndpointNameTagKey) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.LocalEndpointNameTagKey)))
		i += copy(dAtA[i:], m.LocalEndpointNameTagKey)
	}
	if len(m.LocalEndpointIpTagKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.LocalEndpointIpTagKey)))
		i += copy(dAtA[i:], m.LocalEndpointIpTagKey)
	}
	if len(m.RemoteEndpointNameTagKey) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RemoteEndpointNameTagKey)))
		i += copy(dAtA[i:], m.RemoteEndpointNameTagKey)
	}
	if len(m.RemoteEndpointIpTagKey) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RemoteEndpointIpTagKey)))
		i += copy(dAtA[i:], m.RemoteEndpointIpTagKey)
	}
	if m.SwapLocalRemoteEndpoints {
		dAtA[i] = 0x38
		i++
		if m.SwapLocalRemoteEndpoints {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, e := range m.Metrics {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.IngestUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = types.SizeOfStdDuration(m.DatapointInterval)
	n += 1 + l + sovConfig(uint64(l))
	if m.EnableMetrics {
		n += 2
	}
	if m.EnableTracing {
		n += 2
	}
	if m.Tracing != nil {
		l = m.Tracing.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *Params_MetricConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovConfig(uint64(m.Type))
	}
	return n
}

func (m *Params_TracingConfig) Size() (n int) {
	var l int
	_ = l
	if m.BufferSize != 0 {
		n += 1 + sovConfig(uint64(m.BufferSize))
	}
	l = len(m.LocalEndpointNameTagKey)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.LocalEndpointIpTagKey)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.RemoteEndpointNameTagKey)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.RemoteEndpointIpTagKey)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.SwapLocalRemoteEndpoints {
		n += 2
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Metrics:` + strings.Replace(fmt.Sprintf("%v", this.Metrics), "Params_MetricConfig", "Params_MetricConfig", 1) + `,`,
		`IngestUrl:` + fmt.Sprintf("%v", this.IngestUrl) + `,`,
		`AccessToken:` + fmt.Sprintf("%v", this.AccessToken) + `,`,
		`DatapointInterval:` + strings.Replace(strings.Replace(this.DatapointInterval.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`EnableMetrics:` + fmt.Sprintf("%v", this.EnableMetrics) + `,`,
		`EnableTracing:` + fmt.Sprintf("%v", this.EnableTracing) + `,`,
		`Tracing:` + strings.Replace(fmt.Sprintf("%v", this.Tracing), "Params_TracingConfig", "Params_TracingConfig", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_MetricConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_MetricConfig{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_TracingConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_TracingConfig{`,
		`BufferSize:` + fmt.Sprintf("%v", this.BufferSize) + `,`,
		`LocalEndpointNameTagKey:` + fmt.Sprintf("%v", this.LocalEndpointNameTagKey) + `,`,
		`LocalEndpointIpTagKey:` + fmt.Sprintf("%v", this.LocalEndpointIpTagKey) + `,`,
		`RemoteEndpointNameTagKey:` + fmt.Sprintf("%v", this.RemoteEndpointNameTagKey) + `,`,
		`RemoteEndpointIpTagKey:` + fmt.Sprintf("%v", this.RemoteEndpointIpTagKey) + `,`,
		`SwapLocalRemoteEndpoints:` + fmt.Sprintf("%v", this.SwapLocalRemoteEndpoints) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, &Params_MetricConfig{})
			if err := m.Metrics[len(m.Metrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngestUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IngestUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatapointInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdDurationUnmarshal(&m.DatapointInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableMetrics", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableMetrics = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableTracing", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableTracing = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tracing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tracing == nil {
				m.Tracing = &Params_TracingConfig{}
			}
			if err := m.Tracing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_MetricConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Params_MetricConfig_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_TracingConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TracingConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TracingConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferSize", wireType)
			}
			m.BufferSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BufferSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalEndpointNameTagKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalEndpointNameTagKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalEndpointIpTagKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalEndpointIpTagKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteEndpointNameTagKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteEndpointNameTagKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteEndpointIpTagKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteEndpointIpTagKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapLocalRemoteEndpoints", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SwapLocalRemoteEndpoints = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/signalfx/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4f, 0x13, 0x41,
	0x18, 0xc7, 0x77, 0xa0, 0xb4, 0xe5, 0x29, 0x45, 0x9c, 0xf8, 0xb2, 0x14, 0x1d, 0x2a, 0x91, 0xa4,
	0x1a, 0xb3, 0x4d, 0xf0, 0x42, 0x8c, 0x60, 0x04, 0x89, 0x12, 0xa5, 0x98, 0xa5, 0x5c, 0xbc, 0x6c,
	0xa6, 0xed, 0x74, 0x33, 0x61, 0xbb, 0xb3, 0x99, 0x9d, 0x2a, 0xe5, 0xe4, 0x37, 0xd0, 0xa3, 0x1f,
	0xc1, 0x8f, 0xc2, 0x91, 0xa3, 0x27, 0xb5, 0x4b, 0x4c, 0x3c, 0xf2, 0x11, 0x4c, 0x67, 0xb6, 0xd8,
	0x36, 0x92, 0x78, 0xda, 0x9d, 0xff, 0xf3, 0xff, 0xcd, 0xf3, 0x36, 0xf0, 0xb0, 0xc3, 0x8f, 0x99,
	0xac, 0xd2, 0x16, 0x8d, 0x14, 0x93, 0xd5, 0x98, 0xfb, 0x21, 0x0d, 0xda, 0xc7, 0xd5, 0xa6, 0x08,
	0xdb, 0xdc, 0x4f, 0x3f, 0x4e, 0x24, 0x85, 0x12, 0xf8, 0xda, 0x30, 0xea, 0x18, 0xb9, 0x74, 0xc3,
	0x17, 0xbe, 0xd0, 0xb1, 0xea, 0xe0, 0xcf, 0xd8, 0x4a, 0xc4, 0x17, 0xc2, 0x0f, 0x58, 0x55, 0x9f,
	0x1a, 0xdd, 0x76, 0xb5, 0xd5, 0x95, 0x54, 0x71, 0x11, 0x9a, 0xf8, 0xca, 0xaf, 0x2c, 0x64, 0xdf,
	0x52, 0x49, 0x3b, 0x31, 0xde, 0x84, 0x5c, 0x87, 0x29, 0xc9, 0x9b, 0xb1, 0x8d, 0xca, 0xd3, 0x95,
	0xc2, 0xda, 0x7d, 0x67, 0x22, 0x87, 0x63, 0x9c, 0xce, 0x9e, 0xb6, 0x6d, 0x6b, 0xcd, 0x1d, 0x42,
	0xf8, 0x2e, 0x00, 0x0f, 0x7d, 0x16, 0x2b, 0xaf, 0x2b, 0x03, 0x7b, 0xaa, 0x8c, 0x2a, 0xb3, 0xee,
	0xac, 0x51, 0x0e, 0x65, 0x80, 0xef, 0xc1, 0x1c, 0x6d, 0x36, 0x59, 0x1c, 0x7b, 0x4a, 0x1c, 0xb1,
	0xd0, 0x9e, 0xd6, 0x86, 0x82, 0xd1, 0xea, 0x03, 0x09, 0xbb, 0x80, 0x5b, 0x54, 0xd1, 0x48, 0xf0,
	0x50, 0x79, 0x3c, 0x54, 0x4c, 0xbe, 0xa7, 0x81, 0x9d, 0x29, 0xa3, 0x4a, 0x61, 0x6d, 0xd1, 0x31,
	0x9d, 0x38, 0xc3, 0x4e, 0x9c, 0x17, 0x69, 0x27, 0x5b, 0xf9, 0xd3, 0xef, 0xcb, 0xd6, 0x97, 0x1f,
	0xcb, 0xc8, 0xbd, 0x7e, 0x89, 0xef, 0xa6, 0x34, 0x5e, 0x85, 0x79, 0x16, 0xd2, 0x46, 0xc0, 0xbc,
	0x61, 0x73, 0x33, 0x65, 0x54, 0xc9, 0xbb, 0x45, 0xa3, 0xee, 0xa5, 0xc5, 0xff, 0xb5, 0x29, 0x49,
	0x9b, 0x3c, 0xf4, 0xed, 0xec, 0xa8, 0xad, 0x6e, 0x44, 0xfc, 0x0c, 0x72, 0xc3, 0x78, 0x4e, 0x97,
	0xb5, 0x7a, 0xd5, 0x8c, 0x52, 0x62, 0x38, 0xa4, 0x94, 0x2a, 0x7d, 0x42, 0x30, 0x37, 0x3a, 0x3e,
	0x8c, 0x21, 0x13, 0xd2, 0x0e, 0xb3, 0x91, 0x1e, 0x87, 0xfe, 0xc7, 0x1b, 0x90, 0x51, 0xbd, 0x88,
	0xe9, 0xce, 0xe7, 0xd7, 0x1e, 0xfc, 0xcf, 0x1a, 0x9c, 0x7a, 0x2f, 0x62, 0xae, 0xc6, 0x56, 0x1e,
	0x41, 0x66, 0x70, 0xc2, 0x79, 0xc8, 0xd4, 0xf6, 0x6b, 0x3b, 0x0b, 0x16, 0x2e, 0x40, 0x6e, 0x7b,
	0xff, 0xb0, 0x56, 0xdf, 0x71, 0x17, 0x10, 0x2e, 0xc2, 0xec, 0xab, 0xdd, 0x83, 0xfa, 0xfe, 0x4b,
	0xf7, 0xf9, 0xde, 0xc2, 0x54, 0xe9, 0x7c, 0x0a, 0x8a, 0x63, 0xc5, 0xe2, 0x65, 0x28, 0x34, 0xba,
	0xed, 0x36, 0x93, 0x5e, 0xcc, 0x4f, 0x4c, 0x65, 0x45, 0x17, 0x8c, 0x74, 0xc0, 0x4f, 0x18, 0x7e,
	0x0a, 0x4b, 0x81, 0x68, 0xd2, 0xc0, 0x63, 0x61, 0xcb, 0x2c, 0x6b, 0x50, 0xb6, 0xa7, 0xa8, 0xef,
	0x1d, 0xb1, 0x5e, 0xba, 0xd9, 0xdb, 0xda, 0xb2, 0x93, 0x3a, 0x6a, 0xb4, 0xc3, 0xea, 0xd4, 0x7f,
	0xcd, 0x7a, 0x78, 0x1d, 0x16, 0x27, 0x68, 0x1e, 0x5d, 0xb2, 0x19, 0xcd, 0xde, 0x1c, 0x63, 0x77,
	0xa3, 0x94, 0xdc, 0x84, 0x3b, 0x92, 0x75, 0x84, 0x62, 0x57, 0x24, 0x9e, 0xd1, 0xb0, 0x6d, 0x3c,
	0xff, 0xc8, 0xfc, 0x04, 0x4a, 0x93, 0xfc, 0x48, 0xea, 0xac, 0xa6, 0x6f, 0x8d, 0xd3, 0x97, 0xb9,
	0x37, 0x60, 0x29, 0xfe, 0x40, 0x23, 0xcf, 0x94, 0x3e, 0x71, 0x4d, 0xac, 0x5f, 0x43, 0xde, 0xb5,
	0x07, 0x96, 0x37, 0x03, 0x87, 0x3b, 0x76, 0x4b, 0xbc, 0xb5, 0x7e, 0xda, 0x27, 0xd6, 0x59, 0x9f,
	0x58, 0xdf, 0xfa, 0xc4, 0xba, 0xe8, 0x13, 0xeb, 0x63, 0x42, 0xd0, 0xd7, 0x84, 0x58, 0xa7, 0x09,
	0x41, 0x67, 0x09, 0x41, 0x3f, 0x13, 0x82, 0x7e, 0x27, 0xc4, 0xba, 0x48, 0x08, 0xfa, 0x7c, 0x4e,
	0xac, 0x77, 0x59, 0xb3, 0xec, 0x46, 0x56, 0x3f, 0xf8, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x56, 0x60, 0x06, 0x7e, 0x1d, 0x04, 0x00, 0x00,
}
