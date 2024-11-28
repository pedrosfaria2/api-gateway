package ports

import (
	"context"
	"net/http"
)

type Gateway interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Router interface {
	Route(r *http.Request) (*Backend, error)
}

type Backend struct {
	Name      string
	URL       string
	Timeout   int
	Retries   int
	Circuit   CircuitBreaker
	RateLimit RateLimiter
}

type Middleware interface {
	Process(next http.Handler) http.Handler
}

type CircuitBreaker interface {
	AllowRequest() bool
	RecordSuccess()
	RecordFailure(err error)
}

type RateLimiter interface {
	Allow(key string) bool
}
