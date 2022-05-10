package main

import (
	"github.com/gin-gonic/gin"
	"go-embed/public"
	"go-embed/static"
)

func main() {

	router := gin.Default()

	// example: /public/static/js/1.js
	public.Static(router)
	static.Static(router)
	MainStatic(router)
	router.Run(":9080")
}
