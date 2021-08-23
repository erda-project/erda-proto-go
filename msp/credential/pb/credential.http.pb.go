// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: credential.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// AccessKeyServiceHandler is the server API for AccessKeyService service.
type AccessKeyServiceHandler interface {
	// POST /api/msp/credential/access-keys
	CreateAccessKey(context.Context, *CreateAccessKeyRequest) (*CreateAccessKeyResponse, error)
	// DELETE /api/msp/credential/access-keys/{accessKeyId}
	DeleteAccessKey(context.Context, *DeleteAccessKeyRequest) (*DeleteAccessKeyResponse, error)
	// GET /api/msp/credential/access-keys/{accessKeyId}
	GetAccessKey(context.Context, *GetAccessKeyRequest) (*GetAccessKeyResponse, error)
	// POST /api/msp/credential/access-keys/download
	DownloadAccessKeyFile(context.Context, *DownloadAccessKeyFileRequest) (*DownloadAccessKeyFileResponse, error)
}

// RegisterAccessKeyServiceHandler register AccessKeyServiceHandler to http.Router.
func RegisterAccessKeyServiceHandler(r http.Router, srv AccessKeyServiceHandler, opts ...http.HandleOption) {
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

	add_CreateAccessKey := func(method, path string, fn func(context.Context, *CreateAccessKeyRequest) (*CreateAccessKeyResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateAccessKeyRequest))
		}
		var CreateAccessKey_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateAccessKey_info = transport.NewServiceInfo("erda.msp.credential.AccessKeyService", "CreateAccessKey", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateAccessKey_info)
				}
				r = r.WithContext(ctx)
				var in CreateAccessKeyRequest
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

	add_DeleteAccessKey := func(method, path string, fn func(context.Context, *DeleteAccessKeyRequest) (*DeleteAccessKeyResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteAccessKeyRequest))
		}
		var DeleteAccessKey_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteAccessKey_info = transport.NewServiceInfo("erda.msp.credential.AccessKeyService", "DeleteAccessKey", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteAccessKey_info)
				}
				r = r.WithContext(ctx)
				var in DeleteAccessKeyRequest
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
						case "accessKeyId":
							in.AccessKeyId = val
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

	add_GetAccessKey := func(method, path string, fn func(context.Context, *GetAccessKeyRequest) (*GetAccessKeyResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetAccessKeyRequest))
		}
		var GetAccessKey_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetAccessKey_info = transport.NewServiceInfo("erda.msp.credential.AccessKeyService", "GetAccessKey", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetAccessKey_info)
				}
				r = r.WithContext(ctx)
				var in GetAccessKeyRequest
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
						case "accessKeyId":
							in.AccessKeyId = val
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

	add_DownloadAccessKeyFile := func(method, path string, fn func(context.Context, *DownloadAccessKeyFileRequest) (*DownloadAccessKeyFileResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DownloadAccessKeyFileRequest))
		}
		var DownloadAccessKeyFile_info transport.ServiceInfo
		if h.Interceptor != nil {
			DownloadAccessKeyFile_info = transport.NewServiceInfo("erda.msp.credential.AccessKeyService", "DownloadAccessKeyFile", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DownloadAccessKeyFile_info)
				}
				r = r.WithContext(ctx)
				var in DownloadAccessKeyFileRequest
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

	add_CreateAccessKey("POST", "/api/msp/credential/access-keys", srv.CreateAccessKey)
	add_DeleteAccessKey("DELETE", "/api/msp/credential/access-keys/{accessKeyId}", srv.DeleteAccessKey)
	add_GetAccessKey("GET", "/api/msp/credential/access-keys/{accessKeyId}", srv.GetAccessKey)
	add_DownloadAccessKeyFile("POST", "/api/msp/credential/access-keys/download", srv.DownloadAccessKeyFile)
}
