package utils

import (
    "errors"
    "net"
    "testing"
)

// 获取系统ip地址
func TestGetIP(t *testing.T) {
    var (
        addrs   []net.Addr
        addr    net.Addr
        ipNet   *net.IPNet // IP地址
        isIpNet bool
        ipv4    string
        err     error
    )
    // 获取所有网卡
    if addrs, err = net.InterfaceAddrs(); err != nil {
        t.Error(err)
    }
    // 取第一个非lo的网卡IP
    for _, addr = range addrs {
        // 这个网络地址是IP地址: ipv4, ipv6
        if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
            // 跳过IPV6
            if ipNet.IP.To4() != nil {
                ipv4 = ipNet.IP.String() // 192.168.1.1
                t.Log(ipv4)
                return
            }
        }
    }
    err = errors.New("没有找到网卡IP")
    t.Error(err)
    return
}
