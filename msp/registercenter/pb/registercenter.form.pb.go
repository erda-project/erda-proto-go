// Code generated by protoc-gen-form. DO NOT EDIT.
// Source: registercenter.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ListInterfaceRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListInterfaceResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetHTTPServicesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetHTTPServicesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnableHTTPServiceRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnableHTTPServiceResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnableHTTPService)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetRouteRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetRouteRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateRouteRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateRouteRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteRouteRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteRouteRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CetHostRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CetHostRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateHostRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateHostRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteHostRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteHostRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetHostRuntimeRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetHostRuntimeRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateHostRuntimeRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateHostRuntimeRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAllHostRuntimeRulesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAllHostRuntimeRulesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceTimeRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceTimeResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DubboInterfaceTime)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceQPSRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceQPSResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceFailedRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceFailedResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceAvgTimeRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetDubboInterfaceAvgTimeResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Interface)(nil)
var _ urlenc.URLValuesUnmarshaler = (*InterfaceOwner)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RequestRule)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HostRules)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HostRoute)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HostRuntimeRules)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HostRuntimeRule)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HostRuntimeInterfaces)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HTTPServices)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HTTPService)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HTTPServiceItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DubboInterface)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DubboInterfaceResult)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DubboMointorMap)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DubboMointor)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ServiceIpRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ServiceIpInfoResponse)(nil)

// ListInterfaceRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListInterfaceRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "tenantGroup":
				m.TenantGroup = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// ListInterfaceResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListInterfaceResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetHTTPServicesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetHTTPServicesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "nacosID":
				m.NacosID = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantGroup":
				m.TenantGroup = vals[0]
			}
		}
	}
	return nil
}

// GetHTTPServicesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetHTTPServicesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HTTPServices{}
				}
			}
		}
	}
	return nil
}

// EnableHTTPServiceRequest implement urlenc.URLValuesUnmarshaler.
func (m *EnableHTTPServiceRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "nacosID":
				m.NacosID = vals[0]
			case "tenantGroup":
				m.TenantGroup = vals[0]
			case "service":
				if m.Service == nil {
					m.Service = &EnableHTTPService{}
				}
			case "service.serviceName":
				if m.Service == nil {
					m.Service = &EnableHTTPService{}
				}
				m.Service.ServiceName = vals[0]
			case "service.address":
				if m.Service == nil {
					m.Service = &EnableHTTPService{}
				}
				m.Service.Address = vals[0]
			case "service.online":
				if m.Service == nil {
					m.Service = &EnableHTTPService{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Service.Online = val
			case "service.ip":
				if m.Service == nil {
					m.Service = &EnableHTTPService{}
				}
				m.Service.Ip = vals[0]
			case "service.port":
				if m.Service == nil {
					m.Service = &EnableHTTPService{}
				}
				m.Service.Port = vals[0]
			}
		}
	}
	return nil
}

// EnableHTTPServiceResponse implement urlenc.URLValuesUnmarshaler.
func (m *EnableHTTPServiceResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &EnableHTTPService{}
				}
			case "data.serviceName":
				if m.Data == nil {
					m.Data = &EnableHTTPService{}
				}
				m.Data.ServiceName = vals[0]
			case "data.address":
				if m.Data == nil {
					m.Data = &EnableHTTPService{}
				}
				m.Data.Address = vals[0]
			case "data.online":
				if m.Data == nil {
					m.Data = &EnableHTTPService{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Online = val
			case "data.ip":
				if m.Data == nil {
					m.Data = &EnableHTTPService{}
				}
				m.Data.Ip = vals[0]
			case "data.port":
				if m.Data == nil {
					m.Data = &EnableHTTPService{}
				}
				m.Data.Port = vals[0]
			}
		}
	}
	return nil
}

// EnableHTTPService implement urlenc.URLValuesUnmarshaler.
func (m *EnableHTTPService) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "serviceName":
				m.ServiceName = vals[0]
			case "address":
				m.Address = vals[0]
			case "online":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Online = val
			case "ip":
				m.Ip = vals[0]
			case "port":
				m.Port = vals[0]
			}
		}
	}
	return nil
}

// GetRouteRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetRouteRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// GetRouteRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetRouteRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
			case "data.lb_type":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.LbType = vals[0]
			case "data.max_request_per_conn":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxRequestPerConn = vals[0]
			case "data.max_connections":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxConnections = vals[0]
			case "data.max_requests":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxRequests = vals[0]
			}
		}
	}
	return nil
}

// CreateRouteRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateRouteRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			case "rule":
				if m.Rule == nil {
					m.Rule = &RequestRule{}
				}
			case "rule.lb_type":
				if m.Rule == nil {
					m.Rule = &RequestRule{}
				}
				m.Rule.LbType = vals[0]
			case "rule.max_request_per_conn":
				if m.Rule == nil {
					m.Rule = &RequestRule{}
				}
				m.Rule.MaxRequestPerConn = vals[0]
			case "rule.max_connections":
				if m.Rule == nil {
					m.Rule = &RequestRule{}
				}
				m.Rule.MaxConnections = vals[0]
			case "rule.max_requests":
				if m.Rule == nil {
					m.Rule = &RequestRule{}
				}
				m.Rule.MaxRequests = vals[0]
			}
		}
	}
	return nil
}

// CreateRouteRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateRouteRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
			case "data.lb_type":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.LbType = vals[0]
			case "data.max_request_per_conn":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxRequestPerConn = vals[0]
			case "data.max_connections":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxConnections = vals[0]
			case "data.max_requests":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxRequests = vals[0]
			}
		}
	}
	return nil
}

// DeleteRouteRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteRouteRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// DeleteRouteRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteRouteRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
			case "data.lb_type":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.LbType = vals[0]
			case "data.max_request_per_conn":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxRequestPerConn = vals[0]
			case "data.max_connections":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxConnections = vals[0]
			case "data.max_requests":
				if m.Data == nil {
					m.Data = &RequestRule{}
				}
				m.Data.MaxRequests = vals[0]
			}
		}
	}
	return nil
}

// CetHostRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *CetHostRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// CetHostRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *CetHostRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
			case "data.rule":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
			case "data.rule.branch":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
				m.Data.Rule.Branch = vals[0]
			case "data.rule.weight":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Rule.Weight = val
			}
		}
	}
	return nil
}

// CreateHostRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateHostRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			case "rules":
				if m.Rules == nil {
					m.Rules = &HostRules{}
				}
			case "rules.rule":
				if m.Rules == nil {
					m.Rules = &HostRules{}
				}
				if m.Rules.Rule == nil {
					m.Rules.Rule = &HostRoute{}
				}
			case "rules.rule.branch":
				if m.Rules == nil {
					m.Rules = &HostRules{}
				}
				if m.Rules.Rule == nil {
					m.Rules.Rule = &HostRoute{}
				}
				m.Rules.Rule.Branch = vals[0]
			case "rules.rule.weight":
				if m.Rules == nil {
					m.Rules = &HostRules{}
				}
				if m.Rules.Rule == nil {
					m.Rules.Rule = &HostRoute{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Rules.Rule.Weight = val
			}
		}
	}
	return nil
}

// CreateHostRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateHostRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
			case "data.rule":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
			case "data.rule.branch":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
				m.Data.Rule.Branch = vals[0]
			case "data.rule.weight":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Rule.Weight = val
			}
		}
	}
	return nil
}

// DeleteHostRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteHostRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// DeleteHostRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteHostRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
			case "data.rule":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
			case "data.rule.branch":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
				m.Data.Rule.Branch = vals[0]
			case "data.rule.weight":
				if m.Data == nil {
					m.Data = &HostRules{}
				}
				if m.Data.Rule == nil {
					m.Data.Rule = &HostRoute{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Rule.Weight = val
			}
		}
	}
	return nil
}

// GetHostRuntimeRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetHostRuntimeRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "host":
				m.Host = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// GetHostRuntimeRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetHostRuntimeRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HostRuntimeRules{}
				}
			case "data.address":
				if m.Data == nil {
					m.Data = &HostRuntimeRules{}
				}
				m.Data.Address = vals[0]
			}
		}
	}
	return nil
}

// CreateHostRuntimeRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateHostRuntimeRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "host":
				m.Host = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			case "rules":
				if m.Rules == nil {
					m.Rules = &HostRuntimeRules{}
				}
			case "rules.address":
				if m.Rules == nil {
					m.Rules = &HostRuntimeRules{}
				}
				m.Rules.Address = vals[0]
			}
		}
	}
	return nil
}

// CreateHostRuntimeRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateHostRuntimeRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HostRuntimeRules{}
				}
			case "data.address":
				if m.Data == nil {
					m.Data = &HostRuntimeRules{}
				}
				m.Data.Address = vals[0]
			}
		}
	}
	return nil
}

// GetAllHostRuntimeRulesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetAllHostRuntimeRulesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// GetAllHostRuntimeRulesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetAllHostRuntimeRulesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &HostRuntimeInterfaces{}
				}
			case "data.appid":
				if m.Data == nil {
					m.Data = &HostRuntimeInterfaces{}
				}
				m.Data.Appid = vals[0]
			}
		}
	}
	return nil
}

// GetDubboInterfaceTimeRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceTimeRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			}
		}
	}
	return nil
}

// GetDubboInterfaceTimeResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceTimeResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &DubboInterfaceTime{}
				}
			case "data.providerTime":
				if m.Data == nil {
					m.Data = &DubboInterfaceTime{}
				}
				m.Data.ProviderTime = vals[0]
			case "data.consumerTime":
				if m.Data == nil {
					m.Data = &DubboInterfaceTime{}
				}
				m.Data.ConsumerTime = vals[0]
			}
		}
	}
	return nil
}

// DubboInterfaceTime implement urlenc.URLValuesUnmarshaler.
func (m *DubboInterfaceTime) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "providerTime":
				m.ProviderTime = vals[0]
			case "consumerTime":
				m.ConsumerTime = vals[0]
			}
		}
	}
	return nil
}

// GetDubboInterfaceQPSRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceQPSRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			case "point":
				m.Point = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			}
		}
	}
	return nil
}

// GetDubboInterfaceQPSResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceQPSResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
			case "data.title":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
				m.Data.Title = vals[0]
			case "data.total":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Total = val
			}
		}
	}
	return nil
}

// GetDubboInterfaceFailedRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceFailedRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			case "point":
				m.Point = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			}
		}
	}
	return nil
}

// GetDubboInterfaceFailedResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceFailedResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
			case "data.title":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
				m.Data.Title = vals[0]
			case "data.total":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Total = val
			}
		}
	}
	return nil
}

// GetDubboInterfaceAvgTimeRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceAvgTimeRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfaceName":
				m.InterfaceName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "env":
				m.Env = vals[0]
			case "start":
				m.Start = vals[0]
			case "end":
				m.End = vals[0]
			case "point":
				m.Point = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "appID":
				m.AppID = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			}
		}
	}
	return nil
}

// GetDubboInterfaceAvgTimeResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetDubboInterfaceAvgTimeResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
			case "data.title":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
				m.Data.Title = vals[0]
			case "data.total":
				if m.Data == nil {
					m.Data = &DubboInterface{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Total = val
			}
		}
	}
	return nil
}

// Interface implement urlenc.URLValuesUnmarshaler.
func (m *Interface) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "interfacename":
				m.Interfacename = vals[0]
			}
		}
	}
	return nil
}

// InterfaceOwner implement urlenc.URLValuesUnmarshaler.
func (m *InterfaceOwner) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ip":
				m.Ip = vals[0]
			case "owner":
				m.Owner = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "env":
				m.Env = vals[0]
			case "hostIp":
				m.HostIp = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			case "feature":
				m.Feature = vals[0]
			case "serviceName":
				m.ServiceName = vals[0]
			}
		}
	}
	return nil
}

