FROM golang:1.16.5-alpine3.13

#为golang镜像设置必要的环境变量

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

#移动到工作目录
WORKDIR /build

#将代码复制到容器中
COPY . .

#将我们的代码编译成二进制app
RUN go build -o app .


#移动到用于存放生成的二进制/dist目录
#WORKDIR /dist

#将二进制文件从/build 目录复制到这里
#RUN cp /build/app .


EXPOSE 8888

CMD ["/dist/app"]