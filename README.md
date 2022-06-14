# blog-go

```
# 安装grpcurl工具
$ go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl

# 查看注册的grpc服务 (通过反射注册)
$ grpcurl -plaintext localhost:8001 list
grpc.reflection.v1alpha.ServerReflection
proto.TagService

# 查看选择的反射方法中的grpc方法
$ grpcurl -plaintext localhost:8001 list proto.TagService
proto.TagService.GetTagList

# 调用grpc方法
$ grpcurl -plaintext -d '{"name":"Go"}' localhost:8001 proto.TagService.GetTagList  
{
  "list": [
    {
      "id": "1",
      "name": "Go",
      "state": 1
    }
  ],
  "pager": {
    "page": "1",
    "pageSize": "10",
    "totalRows": "1"
  }
}
```


```
# 生成proto代码文件
protoc --go_out=plugins=grpc:. proto/*.proto
# 生成grpc-gateway proto代码文件
protoc -I. -I"$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis" --grpc-gateway_out=logtostderr=true:. proto/*.proto
# 生成 swagger文件
protoc -I. -I"$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis" --swagger_out=logtostderr=true:. proto/*.proto 
```