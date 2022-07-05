package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"go-embed/config"
	"go-embed/public"
	"go-embed/static"
)

func main() {
	Embed.RestoreFolder(".")
	router := gin.Default()

	// example: /public/static/js/1.js
	public.Static(router)
	static.Static(router)
	WithConfStatic(router)
	MainStatic(router)
	conf := config.GetGlobal()
	color.Redln(conf.Mysql.User)
	router.Run(":9080")
}
