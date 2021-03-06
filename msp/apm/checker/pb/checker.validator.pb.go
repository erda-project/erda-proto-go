// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checker.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateCheckerRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateCheckerResponse) Validate() error {
	return nil
}
func (this *UpdateCheckerRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateCheckerResponse) Validate() error {
	return nil
}
func (this *DeleteCheckerRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *DeleteCheckerResponse) Validate() error {
	return nil
}
func (this *ListCheckersRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *ListCheckersResponse) Validate() error {
	return nil
}
func (this *DescribeCheckersRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *DescribeCheckersResponse) Validate() error {
	return nil
}
func (this *DescribeCheckerRequest) Validate() error {
	if this.Scope == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must not be an empty string`, this.Scope))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DescribeCheckerResponse) Validate() error {
	return nil
}
func (this *Checker) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
