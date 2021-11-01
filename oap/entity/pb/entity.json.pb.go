// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: entity.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*Entity)(nil)
var _ json.Unmarshaler = (*Entity)(nil)
var _ json.Marshaler = (*SetEntityRequest)(nil)
var _ json.Unmarshaler = (*SetEntityRequest)(nil)
var _ json.Marshaler = (*SetEntityResponse)(nil)
var _ json.Unmarshaler = (*SetEntityResponse)(nil)
var _ json.Marshaler = (*RemoveEntityRequest)(nil)
var _ json.Unmarshaler = (*RemoveEntityRequest)(nil)
var _ json.Marshaler = (*RemoveEntityResponse)(nil)
var _ json.Unmarshaler = (*RemoveEntityResponse)(nil)
var _ json.Marshaler = (*GetEntityRequest)(nil)
var _ json.Unmarshaler = (*GetEntityRequest)(nil)
var _ json.Marshaler = (*GetEntityResponse)(nil)
var _ json.Unmarshaler = (*GetEntityResponse)(nil)
var _ json.Marshaler = (*ListEntitiesRequest)(nil)
var _ json.Unmarshaler = (*ListEntitiesRequest)(nil)
var _ json.Marshaler = (*ListEntitiesResponse)(nil)
var _ json.Unmarshaler = (*ListEntitiesResponse)(nil)
var _ json.Marshaler = (*EntityList)(nil)
var _ json.Unmarshaler = (*EntityList)(nil)

// Entity implement json.Marshaler.
func (m *Entity) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Entity implement json.Marshaler.
func (m *Entity) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SetEntityRequest implement json.Marshaler.
func (m *SetEntityRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SetEntityRequest implement json.Marshaler.
func (m *SetEntityRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SetEntityResponse implement json.Marshaler.
func (m *SetEntityResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SetEntityResponse implement json.Marshaler.
func (m *SetEntityResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RemoveEntityRequest implement json.Marshaler.
func (m *RemoveEntityRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RemoveEntityRequest implement json.Marshaler.
func (m *RemoveEntityRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RemoveEntityResponse implement json.Marshaler.
func (m *RemoveEntityResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RemoveEntityResponse implement json.Marshaler.
func (m *RemoveEntityResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEntityRequest implement json.Marshaler.
func (m *GetEntityRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEntityRequest implement json.Marshaler.
func (m *GetEntityRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetEntityResponse implement json.Marshaler.
func (m *GetEntityResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetEntityResponse implement json.Marshaler.
func (m *GetEntityResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListEntitiesRequest implement json.Marshaler.
func (m *ListEntitiesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListEntitiesRequest implement json.Marshaler.
func (m *ListEntitiesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListEntitiesResponse implement json.Marshaler.
func (m *ListEntitiesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListEntitiesResponse implement json.Marshaler.
func (m *ListEntitiesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// EntityList implement json.Marshaler.
func (m *EntityList) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// EntityList implement json.Marshaler.
func (m *EntityList) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
