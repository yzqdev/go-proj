package static

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//go:embed index.html  a.js
var indexFs embed.FS

func Assets() http.FileSystem {

	return http.FS(indexFs)
}

func Static(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.Status(200)
		index, _ := Assets().Open("index.html")
		indexHtml, _ := ioutil.ReadAll(index)
		_, _ = c.Writer.WriteString(string(indexHtml))

		c.Writer.Flush()
		c.Writer.WriteHeaderNow()
	})
	router.StaticFS("/static", Assets())

}
