package discovery

import (
    "google.golang.org/grpc/resolver"
    "google.golang.org/grpc/resolver/manual"
)


func NewDefaultResolver() resolver.Builder {
    r, _ := manual.GenerateAndRegisterManualResolver()
    //defer cleanup()
    r.InitialState(resolver.State{
        Addresses: []resolver.Address{
            {Addr: "127.0.0.1:50051"},
            {Addr: "127.0.0.1:50052"},
        },
    })
    return r
}



