# goctl

[English](readme.md) | 简体中文

goctl 使用见文档 https://go-zero.dev/docs/tutorials/cli/overview

# 新增功能
生成rpc 框架代码时支持忽略生成protobuf的代码，如
goctl rpc protoc xx.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. -i

-i, --ignore_pb         Whether to ignore generate pb & grpc files
