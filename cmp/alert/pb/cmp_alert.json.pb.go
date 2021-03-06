// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: cmp_alert.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetAlertConditionsRequest)(nil)
var _ json.Unmarshaler = (*GetAlertConditionsRequest)(nil)
var _ json.Marshaler = (*GetAlertConditionsResponse)(nil)
var _ json.Unmarshaler = (*GetAlertConditionsResponse)(nil)
var _ json.Marshaler = (*GetAlertConditionsValueRequest)(nil)
var _ json.Unmarshaler = (*GetAlertConditionsValueRequest)(nil)
var _ json.Marshaler = (*GetAlertConditionsValueResponse)(nil)
var _ json.Unmarshaler = (*GetAlertConditionsValueResponse)(nil)

// GetAlertConditionsRequest implement json.Marshaler.
func (m *GetAlertConditionsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertConditionsRequest implement json.Marshaler.
func (m *GetAlertConditionsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertConditionsResponse implement json.Marshaler.
func (m *GetAlertConditionsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertConditionsResponse implement json.Marshaler.
func (m *GetAlertConditionsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertConditionsValueRequest implement json.Marshaler.
func (m *GetAlertConditionsValueRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertConditionsValueRequest implement json.Marshaler.
func (m *GetAlertConditionsValueRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertConditionsValueResponse implement json.Marshaler.
func (m *GetAlertConditionsValueResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertConditionsValueResponse implement json.Marshaler.
func (m *GetAlertConditionsValueResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
