package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

func RateLimiter(maxRequests int, duration time.Duration) echo.MiddlewareFunc {
	// Create a map to hold rate limiters for each IP
	var visitors = make(map[string]*rate.Limiter)
	var mutex sync.Mutex

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()

			// Lock the map to avoid race conditions
			mutex.Lock()
			if _, exists := visitors[ip]; !exists {
				visitors[ip] = rate.NewLimiter(rate.Every(duration), maxRequests)
			}

			limiter := visitors[ip]
			mutex.Unlock()

			if !limiter.Allow() {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"error": "Too many requests",
				})
			}

			return next(c)
		}
	}
}