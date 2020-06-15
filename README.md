# 安装

克隆代码到本地

`git clone https://gitee.com/N-age/TARSTOOL`

然后运行 `init.cmd` 

之后就可以打包TARS文件了

# 使用

在项目目录下打开 `cmd`  或 在`IDE`中打开 `Terminal` 工具

输入 `packtar 应用名 服务名` 即可生成打包文件 `.tgz`
在项目目录下新建 tools_config.json
```json
{
  "tars_url" : "http://47.57.119.12:3000",
  "token" : ""
}
```
可用参数 `-u` 上传到 tars `-p` 发布

```powershell
build -u -p im Hello
```

