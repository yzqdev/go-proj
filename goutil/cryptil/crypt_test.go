package cryptil

import (
	"fmt"
	"testing"
)

func TestMd5Crypt(t *testing.T) {
	fmt.Println(Md5Crypt("aaa"))
}
func TestSha1Crypt(t *testing.T) {
	fmt.Println(Sha1Crypt("aaa"))
}