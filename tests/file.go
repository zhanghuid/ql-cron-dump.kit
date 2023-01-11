package tests

import (
	"os"
	"path/filepath"
)

func OpenFile(path string) *os.File {
	s, err := os.Stat(path)

	if err == nil {
		// 是文件
		if !s.IsDir() {
			f2, err2 := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, os.ModeAppend)
			wrapError(err2)
			return f2
		}
	}

	// 是目录，得判断一下目录是否存在
	// 如果不存在
	err3 := os.MkdirAll(filepath.Dir(path), os.FileMode(0775))
	wrapError(err3)
	// 先创建文件
	f, err4 := os.Create(path)
	wrapError(err4)

	return f
}

func wrapError(err error) {
	if err != nil {
		panic(err)
	}
}