// RequestRule implement urlenc.URLValuesUnmarshaler.
func (m *RequestRule) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "lb_type":
				m.LbType = vals[0]
			case "max_request_per_conn":
				m.MaxRequestPerConn = vals[0]
			case "max_connections":
				m.MaxConnections = vals[0]
			case "max_requests":
				m.MaxRequests = vals[0]
			}
		}
	}
	return nil
}

// HostRules implement urlenc.URLValuesUnmarshaler.
func (m *HostRules) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "rule":
				if m.Rule == nil {
					m.Rule = &HostRoute{}
				}
			case "rule.branch":
				if m.Rule == nil {
					m.Rule = &HostRoute{}
				}
				m.Rule.Branch = vals[0]
			case "rule.weight":
				if m.Rule == nil {
					m.Rule = &HostRoute{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Rule.Weight = val
			}
		}
	}
	return nil
}

// HostRoute implement urlenc.URLValuesUnmarshaler.
func (m *HostRoute) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "branch":
				m.Branch = vals[0]
			case "weight":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Weight = val
			}
		}
	}
	return nil
}

// HostRuntimeRules implement urlenc.URLValuesUnmarshaler.
func (m *HostRuntimeRules) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "address":
				m.Address = vals[0]
			}
		}
	}
	return nil
}

// HostRuntimeRule implement urlenc.URLValuesUnmarshaler.
func (m *HostRuntimeRule) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "weight":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Weight = val
			}
		}
	}
	return nil
}

// HostRuntimeInterfaces implement urlenc.URLValuesUnmarshaler.
func (m *HostRuntimeInterfaces) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "appid":
				m.Appid = vals[0]
			}
		}
	}
	return nil
}

// HTTPServices implement urlenc.URLValuesUnmarshaler.
func (m *HTTPServices) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// HTTPService implement urlenc.URLValuesUnmarshaler.
func (m *HTTPService) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "serviceName":
				m.ServiceName = vals[0]
			case "serviceDomain":
				m.ServiceDomain = vals[0]
			}
		}
	}
	return nil
}

// HTTPServiceItem implement urlenc.URLValuesUnmarshaler.
func (m *HTTPServiceItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "address":
				m.Address = vals[0]
			case "online":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Online = val
			}
		}
	}
	return nil
}

// DubboInterface implement urlenc.URLValuesUnmarshaler.
func (m *DubboInterface) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "title":
				m.Title = vals[0]
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

// DubboInterfaceResult implement urlenc.URLValuesUnmarshaler.
func (m *DubboInterfaceResult) UnmarshalURLValues(prefix string, values url.Values) error {
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

// DubboMointorMap implement urlenc.URLValuesUnmarshaler.
func (m *DubboMointorMap) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "value":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
			case "value.name":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
				m.Value.Name = vals[0]
			case "value.tag":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
				m.Value.Tag = vals[0]
			case "value.unit":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
				m.Value.Unit = vals[0]
			case "value.unitType":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
				m.Value.UnitType = vals[0]
			case "value.chartType":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
				m.Value.ChartType = vals[0]
			case "value.axisIndex":
				if m.Value == nil {
					m.Value = &DubboMointor{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Value.AxisIndex = val
			}
		}
	}
	return nil
}

// DubboMointor implement urlenc.URLValuesUnmarshaler.
func (m *DubboMointor) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "tag":
				m.Tag = vals[0]
			case "unit":
				m.Unit = vals[0]
			case "unitType":
				m.UnitType = vals[0]
			case "chartType":
				m.ChartType = vals[0]
			case "axisIndex":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.AxisIndex = val
			}
		}
	}
	return nil
}

// ServiceIpRequest implement urlenc.URLValuesUnmarshaler.
func (m *ServiceIpRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				m.ProjectID = vals[0]
			case "workspace":
				m.Workspace = vals[0]
			case "ip":
				m.Ip = vals[0]
			}
		}
	}
	return nil
}

// ServiceIpInfoResponse implement urlenc.URLValuesUnmarshaler.
func (m *ServiceIpInfoResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "appID":
				m.AppID = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			case "serviceName":
				m.ServiceName = vals[0]
			}
		}
	}
	return nil
}
