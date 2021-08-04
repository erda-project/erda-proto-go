// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: project.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetProjectOverviewRequest)(nil)
var _ json.Unmarshaler = (*GetProjectOverviewRequest)(nil)
var _ json.Marshaler = (*GetProjectOverviewResponse)(nil)
var _ json.Unmarshaler = (*GetProjectOverviewResponse)(nil)
var _ json.Marshaler = (*ProjectOverviewList)(nil)
var _ json.Unmarshaler = (*ProjectOverviewList)(nil)
var _ json.Marshaler = (*ProjectOverview)(nil)
var _ json.Unmarshaler = (*ProjectOverview)(nil)
var _ json.Marshaler = (*GetProjectRequest)(nil)
var _ json.Unmarshaler = (*GetProjectRequest)(nil)
var _ json.Marshaler = (*GetProjectResponse)(nil)
var _ json.Unmarshaler = (*GetProjectResponse)(nil)
var _ json.Marshaler = (*GetProjectsRequest)(nil)
var _ json.Unmarshaler = (*GetProjectsRequest)(nil)
var _ json.Marshaler = (*GetProjectsResponse)(nil)
var _ json.Unmarshaler = (*GetProjectsResponse)(nil)
var _ json.Marshaler = (*DeleteProjectRequest)(nil)
var _ json.Unmarshaler = (*DeleteProjectRequest)(nil)
var _ json.Marshaler = (*DeleteProjectResponse)(nil)
var _ json.Unmarshaler = (*DeleteProjectResponse)(nil)
var _ json.Marshaler = (*CreateProjectRequest)(nil)
var _ json.Unmarshaler = (*CreateProjectRequest)(nil)
var _ json.Marshaler = (*CreateProjectResponse)(nil)
var _ json.Unmarshaler = (*CreateProjectResponse)(nil)
var _ json.Marshaler = (*UpdateProjectRequest)(nil)
var _ json.Unmarshaler = (*UpdateProjectRequest)(nil)
var _ json.Marshaler = (*UpdateProjectResponse)(nil)
var _ json.Unmarshaler = (*UpdateProjectResponse)(nil)
var _ json.Marshaler = (*Project)(nil)
var _ json.Unmarshaler = (*Project)(nil)
var _ json.Marshaler = (*TenantRelationship)(nil)
var _ json.Unmarshaler = (*TenantRelationship)(nil)

// GetProjectOverviewRequest implement json.Marshaler.
func (m *GetProjectOverviewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetProjectOverviewRequest implement json.Marshaler.
func (m *GetProjectOverviewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetProjectOverviewResponse implement json.Marshaler.
func (m *GetProjectOverviewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetProjectOverviewResponse implement json.Marshaler.
func (m *GetProjectOverviewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProjectOverviewList implement json.Marshaler.
func (m *ProjectOverviewList) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProjectOverviewList implement json.Marshaler.
func (m *ProjectOverviewList) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProjectOverview implement json.Marshaler.
func (m *ProjectOverview) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProjectOverview implement json.Marshaler.
func (m *ProjectOverview) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetProjectRequest implement json.Marshaler.
func (m *GetProjectRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetProjectRequest implement json.Marshaler.
func (m *GetProjectRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetProjectResponse implement json.Marshaler.
func (m *GetProjectResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetProjectResponse implement json.Marshaler.
func (m *GetProjectResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetProjectsRequest implement json.Marshaler.
func (m *GetProjectsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetProjectsRequest implement json.Marshaler.
func (m *GetProjectsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetProjectsResponse implement json.Marshaler.
func (m *GetProjectsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetProjectsResponse implement json.Marshaler.
func (m *GetProjectsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteProjectRequest implement json.Marshaler.
func (m *DeleteProjectRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteProjectRequest implement json.Marshaler.
func (m *DeleteProjectRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteProjectResponse implement json.Marshaler.
func (m *DeleteProjectResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteProjectResponse implement json.Marshaler.
func (m *DeleteProjectResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateProjectRequest implement json.Marshaler.
func (m *CreateProjectRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateProjectRequest implement json.Marshaler.
func (m *CreateProjectRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateProjectResponse implement json.Marshaler.
func (m *CreateProjectResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateProjectResponse implement json.Marshaler.
func (m *CreateProjectResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateProjectRequest implement json.Marshaler.
func (m *UpdateProjectRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateProjectRequest implement json.Marshaler.
func (m *UpdateProjectRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateProjectResponse implement json.Marshaler.
func (m *UpdateProjectResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateProjectResponse implement json.Marshaler.
func (m *UpdateProjectResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Project implement json.Marshaler.
func (m *Project) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Project implement json.Marshaler.
func (m *Project) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TenantRelationship implement json.Marshaler.
func (m *TenantRelationship) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TenantRelationship implement json.Marshaler.
func (m *TenantRelationship) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
