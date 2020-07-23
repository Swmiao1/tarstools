package util

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) {
	source, err := os.Open(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "文件不存在")
		return
	}
	destination, _ := os.Create(dst)
	buf := make([]byte, 1024)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return
		}
	}
}
