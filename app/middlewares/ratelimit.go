package middlewares

import (
	"gulmen/app/responses"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.RWMutex
	rate     int           // requests per duration
	duration time.Duration // time window
}

type Visitor struct {
	count    int
	lastSeen time.Time
}

func NewRateLimiter(rate int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		rate:     rate,
		duration: duration,
	}

	// Cleanup goroutine
	go rl.cleanup()

	return rl
}

func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > rl.duration {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) isAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		rl.visitors[ip] = &Visitor{count: 1, lastSeen: time.Now()}
		return true
	}

	if time.Since(v.lastSeen) > rl.duration {
		v.count = 1
		v.lastSeen = time.Now()
		return true
	}

	if v.count >= rl.rate {
		return false
	}

	v.count++
	v.lastSeen = time.Now()
	return true
}

func RateLimitMiddleware(rate int, duration time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(rate, duration)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !limiter.isAllowed(ip) {
			responses.Error(c, 429, "Too many requests", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
