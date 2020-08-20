package tar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func NewFile(dest string) *File {
	var err error
	temp := &File{}
	//创建文件
	temp.FileId, err = os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Gzip = gzip.NewWriter(temp.FileId)
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Tar = tar.NewWriter(temp.Gzip)
	if err != nil {
		fmt.Println(err.Error())
	}
	return temp
}

type File struct {
	FileId *os.File
	Gzip   *gzip.Writer
	Tar    *tar.Writer
}

func (f *File) Close() {
	_ = f.Tar.Close()
	_ = f.Gzip.Close()
	_ = f.FileId.Close()
}

type FileList struct {
	BasePath string // 本地路径
	FillPath string // 压缩路径
}

func (f *File) Compress(list *[]FileList) error {
	for _, file := range *list {
		//获取文件INFO
		fileInfo, err := os.Stat(file.BasePath)
		if err != nil {
			return err
		}
		//获取 HEADER
		header, err := tar.FileInfoHeader(fileInfo, "")
		if err != nil {
			return err
		}
		//转换成linux符号
		header.Name = filepath.ToSlash(file.FillPath)
		//写入文件头
		if err = f.Tar.WriteHeader(header); err != nil {
			return err
		}
		//进行压缩
		if !fileInfo.IsDir() {
			//打开文件
			fmt.Print(file.FillPath, " -")
			buf, err := os.Open(file.BasePath)
			if err != nil {
				fmt.Println("文件打开失败", err)
				panic(err)
			}
			//复制
			if _, err = io.Copy(f.Tar, buf); err != nil {
				return err
			}
			_ = buf.Close()
			fmt.Println("OK")
		}
	}
	fmt.Println("打包完成")
	return nil
}
