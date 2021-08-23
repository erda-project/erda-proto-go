// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: entity.proto

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
var _ urlenc.URLValuesUnmarshaler = (*EntityRow)(nil)

// EntityRow implement urlenc.URLValuesUnmarshaler.
func (m *EntityRow) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "entityID":
				m.EntityID = vals[0]
			case "table":
				m.Table = vals[0]
			case "rowID":
				val, err := base64.StdEncoding.DecodeString(vals[0])
				if err != nil {
					return err
				}
				m.RowID = val
			case "rowData":
				if m.RowData == nil {
					m.RowData = &pb.KeyValueList{}
				}
			case "createTimeUnixNano":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreateTimeUnixNano = val
			case "updateTimeUnixNano":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdateTimeUnixNano = val
			}
		}
	}
	return nil
}
