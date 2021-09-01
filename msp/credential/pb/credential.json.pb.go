// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: credential.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CreateAccessKeyRequest)(nil)
var _ json.Unmarshaler = (*CreateAccessKeyRequest)(nil)
var _ json.Marshaler = (*CreateAccessKeyResponse)(nil)
var _ json.Unmarshaler = (*CreateAccessKeyResponse)(nil)
var _ json.Marshaler = (*DeleteAccessKeyRequest)(nil)
var _ json.Unmarshaler = (*DeleteAccessKeyRequest)(nil)
var _ json.Marshaler = (*DeleteAccessKeyResponse)(nil)
var _ json.Unmarshaler = (*DeleteAccessKeyResponse)(nil)
var _ json.Marshaler = (*GetAccessKeyRequest)(nil)
var _ json.Unmarshaler = (*GetAccessKeyRequest)(nil)
var _ json.Marshaler = (*GetAccessKeyResponse)(nil)
var _ json.Unmarshaler = (*GetAccessKeyResponse)(nil)
var _ json.Marshaler = (*DownloadAccessKeyFileRequest)(nil)
var _ json.Unmarshaler = (*DownloadAccessKeyFileRequest)(nil)
var _ json.Marshaler = (*DownloadAccessKeyFileResponse)(nil)
var _ json.Unmarshaler = (*DownloadAccessKeyFileResponse)(nil)
var _ json.Marshaler = (*QueryAccessKeysRequest)(nil)
var _ json.Unmarshaler = (*QueryAccessKeysRequest)(nil)
var _ json.Marshaler = (*QueryAccessKeysResponse)(nil)
var _ json.Unmarshaler = (*QueryAccessKeysResponse)(nil)
var _ json.Marshaler = (*QueryAccessKeysData)(nil)
var _ json.Unmarshaler = (*QueryAccessKeysData)(nil)

// CreateAccessKeyRequest implement json.Marshaler.
func (m *CreateAccessKeyRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAccessKeyRequest implement json.Marshaler.
func (m *CreateAccessKeyRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateAccessKeyResponse implement json.Marshaler.
func (m *CreateAccessKeyResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAccessKeyResponse implement json.Marshaler.
func (m *CreateAccessKeyResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteAccessKeyRequest implement json.Marshaler.
func (m *DeleteAccessKeyRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteAccessKeyRequest implement json.Marshaler.
func (m *DeleteAccessKeyRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteAccessKeyResponse implement json.Marshaler.
func (m *DeleteAccessKeyResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteAccessKeyResponse implement json.Marshaler.
func (m *DeleteAccessKeyResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAccessKeyRequest implement json.Marshaler.
func (m *GetAccessKeyRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAccessKeyRequest implement json.Marshaler.
func (m *GetAccessKeyRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAccessKeyResponse implement json.Marshaler.
func (m *GetAccessKeyResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAccessKeyResponse implement json.Marshaler.
func (m *GetAccessKeyResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DownloadAccessKeyFileRequest implement json.Marshaler.
func (m *DownloadAccessKeyFileRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DownloadAccessKeyFileRequest implement json.Marshaler.
func (m *DownloadAccessKeyFileRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DownloadAccessKeyFileResponse implement json.Marshaler.
func (m *DownloadAccessKeyFileResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DownloadAccessKeyFileResponse implement json.Marshaler.
func (m *DownloadAccessKeyFileResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAccessKeysRequest implement json.Marshaler.
func (m *QueryAccessKeysRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAccessKeysRequest implement json.Marshaler.
func (m *QueryAccessKeysRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAccessKeysResponse implement json.Marshaler.
func (m *QueryAccessKeysResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAccessKeysResponse implement json.Marshaler.
func (m *QueryAccessKeysResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAccessKeysData implement json.Marshaler.
func (m *QueryAccessKeysData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAccessKeysData implement json.Marshaler.
func (m *QueryAccessKeysData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
