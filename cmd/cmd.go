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

func Build(path string) (bool, error) {
	cmd := exec.Command("build", path)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	ret := string(ConvertToByte(string(buf), "gbk", "utf8"))
	if strings.Contains(ret, "编译成功") {
		return true, nil
	} else {
		return false, fmt.Errorf(ret)
	}
}
