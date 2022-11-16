package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	parser "github.com/KJone1/gophercises-url-shortener/src/parsers"
	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func Redirect(file parser.Destructured) gin.HandlerFunc {
	context := func(c *gin.Context) {
		url := "/" + c.Param("route")

		for _, m := range file.Mapping {
			if url == m.From {
				text := "Redirected to: " + m.To + "from url: " + url
				bold_text := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 1, text)
				log.Println(bold_text)
				c.Redirect(http.StatusMovedPermanently, m.To)

				return
			}
		}

		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	}

	return gin.HandlerFunc(context)
}

func main() {
	debug_mode := flag.Bool("v", false, "Run in debug mode, verbose logging.")
	route_file := flag.String("f", "routeFile.yaml", "Route file to use.")
	flag.Parse()

	if !*debug_mode {
		gin.SetMode(gin.ReleaseMode)
	}

	py := parser.Yaml(*route_file)

	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	router.LoadHTMLGlob("./src/assets/404/404.html")
	router.Static("/assets/404", "./src/assets/404/")
	router.StaticFile("/favicon.ico", "./src/assets/favicon.ico")

	v1 := router.Group("api/v1/")
	{
		v1.GET("/:route", Redirect(py))
		v1.GET("/ping", StatusOK)
	}

	err = router.Run(":4001")
	if err != nil {
		log.Fatal(err)
	}
}
