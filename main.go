package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, makeGreeting(name))
		router.Run(":8080")
	})
}

func makeGreeting(name string) string {
	return fmt.Sprintf("Hello, %d", name)
}