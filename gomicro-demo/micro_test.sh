#!/bin/sh

# 使用micro工具测试我们定义的服务， 首先需要安装 go get github.com/micro/micro 会在gopath/bin下生成micro工具


# 查看当前etcd中一共注册了多少个服务
micro --registry etcd --registry_address 127.0.0.1:2379 list services

# 查看目标服务的元信息
micro --registry etcd --registry_address 127.0.0.1:2379 get service productservice

# 调用具体的服务接口
micro --registry etcd --registry_address 127.0.0.1:2379 call productservice ProductService.GetProdDetail "{\"ID\":3}"
micro --registry etcd --registry_address 127.0.0.1:2379 call productservice ProductService.GetProdList "{\"size\":10}"
