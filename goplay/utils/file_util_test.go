package utils

import (
	"github.com/gookit/color"
	"io/fs"
	"path/filepath"
	"testing"
)

func TestWalkByTime(t *testing.T) {
	WalkByTime("C:\\Users\\yanni\\Optimus", func(path string, info fs.FileInfo, err error) error {
		color.Redln(info.Name())
		return nil
	})
}
func TestWak(t *testing.T) {
	filepath.Walk("C:\\Users\\yanni\\Optimus", func(path string, info fs.FileInfo, err error) error {
		color.Redln(info.Name())
		return nil
	})
}
