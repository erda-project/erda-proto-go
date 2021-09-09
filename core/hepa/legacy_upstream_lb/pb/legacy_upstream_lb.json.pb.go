// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: legacy_upstream_lb.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*UpstreamLb)(nil)
var _ json.Unmarshaler = (*UpstreamLb)(nil)
var _ json.Marshaler = (*TargetOnlineRequest)(nil)
var _ json.Unmarshaler = (*TargetOnlineRequest)(nil)
var _ json.Marshaler = (*TargetOnlineResponse)(nil)
var _ json.Unmarshaler = (*TargetOnlineResponse)(nil)
var _ json.Marshaler = (*TargetOfflineRequest)(nil)
var _ json.Unmarshaler = (*TargetOfflineRequest)(nil)
var _ json.Marshaler = (*TargetOfflineResponse)(nil)
var _ json.Unmarshaler = (*TargetOfflineResponse)(nil)

// UpstreamLb implement json.Marshaler.
func (m *UpstreamLb) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpstreamLb implement json.Marshaler.
func (m *UpstreamLb) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TargetOnlineRequest implement json.Marshaler.
func (m *TargetOnlineRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TargetOnlineRequest implement json.Marshaler.
func (m *TargetOnlineRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TargetOnlineResponse implement json.Marshaler.
func (m *TargetOnlineResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TargetOnlineResponse implement json.Marshaler.
func (m *TargetOnlineResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TargetOfflineRequest implement json.Marshaler.
func (m *TargetOfflineRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TargetOfflineRequest implement json.Marshaler.
func (m *TargetOfflineRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TargetOfflineResponse implement json.Marshaler.
func (m *TargetOfflineResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TargetOfflineResponse implement json.Marshaler.
func (m *TargetOfflineResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
