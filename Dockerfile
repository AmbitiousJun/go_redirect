# 第一阶段：构建阶段
FROM golang:1.21 AS builder

# 设置工作目录
WORKDIR /app

# 设置代理
RUN go env -w GOPROXY=https://goproxy.cn

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源码
COPY . .

# 编译源码
RUN go build -o main .

# 第二阶段：运行阶段
FROM ubuntu:latest

# 设置时区
RUN apt-get update && apt-get install -y tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译后的二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 5555

# 运行应用程序
CMD ["./main"]