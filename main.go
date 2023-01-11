package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

func main() {

	cmd := exec.Command("crontab", "-l")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// outStr, errStr := stdout.String(), stderr.String()
	// fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	// fmt.Println(strings.Split(outStr, "\n"))
	lines := strings.Split(stdout.String(), "\n")
	f := openFile("./rerun.sh")

	defer f.Close()

	write(f, "#!/bin/sh\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		arr := strings.Split(line, " ")

		if strings.Contains(arr[0], "#") {
			fmt.Println(line)
			continue
		}

		_, index, ok := lo.FindIndexOf(arr, func(i string) bool {
			return i == "task"
		})

		if !ok {
			fmt.Printf("无效的任务表达式 %s \n\n", line)
			continue
		}

		sub := lo.Slice(arr, index, len(arr))
		content := fmt.Sprintf("%s now \n", strings.Join(sub, " "))

		write(f, content)
	}

}

func write(f *os.File, content string) {
	_, err := io.WriteString(f, content)

	if err != nil {
		panic(content + fmt.Sprintf("write fail: %v", err))
	}

	// fmt.Printf("写入 %d 个字节\n", n)
}

func openFile(path string) *os.File {
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
