package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	App         string            `json:"app"`
	Service     string            `json:"service"`
	Tag         string            `json:"-"`
	Token       string            `json:"token"`
	TarsUrl     string            `json:"tars_url"`
	IncludeFile map[string]string `json:"include_file"`
	IsUpload    bool              `json:"-"`
	IsPublic    bool              `json:"-"`
	IsClear     bool              `json:"-"`
	IsHelp      bool              `json:"-"`
	Status      bool              `json:"-"`
}

func (c *config) ReadFile() bool {
	//获取配置
	file, err := os.Open("tools_config.json")
	if err != nil {
		c.createConfig()
		return false
	}
	tempConfig := new(config)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(tempConfig)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "解析配置失败:", err)
		return false
	}
	//验证配置是否正确
	if c.App == "" {
		if tempConfig.App == "" {
			_, _ = fmt.Fprintln(os.Stderr, "未设置应用名称")
			return false
		}
		c.App = tempConfig.App
	}
	if c.Service == "" {
		if tempConfig.Service == "" {
			_, _ = fmt.Fprintln(os.Stderr, "未设置应用名称")
			return false
		}
		c.Service = tempConfig.Service
	}
	c.Token = tempConfig.Token
	c.TarsUrl = tempConfig.TarsUrl
	c.IncludeFile = tempConfig.IncludeFile
	return true
}

//创建配置文件
func (c *config) createConfig() {
	fmt.Print("未找到配置文件,是否创建默认文件(y)/n:")
	ret := ""
	_, _ = fmt.Scanln(&ret)
	if ret != "n" {
		configFile, _ := os.Create("tools_config.json")
		defer configFile.Close()
		configTemp, _ := json.Marshal(config{})
		_, _ = io.WriteString(configFile, string(configTemp))
		fmt.Println("创建完成,请修改配置后重新运行")
	}
}

var (
	Config config
)

func init() {
	flag.BoolVar(&Config.IsUpload, "u", false, "是否上传")
	flag.BoolVar(&Config.IsPublic, "p", false, "是否发布")
	flag.BoolVar(&Config.IsClear, "clear", false, "清理tgz")
	flag.BoolVar(&Config.IsHelp, "h", false, "帮助")
	flag.BoolVar(&Config.Status, "status", false, "查看状态")

	flag.StringVar(&Config.Tag, "t", "", "编译标签")

	flag.StringVar(&Config.App, "a", "", "应用名称")
	flag.StringVar(&Config.Service, "s", "", "服务名称")
	flag.Parse()
}
