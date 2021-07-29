// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: release.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ReleaseCreateRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Resources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
			}
		}
	}
	return nil
}
func (this *ReleaseResource) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ReleaseCreateResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ReleaseCreateResponseData) Validate() error {
	return nil
}
func (this *ReleaseUpdateRequest) Validate() error {
	return nil
}
func (this *ReleaseDataResponse) Validate() error {
	return nil
}
func (this *ReleaseUserDataResponse) Validate() error {
	return nil
}
func (this *ReleaseUpdateResponse) Validate() error {
	return nil
}
func (this *ReleaseReferenceUpdateRequest) Validate() error {
	return nil
}
func (this *GetIosPlistRequest) Validate() error {
	return nil
}
func (this *GetIosPlistResponse) Validate() error {
	return nil
}
func (this *ReleaseDeleteRequest) Validate() error {
	return nil
}
func (this *ReleaseDeleteResponse) Validate() error {
	return nil
}
func (this *ReleaseGetRequest) Validate() error {
	return nil
}
func (this *ReleaseGetResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ReleaseGetResponseData) Validate() error {
	for _, item := range this.Resources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *ReleaseListRequest) Validate() error {
	return nil
}
func (this *ReleaseListResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListReleaseNameRequest) Validate() error {
	return nil
}
func (this *ListReleaseNameResponse) Validate() error {
	return nil
}
func (this *ReleaseListResponseData) Validate() error {
	for _, item := range this.Releases {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Releases", err)
			}
		}
	}
	return nil
}
func (this *ReleaseNameListRequest) Validate() error {
	return nil
}
func (this *ReleaseNameListResponse) Validate() error {
	return nil
}
func (this *GetLatestReleasesRequest) Validate() error {
	return nil
}
func (this *GetLatestReleasesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetLatestReleasesResponseData) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *ReleaseGCRequest) Validate() error {
	return nil
}
