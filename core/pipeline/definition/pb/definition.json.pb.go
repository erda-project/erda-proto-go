// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: definition.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*PipelineDefinitionProcessRequest)(nil)
var _ json.Unmarshaler = (*PipelineDefinitionProcessRequest)(nil)
var _ json.Marshaler = (*PipelineDefinitionProcessResponse)(nil)
var _ json.Unmarshaler = (*PipelineDefinitionProcessResponse)(nil)
var _ json.Marshaler = (*PipelineDefinitionProcessVersionRequest)(nil)
var _ json.Unmarshaler = (*PipelineDefinitionProcessVersionRequest)(nil)
var _ json.Marshaler = (*PipelineDefinitionProcessVersionResponse)(nil)
var _ json.Unmarshaler = (*PipelineDefinitionProcessVersionResponse)(nil)

// PipelineDefinitionProcessRequest implement json.Marshaler.
func (m *PipelineDefinitionProcessRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PipelineDefinitionProcessRequest implement json.Marshaler.
func (m *PipelineDefinitionProcessRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// PipelineDefinitionProcessResponse implement json.Marshaler.
func (m *PipelineDefinitionProcessResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PipelineDefinitionProcessResponse implement json.Marshaler.
func (m *PipelineDefinitionProcessResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// PipelineDefinitionProcessVersionRequest implement json.Marshaler.
func (m *PipelineDefinitionProcessVersionRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PipelineDefinitionProcessVersionRequest implement json.Marshaler.
func (m *PipelineDefinitionProcessVersionRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// PipelineDefinitionProcessVersionResponse implement json.Marshaler.
func (m *PipelineDefinitionProcessVersionResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PipelineDefinitionProcessVersionResponse implement json.Marshaler.
func (m *PipelineDefinitionProcessVersionResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}
