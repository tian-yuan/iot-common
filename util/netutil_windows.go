package util

import (
	"crypto/tls"
	"github.com/micro/go-micro/util/log"
	"net"
	"os"
	"reflect"
	"unsafe"
)

const SO_REUSEPORT = 15
const backlog = 512

// addr should be in format '0.0.0.0:1235'
func NewTCPListener(address string, reusePort bool) (net.Listener, error) {
	log.Info("window tcp server listen.")
	listen, err := net.Listen("tcp", address)
	return listen, err
}

const testAddr = "www.baidu.com:443"

func GetConnFd(c net.Conn) int64 {
	if t, ok := c.(*net.TCPConn); ok {
		return getTCPConnFd(t)
	}

	tlsv, _ := c.(*tls.Conn)
	cv := reflect.Indirect(reflect.ValueOf(tlsv)).FieldByName("conn").InterfaceData()
	cc := (*net.TCPConn)(unsafe.Pointer(cv[1]))
	return getTCPConnFd(cc)
}

var fdCheck int32 = 1

func getTCPConnFd(c *net.TCPConn) int64 {
	if fdCheck == 2 {
		fd := reflect.Indirect(reflect.ValueOf(c)).FieldByName("fd")
		return reflect.Indirect(fd).FieldByName("pfd").FieldByName("Sysfd").Int()
	}
	fd := reflect.ValueOf(c).FieldByName("fd")
	return reflect.Indirect(fd).FieldByName("sysfd").Int()
}

func DetectHost() string {
	if h, e := os.Hostname(); e == nil {
		addrs, e := net.LookupHost(h)
		if e == nil && len(addrs) > 0 {
			return addrs[0]
		}
	}

	if eth0, e := net.InterfaceByName("eth0"); e == nil {
		if addrs, e := eth0.Addrs(); e == nil && len(addrs) > 0 {
			return addrs[0].String()
		}
	}

	if eth1, e := net.InterfaceByName("eth1"); e == nil {
		if addrs, e := eth1.Addrs(); e == nil && len(addrs) > 0 {
			return addrs[0].String()
		}
	}

	if eth2, e := net.InterfaceByName("eth2"); e == nil {
		if addrs, e := eth2.Addrs(); e == nil && len(addrs) > 0 {
			return addrs[0].String()
		}
	}
	return ""
}
