// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: notify.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetAllNotifyTemplatesRequest)(nil)
var _ json.Unmarshaler = (*GetAllNotifyTemplatesRequest)(nil)
var _ json.Marshaler = (*GetAllNotifyTemplatesResponse)(nil)
var _ json.Unmarshaler = (*GetAllNotifyTemplatesResponse)(nil)
var _ json.Marshaler = (*Model)(nil)
var _ json.Unmarshaler = (*Model)(nil)
var _ json.Marshaler = (*Metadata)(nil)
var _ json.Unmarshaler = (*Metadata)(nil)
var _ json.Marshaler = (*Behavior)(nil)
var _ json.Unmarshaler = (*Behavior)(nil)
var _ json.Marshaler = (*Templates)(nil)
var _ json.Unmarshaler = (*Templates)(nil)
var _ json.Marshaler = (*Render)(nil)
var _ json.Unmarshaler = (*Render)(nil)
var _ json.Marshaler = (*GetNotifyTemplateRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyTemplateRequest)(nil)
var _ json.Marshaler = (*GetNotifyTemplateResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyTemplateResponse)(nil)
var _ json.Marshaler = (*GetNotifyRes)(nil)
var _ json.Unmarshaler = (*GetNotifyRes)(nil)
var _ json.Marshaler = (*CreateNotifyRequest)(nil)
var _ json.Unmarshaler = (*CreateNotifyRequest)(nil)
var _ json.Marshaler = (*CreateNotifyResponse)(nil)
var _ json.Unmarshaler = (*CreateNotifyResponse)(nil)
var _ json.Marshaler = (*DeleteNotifyRequest)(nil)
var _ json.Unmarshaler = (*DeleteNotifyRequest)(nil)
var _ json.Marshaler = (*DeleteNotifyResponse)(nil)
var _ json.Unmarshaler = (*DeleteNotifyResponse)(nil)
var _ json.Marshaler = (*UpdateNotifyRequest)(nil)
var _ json.Unmarshaler = (*UpdateNotifyRequest)(nil)
var _ json.Marshaler = (*UpdateNotifyResponse)(nil)
var _ json.Unmarshaler = (*UpdateNotifyResponse)(nil)
var _ json.Marshaler = (*GetUserNotifyListRequest)(nil)
var _ json.Unmarshaler = (*GetUserNotifyListRequest)(nil)
var _ json.Marshaler = (*GetUserNotifyListResponse)(nil)
var _ json.Unmarshaler = (*GetUserNotifyListResponse)(nil)
var _ json.Marshaler = (*NotifyRes)(nil)
var _ json.Unmarshaler = (*NotifyRes)(nil)
var _ json.Marshaler = (*NotifyTarget)(nil)
var _ json.Unmarshaler = (*NotifyTarget)(nil)
var _ json.Marshaler = (*TargetValue)(nil)
var _ json.Unmarshaler = (*TargetValue)(nil)
var _ json.Marshaler = (*NotifyEnableRequest)(nil)
var _ json.Unmarshaler = (*NotifyEnableRequest)(nil)
var _ json.Marshaler = (*NotifyEnableResponse)(nil)
var _ json.Unmarshaler = (*NotifyEnableResponse)(nil)
var _ json.Marshaler = (*CreateUserDefineNotifyTemplateRequest)(nil)
var _ json.Unmarshaler = (*CreateUserDefineNotifyTemplateRequest)(nil)
var _ json.Marshaler = (*CreateUserDefineNotifyTemplateResponse)(nil)
var _ json.Unmarshaler = (*CreateUserDefineNotifyTemplateResponse)(nil)
var _ json.Marshaler = (*GetNotifyDetailRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyDetailRequest)(nil)
var _ json.Marshaler = (*GetNotifyDetailResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyDetailResponse)(nil)
var _ json.Marshaler = (*NotifyDetailResponse)(nil)
var _ json.Unmarshaler = (*NotifyDetailResponse)(nil)
var _ json.Marshaler = (*GetAllGroupsRequest)(nil)
var _ json.Unmarshaler = (*GetAllGroupsRequest)(nil)
var _ json.Marshaler = (*GetAllGroupsResponse)(nil)
var _ json.Unmarshaler = (*GetAllGroupsResponse)(nil)
var _ json.Marshaler = (*GetAllGroupData)(nil)
var _ json.Unmarshaler = (*GetAllGroupData)(nil)

