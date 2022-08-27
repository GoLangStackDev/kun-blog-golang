# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.17-alpine3.16 as builder
# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

# 编译代码
RUN go build ./cmd/kserver/kserver.go

RUN mkdir publish && cp kserver publish && \
    cp -r public publish

FROM alpine:3.16

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/publish .

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 80

ENTRYPOINT ["./kserver"]
