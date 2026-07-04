package middlewares

import "portfolio-api/config"

// Middleware aggregates the application middlewares so they can share config.
type Middleware struct {
	Config config.AppConfig
}

// NewMiddleware builds the middleware aggregator.
func NewMiddleware(cfg config.AppConfig) Middleware {
	return Middleware{Config: cfg}
}
