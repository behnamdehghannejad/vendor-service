package metrics

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()

		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}

		method := c.Request.Method
		status := fmt.Sprint(c.Writer.Status())

		HttpRequestsTotal.WithLabelValues(
			path,
			method,
			status,
		).Inc()

		HttpRequestDuration.WithLabelValues(
			path,
			method,
		).Observe(duration)
	}
}
