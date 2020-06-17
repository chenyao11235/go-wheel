package main

import (
    "flag"
    "fmt"
    "github.com/pkg/errors"
    uuid "github.com/satori/go.uuid"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    healthpb "google.golang.org/grpc/health/grpc_health_v1"
    "log"
    "net"
    "time"
    pb "wheel/grpc-demo/proto"
    "wheel/grpc-demo/server/initdata"
    "wheel/grpc-demo/server/service"
)

const (
    system       = "" // empty string represents the health of the system
    SERVICE_NAME = "book"
)

var (
    sleep = flag.Duration("sleep", time.Second*5, "duration between changes in health")
    port  = flag.Int("port", 50051, "the port to serve on")
)

func main() {
    flag.Parse()

    if err := initdata.InitData(); err != nil {
        log.Fatal(err)
    }

    regKey := fmt.Sprintf("/service/%s/%s", SERVICE_NAME, uuid.NewV4().String())

    ip, _ := getLocalIP()
    regValue := fmt.Sprintf("%s:%d", ip, *port)

    rpcServer := grpc.NewServer()
    // 健康检查
    healthcheck := health.NewServer()
    healthpb.RegisterHealthServer(rpcServer, healthcheck)
    // 书籍服务
    bookservice := service.NewBookService()
    pb.RegisterBookServiceServer(rpcServer, bookservice)

    if err := service.InitRegister(regKey, regValue); err != nil {
        log.Fatal(err)
    }

    go func() {
        // asynchronously inspect dependencies and toggle serving status as needed
        next := healthpb.HealthCheckResponse_SERVING

        for {
            healthcheck.SetServingStatus(system, next)

            if next == healthpb.HealthCheckResponse_SERVING {
                next = healthpb.HealthCheckResponse_NOT_SERVING
            } else {
                next = healthpb.HealthCheckResponse_SERVING
            }

            time.Sleep(*sleep)
        }
    }()

    lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    _ = rpcServer.Serve(lis)
}

// 获取本机网卡IP
func getLocalIP() (ipv4 string, err error) {
    var (
        addrs   []net.Addr
        addr    net.Addr
        ipNet   *net.IPNet // IP地址
        isIpNet bool
    )
    // 获取所有网卡
    if addrs, err = net.InterfaceAddrs(); err != nil {
        return
    }
    // 取第一个非lo的网卡IP
    for _, addr = range addrs {
        // 这个网络地址是IP地址: ipv4, ipv6
        if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
            // 跳过IPV6
            if ipNet.IP.To4() != nil {
                ipv4 = ipNet.IP.String() // 192.168.1.1
                return
            }
        }
    }

    err = errors.New("ip address not found")
    return
}
