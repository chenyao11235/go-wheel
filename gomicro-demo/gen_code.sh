#!/bin/sh

# 代码生成
cd protos
protoc --micro_out=./ --go_out=./ product.proto
# 修改tag  使用这个工具 go get -v github.com/favadi/protoc-go-inject-tag
# 可以使得在proto文件中自定义的message中的字段的tag  
# 具体如何自定义tag 见proto文件
protoc-go-inject-tag -input=./product.pb.go
cd ..