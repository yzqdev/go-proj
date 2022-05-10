package goutils

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Print(RandString(12))
}