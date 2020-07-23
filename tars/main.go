package tars

import (
	"encoding/json"
	"fmt"
	"github.com/Swmiao1/tarstools/config"
	"github.com/Swmiao1/tarstools/util"
	"io/ioutil"

	"time"
)

const (
	//上传发布包
	uploadPatchPackage = "/api/upload_patch_package"
)

type Config struct {
	TarsUrl   string `json:"tars_url"`
	Token     string `json:"token"`
	UploadMod int    `json:"Upload_mod"`
}

type uploadPatchPackageOk struct {
	Data    uploadPatchPackageData `json:"data"`
	RetCode int                    `json:"ret_code"`
	ErrMsg  string                 `json:"err_msg"`
}
type uploadPatchPackageData struct {
	Id       int    `json:"id"`
	Server   string `json:"server"`
	Tgz      string `json:"tgz"`
	Comment  string `json:"comment"`
	Posttime string `json:"posttime"`
}

type Tars struct {
	Url   string `json:"tars_url"`
	Token string `json:"token"`
}

func (t *Tars) Upload(file string) {
	url := t.Url + uploadPatchPackage
	clint := util.NewClient(url)
	//设置Token
	clint.Header["Cookie"] = fmt.Sprintf("uid=admin; ticket=%v; dcache=true", config.Config.Token)
	clint.Params["application"] = config.Config.App
	clint.Params["module_name"] = config.Config.Service
	clint.Params["comment"] = "TarsTool Auto upload"
	clint.Params["task_id"] = fmt.Sprintf("%v", time.Now().Nanosecond())
	clint.Files["suse"] = file
	Response := clint.Post()
	body, _ := ioutil.ReadAll(Response.Body)
	if Response.StatusCode == 200 {
		data := uploadPatchPackageOk{}
		_ = json.Unmarshal(body, &data)
		fmt.Printf("上传成功: 编号:%v\n", data.Data.Id)
	} else {
		fmt.Println(body)
	}
	Response.Body.Close()

}
