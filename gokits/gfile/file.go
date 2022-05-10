package gfile

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// FileExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	return e == nil && !!f.IsDir()
}

// IsDir returns true if given path is a directory,
// or returns false when it's a file or does not exist.
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	return e == nil && f.IsDir()
}

// FileSize returns file size in bytes and possible error.
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// GetProcPwd returns the complete directory of the current execution file
func GetProcPwd() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(filepath.Dir(file))
	return path
}

// GetPwd returns the complete directory of the current path
func GetPwd() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// CopyFile is a utility to assist with copying a file from src to dest.
// Note that file permissions are maintained.
func CopyFile(src, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, source)
	if err != nil {
		return err
	}

	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dest, sourceInfo.Mode())
}

// CopyDir is a utility to assist with copying a directory from src to dest.
// Note that directory permissions are not maintained, but the permissions of
// the files in those directories are.
func CopyDir(src, dest string, dirMode os.FileMode) error {
	dir, err := os.Open(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dest, dirMode); err != nil {
		return err
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dest, file.Name())
		if file.IsDir() {
			if err := CopyDir(srcPath, dstPath, dirMode); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}
