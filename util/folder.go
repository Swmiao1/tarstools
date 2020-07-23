package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"tarstools/tar"
)

type folder struct {
	Path string
}

func (f *folder) Make() {
	err := os.MkdirAll(f.Path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
func (f *folder) Del() {
	err := os.RemoveAll(path.Dir(f.Path))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (f *folder) Compress(fileName string) {
	//获取目录下所有文件
	list := f.getFileList()
	tarFile := tar.NewFile(fileName)
	tarFile.Compress(list)
	defer tarFile.Close()
}

//获取文件列表
func (f *folder) getFileList() *[]tar.FileList {
	//查找文件
	var list []tar.FileList
	base := path.Dir(f.Path) + "\\"
	fmt.Println("正在获取列表", base)
	_ = filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return err
		}
		//加入列表
		if path != base {
			fmt.Println(filepath.ToSlash(strings.TrimPrefix(path, base)))
			list = append(list, tar.FileList{
				Info:     info,
				FillPath: path,
				BasePath: strings.TrimPrefix(path, base),
			})
		}
		return nil
	})
	return &list
}

func NewFolder(path string) *folder {
	return &folder{Path: path}
}
