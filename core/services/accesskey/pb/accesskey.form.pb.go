// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: accesskey.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*QueryAccessKeysRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryAccessKeysResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAccessKeysRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAccessKeysResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateAccessKeysRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateAccessKeysResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateAccessKeysRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateAccessKeysResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteAccessKeysRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteAccessKeysResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*AccessKeysItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GenericEnum)(nil)

// QueryAccessKeysRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryAccessKeysRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "status":
			case "subjectType":
			case "subject":
				m.Subject = &vals[0]
			case "accessKey":
				m.AccessKey = &vals[0]
			}
		}
	}
	return nil
}

// QueryAccessKeysResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryAccessKeysResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetAccessKeysRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetAccessKeysRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// GetAccessKeysResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetAccessKeysResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.Id = vals[0]
			case "data.accessKey":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.AccessKey = vals[0]
			case "data.secretKey":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.SecretKey = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
			case "data.subjectType":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
			case "data.subject":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.Subject = vals[0]
			case "data.description":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.Description = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// CreateAccessKeysRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateAccessKeysRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "subjectType":
			case "subject":
				m.Subject = vals[0]
			case "description":
				m.Description = vals[0]
			}
		}
	}
	return nil
}

// CreateAccessKeysResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateAccessKeysResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.Id = vals[0]
			case "data.accessKey":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.AccessKey = vals[0]
			case "data.secretKey":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.SecretKey = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
			case "data.subjectType":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
			case "data.subject":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.Subject = vals[0]
			case "data.description":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				m.Data.Description = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &AccessKeysItem{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// UpdateAccessKeysRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateAccessKeysRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "status":
			case "description":
				m.Description = &vals[0]
			}
		}
	}
	return nil
}

// UpdateAccessKeysResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateAccessKeysResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// DeleteAccessKeysRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteAccessKeysRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// DeleteAccessKeysResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteAccessKeysResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// AccessKeysItem implement urlenc.URLValuesUnmarshaler.
func (m *AccessKeysItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "accessKey":
				m.AccessKey = vals[0]
			case "secretKey":
				m.SecretKey = vals[0]
			case "status":
			case "subjectType":
			case "subject":
				m.Subject = vals[0]
			case "description":
				m.Description = vals[0]
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
			}
		}
	}
	return nil
}

// GenericEnum implement urlenc.URLValuesUnmarshaler.
func (m *GenericEnum) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
