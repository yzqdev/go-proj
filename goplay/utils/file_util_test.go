package utils

import (
	"github.com/gookit/color"
	"io/fs"
	"os/exec"
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
func TestExe(t *testing.T) {
	path := "C:\\Users\\yanni\\AppData\\Roaming\\optimus"
	exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Run()
}
