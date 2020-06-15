package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"tartools/cmd"
	"tartools/tar"
	"tartools/upload"
	"time"
)

var (
	app      string
	service  string
	isUpload bool
	isPublic bool
)

func init() {
	flag.BoolVar(&isUpload, "u", false, "是否上传")
	flag.BoolVar(&isPublic, "p", false, "是否发布")
}

func main() {
	//绑定参数
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Printf("build <应用名> <服务名>")
		return
	}
	app = args[0]
	service = args[1]
	fmt.Printf("app:%v service:%v\n", app, service)
	//生成临时目录
	tempPath := fmt.Sprintf("temp_%v_%v\\", app, time.Now().Nanosecond())
	//创建文件夹
	makeTempDir(tempPath)
	//删除目录
	defer DelTempDir(tempPath)
	//编译
	build(tempPath)
	//生成随机文件名
	tgzPath := fmt.Sprintf("%v_%v_%v.tgz", app, service, time.Now().Format("01_02_15_04_05"))
	//打包压缩文件
	tar.Compose2(tempPath, tgzPath)
	//上传
	if isUpload {
		//获取配置
		file, _ := os.Open("tools_config.json")
		defer file.Close()
		//读取文件
		decoder := json.NewDecoder(file)
		conf := upload.Config{}
		err := decoder.Decode(&conf)
		if err != nil {
			fmt.Println("json:", err)
			return
		}
		println("正在上传:", tgzPath)
		if isPublic {
			conf.UploadMod = 1
		}
		conf.Upload(app, service, tgzPath)
	}
}

//创建临时目录
func makeTempDir(path string) {
	err := os.MkdirAll(path+service, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

//创建临时目录
func DelTempDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//编译程序
func build(path string) {
	println("正在编译")
	ok, err := cmd.Build(path + service + "\\" + service)
	if !ok {
		panic(err)
	}
	println("编译成功")
}
