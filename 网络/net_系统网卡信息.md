# 获取网络连接信息

net包对于网络I/O提供了便携式接口，包括TCP/IP,UDP，域名解析以及Unix Socket。
尽管net包提供了大量访问底层的接口，但是大多数情况下，客户端仅仅只需要最基本的接口，例如Dial，LIsten，Accepte以及分配的conn连接和listener接口。 
crypto/tls包使用相同的接口以及类似的Dial和Listen函数


提供了对以下协议的连接支持
"ip","tcp","tcp4","tcp6","udp","udp4","udp6","unix"
每一种通信协议都使用xxConn结构体来表示，诸如IPConn、TCPConn等，这些结构体都实现了Conn接口，Conn接口实现了基本的读、写、关闭、获取远程和本地地址、设置timeout等功能

## 连接信息net.Conn(interface)

type Conn interface{
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}

```go
//很多时候，服务器会处理多个端口的监听！可以使用select轮询处理这种情况
package main

import (
    "net"
    "fmt"
    "os"
)

func main() {
	lsr, err := net.Listen("tcp", ":7070")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	for {
		conn , err := lsr.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			continue
		}
		go connHandler(conn)
	}
	fmt.Println("Done !")
}

func connHandler(conn net.Conn) {
	defer conn.Close()

	var buf[512]byte
	for {
		n , err := conn.Read(buf[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
	}
}
```

## 获取本地硬件信息

```go
//net.Interfaces()//获取硬件的网卡信息
//net.InterfaceAddrs()//获取硬件的ip信息
ips,_:=net.Interfaces()//获取硬件的网卡信息
//fmt.Printf("%#v",ips)
for k,v:=range ips{
  fmt.Println(k,v)
}
netIp,_:=net.InterfaceAddrs()//获取硬件的ip信息
fmt.Println(netIp)
```

## 对域名和公网ip的解析

>func LookupAddr(addr string) (names []string, err error)
>func LookupCNAME(host string) (cname string, err error)
>func LookupHost(host string) (addrs []string, err error)
>func LookupPort(network, service string) (port int, err error)
>func LookupTXT(name string) ([]string, error)
>net.JoinHostPort("127.0.0.1", "8080")



## 解析ip地址

func ResolveIPAddr(network, address string) (*IPAddr, error)
func (a *IPAddr) Network() string
func (a *IPAddr) String() string
func ResolveTCPAddr(network, address string) (*TCPAddr, error)
func ResolveUDPAddr(network, address string) (*UDPAddr, error)
func ResolveUnixAddr(network, address string) (*UnixAddr, error)

## 解析掩码(mask)

func CIDRMask(ones, bits int) IPMask
func IPv4Mask(a, b, c, d byte) IPMask
func (m IPMask) Size() (ones, bits int)

## 获取ip信息

```go
name := "192.168.1.97"
ip := net.ParseIP(name)
fmt.Printf("IP: %s\n",ip.String())

defaultMask := ip.DefaultMask()
fmt.Printf("DefaultMask: %s\n",defaultMask.String())

ones, bits := defaultMask.Size()
fmt.Printf("ones: %d bits: %d\n", ones, bits)

network:=ip.Mask(defaultMask)
fmt.Println(network)
```