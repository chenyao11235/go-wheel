package client

import (
    "context"
    "time"
)

type Client interface {
    Invoke(ctx context.Context, req, rsp interface{}, path string, opts ...Option) error
}

type defaultClient struct {
    opts *Options
}

func (c *defaultClient) Call(ctx context.Context, servicePath string, req interface{}, rsp interface{}, opts ...Option) error {
    callOpts := make([]Option, 0, len(opts)+1)
    callOpts = append(callOpts, opts...)
    callOpts = append(callOpts, withMethod())

    err := c.Invoke(ctx, req, rsp, servicePath, callOpts)
    if err != nil {
        return err
    }

    return nil
}


var DefaultClient = New()

var New = func() *defaultClient {
    return &defaultClient{
        opts: &Options{
            protocol: "proto",
        },
    }
}
