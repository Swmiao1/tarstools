@echo off
set GOOS=linux
if  "%1"== "" (set FILE_NAME=) else set FILE_NAME=-o %1
if  "%2"== "" (set FILE_NAME=) else set BUILD_TGA=-tags %2
echo %cd%
echo ���ڱ��� go build -ldflags "-s -w" %BUILD_TGA% %FILE_NAME%
go build -ldflags "-s -w" %BUILD_TGA% %FILE_NAME%&& echo ����ɹ� && goto OK
echo ����ʧ��
:OK