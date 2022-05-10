package gfile

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExist(t *testing.T) {
	var f os.FileInfo
	f, _ = os.Stat("./file.go")
	fmt.Println(f.IsDir())
	flag, _ := FileSize("./file.go")
	fmt.Println(flag)
}
func TestGetPwd(t *testing.T) {
	fmt.Println(GetPwd())
}
