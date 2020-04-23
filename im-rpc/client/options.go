package client

import (
    "time"
    "wheel/cyrpc/interceptor"
    "wheel/cyrpc/transport"
)

type Options struct {
    serviceName       string
    method            string
    target            string
    timeout           time.Duration
    network           string
    protocol          string
    serializationType string
    transportOpts     transport.ClientTransportOptions
    interceptors      []interceptor.ClientInterceptor
    selectorName      string
}

type Option func(*Options)

func WithServiceName(serviceName string) Option {
    return func(options *Options) {
        options.serviceName = serviceName
    }
}

func withMethod(method string) Option {
    return func(options *Options) {
        options.method = method
    }
}
