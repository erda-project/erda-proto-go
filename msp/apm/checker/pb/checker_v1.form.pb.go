// Code generated by protoc-gen-form. DO NOT EDIT.
// Source: checker_v1.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CreateCheckerV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateCheckerV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateCheckerV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateCheckerV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteCheckerV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteCheckerV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckersV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckersV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckerV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckerV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetCheckerStatusV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetCheckerStatusV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CheckerStatusV1)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetCheckerIssuesV1Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetCheckerIssuesV1Response)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CheckerV1)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeResultV1)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeItemV1)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CheckerChartV1)(nil)

// CreateCheckerV1Request implement urlenc.URLValuesUnmarshaler.
func (m *CreateCheckerV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "data":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
			case "data.name":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Name = vals[0]
			case "data.mode":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Mode = vals[0]
			case "data.url":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Url = vals[0]
			case "data.env":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Env = vals[0]
			}
		}
	}
	return nil
}

// CreateCheckerV1Response implement urlenc.URLValuesUnmarshaler.
func (m *CreateCheckerV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// UpdateCheckerV1Request implement urlenc.URLValuesUnmarshaler.
func (m *UpdateCheckerV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "data":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
			case "data.name":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Name = vals[0]
			case "data.mode":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Mode = vals[0]
			case "data.url":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Url = vals[0]
			case "data.env":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Env = vals[0]
			}
		}
	}
	return nil
}

// UpdateCheckerV1Response implement urlenc.URLValuesUnmarshaler.
func (m *UpdateCheckerV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// DeleteCheckerV1Request implement urlenc.URLValuesUnmarshaler.
func (m *DeleteCheckerV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// DeleteCheckerV1Response implement urlenc.URLValuesUnmarshaler.
func (m *DeleteCheckerV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
			case "data.name":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Name = vals[0]
			case "data.mode":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Mode = vals[0]
			case "data.url":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Url = vals[0]
			case "data.env":
				if m.Data == nil {
					m.Data = &CheckerV1{}
				}
				m.Data.Env = vals[0]
			}
		}
	}
	return nil
}

// DescribeCheckersV1Request implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckersV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "env":
				m.Env = vals[0]
			}
		}
	}
	return nil
}

// DescribeCheckersV1Response implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckersV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &DescribeResultV1{}
				}
			case "data.downCount":
				if m.Data == nil {
					m.Data = &DescribeResultV1{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.DownCount = val
			}
		}
	}
	return nil
}

// DescribeCheckerV1Request implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckerV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "period":
				m.Period = vals[0]
			}
		}
	}
	return nil
}

// DescribeCheckerV1Response implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckerV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &DescribeResultV1{}
				}
			case "data.downCount":
				if m.Data == nil {
					m.Data = &DescribeResultV1{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.DownCount = val
			}
		}
	}
	return nil
}

// GetCheckerStatusV1Request implement urlenc.URLValuesUnmarshaler.
func (m *GetCheckerStatusV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// GetCheckerStatusV1Response implement urlenc.URLValuesUnmarshaler.
func (m *GetCheckerStatusV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &CheckerStatusV1{}
				}
			case "data.time":
				if m.Data == nil {
					m.Data = &CheckerStatusV1{}
				}
				list := make([]int64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseInt(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Data.Time = list
			case "data.status":
				if m.Data == nil {
					m.Data = &CheckerStatusV1{}
				}
				m.Data.Status = vals
			}
		}
	}
	return nil
}

// CheckerStatusV1 implement urlenc.URLValuesUnmarshaler.
func (m *CheckerStatusV1) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "time":
				list := make([]int64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseInt(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Time = list
			case "status":
				m.Status = vals
			}
		}
	}
	return nil
}

// GetCheckerIssuesV1Request implement urlenc.URLValuesUnmarshaler.
func (m *GetCheckerIssuesV1Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// GetCheckerIssuesV1Response implement urlenc.URLValuesUnmarshaler.
func (m *GetCheckerIssuesV1Response) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CheckerV1 implement urlenc.URLValuesUnmarshaler.
func (m *CheckerV1) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "url":
				m.Url = vals[0]
			case "env":
				m.Env = vals[0]
			}
		}
	}
	return nil
}

// DescribeResultV1 implement urlenc.URLValuesUnmarshaler.
func (m *DescribeResultV1) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "downCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.DownCount = val
			}
		}
	}
	return nil
}

// DescribeItemV1 implement urlenc.URLValuesUnmarshaler.
func (m *DescribeItemV1) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "url":
				m.Url = vals[0]
			case "status":
				m.Status = vals[0]
			case "uptime":
				m.Uptime = vals[0]
			case "downtime":
				m.Downtime = vals[0]
			case "downDuration":
				m.DownDuration = vals[0]
			case "latency":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Latency = val
			case "requestId":
				m.RequestId = vals[0]
			case "apdex":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Apdex = val
			case "avg":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Avg = val
			case "max":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Max = val
			case "min":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Min = val
			case "chart":
				if m.Chart == nil {
					m.Chart = &CheckerChartV1{}
				}
			case "chart.latency":
				if m.Chart == nil {
					m.Chart = &CheckerChartV1{}
				}
				list := make([]float64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseFloat(text, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Chart.Latency = list
			case "chart.time":
				if m.Chart == nil {
					m.Chart = &CheckerChartV1{}
				}
				list := make([]int64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseInt(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Chart.Time = list
			case "chart.status":
				if m.Chart == nil {
					m.Chart = &CheckerChartV1{}
				}
				m.Chart.Status = vals
			}
		}
	}
	return nil
}

// CheckerChartV1 implement urlenc.URLValuesUnmarshaler.
func (m *CheckerChartV1) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "latency":
				list := make([]float64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseFloat(text, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Latency = list
			case "time":
				list := make([]int64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseInt(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Time = list
			case "status":
				m.Status = vals
			}
		}
	}
	return nil
}
