package middleware

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/thvinhtruong/legoha/pkg/conversion"
)

var (
	corsOnce        sync.Once
	cors_middleware fiber.Handler
)

var (
	localhost      = []string{"http://localhost:3000", "http://localhost:5001", "http://localhost:8080"}
	methods        = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	headers        = []string{"Accept", "Origin", "Content-Type"}
	expose_headers = []string{"Content-Length", "content-type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
)

func CorsMiddleware() fiber.Handler {
	corsOnce.Do(func() {
		cors_middleware = cors.New(cors.Config{
			AllowOrigins:     conversion.MergeString(localhost...),
			AllowMethods:     conversion.MergeString(methods...),
			AllowHeaders:     conversion.MergeString(headers...),
			ExposeHeaders:    conversion.MergeString(expose_headers...),
			AllowCredentials: true,
			MaxAge:           86400,
		})
	})

	return cors_middleware
}
