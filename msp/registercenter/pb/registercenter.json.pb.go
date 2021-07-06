// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: registercenter.proto

package pb

import (
	json "encoding/json"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ListInterfaceRequest)(nil)
var _ json.Unmarshaler = (*ListInterfaceRequest)(nil)
var _ json.Marshaler = (*ListInterfaceResponse)(nil)
var _ json.Unmarshaler = (*ListInterfaceResponse)(nil)
var _ json.Marshaler = (*GetHTTPServicesRequest)(nil)
var _ json.Unmarshaler = (*GetHTTPServicesRequest)(nil)
var _ json.Marshaler = (*GetHTTPServicesResponse)(nil)
var _ json.Unmarshaler = (*GetHTTPServicesResponse)(nil)
var _ json.Marshaler = (*EnableHTTPServiceRequest)(nil)
var _ json.Unmarshaler = (*EnableHTTPServiceRequest)(nil)
var _ json.Marshaler = (*EnableHTTPServiceResponse)(nil)
var _ json.Unmarshaler = (*EnableHTTPServiceResponse)(nil)
var _ json.Marshaler = (*EnableHTTPService)(nil)
var _ json.Unmarshaler = (*EnableHTTPService)(nil)
var _ json.Marshaler = (*GetRouteRuleRequest)(nil)
var _ json.Unmarshaler = (*GetRouteRuleRequest)(nil)
var _ json.Marshaler = (*GetRouteRuleResponse)(nil)
var _ json.Unmarshaler = (*GetRouteRuleResponse)(nil)
var _ json.Marshaler = (*CreateRouteRuleRequest)(nil)
var _ json.Unmarshaler = (*CreateRouteRuleRequest)(nil)
var _ json.Marshaler = (*CreateRouteRuleResponse)(nil)
var _ json.Unmarshaler = (*CreateRouteRuleResponse)(nil)
var _ json.Marshaler = (*DeleteRouteRuleRequest)(nil)
var _ json.Unmarshaler = (*DeleteRouteRuleRequest)(nil)
var _ json.Marshaler = (*DeleteRouteRuleResponse)(nil)
var _ json.Unmarshaler = (*DeleteRouteRuleResponse)(nil)
var _ json.Marshaler = (*CetHostRuleRequest)(nil)
var _ json.Unmarshaler = (*CetHostRuleRequest)(nil)
var _ json.Marshaler = (*CetHostRuleResponse)(nil)
var _ json.Unmarshaler = (*CetHostRuleResponse)(nil)
var _ json.Marshaler = (*CreateHostRuleRequest)(nil)
var _ json.Unmarshaler = (*CreateHostRuleRequest)(nil)
var _ json.Marshaler = (*CreateHostRuleResponse)(nil)
var _ json.Unmarshaler = (*CreateHostRuleResponse)(nil)
var _ json.Marshaler = (*DeleteHostRuleRequest)(nil)
var _ json.Unmarshaler = (*DeleteHostRuleRequest)(nil)
var _ json.Marshaler = (*DeleteHostRuleResponse)(nil)
var _ json.Unmarshaler = (*DeleteHostRuleResponse)(nil)
var _ json.Marshaler = (*GetHostRuntimeRuleRequest)(nil)
var _ json.Unmarshaler = (*GetHostRuntimeRuleRequest)(nil)
var _ json.Marshaler = (*GetHostRuntimeRuleResponse)(nil)
var _ json.Unmarshaler = (*GetHostRuntimeRuleResponse)(nil)
var _ json.Marshaler = (*CreateHostRuntimeRuleRequest)(nil)
var _ json.Unmarshaler = (*CreateHostRuntimeRuleRequest)(nil)
var _ json.Marshaler = (*CreateHostRuntimeRuleResponse)(nil)
var _ json.Unmarshaler = (*CreateHostRuntimeRuleResponse)(nil)
var _ json.Marshaler = (*GetAllHostRuntimeRulesRequest)(nil)
var _ json.Unmarshaler = (*GetAllHostRuntimeRulesRequest)(nil)
var _ json.Marshaler = (*GetAllHostRuntimeRulesResponse)(nil)
var _ json.Unmarshaler = (*GetAllHostRuntimeRulesResponse)(nil)
var _ json.Marshaler = (*GetDubboInterfaceTimeRequest)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceTimeRequest)(nil)
var _ json.Marshaler = (*GetDubboInterfaceTimeResponse)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceTimeResponse)(nil)
var _ json.Marshaler = (*DubboInterfaceTime)(nil)
var _ json.Unmarshaler = (*DubboInterfaceTime)(nil)
var _ json.Marshaler = (*GetDubboInterfaceQPSRequest)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceQPSRequest)(nil)
var _ json.Marshaler = (*GetDubboInterfaceQPSResponse)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceQPSResponse)(nil)
var _ json.Marshaler = (*GetDubboInterfaceFailedRequest)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceFailedRequest)(nil)
var _ json.Marshaler = (*GetDubboInterfaceFailedResponse)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceFailedResponse)(nil)
var _ json.Marshaler = (*GetDubboInterfaceAvgTimeRequest)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceAvgTimeRequest)(nil)
var _ json.Marshaler = (*GetDubboInterfaceAvgTimeResponse)(nil)
var _ json.Unmarshaler = (*GetDubboInterfaceAvgTimeResponse)(nil)
var _ json.Marshaler = (*Interface)(nil)
var _ json.Unmarshaler = (*Interface)(nil)
var _ json.Marshaler = (*InterfaceOwner)(nil)
var _ json.Unmarshaler = (*InterfaceOwner)(nil)
var _ json.Marshaler = (*RequestRule)(nil)
var _ json.Unmarshaler = (*RequestRule)(nil)
var _ json.Marshaler = (*HostRules)(nil)
var _ json.Unmarshaler = (*HostRules)(nil)
var _ json.Marshaler = (*HostRoute)(nil)
var _ json.Unmarshaler = (*HostRoute)(nil)
var _ json.Marshaler = (*HostRuntimeRules)(nil)
var _ json.Unmarshaler = (*HostRuntimeRules)(nil)
var _ json.Marshaler = (*HostRuntimeRule)(nil)
var _ json.Unmarshaler = (*HostRuntimeRule)(nil)
var _ json.Marshaler = (*HostRuntimeInterfaces)(nil)
var _ json.Unmarshaler = (*HostRuntimeInterfaces)(nil)
var _ json.Marshaler = (*HTTPServices)(nil)
var _ json.Unmarshaler = (*HTTPServices)(nil)
var _ json.Marshaler = (*HTTPService)(nil)
var _ json.Unmarshaler = (*HTTPService)(nil)
var _ json.Marshaler = (*HTTPServiceItem)(nil)
var _ json.Unmarshaler = (*HTTPServiceItem)(nil)
var _ json.Marshaler = (*DubboInterface)(nil)
var _ json.Unmarshaler = (*DubboInterface)(nil)
var _ json.Marshaler = (*DubboInterfaceResult)(nil)
var _ json.Unmarshaler = (*DubboInterfaceResult)(nil)
var _ json.Marshaler = (*DubboMointorMap)(nil)
var _ json.Unmarshaler = (*DubboMointorMap)(nil)
var _ json.Marshaler = (*DubboMointor)(nil)
var _ json.Unmarshaler = (*DubboMointor)(nil)
var _ json.Marshaler = (*ServiceIpRequest)(nil)
var _ json.Unmarshaler = (*ServiceIpRequest)(nil)
var _ json.Marshaler = (*ServiceIpInfoResponse)(nil)
var _ json.Unmarshaler = (*ServiceIpInfoResponse)(nil)

