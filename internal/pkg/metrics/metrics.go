package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTPRequestsTotal é um contador que registra o número total de requisições HTTP
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP",
		},
		[]string{"method", "endpoint", "status"},
	)

	// HTTPRequestDuration é um histograma que registra a duração das requisições HTTP
	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duração das requisições HTTP em segundos",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// DatabaseOperationsTotal é um contador que registra o número total de operações no banco de dados
	DatabaseOperationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_operations_total",
			Help: "Total de operações no banco de dados",
		},
		[]string{"operation", "table"},
	)

	// DatabaseOperationDuration é um histograma que registra a duração das operações no banco de dados
	DatabaseOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "database_operation_duration_seconds",
			Help:    "Duração das operações no banco de dados em segundos",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation", "table"},
	)
)
