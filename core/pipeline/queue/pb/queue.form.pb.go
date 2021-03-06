// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: queue.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/common/pb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*QueueDeleteRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueDeleteResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueUpdateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueUpdateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueuePagingRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueuePagingResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueGetRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueGetResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueCreateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueCreateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Queue)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueUsage)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueueUsageItem)(nil)

// QueueDeleteRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueueDeleteRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "queueID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.QueueID = val
			}
		}
	}
	return nil
}

// QueueDeleteResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueueDeleteResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// QueueUpdateRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueueUpdateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "queueID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.QueueID = val
			case "name":
				m.Name = vals[0]
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "scheduleStrategy":
				m.ScheduleStrategy = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "priority":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Priority = val
			case "concurrency":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Concurrency = val
			case "maxCPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxCPU = val
			case "maxMemoryMB":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxMemoryMB = val
			case "identityInfo":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
			case "identityInfo.userID":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.UserID = vals[0]
			case "identityInfo.internalClient":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.InternalClient = vals[0]
			case "identityInfo.orgID":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.OrgID = vals[0]
			}
		}
	}
	return nil
}

// QueueUpdateResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueueUpdateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Queue{}
				}
			case "data.ID":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ID = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.Name = vals[0]
			case "data.pipelineSource":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.PipelineSource = vals[0]
			case "data.clusterName":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.ClusterName = vals[0]
			case "data.scheduleStrategy":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.ScheduleStrategy = vals[0]
			case "data.mode":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.Mode = vals[0]
			case "data.priority":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Priority = val
			case "data.concurrency":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Concurrency = val
			case "data.maxCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MaxCPU = val
			case "data.maxMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MaxMemoryMB = val
			case "data.timeCreated":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
			case "data.timeCreated.seconds":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Seconds = val
			case "data.timeCreated.nanos":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Nanos = int32(val)
			case "data.timeUpdated":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "data.timeUpdated.seconds":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Seconds = val
			case "data.timeUpdated.nanos":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Nanos = int32(val)
			case "data.usage":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
			case "data.usage.inUseCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.InUseCPU = val
			case "data.usage.inUseMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.InUseMemoryMB = val
			case "data.usage.remainingCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.RemainingCPU = val
			case "data.usage.remainingMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.RemainingMemoryMB = val
			case "data.usage.processingCount":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Usage.ProcessingCount = val
			case "data.usage.pendingCount":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Usage.PendingCount = val
			}
		}
	}
	return nil
}

// QueuePagingRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueuePagingRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "pipelineSources":
				m.PipelineSources = vals
			case "clusterName":
				m.ClusterName = vals[0]
			case "scheduleStrategy":
				m.ScheduleStrategy = vals[0]
			case "priority":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Priority = val
			case "concurrency":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Concurrency = val
			case "mustMatchLabels":
				m.MustMatchLabels = vals
			case "anyMatchLabels":
				m.AnyMatchLabels = vals
			case "allowNoPipelineSources":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.AllowNoPipelineSources = val
			case "orderByTargetIDAsc":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.OrderByTargetIDAsc = val
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			}
		}
	}
	return nil
}

// QueuePagingResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueuePagingResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Total = val
			}
		}
	}
	return nil
}

// QueueGetRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueueGetRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "queueID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.QueueID = val
			}
		}
	}
	return nil
}

// QueueGetResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueueGetResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Queue{}
				}
			case "data.ID":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ID = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.Name = vals[0]
			case "data.pipelineSource":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.PipelineSource = vals[0]
			case "data.clusterName":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.ClusterName = vals[0]
			case "data.scheduleStrategy":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.ScheduleStrategy = vals[0]
			case "data.mode":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.Mode = vals[0]
			case "data.priority":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Priority = val
			case "data.concurrency":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Concurrency = val
			case "data.maxCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MaxCPU = val
			case "data.maxMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MaxMemoryMB = val
			case "data.timeCreated":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
			case "data.timeCreated.seconds":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Seconds = val
			case "data.timeCreated.nanos":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Nanos = int32(val)
			case "data.timeUpdated":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "data.timeUpdated.seconds":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Seconds = val
			case "data.timeUpdated.nanos":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Nanos = int32(val)
			case "data.usage":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
			case "data.usage.inUseCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.InUseCPU = val
			case "data.usage.inUseMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.InUseMemoryMB = val
			case "data.usage.remainingCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.RemainingCPU = val
			case "data.usage.remainingMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.RemainingMemoryMB = val
			case "data.usage.processingCount":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Usage.ProcessingCount = val
			case "data.usage.pendingCount":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Usage.PendingCount = val
			}
		}
	}
	return nil
}

// QueueCreateRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueueCreateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "scheduleStrategy":
				m.ScheduleStrategy = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "priority":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Priority = val
			case "concurrency":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Concurrency = val
			case "maxCPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxCPU = val
			case "maxMemoryMB":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxMemoryMB = val
			case "identityInfo":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
			case "identityInfo.userID":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.UserID = vals[0]
			case "identityInfo.internalClient":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.InternalClient = vals[0]
			case "identityInfo.orgID":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.OrgID = vals[0]
			}
		}
	}
	return nil
}

// QueueCreateResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueueCreateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Queue{}
				}
			case "data.ID":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ID = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.Name = vals[0]
			case "data.pipelineSource":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.PipelineSource = vals[0]
			case "data.clusterName":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.ClusterName = vals[0]
			case "data.scheduleStrategy":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.ScheduleStrategy = vals[0]
			case "data.mode":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				m.Data.Mode = vals[0]
			case "data.priority":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Priority = val
			case "data.concurrency":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Concurrency = val
			case "data.maxCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MaxCPU = val
			case "data.maxMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MaxMemoryMB = val
			case "data.timeCreated":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
			case "data.timeCreated.seconds":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Seconds = val
			case "data.timeCreated.nanos":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Nanos = int32(val)
			case "data.timeUpdated":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "data.timeUpdated.seconds":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Seconds = val
			case "data.timeUpdated.nanos":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Nanos = int32(val)
			case "data.usage":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
			case "data.usage.inUseCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.InUseCPU = val
			case "data.usage.inUseMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.InUseMemoryMB = val
			case "data.usage.remainingCPU":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.RemainingCPU = val
			case "data.usage.remainingMemoryMB":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.Usage.RemainingMemoryMB = val
			case "data.usage.processingCount":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Usage.ProcessingCount = val
			case "data.usage.pendingCount":
				if m.Data == nil {
					m.Data = &Queue{}
				}
				if m.Data.Usage == nil {
					m.Data.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Usage.PendingCount = val
			}
		}
	}
	return nil
}

// Queue implement urlenc.URLValuesUnmarshaler.
func (m *Queue) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ID = val
			case "name":
				m.Name = vals[0]
			case "pipelineSource":
				m.PipelineSource = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "scheduleStrategy":
				m.ScheduleStrategy = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "priority":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Priority = val
			case "concurrency":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Concurrency = val
			case "maxCPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxCPU = val
			case "maxMemoryMB":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxMemoryMB = val
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
			case "usage":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
			case "usage.inUseCPU":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Usage.InUseCPU = val
			case "usage.inUseMemoryMB":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Usage.InUseMemoryMB = val
			case "usage.remainingCPU":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Usage.RemainingCPU = val
			case "usage.remainingMemoryMB":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Usage.RemainingMemoryMB = val
			case "usage.processingCount":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Usage.ProcessingCount = val
			case "usage.pendingCount":
				if m.Usage == nil {
					m.Usage = &QueueUsage{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Usage.PendingCount = val
			}
		}
	}
	return nil
}

// QueueUsage implement urlenc.URLValuesUnmarshaler.
func (m *QueueUsage) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "inUseCPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.InUseCPU = val
			case "inUseMemoryMB":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.InUseMemoryMB = val
			case "remainingCPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.RemainingCPU = val
			case "remainingMemoryMB":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.RemainingMemoryMB = val
			case "processingCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProcessingCount = val
			case "pendingCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PendingCount = val
			}
		}
	}
	return nil
}

// QueueUsageItem implement urlenc.URLValuesUnmarshaler.
func (m *QueueUsageItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pipelineID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PipelineID = val
			case "requestsCPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.RequestsCPU = val
			case "requestsMemoryMB":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.RequestsMemoryMB = val
			case "index":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Index = val
			case "priority":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Priority = val
			case "addedTime":
				if m.AddedTime == nil {
					m.AddedTime = &timestamppb.Timestamp{}
				}
			case "addedTime.seconds":
				if m.AddedTime == nil {
					m.AddedTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.AddedTime.Seconds = val
			case "addedTime.nanos":
				if m.AddedTime == nil {
					m.AddedTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.AddedTime.Nanos = int32(val)
			}
		}
	}
	return nil
}
