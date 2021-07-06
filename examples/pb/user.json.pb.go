// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: user.proto

package pb

import (
	json "encoding/json"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*User)(nil)
var _ json.Unmarshaler = (*User)(nil)
var _ json.Marshaler = (*GetUserRequest)(nil)
var _ json.Unmarshaler = (*GetUserRequest)(nil)
var _ json.Marshaler = (*GetUserResponse)(nil)
var _ json.Unmarshaler = (*GetUserResponse)(nil)
var _ json.Marshaler = (*UpdateUserRequest)(nil)
var _ json.Unmarshaler = (*UpdateUserRequest)(nil)
var _ json.Marshaler = (*UpdateUserResponse)(nil)
var _ json.Unmarshaler = (*UpdateUserResponse)(nil)

// User implement json.Marshaler.
func (m *User) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// User implement json.Marshaler.
func (m *User) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetUserRequest implement json.Marshaler.
func (m *GetUserRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetUserRequest implement json.Marshaler.
func (m *GetUserRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetUserResponse implement json.Marshaler.
func (m *GetUserResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetUserResponse implement json.Marshaler.
func (m *GetUserResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// UpdateUserRequest implement json.Marshaler.
func (m *UpdateUserRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// UpdateUserRequest implement json.Marshaler.
func (m *UpdateUserRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// UpdateUserResponse implement json.Marshaler.
func (m *UpdateUserResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// UpdateUserResponse implement json.Marshaler.
func (m *UpdateUserResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}
