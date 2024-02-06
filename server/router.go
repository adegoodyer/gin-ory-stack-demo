package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	// middlewares
	router.Use(func(ctx *gin.Context) {
		// custom prometheus metric - count total number of HTTP requests
		httpRequestsTotal.WithLabelValues(ctx.Request.Method, ctx.FullPath()).Inc()
		ctx.Next()
	})

	// routes
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
