package main

import (
	"flag"
	"fmt"
	"os"
	"tartools/tar"
)

var (
	path string
	o    string
)

func init() {
	flag.StringVar(&path, "p", "nil", "打包路径")
	flag.StringVar(&o, "o", "test.tgz", "输出文件名")
}

func main() {
	flag.Parse()
	if path == "nil" {
		path, _ = os.Getwd()
	}
	path += "\\"
	fmt.Println(path)
	tar.Compose2(path, o)
}
