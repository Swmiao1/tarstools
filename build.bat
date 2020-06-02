@echo off
set APP=%1
set SERVER=%2
set GOOS=linux
set TIME=%time:~0,2%_%time:~3,2%_%time:~6,2%
set TAR=%APP%_%SERVER%_%date:~5,2%%date:~8,2%
mkdir temp_%TIME%\%SERVER%
go build -o temp_%TIME%\%SERVER%\%SERVER% && packtar -p=%cd%\temp_%TIME% -o=%TAR%_%TIME%.tgz
rd /S /Q %cd%\temp_%TIME%
