package util

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type Client struct {
	Url    string
	Params map[string]string
	Files  map[string]string
	Header map[string]string
}

func NewClient(url string) *Client {
	return &Client{
		Url:    url,
		Params: make(map[string]string),
		Files:  make(map[string]string),
		Header: make(map[string]string),
	}
}

func (c *Client) Post() *http.Response {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)
	fmt.Println("POST:", c.Url)
	//创建请求数据
	request, err = c.createRequest()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil
	}
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return response
}

func (c *Client) Get() *http.Response {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)
	fmt.Println("POST:", c.Url)
	//创建请求数据
	request, err = c.createRequest(true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil
	}
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return response
}

func (c *Client) createRequest(Get ...bool) (Request *http.Request, err error) {
	//创建body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//写入文件
	for key, val := range c.Files {
		//打开文件
		file, err := os.Open(val)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return nil, err
		}
		//创建文件
		part, err := writer.CreateFormFile(key, val)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		//复制文件
		_, err = io.Copy(part, file)
		_ = file.Close()
	}
	//写入普通参数
	for key, val := range c.Params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()
	if len(Get) > 0 {
		Request, err = http.NewRequest("GET", c.Url, body)
	} else {
		Request, err = http.NewRequest("POST", c.Url, body)
	}

	if err != nil {
		return nil, err
	}
	//写入header
	c.Header["Content-Type"] = writer.FormDataContentType()
	for key, val := range c.Header {
		//写入header
		Request.Header.Set(key, val)
	}
	return
}
