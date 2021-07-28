// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: details.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QuerySystemPodMetricsRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.ClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.ClusterName))
	}
	if !(this.Timestamp > 1800000) {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf(`value '%v' must be greater than '1800000'`, this.Timestamp))
	}
	return nil
}
func (this *QuerySystemPodMetricsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PodInfo) Validate() error {
	if this.Summary != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Summary); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Summary", err)
		}
	}
	for _, item := range this.Instances {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Instances", err)
			}
		}
	}
	return nil
}
func (this *PodInfoSummary) Validate() error {
	return nil
}
func (this *PodInfoInstanse) Validate() error {
	return nil
}
