package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// ---------------------------------------------------------
// 1. PROMETHEUS METRICS SIMULATION (Counter & Histogram)
// ---------------------------------------------------------

type PrometheusMetrics struct {
	httpRequestsTotal map[string]int
	requestDurationMs []float64
}

func NewPrometheusMetrics() *PrometheusMetrics {
	return &PrometheusMetrics{
		httpRequestsTotal: make(map[string]int),
	}
}

func (m *PrometheusMetrics) RecordRequest(method, path string, statusCode int, durationMs float64) {
	key := fmt.Sprintf("%s %s %d", method, path, statusCode)
	m.httpRequestsTotal[key]++
	m.requestDurationMs = append(m.requestDurationMs, durationMs)
}

func main() {
	// ---------------------------------------------------------
	// 2. GO 1.21+ STRUCTURED LOGGING (slog JSON Handler)
	// ---------------------------------------------------------
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)

	fmt.Println("==================================================")
	fmt.Println(" 📊 Observability, Prometheus Metrics & slog JSON")
	fmt.Println("==================================================")

	metrics := NewPrometheusMetrics()
	ctx := context.Background()

	// Simulate Production API Requests with Structured Logging
	logger.InfoContext(ctx, "HTTP Server starting...", "port", 8080, "env", "production")

	// Request 1: Success
	start1 := time.Now()
	time.Sleep(15 * time.Millisecond)
	duration1 := float64(time.Since(start1).Milliseconds())
	metrics.RecordRequest("GET", "/api/books", 200, duration1)

	logger.InfoContext(ctx, "HTTP Request Processed",
		"method", "GET",
		"path", "/api/books",
		"status", 200,
		"duration_ms", duration1,
		"user_id", 101,
	)

	// Request 2: Error
	start2 := time.Now()
	time.Sleep(45 * time.Millisecond)
	duration2 := float64(time.Since(start2).Milliseconds())
	metrics.RecordRequest("POST", "/api/login", 401, duration2)

	logger.ErrorContext(ctx, "Authentication Failed",
		"method", "POST",
		"path", "/api/login",
		"status", 401,
		"duration_ms", duration2,
		"error", "invalid credentials",
	)

	fmt.Println("\n📈 Recorded Prometheus Metrics Endpoint (/metrics):")
	for k, count := range metrics.httpRequestsTotal {
		fmt.Printf("   http_requests_total{endpoint=\"%s\"} %d\n", k, count)
	}
}
