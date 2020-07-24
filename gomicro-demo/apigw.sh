#!/bin/sh

export MICOR_REGISTRY=etcd
export MICOR_REGISTRY_ADDRESS=127.0.0.1:2379
export MICRO_API_NAMESPACE=api.product.com
export MICRO_API_HANDLER=rcp
micro api