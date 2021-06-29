// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checker_v1.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateCheckerV1Request) Validate() error {
	if !(this.ProjectID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ProjectID))
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
func (this *CreateCheckerV1Response) Validate() error {
	return nil
}
func (this *UpdateCheckerV1Request) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
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
func (this *UpdateCheckerV1Response) Validate() error {
	return nil
}
func (this *DeleteCheckerV1Request) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteCheckerV1Response) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DescribeCheckersV1Request) Validate() error {
	if !(this.ProjectID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ProjectID))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *DescribeCheckersV1Response) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DescribeCheckerV1Request) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DescribeCheckerV1Response) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetCheckerStatusV1Request) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetCheckerStatusV1Response) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CheckerStatusV1) Validate() error {
	return nil
}
func (this *GetCheckerIssuesV1Request) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetCheckerIssuesV1Response) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CheckerV1) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Mode == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Mode", fmt.Errorf(`value '%v' must not be an empty string`, this.Mode))
	}
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`value '%v' must not be an empty string`, this.Url))
	}
	if this.Env == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Env", fmt.Errorf(`value '%v' must not be an empty string`, this.Env))
	}
	return nil
}
func (this *DescribeResultV1) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *DescribeItemV1) Validate() error {
	if this.Chart != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Chart); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Chart", err)
		}
	}
	return nil
}
func (this *CheckerChartV1) Validate() error {
	return nil
}
