package main

import (
	"flag"
	"fmt"
	"net/http"

	parser "github.com/KJone1/gophercises-url-shortener/parsers"
	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func Redirect(file parser.Destructured) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		url := "/" + c.Param("route")

		for _, m := range file {
			if url == m.From {
				fmt.Printf("redirected to: %s from url: %s\n", m.To, url)
				c.Redirect(http.StatusMovedPermanently, m.To)
				return
			}
		}

		c.HTML(http.StatusNotFound, "404.html", gin.H{})

	}

	return gin.HandlerFunc(fn)

}

func main() {

	debug_mode := flag.Bool("v", false, "Run in debug mode, verbose logging.")
	route_file := flag.String("f", "routeFile.yaml", "Route file to use.")
	flag.Parse()

	if *debug_mode == false {
		gin.SetMode(gin.ReleaseMode)
	}

	py := parser.Yaml(*route_file)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("./assets/404/404.html")
	router.Static("/assets/404", "./assets/404/")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	v1 := router.Group("api/v1/")
	{
		v1.GET("/:route", Redirect(py))
		v1.GET("/ping", StatusOK)
	}

	router.Run(":3000")
}
