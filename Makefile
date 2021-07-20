.PHONY: all build run gotool clean help

BINARY="eustoma"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

run:
	@go run ./main.go conf/systeminfo.ini

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 GO 代码 ，并编译生成二进制文件"
	@echo "make build - 编译 GO 代码，生成二进制文件"
	@echo "make run - 直接运行 GO 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 go 工具 'fmt' and 'vet' 格式化和检查代码"

