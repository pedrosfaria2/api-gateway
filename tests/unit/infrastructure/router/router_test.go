package router_test

import (
	"net/http"
	"testing"

	"github.com/pedrosfaria2/api-gateway/internal/core/ports"
	"github.com/pedrosfaria2/api-gateway/internal/infrastructure/router"
)

func TestRouter_Route(t *testing.T) {
	backends := []ports.Backend{
		{
			Name: "users",
			URL:  "http://users-service",
			Routes: []ports.Route{
				{
					Path:    "/api/v1/users",
					Methods: []string{"GET", "POST"},
				},
				{
					Path:    "/api/v1/users/**",
					Methods: []string{"GET"},
				},
			},
		},
		{
			Name: "orders",
			URL:  "http://orders-service",
			Routes: []ports.Route{
				{
					Path:    "/api/v1/orders",
					Methods: []string{"POST"},
				},
			},
		},
	}

	r := router.New(backends)

	tests := []struct {
		name    string
		path    string
		method  string
		want    string
		wantErr bool
	}{
		{
			name:    "exact path and method match",
			path:    "/api/v1/users",
			method:  "GET",
			want:    "users",
			wantErr: false,
		},
		{
			name:    "wildcard path match",
			path:    "/api/v1/users/123",
			method:  "GET",
			want:    "users",
			wantErr: false,
		},
		{
			name:    "method not allowed",
			path:    "/api/v1/orders",
			method:  "GET",
			wantErr: true,
		},
		{
			name:    "path not found",
			path:    "/api/v1/invalid",
			method:  "GET",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			got, err := r.Route(req)

			if (err != nil) != tt.wantErr {
				t.Errorf("Router.Route() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got.Name != tt.want {
				t.Errorf("Router.Route() = %v, want %v", got.Name, tt.want)
			}
		})
	}
}
