go env -w GOOS=windows
go build -o spider -ldflags "-s -w" main.go