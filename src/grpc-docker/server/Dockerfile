# Multi stage build

# Build stage I : Go lang and Alpine Linux is only needed to build the program
#FROM golang:1.11-alpine AS build
# 构建程序只需要Go语言和Alpine Linux
FROM golang AS build

ENV location /grpc-up-and-running/src/grpc-docker

WORKDIR ${location}/server

ADD ./server ${location}/server
ADD ./proto-gen ${location}/proto-gen

#ADD main.go ${location}/server
#ADD ../proto-gen ${location}/proto-gen

# Download all the dependencies
# 下载所有的依赖项。
RUN go get -d ./...
# Install the package
# 安装所有的包。
RUN go install ./...

# 构建服务器端应用程序。
RUN CGO_ENABLED=0 go build -o /bin/grpc-productinfo-server

# Build stage II : Go binaries are self-contained executables.
# Go二进制文件是自包含的可执行文件。
FROM scratch
# 将在上一阶段构建的二进制文件复制到新的位置。
COPY --from=build /bin/grpc-productinfo-server /bin/grpc-productinfo-server

ENTRYPOINT ["/bin/grpc-productinfo-server"]
EXPOSE 50051