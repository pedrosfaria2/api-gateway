package config

import (
	"time"
)

type Config struct {
	Server   ServerConfig    `yaml:"server"`
	Backends []BackendConfig `yaml:"backends"`
	Plugins  PluginConfig    `yaml:"plugins"`
}

type ServerConfig struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
}

type BackendConfig struct {
	Name      string            `yaml:"name"`
	URL       string            `yaml:"url"`
	Timeout   time.Duration     `yaml:"timeout"`
	Retries   int               `yaml:"retries"`
	RateLimit RateLimitConfig   `yaml:"rateLimit"`
	Circuit   CircuitConfig     `yaml:"circuit"`
	Routes    []RouteConfig     `yaml:"routes"`
	Headers   map[string]string `yaml:"headers"`
}

type RouteConfig struct {
	Path    string   `yaml:"path"`
	Methods []string `yaml:"methods"`
	Strip   bool     `yaml:"strip"`
}

type RateLimitConfig struct {
	Enabled  bool          `yaml:"enabled"`
	Requests int           `yaml:"requests"`
	Duration time.Duration `yaml:"duration"`
}

type CircuitConfig struct {
	Enabled     bool          `yaml:"enabled"`
	Threshold   float64       `yaml:"threshold"`
	Timeout     time.Duration `yaml:"timeout"`
	MaxRequests int           `yaml:"maxRequests"`
}

type PluginConfig struct {
	Auth    AuthConfig    `yaml:"auth"`
	Metrics MetricsConfig `yaml:"metrics"`
	Tracing TracingConfig `yaml:"tracing"`
}

type AuthConfig struct {
	Enabled bool   `yaml:"enabled"`
	Type    string `yaml:"type"`
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

type TracingConfig struct {
	Enabled bool   `yaml:"enabled"`
	Type    string `yaml:"type"`
}
