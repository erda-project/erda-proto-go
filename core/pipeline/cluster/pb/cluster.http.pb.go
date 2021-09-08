// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: cluster.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ClusterServiceHandler is the server API for ClusterService service.
type ClusterServiceHandler interface {
	// POST /api/pipeline-clusters/actions/hook
	ClusterHook(context.Context, *ClusterHookRequest) (*ClusterHookResponse, error)
}

// RegisterClusterServiceHandler register ClusterServiceHandler to http.Router.
func RegisterClusterServiceHandler(r http.Router, srv ClusterServiceHandler, opts ...http.HandleOption) {
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

	add_ClusterHook := func(method, path string, fn func(context.Context, *ClusterHookRequest) (*ClusterHookResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ClusterHookRequest))
		}
		var ClusterHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			ClusterHook_info = transport.NewServiceInfo("erda.core.pipeline.cluster.ClusterService", "ClusterHook", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ClusterHook_info)
				}
				r = r.WithContext(ctx)
				var in ClusterHookRequest
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

	add_ClusterHook("POST", "/api/pipeline-clusters/actions/hook", srv.ClusterHook)
}
