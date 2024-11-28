package router

import (
	"net/http"
	"path"
	"strings"

	"github.com/pedrosfaria2/api-gateway/internal/core/ports"
)

type Router struct {
	backends []ports.Backend
}

func New(backends []ports.Backend) *Router {
	return &Router{
		backends: backends,
	}
}

func (r *Router) Route(req *http.Request) (*ports.Backend, error) {
	reqPath := path.Clean(req.URL.Path)
	reqMethod := req.Method

	for _, backend := range r.backends {
		for _, route := range backend.Routes {
			if matchPath(reqPath, route.Path) && matchMethod(reqMethod, route.Methods) {
				return &backend, nil
			}
		}
	}

	return nil, ports.ErrBackendNotFound
}

func matchPath(reqPath, routePath string) bool {
	if strings.HasSuffix(routePath, "/**") {
		basePattern := strings.TrimSuffix(routePath, "/**")
		return strings.HasPrefix(reqPath, basePattern)
	}

	return reqPath == routePath
}

func matchMethod(reqMethod string, routeMethods []string) bool {
	if len(routeMethods) == 0 {
		return true
	}

	for _, method := range routeMethods {
		if method == reqMethod {
			return true
		}
	}
	return false
}
