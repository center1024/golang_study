setlocal
if exist install.bat goto ok
echo install.bat must be run from its folder
goto end
:ok
set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0
gofmt -w src
go install lession_code
:end
echo finished