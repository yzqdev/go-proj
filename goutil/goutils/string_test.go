package goutils

import (
	"fmt"
	"testing"
)

func TestUnderscoreName(t *testing.T) {
	fmt.Print(UnderscoreName("ConvertTest"))
}
func TestCamelName(t *testing.T) {
	fmt.Print(CamelName("test_util"))
}