// Code generated by protoc-gen-form. DO NOT EDIT.
// Source: metric.proto

package pb

import (
	base64 "encoding/base64"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	anypb "google.golang.org/protobuf/types/known/anypb"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*QueryWithInfluxFormatRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryWithInfluxFormatResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Series)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Results)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryWithTableFormatRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryWithTableFormatResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TableResult)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TableColumn)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TableRow)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Filter)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MapValue)(nil)

// QueryWithInfluxFormatRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithInfluxFormatRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			case "q":
				m.Q = vals[0]
			}
		}
	}
	return nil
}

// QueryWithInfluxFormatResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithInfluxFormatResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "results":
				if m.Results == nil {
					m.Results = &Results{}
				}
			case "results.statement_id":
				if m.Results == nil {
					m.Results = &Results{}
				}
				m.Results.StatementId = vals[0]
			}
		}
	}
	return nil
}

// Series implement urlenc.URLValuesUnmarshaler.
func (m *Series) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// Results implement urlenc.URLValuesUnmarshaler.
func (m *Results) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "statement_id":
				m.StatementId = vals[0]
			}
		}
	}
	return nil
}

// QueryWithTableFormatRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithTableFormatRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			}
		}
	}
	return nil
}

// QueryWithTableFormatResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryWithTableFormatResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TableResult{}
				}
			case "data.interval":
				if m.Data == nil {
					m.Data = &TableResult{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Interval = val
			}
		}
	}
	return nil
}

// TableResult implement urlenc.URLValuesUnmarshaler.
func (m *TableResult) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interval":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Interval = val
			}
		}
	}
	return nil
}

// TableColumn implement urlenc.URLValuesUnmarshaler.
func (m *TableColumn) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			case "flag":
				m.Flag = vals[0]
			}
		}
	}
	return nil
}

// TableRow implement urlenc.URLValuesUnmarshaler.
func (m *TableRow) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Filter implement urlenc.URLValuesUnmarshaler.
func (m *Filter) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "op":
				m.Op = vals[0]
			case "value":
				if m.Value == nil {
					m.Value = &anypb.Any{}
				}
			case "value.type_url":
				if m.Value == nil {
					m.Value = &anypb.Any{}
				}
				m.Value.TypeUrl = vals[0]
			case "value.value":
				if m.Value == nil {
					m.Value = &anypb.Any{}
				}
				val, err := base64.StdEncoding.DecodeString(vals[0])
				if err != nil {
					return err
				}
				m.Value.Value = val
			}
		}
	}
	return nil
}

// MapValue implement urlenc.URLValuesUnmarshaler.
func (m *MapValue) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
