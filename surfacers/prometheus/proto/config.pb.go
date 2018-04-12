// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/prometheus/proto/config.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/surfacers/prometheus/proto/config.proto

It has these top-level messages:
	SurfacerConf
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SurfacerConf struct {
	// How many metrics entries (EventMetrics) to buffer. Incoming metrics
	// processing is paused while serving data to prometheus. This buffer is to
	// make writes to prometheus surfacer non-blocking.
	MetricsBufferSize *int64 `protobuf:"varint,1,opt,name=metrics_buffer_size,json=metricsBufferSize,def=10000" json:"metrics_buffer_size,omitempty"`
	// Whether to include timestamps in metrics. If enabled (default) each metric
	// string includes the metric timestamp as recorded in the EventMetric.
	// Prometheus associates the scraped values with this timestamp. If disabled,
	// i.e. timestamps are not exported, prometheus associates scraped values with
	// scrape timestamp.
	IncludeTimestamp *bool `protobuf:"varint,2,opt,name=include_timestamp,json=includeTimestamp,def=1" json:"include_timestamp,omitempty"`
	// URL that prometheus scrapes metrics from.
	MetricsUrl       *string `protobuf:"bytes,3,opt,name=metrics_url,json=metricsUrl,def=/metrics" json:"metrics_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SurfacerConf) Reset()                    { *m = SurfacerConf{} }
func (m *SurfacerConf) String() string            { return proto1.CompactTextString(m) }
func (*SurfacerConf) ProtoMessage()               {}
func (*SurfacerConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_SurfacerConf_MetricsBufferSize int64 = 10000
const Default_SurfacerConf_IncludeTimestamp bool = true
const Default_SurfacerConf_MetricsUrl string = "/metrics"

func (m *SurfacerConf) GetMetricsBufferSize() int64 {
	if m != nil && m.MetricsBufferSize != nil {
		return *m.MetricsBufferSize
	}
	return Default_SurfacerConf_MetricsBufferSize
}

func (m *SurfacerConf) GetIncludeTimestamp() bool {
	if m != nil && m.IncludeTimestamp != nil {
		return *m.IncludeTimestamp
	}
	return Default_SurfacerConf_IncludeTimestamp
}

func (m *SurfacerConf) GetMetricsUrl() string {
	if m != nil && m.MetricsUrl != nil {
		return *m.MetricsUrl
	}
	return Default_SurfacerConf_MetricsUrl
}

func init() {
	proto1.RegisterType((*SurfacerConf)(nil), "cloudprober.surfacer.prometheus.SurfacerConf")
}

func init() {
	proto1.RegisterFile("github.com/google/cloudprober/surfacers/prometheus/proto/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8d, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x0a, 0x52, 0x31, 0x0c, 0x34, 0x2c, 0xd9, 0x88, 0x98, 0xc2, 0x12, 0xa7, 0x03,
	0x4b, 0x46, 0x10, 0x2f, 0x90, 0xc2, 0x1c, 0x35, 0xee, 0xd9, 0xb5, 0x64, 0xe7, 0xa2, 0xb3, 0x6f,
	0xe9, 0xeb, 0xf0, 0xa2, 0x48, 0x91, 0x23, 0xba, 0xfd, 0xf7, 0xff, 0xdf, 0xa7, 0x93, 0x5f, 0xd6,
	0xa5, 0x33, 0x8f, 0x8d, 0xc6, 0xa0, 0x2c, 0xa2, 0xf5, 0xa0, 0xb4, 0x47, 0x3e, 0xcd, 0x84, 0x23,
	0x90, 0x8a, 0x4c, 0xe6, 0xa8, 0x81, 0xa2, 0x9a, 0x09, 0x03, 0xa4, 0x33, 0xf0, 0x12, 0x13, 0x2a,
	0x8d, 0x93, 0x71, 0xb6, 0x59, 0x8e, 0xe2, 0xe5, 0x4a, 0x6a, 0x56, 0xa9, 0xf9, 0x77, 0x5e, 0x7f,
	0x85, 0x7c, 0x3c, 0xe4, 0xfe, 0x13, 0x27, 0x53, 0xbc, 0xcb, 0xe7, 0x00, 0x89, 0x9c, 0x8e, 0xc3,
	0xc8, 0xc6, 0x00, 0x0d, 0xd1, 0x5d, 0xa0, 0x14, 0x95, 0xa8, 0x37, 0xdd, 0xdd, 0xbe, 0x6d, 0xdb,
	0xb6, 0xdf, 0x65, 0xe2, 0x63, 0x01, 0x0e, 0xee, 0x02, 0xc5, 0x5e, 0xee, 0xdc, 0xa4, 0x3d, 0x9f,
	0x60, 0x48, 0x2e, 0x40, 0x4c, 0xc7, 0x30, 0x97, 0x37, 0x95, 0xa8, 0xb7, 0xdd, 0x6d, 0x22, 0x86,
	0xfe, 0x29, 0xcf, 0xdf, 0xeb, 0x5a, 0xbc, 0xc9, 0x87, 0xf5, 0x13, 0x93, 0x2f, 0x37, 0x95, 0xa8,
	0xef, 0xbb, 0xad, 0xca, 0x5d, 0x2f, 0x73, 0xf8, 0x21, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0xab, 0xab, 0xe4, 0x0e, 0x01, 0x00, 0x00,
}