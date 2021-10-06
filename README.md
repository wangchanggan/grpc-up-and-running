# gPRC-go与云原生应用开发
Source Code From https://github.com/grpc-up-and-running/samples/archive/refs/tags/v1.0.0.tar.gz

参考gRPC与云原生应用开发（gRPC:Up and Running）（张卫滨/译）

## 目录
-   [gPRC-go与云原生应用开发](#gprc-go与云原生应用开发)
    -   [目录](#目录)
    -   [gRPC入门](#grpc入门)
    -   [开始使用gRPC](#开始使用grpc)
    -   [gRPC的通信模式](#grpc的通信模式)
    -   [gRPC的底层原理](#grpc的底层原理)
    -   [gRPC超越基础知识](#grpc超越基础知识)
        -   [拦截器](#拦截器)
        -   [截止时间](#截止时间)
        -   [取消](#取消)
        -   [错误处理](#错误处理)
        -   [多路复用](#多路复用)
        -   [元数据](#元数据)
        -   [负载均衡](#负载均衡)
    -   [安全的gRPC](#安全的grpc)
        -   [使用TLS认证gRPC通道](#使用tls认证grpc通道)
            -   [启用单向安全连接](#启用单向安全连接)
            -   [启用mTLS保护的连接](#启用mtls保护的连接)
        -   [对gRPC调用进行认证](#对grpc调用进行认证)
            -   [使用basic认证](#使用basic认证)
            -   [使用OAuth2.0、JWT和基于令牌的谷歌认证](#使用oauth2.0jwt和基于令牌的谷歌认证)
    -   [在生产环境中运行gRPC](#在生产环境中运行grpc)
        -   [测试gRPC应用程序](#测试grpc应用程序)
        -   [部署](#部署)
            -   [部署到Docker上](#部署到docker上)
            -   [部署到Kubernetes上](#部署到kubernetes上)
        -   [可观察性](#可观察性)
            -   [度量指标](#度量指标)
                -   [在gRPC中使用OpenCensus](#在grpc中使用opencensus)
                -   [在gRPC中使用Prometheus](#在grpc中使用prometheus)
            -   [跟踪](#跟踪)
    -   [gRPC的生态系统](#grpc的生态系统)
        -   [gRPC网关](#grpc网关)
        -   [gRPC服务器端反射协议](#grpc服务器端反射协议)
        -   [健康检查协议](#健康检查协议)

## gRPC入门
docs/gRPC入门.doc

## 开始使用gRPC
docs/开始使用gRPC.doc

src/productinfo

## gRPC的通信模式
docs/gRPC的通信模式.doc

src/ordermgt

## gRPC的底层原理
docs/gRPC的底层原理.doc

## gRPC超越基础知识
docs/gRPC超越基础知识.doc

### 拦截器
src/interceptors

### 截止时间
src/deadlines

### 取消
src/cancellation

### 错误处理
src/error-handling

### 多路复用
src/multiplexing

### 元数据
src/metadata

### 负载均衡
src/loadbalancing

## 安全的gRPC
docs/安全的gRPC.doc

### 使用TLS认证gRPC通道

#### 启用单向安全连接
src/secure-channel

#### 启用mTLS保护的连接
src/mutual-tls-channel

### 对gRPC调用进行认证

#### 使用basic认证
src/basic-authentication

#### 使用OAuth2.0、JWT和基于令牌的谷歌认证
src/token-based-authentication

## 在生产环境中运行gRPC
docs/在生产环境中运行gRPC.doc

### 测试gRPC应用程序
src/grpc-continous-integration

### 部署
#### 部署到Docker上
src/grpc-docker

#### 部署到Kubernetes上
src/grpc-kubernetes

### 可观察性
#### 度量指标
##### 在gRPC中使用OpenCensus
src/grpc-opencensus

##### 在gRPC中使用Prometheus
src/grpc-prometheus

#### 跟踪
src/grpc-opencensus-tracing

src/grpc-opentracing

## gRPC的生态系统
docs/gRPC的生态系统.doc

### gRPC网关
src/grpc-gateway

### gRPC服务器端反射协议
src/server-reflection

### 健康检查协议
src/health-check
