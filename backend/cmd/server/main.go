package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/leondevpt/decode-ethtx/backend/handler"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leondevpt/decode-ethtx/backend"
)

var port int

func init() {
	flag.IntVar(&port, "p", 3000, "listen port")
}

func main() {
	flag.Parse()

	fmt.Printf("listen on port:%d\n", port)
	r := gin.Default()
	r.Use(cors.Default())

	// web
	webRoot := backend.EmbedFolder(backend.StaticFiles, "web")
	staticServer := static.Serve("/", webRoot)

	r.Use(staticServer)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	// api
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/ping", handler.PingHandler)
	apiV1.POST("/decode", handler.DecodeHandler)

	r.Run(fmt.Sprintf(":%d", port))
}

func GetEnv(name string, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
