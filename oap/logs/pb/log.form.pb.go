// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: log.proto

package pb

import (
	base64 "encoding/base64"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/oap/common/pb"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*Log)(nil)

// Log implement urlenc.URLValuesUnmarshaler.
func (m *Log) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "timeUnixNano":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUnixNano = val
			case "name":
				m.Name = vals[0]
			case "severity":
				m.Severity = vals[0]
			case "relations":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
			case "relations.traceID":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				val, err := base64.StdEncoding.DecodeString(vals[0])
				if err != nil {
					return err
				}
				m.Relations.TraceID = val
			case "relations.resID":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.ResID = vals[0]
			case "relations.resType":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.ResType = vals[0]
			case "relations.resourceKeys":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.ResourceKeys = vals
			case "content":
				val, err := base64.StdEncoding.DecodeString(vals[0])
				if err != nil {
					return err
				}
				m.Content = val
			}
		}
	}
	return nil
}
