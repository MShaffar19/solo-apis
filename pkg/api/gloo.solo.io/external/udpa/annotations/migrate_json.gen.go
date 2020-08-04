// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/migrate.proto

package annotations

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for MigrateAnnotation
func (this *MigrateAnnotation) MarshalJSON() ([]byte, error) {
	str, err := MigrateMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MigrateAnnotation
func (this *MigrateAnnotation) UnmarshalJSON(b []byte) error {
	return MigrateUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FieldMigrateAnnotation
func (this *FieldMigrateAnnotation) MarshalJSON() ([]byte, error) {
	str, err := MigrateMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FieldMigrateAnnotation
func (this *FieldMigrateAnnotation) UnmarshalJSON(b []byte) error {
	return MigrateUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FileMigrateAnnotation
func (this *FileMigrateAnnotation) MarshalJSON() ([]byte, error) {
	str, err := MigrateMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FileMigrateAnnotation
func (this *FileMigrateAnnotation) UnmarshalJSON(b []byte) error {
	return MigrateUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MigrateMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MigrateUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)