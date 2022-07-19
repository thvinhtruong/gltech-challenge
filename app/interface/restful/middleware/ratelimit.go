package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var (
	rateLimitOnce        sync.Once
	rateLimit_middleware fiber.Handler
)

func RateLimitMiddleware() fiber.Handler {
	rateLimitOnce.Do(func() {
		rateLimit_middleware = limiter.New(limiter.Config{
			Max:        20,
			Expiration: 30 * time.Second,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(429).SendString("Too Many Requests")
			},
			LimiterMiddleware: limiter.SlidingWindow{},
		})
	})

	return rateLimit_middleware

}
