package tar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type fileInfo struct {
	info     os.FileInfo
	BasePath string
	path     string
}

func Compose2(path string, output string) {

	//获取文件
	fmt.Println("正在获取列表")
	list := SearchFile(path)
	compress(output, list)
}

//获取文件列表
func SearchFile(BasePath string) *[]fileInfo {
	var list []fileInfo
	filepath.Walk(BasePath, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if path != BasePath {
			list = append(list, fileInfo{
				info:     info,
				path:     path,
				BasePath: strings.TrimPrefix(path, BasePath),
			})
		}
		return nil
	})
	return &list
}

//压缩 使用gzip压缩成tar.gz
func compress(dest string, list *[]fileInfo) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gzipWriter := gzip.NewWriter(d)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	for _, file := range *list {
		err := packOne(&file, tarWriter)
		if err != nil {
			return err
		}
	}
	fmt.Println()
	return nil
}

func packOne(data *fileInfo, tarWriter *tar.Writer) error {
	println(data.BasePath)
	if data.info.IsDir() {
		if data.BasePath == "" {
			return nil
		}
		header, err := tar.FileInfoHeader(data.info, "")
		if err != nil {
			return err
		}
		header.Name = data.BasePath
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}
		//os.Mkdir(data.BasePath, os.ModeDir)
	} else {
		//打开文件
		buf, err := os.Open(data.path)
		if err != nil {
			fmt.Println("文件打开失败", err)
			panic(err)
		}
		defer buf.Close()
		header, err := tar.FileInfoHeader(data.info, "")
		if err != nil {
			return err
		}
		header.Name = data.BasePath
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}
		//复制
		if _, err = io.Copy(tarWriter, buf); err != nil {
			return err
		}
	}
	return nil
}
