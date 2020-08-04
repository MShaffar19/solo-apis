// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/http_uri.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for HttpUri
func (this *HttpUri) MarshalJSON() ([]byte, error) {
	str, err := HttpUriMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpUri
func (this *HttpUri) UnmarshalJSON(b []byte) error {
	return HttpUriUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	HttpUriMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	HttpUriUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)