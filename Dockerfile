FROM golang:latest

# 将当前工作目录设置为应用程序的目录
WORKDIR /app

# 将本地文件复制到容器中
COPY . .

# 构建应用程序
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o chat-proxy main.go

# 设置环境变量
ENV PORT 8080

# 运行应用程序
CMD ["./chat-proxy"]