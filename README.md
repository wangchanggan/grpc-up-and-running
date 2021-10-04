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