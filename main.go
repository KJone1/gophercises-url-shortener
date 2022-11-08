package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", Redirect)

	router.GET("/ping", StatusOK)

	return router
}

func main() {
	debug_mode := flag.Bool("v", false, "Debug mode, verbose logging")
	flag.Parse()
	if !*debug_mode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := SetupRouter()
	r.Run(":3000")

}
