// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: checker.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// CheckerServiceHandler is the server API for CheckerService service.
type CheckerServiceHandler interface {
	// POST /api/msp/checkers/{scope}/{scopeID}
	CreateChecker(context.Context, *CreateCheckerRequest) (*CreateCheckerResponse, error)
	// PUT /api/msp/checkers/{scope}/{scopeID}/{id}
	UpdateChecker(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error)
	// DELETE /api/msp/checkers/{scope}/{scopeID}/{id}
	DeleteChecker(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error)
	// GET /api/msp/checkers/{scope}/{scopeID}
	ListCheckers(context.Context, *ListCheckersRequest) (*ListCheckersResponse, error)
	// GET /api/msp/checker-descriptions/{scope}/{scopeID}
	DescribeCheckers(context.Context, *DescribeCheckersRequest) (*DescribeCheckersResponse, error)
	// GET /api/msp/checker-descriptions/{scope}/{scopeID}/{id}
	DescribeChecker(context.Context, *DescribeCheckerRequest) (*DescribeCheckerResponse, error)
}

// RegisterCheckerServiceHandler register CheckerServiceHandler to http.Router.
func RegisterCheckerServiceHandler(r http.Router, srv CheckerServiceHandler, opts ...http.HandleOption) {
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

	add_CreateChecker := func(method, path string, fn func(context.Context, *CreateCheckerRequest) (*CreateCheckerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateCheckerRequest))
		}
		var CreateChecker_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "CreateChecker", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateChecker_info)
				}
				r = r.WithContext(ctx)
				var in CreateCheckerRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "scope":
							in.Scope = val
						case "scopeID":
							in.ScopeID = val
						}
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

	add_UpdateChecker := func(method, path string, fn func(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateCheckerRequest))
		}
		var UpdateChecker_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "UpdateChecker", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateChecker_info)
				}
				r = r.WithContext(ctx)
				var in UpdateCheckerRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "scope":
							in.Scope = val
						case "scopeID":
							in.ScopeID = val
						case "id":
							in.Id = val
						}
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

	add_DeleteChecker := func(method, path string, fn func(context.Context, *UpdateCheckerRequest) (*UpdateCheckerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateCheckerRequest))
		}
		var DeleteChecker_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "DeleteChecker", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteChecker_info)
				}
				r = r.WithContext(ctx)
				var in UpdateCheckerRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "scope":
							in.Scope = val
						case "scopeID":
							in.ScopeID = val
						case "id":
							in.Id = val
						}
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

	add_ListCheckers := func(method, path string, fn func(context.Context, *ListCheckersRequest) (*ListCheckersResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListCheckersRequest))
		}
		var ListCheckers_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListCheckers_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "ListCheckers", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListCheckers_info)
				}
				r = r.WithContext(ctx)
				var in ListCheckersRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "scope":
							in.Scope = val
						case "scopeID":
							in.ScopeID = val
						}
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

	add_DescribeCheckers := func(method, path string, fn func(context.Context, *DescribeCheckersRequest) (*DescribeCheckersResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DescribeCheckersRequest))
		}
		var DescribeCheckers_info transport.ServiceInfo
		if h.Interceptor != nil {
			DescribeCheckers_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "DescribeCheckers", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DescribeCheckers_info)
				}
				r = r.WithContext(ctx)
				var in DescribeCheckersRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "scope":
							in.Scope = val
						case "scopeID":
							in.ScopeID = val
						}
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

	add_DescribeChecker := func(method, path string, fn func(context.Context, *DescribeCheckerRequest) (*DescribeCheckerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DescribeCheckerRequest))
		}
		var DescribeChecker_info transport.ServiceInfo
		if h.Interceptor != nil {
			DescribeChecker_info = transport.NewServiceInfo("erda.msp.apm.checker.CheckerService", "DescribeChecker", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DescribeChecker_info)
				}
				r = r.WithContext(ctx)
				var in DescribeCheckerRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "scope":
							in.Scope = val
						case "scopeID":
							in.ScopeID = val
						case "id":
							val, err := strconv.ParseInt(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.Id = val
						}
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

	add_CreateChecker("POST", "/api/msp/checkers/{scope}/{scopeID}", srv.CreateChecker)
	add_UpdateChecker("PUT", "/api/msp/checkers/{scope}/{scopeID}/{id}", srv.UpdateChecker)
	add_DeleteChecker("DELETE", "/api/msp/checkers/{scope}/{scopeID}/{id}", srv.DeleteChecker)
	add_ListCheckers("GET", "/api/msp/checkers/{scope}/{scopeID}", srv.ListCheckers)
	add_DescribeCheckers("GET", "/api/msp/checker-descriptions/{scope}/{scopeID}", srv.DescribeCheckers)
	add_DescribeChecker("GET", "/api/msp/checker-descriptions/{scope}/{scopeID}/{id}", srv.DescribeChecker)
}
