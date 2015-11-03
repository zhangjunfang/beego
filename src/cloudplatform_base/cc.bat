set CGO_ENABLED=0
:::::::::::::::::::::x86:::::::::::::::::::::::::::::::::::::::
set GOARCH=386
 
set GOOS=windows
call make.bat --no-clean
  
set GOOS=linux
call make.bat --no-clean
  
::set GOOS=freebsd
::call make.bat --no-clean
  
::set GOOS=darwin
::call make.bat --no-clean
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
  
::::::::::::::::::::::x64:::::::::::::::::::::::::::::::::::::::
set GOARCH=amd64
 
set GOOS=linux
call make.bat --no-clean
 
::set GOOS=freebsd
::call make.bat --no-clean
  
::set GOOS=darwin
::call make.bat --no-clean
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
  
:::::::::::::::::::::::arm::::::::::::::::::::::::::::::::::::::
::set GOARCH=arm
::set GOOS=linux
::call make.bat --no-clean
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
  
::::::::::::::::::::::install gocode::::::::::::::::::::::::::::
set GOARCH=386
set GOOS=windows
go get github.com/nsf/gocode
pause