package server

import (
	"encoding/gob"
	"gin-ory-stack-demo/handlers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	// gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default() // default includes logger and crash recovery middlewares

	// register custom types to store in cookies
	gob.Register(map[string]interface{}{})

	// session and cookie management
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	// middlewares
	router.Use(func(ctx *gin.Context) {
		// custom prometheus metric - count total number of HTTP requests
		httpRequestsTotal.WithLabelValues(ctx.Request.Method, ctx.FullPath()).Inc()
		ctx.Next()
	})

	// static routes
	router.Static("/static", "web/static")
	router.LoadHTMLGlob("web/template/**/*")

	// app routes
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/dashboard")
	})
	router.GET("/dashboard", handlers.DashboardHandler)
	// router.GET("/login", handlers.LoginHandler)
	router.GET("/sign-up", handlers.AccountRegistrationHandler)
	router.GET("/account-recovery", handlers.AccountRecoveryHandler)
	router.GET("/account-verification", handlers.AccountRecoveryHandler)
	router.GET("/profile-settings", handlers.ProfileSettingsHandler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// default error handling
	// router.NoRoute(func(ctx *gin.Context) {
	// 	ctx.Redirect(http.StatusMovedPermanently, "/error")
	// })
	// router.NoMethod(func(ctx *gin.Context) {
	// 	ctx.Redirect(http.StatusMovedPermanently, "/error")
	// })

	return router
}
