package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().After(fis[j].ModTime())
}
func walk(root string, info fs.FileInfo, fn filepath.WalkFunc) {
	f, _ := os.Open(root)
	fis, _ := f.Readdir(-1)
	defer f.Close()
	sort.Sort(ByModTime(fis))
	//for _, fi := range fis {
	//	color.Redln(fi.Name(), fi.ModTime())
	//}
}
func WalkByTime(root string, fn filepath.WalkFunc) error {

	info, err := os.Lstat(root)

	if err != nil {
		err = fn(root, nil, err)
	} else {
		walk(root, info, fn)
	}
	if err == filepath.SkipDir {
		return nil
	}
	return err
}
