go env -w GOOS=linux
go build -o spider -ldflags "-s -w" main.go