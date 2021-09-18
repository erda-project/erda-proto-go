// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: query.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*LogItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetLogRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetLogByRuntimeRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetLogByOrganizationRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetLogResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetLogByRuntimeResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetLogByOrganizationResponse)(nil)

// LogItem implement urlenc.URLValuesUnmarshaler.
func (m *LogItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "source":
				m.Source = vals[0]
			case "stream":
				m.Stream = vals[0]
			case "timeBucket":
				m.TimeBucket = vals[0]
			case "timestamp":
				m.Timestamp = vals[0]
			case "offset":
				m.Offset = vals[0]
			case "content":
				m.Content = vals[0]
			case "level":
				m.Level = vals[0]
			case "requestId":
				m.RequestId = vals[0]
			case "pattern":
				m.Pattern = vals[0]
			}
		}
	}
	return nil
}

// GetLogRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetLogRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "source":
				m.Source = vals[0]
			case "stream":
				m.Stream = vals[0]
			case "requestId":
				m.RequestId = vals[0]
			case "start":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Start = val
			case "end":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.End = val
			case "count":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Count = val
			case "pattern":
				m.Pattern = vals[0]
			}
		}
	}
	return nil
}

// GetLogByRuntimeRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetLogByRuntimeRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "source":
				m.Source = vals[0]
			case "stream":
				m.Stream = vals[0]
			case "requestId":
				m.RequestId = vals[0]
			case "start":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Start = val
			case "end":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.End = val
			case "count":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Count = val
			case "applicationId":
				m.ApplicationId = vals[0]
			case "pattern":
				m.Pattern = vals[0]
			}
		}
	}
	return nil
}

// GetLogByOrganizationRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetLogByOrganizationRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "source":
				m.Source = vals[0]
			case "stream":
				m.Stream = vals[0]
			case "requestId":
				m.RequestId = vals[0]
			case "start":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Start = val
			case "end":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.End = val
			case "count":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Count = val
			case "clusterName":
				m.ClusterName = vals[0]
			case "pattern":
				m.Pattern = vals[0]
			}
		}
	}
	return nil
}

// GetLogResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetLogResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetLogByRuntimeResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetLogByRuntimeResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetLogByOrganizationResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetLogByOrganizationResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
