// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: runtime.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetRuntimeRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Resources)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Deployments)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Service)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StatusMap)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Extra)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ErrorResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RuntimeInspect)(nil)

// GetRuntimeRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetRuntimeRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "nameOrID":
				m.NameOrID = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "workspace":
				m.Workspace = vals[0]
			}
		}
	}
	return nil
}

// Resources implement urlenc.URLValuesUnmarshaler.
func (m *Resources) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "cpu":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Cpu = val
			case "mem":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Mem = val
			case "disk":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Disk = val
			}
		}
	}
	return nil
}

// Deployments implement urlenc.URLValuesUnmarshaler.
func (m *Deployments) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "replicas":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Replicas = val
			}
		}
	}
	return nil
}

// Service implement urlenc.URLValuesUnmarshaler.
func (m *Service) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "status":
				m.Status = vals[0]
			case "deployments":
				if m.Deployments == nil {
					m.Deployments = &Deployments{}
				}
			case "deployments.replicas":
				if m.Deployments == nil {
					m.Deployments = &Deployments{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Deployments.Replicas = val
			case "resources":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
			case "resources.cpu":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Resources.Cpu = val
			case "resources.mem":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Resources.Mem = val
			case "resources.disk":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Resources.Disk = val
			case "addrs":
				m.Addrs = vals
			case "expose":
				m.Expose = vals
			}
		}
	}
	return nil
}

// StatusMap implement urlenc.URLValuesUnmarshaler.
func (m *StatusMap) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "msg":
				m.Msg = vals[0]
			case "reason":
				m.Reason = vals[0]
			}
		}
	}
	return nil
}

// Extra implement urlenc.URLValuesUnmarshaler.
func (m *Extra) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "applicationId":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ApplicationId = val
			case "buildId":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.BuildId = val
			case "workspace":
				m.Workspace = vals[0]
			}
		}
	}
	return nil
}

// ErrorResponse implement urlenc.URLValuesUnmarshaler.
func (m *ErrorResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "code":
				m.Code = vals[0]
			case "msg":
				m.Msg = vals[0]
			case "ctx":
				m.Ctx = vals[0]
			}
		}
	}
	return nil
}

// RuntimeInspect implement urlenc.URLValuesUnmarshaler.
func (m *RuntimeInspect) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "name":
				m.Name = vals[0]
			case "serviceGroupName":
				m.ServiceGroupName = vals[0]
			case "serviceGroupNamespace":
				m.ServiceGroupNamespace = vals[0]
			case "source":
				m.Source = vals[0]
			case "status":
				m.Status = vals[0]
			case "deployStatus":
				m.DeployStatus = vals[0]
			case "deleteStatus":
				m.DeleteStatus = vals[0]
			case "releaseId":
				m.ReleaseId = vals[0]
			case "clusterId":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ClusterId = val
			case "clusterName":
				m.ClusterName = vals[0]
			case "clusterType":
				m.ClusterType = vals[0]
			case "resources":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
			case "resources.cpu":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Resources.Cpu = val
			case "resources.mem":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Resources.Mem = val
			case "resources.disk":
				if m.Resources == nil {
					m.Resources = &Resources{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Resources.Disk = val
			case "extra":
				if m.Extra == nil {
					m.Extra = &Extra{}
				}
			case "extra.applicationId":
				if m.Extra == nil {
					m.Extra = &Extra{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Extra.ApplicationId = val
			case "extra.buildId":
				if m.Extra == nil {
					m.Extra = &Extra{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Extra.BuildId = val
			case "extra.workspace":
				if m.Extra == nil {
					m.Extra = &Extra{}
				}
				m.Extra.Workspace = vals[0]
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "timeCreated":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
			case "timeCreated.seconds":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeCreated.Seconds = val
			case "timeCreated.nanos":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeCreated.Nanos = int32(val)
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "updatedAt":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "updatedAt.seconds":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedAt.Seconds = val
			case "updatedAt.nanos":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}
