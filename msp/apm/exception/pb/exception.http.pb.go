// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: exception.proto

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

// ExceptionServiceHandler is the server API for ExceptionService service.
type ExceptionServiceHandler interface {
	// GET /api/msp/apm/exceptions
	GetExceptions(context.Context, *GetExceptionsRequest) (*GetExceptionsResponse, error)
	// GET /api/msp/apm/exceptions/event-ids
	GetExceptionEventIds(context.Context, *GetExceptionEventIdsRequest) (*GetExceptionEventIdsResponse, error)
	// GET /api/msp/apm/exceptions/events
	GetExceptionEvent(context.Context, *GetExceptionEventRequest) (*GetExceptionEventResponse, error)
}

// RegisterExceptionServiceHandler register ExceptionServiceHandler to http.Router.
func RegisterExceptionServiceHandler(r http.Router, srv ExceptionServiceHandler, opts ...http.HandleOption) {
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

	add_GetExceptions := func(method, path string, fn func(context.Context, *GetExceptionsRequest) (*GetExceptionsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetExceptionsRequest))
		}
		var GetExceptions_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetExceptions_info = transport.NewServiceInfo("erda.msp.apm.exception.ExceptionService", "GetExceptions", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetExceptions_info)
				}
				r = r.WithContext(ctx)
				var in GetExceptionsRequest
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
				if vals := params["scopeId"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetExceptionEventIds := func(method, path string, fn func(context.Context, *GetExceptionEventIdsRequest) (*GetExceptionEventIdsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetExceptionEventIdsRequest))
		}
		var GetExceptionEventIds_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetExceptionEventIds_info = transport.NewServiceInfo("erda.msp.apm.exception.ExceptionService", "GetExceptionEventIds", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetExceptionEventIds_info)
				}
				r = r.WithContext(ctx)
				var in GetExceptionEventIdsRequest
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
				if vals := params["exceptionId"]; len(vals) > 0 {
					in.ExceptionID = vals[0]
				}
				if vals := params["scopeId"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetExceptionEvent := func(method, path string, fn func(context.Context, *GetExceptionEventRequest) (*GetExceptionEventResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetExceptionEventRequest))
		}
		var GetExceptionEvent_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetExceptionEvent_info = transport.NewServiceInfo("erda.msp.apm.exception.ExceptionService", "GetExceptionEvent", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetExceptionEvent_info)
				}
				r = r.WithContext(ctx)
				var in GetExceptionEventRequest
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
				if vals := params["exceptionEventId"]; len(vals) > 0 {
					in.ExceptionEventID = vals[0]
				}
				if vals := params["scopeId"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetExceptions("GET", "/api/msp/apm/exceptions", srv.GetExceptions)
	add_GetExceptionEventIds("GET", "/api/msp/apm/exceptions/event-ids", srv.GetExceptionEventIds)
	add_GetExceptionEvent("GET", "/api/msp/apm/exceptions/events", srv.GetExceptionEvent)
}
