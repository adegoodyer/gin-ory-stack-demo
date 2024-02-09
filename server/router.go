package server

import (
	"encoding/gob"
	accountrecovery "gin-ory-stack-demo/controllers/account-recovery"
	accountverification "gin-ory-stack-demo/controllers/account-verification"
	"gin-ory-stack-demo/controllers/dashboard"
	"gin-ory-stack-demo/controllers/error"
	"gin-ory-stack-demo/controllers/login"
	profilesettings "gin-ory-stack-demo/controllers/profile-settings"
	signup "gin-ory-stack-demo/controllers/sign-up"
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
	router.GET("/dashboard", dashboard.Handler)
	router.GET("/login", login.Handler)
	router.GET("/sign-up", signup.Handler)
	router.GET("/account-recovery", accountrecovery.Handler)
	router.GET("/account-verification", accountverification.Handler)
	router.GET("/profile-settings", profilesettings.Handler)
	router.GET("/error", error.Handler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// error handling
	// router.NoRoute(func(ctx *gin.Context) {
	// 	ctx.Redirect(http.StatusMovedPermanently, "/error")
	// })
	// router.NoMethod(func(ctx *gin.Context) {
	// 	ctx.Redirect(http.StatusMovedPermanently, "/error")
	// })

	return router
}
