package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed   config.yml
var confFile embed.FS

func WithConfStatic(router *gin.Engine) {

	router.StaticFS("/f", http.FS(confFile))

}
