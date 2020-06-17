package my_resolver

import (
    "context"
    mvccpb2 "github.com/coreos/etcd/mvcc/mvccpb"
    "github.com/pkg/errors"
    "go.etcd.io/etcd/clientv3"
    "google.golang.org/grpc/resolver"
    "log"
    "time"
)

var etcdAddrs = []string{"127.0.0.1:2379"}

type DynamicResolverBuilder struct {
    etcdAddrs []string
}

func NewDynamicResolverBuilder() resolver.Builder {
    return &DynamicResolverBuilder{etcdAddrs: etcdAddrs}
}

func (d DynamicResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
    r := &DynamicResolver{
        etcdAddrs: d.etcdAddrs,
        target:    target,
        cc:        cc,
    }

    err := r.start()
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (d DynamicResolverBuilder) Scheme() string {
    return "dynamic"
}

type DynamicResolver struct {
    etcdAddrs  []string
    target     resolver.Target
    cc         resolver.ClientConn
    addrsStore map[string][]string
}

func (d *DynamicResolver) start() (err error) {
    var (
        client      *clientv3.Client
        config      clientv3.Config
        serverAddrs []resolver.Address
    )
    // 连接etcd
    config = clientv3.Config{
        Endpoints:   d.etcdAddrs,
        DialTimeout: time.Second * 3,
    }
    client, err = clientv3.New(config)
    if err != nil {
        err = errors.Wrap(err, "connect etcd fail")
    }
    // 监控etcd操作，动态更新resolver
    go func() {
        getResp, err := client.Get(context.Background(), d.target.Endpoint, clientv3.WithPrefix())
        if err != nil {
            log.Println(errors.Wrap(err, "获取节点列表失败"))
            return
        }

        for _, kvpair := range getResp.Kvs {
            addr := resolver.Address{Addr: string(kvpair.Value)}
            serverAddrs = append(serverAddrs, addr)
        }

        // 更新resolver中的服务地址
        d.cc.UpdateState(resolver.State{
            Addresses: serverAddrs,
        })
        // 监听服务端的节点变化
        watchStartVersion := getResp.Header.Revision + 1 // 从get时刻的后续版本进行watch
        watchChan := client.Watch(context.Background(), d.target.Endpoint, clientv3.WithRev(watchStartVersion), clientv3.WithPrefix())
        for watchResp := range watchChan {
            for _, watchEvent := range watchResp.Events {
                addr := string(watchEvent.Kv.Value)
                switch watchEvent.Type {
                case mvccpb2.PUT:
                    if !exist(serverAddrs, addr) {
                        serverAddrs = append(serverAddrs, resolver.Address{Addr: addr})
                        d.cc.UpdateState(resolver.State{
                            Addresses: serverAddrs,
                        })
                    }
                case mvccpb2.DELETE:
                    if s, ok := remove(serverAddrs, addr); ok {
                        serverAddrs = s
                        d.cc.UpdateState(resolver.State{
                            Addresses: serverAddrs,
                        })
                    }
                }
            }
        }
    }()
    return
}

func (d DynamicResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (d DynamicResolver) Close() {}



func exist(l []resolver.Address, addr string) bool {
    for i := range l {
        if l[i].Addr == addr {
            return true
        }
    }
    return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
    for i := range s {
        if s[i].Addr == addr {
            s[i] = s[len(s)-1]
            return s[:len(s)-1], true
        }
    }
    return nil, false
}
