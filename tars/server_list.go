package tars

import (
	"encoding/json"
	"fmt"
	"github.com/Swmiao1/tarstools/util"
	"io/ioutil"
	"os"
)

type serverListResponse struct {
	Data    []serverListResponseData
	RetCode int    `json:"ret_code"`
	ErrMsg  string `json:"err_msg"`
}
type serverListResponseData struct {
	Id           int    `json:"id"`
	Application  string `json:"application"`
	ServerName   string `json:"server_name"`
	NodeName     string `json:"node_name"`
	SettingState bool   `json:"setting_state"`
	PresentState bool   `json:"present_state"`
	PatchVersion string `json:"patch_version"`
}

func (t *Tars) ServerList(app string, service string) {
	url := t.Url + serverList
	clint := util.NewClient(url)
	//设置Token
	clint.Header["Cookie"] = fmt.Sprintf("uid=admin; ticket=%v; dcache=true", t.Token)
	//设置参数
	clint.Params["tree_node_id"] = fmt.Sprintf("1%v.5%v", app, service)
	Response := clint.Post()
	defer Response.Body.Close()
	body, _ := ioutil.ReadAll(Response.Body)
	if Response.StatusCode == 200 {
		data := serverListResponse{}
		_ = json.Unmarshal(body, &data)
		if data.RetCode == 200 {
			if len(data.Data) < 1 {
				fmt.Fprintln(os.Stderr, "服务不存在")
			} else {
				fmt.Printf("服务状态:%v,设置状态:%v,请检查服务状态\n", data.Data[0].PresentState, data.Data[0].SettingState)
			}
		} else {
			fmt.Fprintln(os.Stderr, data.ErrMsg)
		}
	} else {
		fmt.Println(string(body))
	}

}
