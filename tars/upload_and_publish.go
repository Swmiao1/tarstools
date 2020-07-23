package tars

import (
	"encoding/json"
	"fmt"
	"github.com/Swmiao1/tarstools/config"
	"github.com/Swmiao1/tarstools/util"
	"io/ioutil"
	"strings"
)

func (t *Tars) UploadAndPublic(file string) {
	url := t.Url + uploadAndPublish
	clint := util.NewClient(url)
	//设置Token
	clint.Header["Cookie"] = fmt.Sprintf("uid=admin; ticket=%v; dcache=true", t.Token)
	clint.Params["application"] = config.Config.App
	clint.Params["module_name"] = config.Config.Service
	clint.Params["comment"] = "TarsTool Auto upload"
	clint.Files["suse"] = file
	Response := clint.Post()
	body, _ := ioutil.ReadAll(Response.Body)
	if Response.StatusCode == 200 {
		data := uploadPatchPackageResponse{}
		_ = json.Unmarshal(body, &data)
		ResponseStr := string(body)
		if strings.Contains(ResponseStr, "EM_I_SUCCESS") {
			fmt.Printf("上传成功,服务正在激活,请检查服务状态\n")
		} else {
			temp := strings.Split(ResponseStr, "\n")
			l := len(temp)
			if l > 1 {
				fmt.Printf("上传失败,%v\n", temp[l-1])
			} else {
				fmt.Printf("上传失败,%v\n", ResponseStr)
			}

		}

	} else {
		fmt.Println(string(body))
	}
	Response.Body.Close()

}
