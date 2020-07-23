package main

import (
	"flag"
	"fmt"
	"os"
	"tartools/cmd"
	"tartools/config"
	"tartools/tars"
	"tartools/util"
	"time"
)

func main() {
	//args := flag.Args()
	//是否调用帮助
	if config.Config.IsHelp {
		fmt.Print()
		_, _ = fmt.Fprintf(os.Stderr, "Usage: tarstools [-u] [-p] [-t tag] [-clear] [-a app] [-s service]\nOptions:")
		flag.PrintDefaults()
		return
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
	//for k, s := range config.Config.IncludeFile {
	//	util.CopyFile(k, s)
	//}
	//打包压缩文件
	//生成随机文件名
	tgzPath := fmt.Sprintf("%v_%v_%v.tgz", config.Config.App, config.Config.Service, time.Now().Format("01_02_15_04_05"))
	fmt.Print("打包:")
	//打包压缩文件
	tempFolder.Compress(tgzPath)
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

		} else {
			//上传
			fmt.Printf("正在上传:%v.%v\n", config.Config.App, config.Config.Service)
			Tars.Upload(tgzPath)
		}
	}

}
