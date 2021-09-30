// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: org_client.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ChangeClientLimitResponse)(nil)
var _ json.Unmarshaler = (*ChangeClientLimitResponse)(nil)
var _ json.Marshaler = (*ChangeClientLimitRequest)(nil)
var _ json.Unmarshaler = (*ChangeClientLimitRequest)(nil)
var _ json.Marshaler = (*GrantEndpointRequest)(nil)
var _ json.Unmarshaler = (*GrantEndpointRequest)(nil)
var _ json.Marshaler = (*GrantEndpointResponse)(nil)
var _ json.Unmarshaler = (*GrantEndpointResponse)(nil)
var _ json.Marshaler = (*RevokeEndpointRequest)(nil)
var _ json.Unmarshaler = (*RevokeEndpointRequest)(nil)
var _ json.Marshaler = (*RevokeEndpointResponse)(nil)
var _ json.Unmarshaler = (*RevokeEndpointResponse)(nil)
var _ json.Marshaler = (*UpdateCredentialsResponse)(nil)
var _ json.Unmarshaler = (*UpdateCredentialsResponse)(nil)
var _ json.Marshaler = (*UpdateCredentialsRequest)(nil)
var _ json.Unmarshaler = (*UpdateCredentialsRequest)(nil)
var _ json.Marshaler = (*GetCredentialsResponse)(nil)
var _ json.Unmarshaler = (*GetCredentialsResponse)(nil)
var _ json.Marshaler = (*GetCredentialsRequest)(nil)
var _ json.Unmarshaler = (*GetCredentialsRequest)(nil)
var _ json.Marshaler = (*DeleteClientRequest)(nil)
var _ json.Unmarshaler = (*DeleteClientRequest)(nil)
var _ json.Marshaler = (*DeleteClientResponse)(nil)
var _ json.Unmarshaler = (*DeleteClientResponse)(nil)
var _ json.Marshaler = (*ClientInfo)(nil)
var _ json.Unmarshaler = (*ClientInfo)(nil)
var _ json.Marshaler = (*CreateClientRequest)(nil)
var _ json.Unmarshaler = (*CreateClientRequest)(nil)
var _ json.Marshaler = (*CreateClientResponse)(nil)
var _ json.Unmarshaler = (*CreateClientResponse)(nil)

// ChangeClientLimitResponse implement json.Marshaler.
func (m *ChangeClientLimitResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ChangeClientLimitResponse implement json.Marshaler.
func (m *ChangeClientLimitResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ChangeClientLimitRequest implement json.Marshaler.
func (m *ChangeClientLimitRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ChangeClientLimitRequest implement json.Marshaler.
func (m *ChangeClientLimitRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GrantEndpointRequest implement json.Marshaler.
func (m *GrantEndpointRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GrantEndpointRequest implement json.Marshaler.
func (m *GrantEndpointRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GrantEndpointResponse implement json.Marshaler.
func (m *GrantEndpointResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GrantEndpointResponse implement json.Marshaler.
func (m *GrantEndpointResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RevokeEndpointRequest implement json.Marshaler.
func (m *RevokeEndpointRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RevokeEndpointRequest implement json.Marshaler.
func (m *RevokeEndpointRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RevokeEndpointResponse implement json.Marshaler.
func (m *RevokeEndpointResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RevokeEndpointResponse implement json.Marshaler.
func (m *RevokeEndpointResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCredentialsResponse implement json.Marshaler.
func (m *UpdateCredentialsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCredentialsResponse implement json.Marshaler.
func (m *UpdateCredentialsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCredentialsRequest implement json.Marshaler.
func (m *UpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCredentialsRequest implement json.Marshaler.
func (m *UpdateCredentialsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCredentialsResponse implement json.Marshaler.
func (m *GetCredentialsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCredentialsResponse implement json.Marshaler.
func (m *GetCredentialsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCredentialsRequest implement json.Marshaler.
func (m *GetCredentialsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCredentialsRequest implement json.Marshaler.
func (m *GetCredentialsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteClientRequest implement json.Marshaler.
func (m *DeleteClientRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteClientRequest implement json.Marshaler.
func (m *DeleteClientRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteClientResponse implement json.Marshaler.
func (m *DeleteClientResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteClientResponse implement json.Marshaler.
func (m *DeleteClientResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClientInfo implement json.Marshaler.
func (m *ClientInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClientInfo implement json.Marshaler.
func (m *ClientInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateClientRequest implement json.Marshaler.
func (m *CreateClientRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateClientRequest implement json.Marshaler.
func (m *CreateClientRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateClientResponse implement json.Marshaler.
func (m *CreateClientResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateClientResponse implement json.Marshaler.
func (m *CreateClientResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
