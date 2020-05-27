@echo off
go build -o packtar.exe
copy build.bat %GOPATH%\bin
move packtar.exe %GOPATH%\bin
@pause
