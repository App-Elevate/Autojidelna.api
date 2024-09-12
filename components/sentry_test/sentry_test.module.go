package sentrytest

import "github.com/gin-gonic/gin"

// Api for testing Sentry is setup correctly
func Register(router *gin.Engine) {
	app := router.Group("sentry")
	app.GET("/crash", testCrash)
}
