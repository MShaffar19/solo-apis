// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/tracing/v3/custom_tag.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/type/metadata/v3"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for CustomTag
func (this *CustomTag) MarshalJSON() ([]byte, error) {
	str, err := CustomTagMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CustomTag
func (this *CustomTag) UnmarshalJSON(b []byte) error {
	return CustomTagUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CustomTag_Literal
func (this *CustomTag_Literal) MarshalJSON() ([]byte, error) {
	str, err := CustomTagMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CustomTag_Literal
func (this *CustomTag_Literal) UnmarshalJSON(b []byte) error {
	return CustomTagUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CustomTag_Environment
func (this *CustomTag_Environment) MarshalJSON() ([]byte, error) {
	str, err := CustomTagMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CustomTag_Environment
func (this *CustomTag_Environment) UnmarshalJSON(b []byte) error {
	return CustomTagUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CustomTag_Header
func (this *CustomTag_Header) MarshalJSON() ([]byte, error) {
	str, err := CustomTagMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CustomTag_Header
func (this *CustomTag_Header) UnmarshalJSON(b []byte) error {
	return CustomTagUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CustomTag_Metadata
func (this *CustomTag_Metadata) MarshalJSON() ([]byte, error) {
	str, err := CustomTagMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CustomTag_Metadata
func (this *CustomTag_Metadata) UnmarshalJSON(b []byte) error {
	return CustomTagUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CustomTagMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	CustomTagUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)