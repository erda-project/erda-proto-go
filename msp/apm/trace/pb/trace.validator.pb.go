// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/descriptorpb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetTraceDebugRequest) Validate() error {
	if this.RequestID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RequestID", fmt.Errorf(`value '%v' must not be an empty string`, this.RequestID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *CreateTraceDebugRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *StopTraceDebugRequest) Validate() error {
	if this.RequestID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RequestID", fmt.Errorf(`value '%v' must not be an empty string`, this.RequestID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetTraceDebugStatusByRequestIDRequest) Validate() error {
	if this.RequestID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RequestID", fmt.Errorf(`value '%v' must not be an empty string`, this.RequestID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetTraceDebugHistoriesRequest) Validate() error {
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetSpansRequest) Validate() error {
	if this.TraceID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TraceID", fmt.Errorf(`value '%v' must not be an empty string`, this.TraceID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetTracesRequest) Validate() error {
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetSpansResponse) Validate() error {
	for _, item := range this.Spans {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Spans", err)
			}
		}
	}
	return nil
}
func (this *GetTracesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTraceDebugHistoriesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTraceDebugResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTraceDebugResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *StopTraceDebugResponse) Validate() error {
	return nil
}
func (this *GetTraceDebugStatusByRequestIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TraceDebug) Validate() error {
	for _, item := range this.History {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("History", err)
			}
		}
	}
	return nil
}
func (this *TraceDebugStatus) Validate() error {
	return nil
}
func (this *TraceDebugHistory) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Span) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Trace) Validate() error {
	return nil
}
