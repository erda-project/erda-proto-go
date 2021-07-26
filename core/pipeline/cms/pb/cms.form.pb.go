// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: cms.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*PipelineCmsNs)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineCmsConfigValue)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineCmsConfigOperations)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineCmsConfigKey)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineCmsConfig)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsCreateNsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsCreateNsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsListNsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsListNsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsUpdateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsUpdateV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsUpdateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsDeleteRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsDeleteResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsGetRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CmsNsConfigsGetResponse)(nil)

// PipelineCmsNs implement urlenc.URLValuesUnmarshaler.
func (m *PipelineCmsNs) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pipeline_source":
				m.PipelineSource = vals[0]
			case "ns":
				m.Ns = vals[0]
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
			case "timeUpdated":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "timeUpdated.seconds":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUpdated.Seconds = val
			case "timeUpdated.nanos":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeUpdated.Nanos = int32(val)
			}
		}
	}
	return nil
}

// PipelineCmsConfigValue implement urlenc.URLValuesUnmarshaler.
func (m *PipelineCmsConfigValue) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "value":
				m.Value = vals[0]
			case "encryptInDB":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.EncryptInDB = val
			case "type":
				m.Type = vals[0]
			case "operations":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
			case "operations.canDownload":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Operations.CanDownload = val
			case "operations.canEdit":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Operations.CanEdit = val
			case "operations.canDelete":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Operations.CanDelete = val
			case "comment":
				m.Comment = vals[0]
			case "from":
				m.From = vals[0]
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
			case "timeUpdated":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "timeUpdated.seconds":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUpdated.Seconds = val
			case "timeUpdated.nanos":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeUpdated.Nanos = int32(val)
			}
		}
	}
	return nil
}

// PipelineCmsConfigOperations implement urlenc.URLValuesUnmarshaler.
func (m *PipelineCmsConfigOperations) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "canDownload":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.CanDownload = val
			case "canEdit":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.CanEdit = val
			case "canDelete":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.CanDelete = val
			}
		}
	}
	return nil
}

// PipelineCmsConfigKey implement urlenc.URLValuesUnmarshaler.
func (m *PipelineCmsConfigKey) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "decrypt":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Decrypt = val
			case "showEncryptedValue":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.ShowEncryptedValue = val
			}
		}
	}
	return nil
}

// PipelineCmsConfig implement urlenc.URLValuesUnmarshaler.
func (m *PipelineCmsConfig) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "value":
				m.Value = vals[0]
			case "encryptInDB":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.EncryptInDB = val
			case "type":
				m.Type = vals[0]
			case "operations":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
			case "operations.canDownload":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Operations.CanDownload = val
			case "operations.canEdit":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Operations.CanEdit = val
			case "operations.canDelete":
				if m.Operations == nil {
					m.Operations = &PipelineCmsConfigOperations{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Operations.CanDelete = val
			case "comment":
				m.Comment = vals[0]
			case "from":
				m.From = vals[0]
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
			case "timeUpdated":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "timeUpdated.seconds":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUpdated.Seconds = val
			case "timeUpdated.nanos":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeUpdated.Nanos = int32(val)
			}
		}
	}
	return nil
}

// CmsCreateNsRequest implement urlenc.URLValuesUnmarshaler.
func (m *CmsCreateNsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "ns":
				m.Ns = vals[0]
			}
		}
	}
	return nil
}

// CmsCreateNsResponse implement urlenc.URLValuesUnmarshaler.
func (m *CmsCreateNsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CmsListNsRequest implement urlenc.URLValuesUnmarshaler.
func (m *CmsListNsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "nsPrefix":
				m.NsPrefix = vals[0]
			}
		}
	}
	return nil
}

// CmsListNsResponse implement urlenc.URLValuesUnmarshaler.
func (m *CmsListNsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CmsNsConfigsUpdateRequest implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsUpdateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ns":
				m.Ns = vals[0]
			case "pipeline_source":
				m.PipelineSource = vals[0]
			}
		}
	}
	return nil
}

// CmsNsConfigsUpdateV1Request implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsUpdateV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ns":
				m.Ns = vals[0]
			}
		}
	}
	return nil
}

// CmsNsConfigsUpdateResponse implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsUpdateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CmsNsConfigsDeleteRequest implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsDeleteRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ns":
				m.Ns = vals[0]
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "deleteNs":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.DeleteNs = val
			case "deleteForce":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.DeleteForce = val
			case "deleteKeys":
				m.DeleteKeys = vals
			}
		}
	}
	return nil
}

// CmsNsConfigsDeleteResponse implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsDeleteResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CmsNsConfigsGetRequest implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsGetRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ns":
				m.Ns = vals[0]
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "globalDecrypt":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.GlobalDecrypt = val
			}
		}
	}
	return nil
}

// CmsNsConfigsGetResponse implement urlenc.URLValuesUnmarshaler.
func (m *CmsNsConfigsGetResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
