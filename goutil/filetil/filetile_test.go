package filetil

import (
	"fmt"
	"testing"
)

func TestScanFiles(t *testing.T) {
	var fl []FileList
	fl, _ = ScanFiles("./")
	fmt.Print(fl)
}