// ListInterfaceRequest implement json.Marshaler.
func (m *ListInterfaceRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// ListInterfaceRequest implement json.Marshaler.
func (m *ListInterfaceRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// ListInterfaceResponse implement json.Marshaler.
func (m *ListInterfaceResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// ListInterfaceResponse implement json.Marshaler.
func (m *ListInterfaceResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetHTTPServicesRequest implement json.Marshaler.
func (m *GetHTTPServicesRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetHTTPServicesRequest implement json.Marshaler.
func (m *GetHTTPServicesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetHTTPServicesResponse implement json.Marshaler.
func (m *GetHTTPServicesResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetHTTPServicesResponse implement json.Marshaler.
func (m *GetHTTPServicesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// EnableHTTPServiceRequest implement json.Marshaler.
func (m *EnableHTTPServiceRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// EnableHTTPServiceRequest implement json.Marshaler.
func (m *EnableHTTPServiceRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// EnableHTTPServiceResponse implement json.Marshaler.
func (m *EnableHTTPServiceResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// EnableHTTPServiceResponse implement json.Marshaler.
func (m *EnableHTTPServiceResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// EnableHTTPService implement json.Marshaler.
func (m *EnableHTTPService) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// EnableHTTPService implement json.Marshaler.
func (m *EnableHTTPService) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetRouteRuleRequest implement json.Marshaler.
func (m *GetRouteRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetRouteRuleRequest implement json.Marshaler.
func (m *GetRouteRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetRouteRuleResponse implement json.Marshaler.
func (m *GetRouteRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetRouteRuleResponse implement json.Marshaler.
func (m *GetRouteRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CreateRouteRuleRequest implement json.Marshaler.
func (m *CreateRouteRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CreateRouteRuleRequest implement json.Marshaler.
func (m *CreateRouteRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CreateRouteRuleResponse implement json.Marshaler.
func (m *CreateRouteRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CreateRouteRuleResponse implement json.Marshaler.
func (m *CreateRouteRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DeleteRouteRuleRequest implement json.Marshaler.
func (m *DeleteRouteRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DeleteRouteRuleRequest implement json.Marshaler.
func (m *DeleteRouteRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DeleteRouteRuleResponse implement json.Marshaler.
func (m *DeleteRouteRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DeleteRouteRuleResponse implement json.Marshaler.
func (m *DeleteRouteRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CetHostRuleRequest implement json.Marshaler.
func (m *CetHostRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CetHostRuleRequest implement json.Marshaler.
func (m *CetHostRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CetHostRuleResponse implement json.Marshaler.
func (m *CetHostRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CetHostRuleResponse implement json.Marshaler.
func (m *CetHostRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CreateHostRuleRequest implement json.Marshaler.
func (m *CreateHostRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CreateHostRuleRequest implement json.Marshaler.
func (m *CreateHostRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CreateHostRuleResponse implement json.Marshaler.
func (m *CreateHostRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CreateHostRuleResponse implement json.Marshaler.
func (m *CreateHostRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DeleteHostRuleRequest implement json.Marshaler.
func (m *DeleteHostRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DeleteHostRuleRequest implement json.Marshaler.
func (m *DeleteHostRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DeleteHostRuleResponse implement json.Marshaler.
func (m *DeleteHostRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DeleteHostRuleResponse implement json.Marshaler.
func (m *DeleteHostRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetHostRuntimeRuleRequest implement json.Marshaler.
func (m *GetHostRuntimeRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetHostRuntimeRuleRequest implement json.Marshaler.
func (m *GetHostRuntimeRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetHostRuntimeRuleResponse implement json.Marshaler.
func (m *GetHostRuntimeRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetHostRuntimeRuleResponse implement json.Marshaler.
func (m *GetHostRuntimeRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CreateHostRuntimeRuleRequest implement json.Marshaler.
func (m *CreateHostRuntimeRuleRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CreateHostRuntimeRuleRequest implement json.Marshaler.
func (m *CreateHostRuntimeRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// CreateHostRuntimeRuleResponse implement json.Marshaler.
func (m *CreateHostRuntimeRuleResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// CreateHostRuntimeRuleResponse implement json.Marshaler.
func (m *CreateHostRuntimeRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetAllHostRuntimeRulesRequest implement json.Marshaler.
func (m *GetAllHostRuntimeRulesRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetAllHostRuntimeRulesRequest implement json.Marshaler.
func (m *GetAllHostRuntimeRulesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetAllHostRuntimeRulesResponse implement json.Marshaler.
func (m *GetAllHostRuntimeRulesResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetAllHostRuntimeRulesResponse implement json.Marshaler.
func (m *GetAllHostRuntimeRulesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceTimeRequest implement json.Marshaler.
func (m *GetDubboInterfaceTimeRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceTimeRequest implement json.Marshaler.
func (m *GetDubboInterfaceTimeRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceTimeResponse implement json.Marshaler.
func (m *GetDubboInterfaceTimeResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceTimeResponse implement json.Marshaler.
func (m *GetDubboInterfaceTimeResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DubboInterfaceTime implement json.Marshaler.
func (m *DubboInterfaceTime) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DubboInterfaceTime implement json.Marshaler.
func (m *DubboInterfaceTime) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceQPSRequest implement json.Marshaler.
func (m *GetDubboInterfaceQPSRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceQPSRequest implement json.Marshaler.
func (m *GetDubboInterfaceQPSRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceQPSResponse implement json.Marshaler.
func (m *GetDubboInterfaceQPSResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceQPSResponse implement json.Marshaler.
func (m *GetDubboInterfaceQPSResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceFailedRequest implement json.Marshaler.
func (m *GetDubboInterfaceFailedRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceFailedRequest implement json.Marshaler.
func (m *GetDubboInterfaceFailedRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceFailedResponse implement json.Marshaler.
func (m *GetDubboInterfaceFailedResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceFailedResponse implement json.Marshaler.
func (m *GetDubboInterfaceFailedResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceAvgTimeRequest implement json.Marshaler.
func (m *GetDubboInterfaceAvgTimeRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceAvgTimeRequest implement json.Marshaler.
func (m *GetDubboInterfaceAvgTimeRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// GetDubboInterfaceAvgTimeResponse implement json.Marshaler.
func (m *GetDubboInterfaceAvgTimeResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// GetDubboInterfaceAvgTimeResponse implement json.Marshaler.
func (m *GetDubboInterfaceAvgTimeResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// Interface implement json.Marshaler.
func (m *Interface) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// Interface implement json.Marshaler.
func (m *Interface) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// InterfaceOwner implement json.Marshaler.
func (m *InterfaceOwner) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// InterfaceOwner implement json.Marshaler.
func (m *InterfaceOwner) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// RequestRule implement json.Marshaler.
func (m *RequestRule) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// RequestRule implement json.Marshaler.
func (m *RequestRule) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HostRules implement json.Marshaler.
func (m *HostRules) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HostRules implement json.Marshaler.
func (m *HostRules) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HostRoute implement json.Marshaler.
func (m *HostRoute) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HostRoute implement json.Marshaler.
func (m *HostRoute) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HostRuntimeRules implement json.Marshaler.
func (m *HostRuntimeRules) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HostRuntimeRules implement json.Marshaler.
func (m *HostRuntimeRules) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HostRuntimeRule implement json.Marshaler.
func (m *HostRuntimeRule) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HostRuntimeRule implement json.Marshaler.
func (m *HostRuntimeRule) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HostRuntimeInterfaces implement json.Marshaler.
func (m *HostRuntimeInterfaces) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HostRuntimeInterfaces implement json.Marshaler.
func (m *HostRuntimeInterfaces) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HTTPServices implement json.Marshaler.
func (m *HTTPServices) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HTTPServices implement json.Marshaler.
func (m *HTTPServices) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HTTPService implement json.Marshaler.
func (m *HTTPService) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HTTPService implement json.Marshaler.
func (m *HTTPService) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// HTTPServiceItem implement json.Marshaler.
func (m *HTTPServiceItem) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// HTTPServiceItem implement json.Marshaler.
func (m *HTTPServiceItem) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DubboInterface implement json.Marshaler.
func (m *DubboInterface) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DubboInterface implement json.Marshaler.
func (m *DubboInterface) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DubboInterfaceResult implement json.Marshaler.
func (m *DubboInterfaceResult) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DubboInterfaceResult implement json.Marshaler.
func (m *DubboInterfaceResult) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DubboMointorMap implement json.Marshaler.
func (m *DubboMointorMap) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DubboMointorMap implement json.Marshaler.
func (m *DubboMointorMap) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// DubboMointor implement json.Marshaler.
func (m *DubboMointor) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// DubboMointor implement json.Marshaler.
func (m *DubboMointor) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// ServiceIpRequest implement json.Marshaler.
func (m *ServiceIpRequest) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// ServiceIpRequest implement json.Marshaler.
func (m *ServiceIpRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}

// ServiceIpInfoResponse implement json.Marshaler.
func (m *ServiceIpInfoResponse) MarshalJSON() ([]byte, error) {
	return (&protojson.MarshalOptions{
		UseProtoNames:   false,
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}).Marshal(m)
}

// ServiceIpInfoResponse implement json.Marshaler.
func (m *ServiceIpInfoResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}).Unmarshal(b, m)
}
