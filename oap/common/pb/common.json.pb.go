// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: common.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*KeyValue)(nil)
var _ json.Unmarshaler = (*KeyValue)(nil)
var _ json.Marshaler = (*KeyValueList)(nil)
var _ json.Unmarshaler = (*KeyValueList)(nil)
var _ json.Marshaler = (*AnyValue)(nil)
var _ json.Unmarshaler = (*AnyValue)(nil)
var _ json.Marshaler = (*ArrayValue)(nil)
var _ json.Unmarshaler = (*ArrayValue)(nil)
var _ json.Marshaler = (*Relation)(nil)
var _ json.Unmarshaler = (*Relation)(nil)

// KeyValue implement json.Marshaler.
func (m *KeyValue) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// KeyValue implement json.Marshaler.
func (m *KeyValue) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// KeyValueList implement json.Marshaler.
func (m *KeyValueList) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// KeyValueList implement json.Marshaler.
func (m *KeyValueList) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AnyValue implement json.Marshaler.
func (m *AnyValue) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AnyValue implement json.Marshaler.
func (m *AnyValue) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ArrayValue implement json.Marshaler.
func (m *ArrayValue) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ArrayValue implement json.Marshaler.
func (m *ArrayValue) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Relation implement json.Marshaler.
func (m *Relation) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Relation implement json.Marshaler.
func (m *Relation) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
