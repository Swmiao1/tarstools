@echo off
set GOOS=linux
if  "%1"== "" (set FILE_NAME=) else set FILE_NAME= -o %1
echo %cd%
echo ���ڱ��� go build %cd%%FILE_NAME%
go build%FILE_NAME%&& echo ����ɹ� && goto OK
echo ����ʧ��
:OK