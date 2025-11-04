go env -w GOOS=windows
go build -o spider.exe -ldflags "-s -w" main.go