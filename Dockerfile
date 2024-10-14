FROM golang:latest
 
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
 
# 移动到工作目录：/usr/local/goproject
WORKDIR /app
 
# 将代码复制到容器中
COPY . .

RUN go mod download

# 将我们的代码编译成二进制可执行文件app
RUN go build -o main .
 
# 声明服务端口
EXPOSE 8080
 
# 启动容器时运行的命令
CMD ["/app/main"]
