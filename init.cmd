@echo off
SET GO111MODULE=on
SET GOPROXY=http://goproxy.cn
go build -o packtar.exe
copy build.bat %GOPATH%\bin
move packtar.exe %GOPATH%\bin
@pause
