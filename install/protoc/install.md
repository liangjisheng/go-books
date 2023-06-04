# protoc

[quickstart](https://grpc.io/docs/languages/go/quickstart/)
[install](https://www.cnblogs.com/ayanmw/p/16525561.html)

## install protoc

打开[仓库](https://github.com/protocolbuffers/protobuf)的 release 页面直接下载二进制文件

## install protoc-gen-go

```shell
# 这个后面对应的是 https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go
# 至于为啥这样安装, google 就是这样的
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install github.com/protocolbuffers/protobuf-go/cmd/protoc-gen-go@latest # 这是错误的

# 下面的这个就不用了
# go install github.com/golang/protobuf/protoc-gen-go@latest
```

版本报错: The import path must contain at least one forward slash ('/') character.  
参考这个文章: [https://www.jianshu.com/p/bbf2eb22b021](https://www.jianshu.com/p/bbf2eb22b021)

## install protoc-gen-go-grpc

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## install protoc-gen-grpc-gateway

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

## install protoc-gen-openapiv2

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

## install protoc-go-inject-tag

```shell
go install github.com/favadi/protoc-go-inject-tag@latest
```
