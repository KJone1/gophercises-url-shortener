package main

import (
	"flag"
	"fmt"
	"net/http"

	parser "github.com/KJone1/gophercises-url-shortener/parsers"
	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func Redirect(c *gin.Context) {
	url := c.Param("route")
	url = "/" + url
	f := parser.Yaml("./routeFile.yaml")

	for _, m := range f {
		if url == m.From {
			fmt.Printf("redirected to: %s from url: %s\n", m.To, url)
			c.Redirect(http.StatusMovedPermanently, m.To)
			return
		}
	}
	c.String(http.StatusBadRequest, "400 BadRequest")

}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/:route", Redirect)

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