// GetAllNotifyTemplatesRequest implement json.Marshaler.
func (m *GetAllNotifyTemplatesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAllNotifyTemplatesRequest implement json.Marshaler.
func (m *GetAllNotifyTemplatesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAllNotifyTemplatesResponse implement json.Marshaler.
func (m *GetAllNotifyTemplatesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAllNotifyTemplatesResponse implement json.Marshaler.
func (m *GetAllNotifyTemplatesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Model implement json.Marshaler.
func (m *Model) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Model implement json.Marshaler.
func (m *Model) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Metadata implement json.Marshaler.
func (m *Metadata) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Metadata implement json.Marshaler.
func (m *Metadata) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Behavior implement json.Marshaler.
func (m *Behavior) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Behavior implement json.Marshaler.
func (m *Behavior) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Templates implement json.Marshaler.
func (m *Templates) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Templates implement json.Marshaler.
func (m *Templates) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Render implement json.Marshaler.
func (m *Render) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Render implement json.Marshaler.
func (m *Render) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyTemplateRequest implement json.Marshaler.
func (m *GetNotifyTemplateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyTemplateRequest implement json.Marshaler.
func (m *GetNotifyTemplateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyTemplateResponse implement json.Marshaler.
func (m *GetNotifyTemplateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyTemplateResponse implement json.Marshaler.
func (m *GetNotifyTemplateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyRes implement json.Marshaler.
func (m *GetNotifyRes) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyRes implement json.Marshaler.
func (m *GetNotifyRes) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateNotifyRequest implement json.Marshaler.
func (m *CreateNotifyRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateNotifyRequest implement json.Marshaler.
func (m *CreateNotifyRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateNotifyResponse implement json.Marshaler.
func (m *CreateNotifyResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateNotifyResponse implement json.Marshaler.
func (m *CreateNotifyResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteNotifyRequest implement json.Marshaler.
func (m *DeleteNotifyRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteNotifyRequest implement json.Marshaler.
func (m *DeleteNotifyRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteNotifyResponse implement json.Marshaler.
func (m *DeleteNotifyResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteNotifyResponse implement json.Marshaler.
func (m *DeleteNotifyResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateNotifyRequest implement json.Marshaler.
func (m *UpdateNotifyRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateNotifyRequest implement json.Marshaler.
func (m *UpdateNotifyRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateNotifyResponse implement json.Marshaler.
func (m *UpdateNotifyResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateNotifyResponse implement json.Marshaler.
func (m *UpdateNotifyResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetUserNotifyListRequest implement json.Marshaler.
func (m *GetUserNotifyListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetUserNotifyListRequest implement json.Marshaler.
func (m *GetUserNotifyListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetUserNotifyListResponse implement json.Marshaler.
func (m *GetUserNotifyListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetUserNotifyListResponse implement json.Marshaler.
func (m *GetUserNotifyListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyRes implement json.Marshaler.
func (m *NotifyRes) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyRes implement json.Marshaler.
func (m *NotifyRes) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyTarget implement json.Marshaler.
func (m *NotifyTarget) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyTarget implement json.Marshaler.
func (m *NotifyTarget) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TargetValue implement json.Marshaler.
func (m *TargetValue) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TargetValue implement json.Marshaler.
func (m *TargetValue) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyEnableRequest implement json.Marshaler.
func (m *NotifyEnableRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyEnableRequest implement json.Marshaler.
func (m *NotifyEnableRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyEnableResponse implement json.Marshaler.
func (m *NotifyEnableResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyEnableResponse implement json.Marshaler.
func (m *NotifyEnableResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateUserDefineNotifyTemplateRequest implement json.Marshaler.
func (m *CreateUserDefineNotifyTemplateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateUserDefineNotifyTemplateRequest implement json.Marshaler.
func (m *CreateUserDefineNotifyTemplateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateUserDefineNotifyTemplateResponse implement json.Marshaler.
func (m *CreateUserDefineNotifyTemplateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateUserDefineNotifyTemplateResponse implement json.Marshaler.
func (m *CreateUserDefineNotifyTemplateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyDetailRequest implement json.Marshaler.
func (m *GetNotifyDetailRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyDetailRequest implement json.Marshaler.
func (m *GetNotifyDetailRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyDetailResponse implement json.Marshaler.
func (m *GetNotifyDetailResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyDetailResponse implement json.Marshaler.
func (m *GetNotifyDetailResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyDetailResponse implement json.Marshaler.
func (m *NotifyDetailResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyDetailResponse implement json.Marshaler.
func (m *NotifyDetailResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAllGroupsRequest implement json.Marshaler.
func (m *GetAllGroupsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAllGroupsRequest implement json.Marshaler.
func (m *GetAllGroupsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAllGroupsResponse implement json.Marshaler.
func (m *GetAllGroupsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAllGroupsResponse implement json.Marshaler.
func (m *GetAllGroupsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAllGroupData implement json.Marshaler.
func (m *GetAllGroupData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAllGroupData implement json.Marshaler.
func (m *GetAllGroupData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
