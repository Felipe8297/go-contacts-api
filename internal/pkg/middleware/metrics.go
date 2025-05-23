package middleware

import (
	"fmt"
	"time"

	"github.com/Felipe8297/go-contacts-api/internal/pkg/metrics"
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware é um middleware que registra métricas para cada requisição HTTP
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Processa a requisição
		c.Next()

		// Registra a duração da requisição
		duration := time.Since(start).Seconds()
		metrics.HTTPRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)

		// Registra o total de requisições
		metrics.HTTPRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			fmt.Sprintf("%d", c.Writer.Status()),
		).Inc()
	}
}
