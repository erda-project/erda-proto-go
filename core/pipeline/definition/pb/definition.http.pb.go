// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: definition.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// DefinitionServiceHandler is the server API for DefinitionService service.
type DefinitionServiceHandler interface {
	// POST /api/pipeline-definitions/actions/process
	Process(context.Context, *PipelineDefinitionProcessRequest) (*PipelineDefinitionProcessResponse, error)
	// GET /api/pipeline-definitions/actions/version
	Version(context.Context, *PipelineDefinitionProcessVersionRequest) (*PipelineDefinitionProcessVersionResponse, error)
}

// RegisterDefinitionServiceHandler register DefinitionServiceHandler to http.Router.
func RegisterDefinitionServiceHandler(r http.Router, srv DefinitionServiceHandler, opts ...http.HandleOption) {
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

	add_Process := func(method, path string, fn func(context.Context, *PipelineDefinitionProcessRequest) (*PipelineDefinitionProcessResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineDefinitionProcessRequest))
		}
		var Process_info transport.ServiceInfo
		if h.Interceptor != nil {
			Process_info = transport.NewServiceInfo("erda.core.pipeline.cms.DefinitionService", "Process", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Process_info)
				}
				r = r.WithContext(ctx)
				var in PipelineDefinitionProcessRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Version := func(method, path string, fn func(context.Context, *PipelineDefinitionProcessVersionRequest) (*PipelineDefinitionProcessVersionResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineDefinitionProcessVersionRequest))
		}
		var Version_info transport.ServiceInfo
		if h.Interceptor != nil {
			Version_info = transport.NewServiceInfo("erda.core.pipeline.cms.DefinitionService", "Version", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Version_info)
				}
				r = r.WithContext(ctx)
				var in PipelineDefinitionProcessVersionRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Process("POST", "/api/pipeline-definitions/actions/process", srv.Process)
	add_Version("GET", "/api/pipeline-definitions/actions/version", srv.Version)
}
