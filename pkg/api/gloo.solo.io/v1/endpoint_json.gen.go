// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/endpoint.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Endpoint
func (this *Endpoint) MarshalJSON() ([]byte, error) {
	str, err := EndpointMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Endpoint
func (this *Endpoint) UnmarshalJSON(b []byte) error {
	return EndpointUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HealthCheckConfig
func (this *HealthCheckConfig) MarshalJSON() ([]byte, error) {
	str, err := EndpointMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HealthCheckConfig
func (this *HealthCheckConfig) UnmarshalJSON(b []byte) error {
	return EndpointUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	EndpointMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	EndpointUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)