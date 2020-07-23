package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Creates a new file upload http request with optional extra params
func (c *Config) newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	//写入参数
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	//写入header
	request.Header.Add("Content-Type", writer.FormDataContentType())
	//request.Header.Set("Cookie","uid=admin; ticket=ea201910-a963-11ea-9406-6d4c4043fa2b; dcache=true")
	request.Header.Set("Cookie", fmt.Sprintf("uid=admin; ticket=%v; dcache=true", c.Token))
	return request, err
}

//上传文件到tars
func (c *Config) Upload(app string, service string, file string) {

	extraParams := map[string]string{
		"application": app,
		"module_name": service,
		"comment":     "tools-auto-upload",
		"task_id":     string(time.Now().Nanosecond()),
	}
	var url string
	if c.UploadMod == 1 {
		//upload_and_publish 上传并发布
		url = c.TarsUrl + "/api/upload_and_publish"
	} else {
		//upload_patch_package 上传
		url = c.TarsUrl + "/pages/server/api/upload_patch_package"
	}
	fmt.Println("post:" + url)
	request, err := c.newfileUploadRequest(url, extraParams, "suse", file)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	//请求
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	if c.UploadMod == 1 {
		println(string(body))
	} else {
		data := uploadPatchPackageOk{}
		json.Unmarshal(body, &data)
		if data.RetCode == 200 {
			fmt.Printf("上传成功: 编号:%v", data.Data.Id)
		} else {
			fmt.Printf("上传失败:%v", data.ErrMsg)
		}

	}
}

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
