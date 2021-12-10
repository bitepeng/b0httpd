set GOARCH=amd64
set GOOS=windows
go build -o bin/win64/easyhttpd.exe .\main.go

set GOARCH=amd64
set GOOS=linux
go build -o bin/linux/easyhttpd .\main.go

set GOARCH=amd64
set GOOS=darwin
go build -o bin/mac/easyhttpd .\main.go