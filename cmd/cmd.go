package cmd

import (
	"fmt"
	"github.com/axgle/mahonia"
	"os/exec"
	"strings"
)

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func Build(path string, tag string) (bool, error) {
	cmd := exec.Command("build", path, tag)
	buf, err := cmd.Output()
	ret := string(ConvertToByte(string(buf), "gbk", "utf8"))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(ret)
	}
	if strings.Contains(ret, "编译成功") {
		return true, nil
	} else {
		return false, fmt.Errorf(ret)
	}
}
