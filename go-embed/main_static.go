package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed   templates
var f embed.FS

func MainAssets() http.FileSystem {
	//这一步是必须的
	pub, _ := fs.Sub(f, "templates")
	return http.FS(pub)
}

func MainStatic(router *gin.Engine) {

	//router.StaticFS("/main", http.FS(f))//这个不对
	router.StaticFS("/main", MainAssets())

}
