@echo off
set GOOS=linux
if  "%1"== "" (set FILE_NAME=) else set FILE_NAME= -o %1
echo %cd%
echo 正在编译 go build %cd%%FILE_NAME%
go build%FILE_NAME%&& echo 编译成功 && goto OK
echo 编译失败
:OK