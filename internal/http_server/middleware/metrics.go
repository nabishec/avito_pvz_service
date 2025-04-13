package middleware

import (
	"net/http"
	"time"

	"github.com/nabishec/avito_pvz_service/internal/metrics"
)

func MetricsRecorder(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()

		metrics.HTTPRequestsTotal.WithLabelValues(r.URL.Path).Inc()

		metrics.HTTPRequestTime.WithLabelValues(r.URL.Path).Observe(duration)
	}
	return http.HandlerFunc(fn)
}
