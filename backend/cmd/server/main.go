package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leondevpt/decode-ethtx/backend/handler"
	"net/http"
	"os"
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/decode", handler.DecodeHandler)

	r.Run(fmt.Sprintf(":%d", port))
}

func GetEnv(name string, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
