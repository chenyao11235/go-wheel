package discovery

import (
    "google.golang.org/grpc/balancer"
    "google.golang.org/grpc/balancer/base"
)

// grpc内置的有pick_first和round_robin两种负载均衡算法，这里自定义一个权重负载均衡器

const Name = "weighted_round_robin"


// AddrInfo will be stored inside Address metadata in order to use weighted roundrobin
// balancer.
type AddrInfo struct {
    Weight uint32
}

type wrrPickerBuilder struct {
}

func (r wrrPickerBuilder) Build(info base.PickerBuildInfo) balancer.V2Picker {
    panic("implement me")
}

type wrrPicker struct {
}

func (w wrrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
    panic("implement me")
}
