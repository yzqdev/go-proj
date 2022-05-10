package htmltil

import "testing"

func TestOpenByBrowser(t *testing.T) {
	_ = OpenByBrowser("http://www.baidu.com")
}