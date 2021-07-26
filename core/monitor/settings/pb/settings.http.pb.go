// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: settings.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// SettingsServiceHandler is the server API for SettingsService service.
type SettingsServiceHandler interface {
	// GET /api/global/settings
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	// PUT /api/global/settings
	PutSettings(context.Context, *PutSettingsRequest) (*PutSettingsResponse, error)
	// PUT /api/config/register
	RegisterMonitorConfig(context.Context, *RegisterMonitorConfigRequest) (*RegisterMonitorConfigResponse, error)
}

// RegisterSettingsServiceHandler register SettingsServiceHandler to http.Router.
func RegisterSettingsServiceHandler(r http.Router, srv SettingsServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_GetSettings := func(method, path string, fn func(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetSettingsRequest))
		}
		var GetSettings_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetSettings_info = transport.NewServiceInfo("erda.core.monitor.settings.SettingsService", "GetSettings", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetSettingsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["org_id"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.OrgID = val
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetSettings_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_PutSettings := func(method, path string, fn func(context.Context, *PutSettingsRequest) (*PutSettingsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PutSettingsRequest))
		}
		var PutSettings_info transport.ServiceInfo
		if h.Interceptor != nil {
			PutSettings_info = transport.NewServiceInfo("erda.core.monitor.settings.SettingsService", "PutSettings", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in PutSettingsRequest
				if err := h.Decode(r, &in.Data); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["org_id"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.OrgID = val
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PutSettings_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_RegisterMonitorConfig := func(method, path string, fn func(context.Context, *RegisterMonitorConfigRequest) (*RegisterMonitorConfigResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RegisterMonitorConfigRequest))
		}
		var RegisterMonitorConfig_info transport.ServiceInfo
		if h.Interceptor != nil {
			RegisterMonitorConfig_info = transport.NewServiceInfo("erda.core.monitor.settings.SettingsService", "RegisterMonitorConfig", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in RegisterMonitorConfigRequest
				if err := h.Decode(r, &in.Data); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RegisterMonitorConfig_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetSettings("GET", "/api/global/settings", srv.GetSettings)
	add_PutSettings("PUT", "/api/global/settings", srv.PutSettings)
	add_RegisterMonitorConfig("PUT", "/api/config/register", srv.RegisterMonitorConfig)
}
