// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: settings.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetSettingsRequest)(nil)
var _ json.Unmarshaler = (*GetSettingsRequest)(nil)
var _ json.Marshaler = (*GetSettingsResponse)(nil)
var _ json.Unmarshaler = (*GetSettingsResponse)(nil)
var _ json.Marshaler = (*PutSettingsRequest)(nil)
var _ json.Unmarshaler = (*PutSettingsRequest)(nil)
var _ json.Marshaler = (*PutSettingsResponse)(nil)
var _ json.Unmarshaler = (*PutSettingsResponse)(nil)
var _ json.Marshaler = (*RegisterMonitorConfigRequest)(nil)
var _ json.Unmarshaler = (*RegisterMonitorConfigRequest)(nil)
var _ json.Marshaler = (*MonitorConfig)(nil)
var _ json.Unmarshaler = (*MonitorConfig)(nil)
var _ json.Marshaler = (*RegisterMonitorConfigResponse)(nil)
var _ json.Unmarshaler = (*RegisterMonitorConfigResponse)(nil)
var _ json.Marshaler = (*ConfigGroups)(nil)
var _ json.Unmarshaler = (*ConfigGroups)(nil)
var _ json.Marshaler = (*ConfigGroup)(nil)
var _ json.Unmarshaler = (*ConfigGroup)(nil)
var _ json.Marshaler = (*ConfigItem)(nil)
var _ json.Unmarshaler = (*ConfigItem)(nil)

// GetSettingsRequest implement json.Marshaler.
func (m *GetSettingsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSettingsRequest implement json.Marshaler.
func (m *GetSettingsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetSettingsResponse implement json.Marshaler.
func (m *GetSettingsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSettingsResponse implement json.Marshaler.
func (m *GetSettingsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// PutSettingsRequest implement json.Marshaler.
func (m *PutSettingsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PutSettingsRequest implement json.Marshaler.
func (m *PutSettingsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// PutSettingsResponse implement json.Marshaler.
func (m *PutSettingsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PutSettingsResponse implement json.Marshaler.
func (m *PutSettingsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RegisterMonitorConfigRequest implement json.Marshaler.
func (m *RegisterMonitorConfigRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RegisterMonitorConfigRequest implement json.Marshaler.
func (m *RegisterMonitorConfigRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// MonitorConfig implement json.Marshaler.
func (m *MonitorConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// MonitorConfig implement json.Marshaler.
func (m *MonitorConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RegisterMonitorConfigResponse implement json.Marshaler.
func (m *RegisterMonitorConfigResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RegisterMonitorConfigResponse implement json.Marshaler.
func (m *RegisterMonitorConfigResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ConfigGroups implement json.Marshaler.
func (m *ConfigGroups) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ConfigGroups implement json.Marshaler.
func (m *ConfigGroups) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ConfigGroup implement json.Marshaler.
func (m *ConfigGroup) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ConfigGroup implement json.Marshaler.
func (m *ConfigGroup) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ConfigItem implement json.Marshaler.
func (m *ConfigItem) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ConfigItem implement json.Marshaler.
func (m *ConfigItem) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
