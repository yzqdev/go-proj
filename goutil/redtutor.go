package main

import (
	"net/http"
	"runtime"
)

func main() {
	color.Red(runtime.GOOS)
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		color.Cyan("失败了")
	}
	color.Red(resp)
	defer resp.Body.Close()
}
