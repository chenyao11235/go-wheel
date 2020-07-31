#!/bin/sh

# 默认的namespace是 go.micro.api
micro --registry="etcd" --registry_address=127.0.0.1:2379 api