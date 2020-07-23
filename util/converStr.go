package util

import "github.com/axgle/mahonia"

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	//新建
	srcCoder := mahonia.NewDecoder(srcCode)
	//填入数据
	srcResult := srcCoder.ConvertString(src)
	//新建
	tagCoder := mahonia.NewDecoder(targetCode)
	//转换
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
