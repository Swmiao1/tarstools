package main

import (
	"flag"
	"fmt"
	"github.com/Swmiao1/tarstools/cmd"
	"github.com/Swmiao1/tarstools/config"
	"github.com/Swmiao1/tarstools/tars"
	"github.com/Swmiao1/tarstools/util"
	"os"
	"time"
)

const version = "v0.0.5"

func main() {
	fmt.Println("version:", version)
	//是否调用帮助
	if config.Config.IsHelp {
		fmt.Print()
		_, _ = fmt.Fprintf(os.Stderr, "Usage: tarstools [-u] [-p] [-t tag] [-clear] [-a app] [-s service]\nOptions:")
		flag.PrintDefaults()
		return
	}
	//是否是清理 tgz
	if config.Config.IsClear {
		util.ClearTgz()
	}
	//读取配置文件
	if !config.Config.ReadFile() {
		return
	}
	//生成临时目录
	tempFolder := util.NewFolder(fmt.Sprintf("temp_%v_%v/%v", config.Config.Service, time.Now().Nanosecond(), config.Config.Service))
	tempFolder.Make()
	defer tempFolder.Del()
	//编译文件
	Exec := cmd.NewCmd()
	//Exec.Log()
	//设置环境变量 GOOS=linux
	Exec.Input("set GOOS=linux")
	Exec.Input("set GO111MODULE=on")
	//编译
	var tags = ""
	if len(config.Config.Tag) > 0 {
		tags = " -tags " + config.Config.Tag
	}
	buildStr := fmt.Sprintf("go build -ldflags \"-s -w\"%v -o %v", tags, tempFolder.Path+"/"+config.Config.Service)
	fmt.Println("正在编译:", buildStr)
	Exec.Input(buildStr)
	Exec.Input("exit")
	_ = Exec.Cmd.Wait()
	if Exec.IsErr {
		return
	}
	//移入包含文件
	for k, s := range config.Config.IncludeFile {
		util.CopyFile(k, tempFolder.Path+s)
	}
	//打包压缩文件
	//生成随机文件名
	tgzPath := fmt.Sprintf("%v_%v_%v.tgz", config.Config.App, config.Config.Service, time.Now().Format("01_02_15_04_05"))
	fmt.Print("打包:")
	//打包压缩文件
	tempFolder.Compress(tgzPath)
	//删除临时文件夹
	tempFolder.Del()
	//上传文件
	if config.Config.IsUpload {
		//判断配置
		if config.Config.TarsUrl == "" || config.Config.Token == "" {
			fmt.Fprintln(os.Stderr, "请配置TarsUrl和Token")
			return
		}
		Tars := tars.Tars{
			Url:   config.Config.TarsUrl,
			Token: config.Config.Token,
		}
		if config.Config.IsPublic {
			//发布
			fmt.Printf("上传并发布至:%v.%v\n", config.Config.App, config.Config.Service)
			Tars.UploadAndPublic(tgzPath)
		} else {
			//上传
			fmt.Printf("上传至:%v.%v\n", config.Config.App, config.Config.Service)
			Tars.Upload(tgzPath)
		}
		//获取服务状态
		Tars.ServerList(config.Config.App, config.Config.Service)
	}

}
