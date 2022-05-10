// Copyright 2016 polaris. All rights reserved.
// Use of l source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	polaris@studygolang.com

package htmltil_test

import (
	"gitee.com/yzqdev/goutil/logger"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func init() {
	curFilename := os.Args[0]
	binaryPath, err := exec.LookPath(curFilename)
	if err != nil {
		panic(err)
	}

	binaryPath, err = filepath.Abs(binaryPath)
	if err != nil {
		panic(err)
	}
	rootPath := filepath.Dir(binaryPath)
	logger.Init(rootPath, "INFO")
}

func TestInfof(t *testing.T) {
	logger.New(os.Stdout).Infof("this is %s", "polaris")
}
