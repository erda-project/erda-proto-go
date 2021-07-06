// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: exception.proto

package pb

import (
	json "encoding/json"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetExceptionsRequest)(nil)
var _ json.Unmarshaler = (*GetExceptionsRequest)(nil)
var _ json.Marshaler = (*GetExceptionsResponse)(nil)
var _ json.Unmarshaler = (*GetExceptionsResponse)(nil)
var _ json.Marshaler = (*GetExceptionEventIdsRequest)(nil)
var _ json.Unmarshaler = (*GetExceptionEventIdsRequest)(nil)
var _ json.Marshaler = (*GetExceptionEventIdsResponse)(nil)
var _ json.Unmarshaler = (*GetExceptionEventIdsResponse)(nil)
var _ json.Marshaler = (*GetExceptionEventRequest)(nil)
var _ json.Unmarshaler = (*GetExceptionEventRequest)(nil)
var _ json.Marshaler = (*GetExceptionEventResponse)(nil)
var _ json.Unmarshaler = (*GetExceptionEventResponse)(nil)
var _ json.Marshaler = (*Exception)(nil)
var _ json.Unmarshaler = (*Exception)(nil)
var _ json.Marshaler = (*Stacks)(nil)
var _ json.Unmarshaler = (*Stacks)(nil)
var _ json.Marshaler = (*ExceptionEvent)(nil)
var _ json.Unmarshaler = (*ExceptionEvent)(nil)

// GetExceptionsRequest implement json.Marshaler.
func (m *GetExceptionsRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetExceptionsRequest implement json.Marshaler.
func (m *GetExceptionsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetExceptionsResponse implement json.Marshaler.
func (m *GetExceptionsResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetExceptionsResponse implement json.Marshaler.
func (m *GetExceptionsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetExceptionEventIdsRequest implement json.Marshaler.
func (m *GetExceptionEventIdsRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetExceptionEventIdsRequest implement json.Marshaler.
func (m *GetExceptionEventIdsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetExceptionEventIdsResponse implement json.Marshaler.
func (m *GetExceptionEventIdsResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetExceptionEventIdsResponse implement json.Marshaler.
func (m *GetExceptionEventIdsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetExceptionEventRequest implement json.Marshaler.
func (m *GetExceptionEventRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetExceptionEventRequest implement json.Marshaler.
func (m *GetExceptionEventRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetExceptionEventResponse implement json.Marshaler.
func (m *GetExceptionEventResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetExceptionEventResponse implement json.Marshaler.
func (m *GetExceptionEventResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Exception implement json.Marshaler.
func (m *Exception) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// Exception implement json.Marshaler.
func (m *Exception) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Stacks implement json.Marshaler.
func (m *Stacks) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// Stacks implement json.Marshaler.
func (m *Stacks) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ExceptionEvent implement json.Marshaler.
func (m *ExceptionEvent) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// ExceptionEvent implement json.Marshaler.
func (m *ExceptionEvent) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
