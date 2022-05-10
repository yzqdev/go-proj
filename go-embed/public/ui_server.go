package public

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"io/fs"
	"net/http"
)

//go:embed   dist
var f embed.FS

//go:embed a.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var txt embed.FS

func Assets() http.FileSystem {
	pub, _ := fs.Sub(f, "dist")
	return http.FS(pub)
}

func Static(router *gin.Engine) {

	router.StaticFS("/atools", Assets())
	color.Redln("这是文本")
	color.Redln(s)
	color.Cyanln("这是byte")
	color.Cyanln(b)
	color.Yellowln("这是fs.FS")
	data, _ := txt.ReadFile("hello.txt")
	color.Yellowln(string(data))
}
