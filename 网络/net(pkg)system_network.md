# net

[TOC]

## 系统网络信息

 [mac_ip.go](./code/mac_ip.go)

### type Interface [Top](#net)

```go
type Interface struct {
    Index        int          // 索引，>=1的整数
    MTU          int          // 最大传输单元
    Name         string       // 接口名，例如"en0"、"lo0"、"eth0.100"
    HardwareAddr HardwareAddr // 硬件地址，IEEE MAC-48、EUI-48或EUI-64格式
    Flags        Flags        // 接口的属性，例如FlagUp、FlagLoopback、FlagMulticast
}
```

#### func Interfaces() ([]Interface, error)

#### InterfaceByIndex InterfaceByName

func InterfaceByIndex(index int) (*Interface, error)
func InterfaceByName(name string) (*Interface, error)

#### Addrs MulticastAddrs(用ByIndex或ByName)

 _因func Interfaces()返回的是slice,此两个函数需用ByIndex或ByName调用_

func (ifi *Interface) Addrs() ([]Addr, error)//返回网络接口ifi的一或多个接口地址
func (ifi *Interface) MulticastAddrs() ([]Addr, error)//返回网络接口ifi加入的多播组地址

### type HardwareAddr [Top](#net)

func (a HardwareAddr) String() string

#### func ParseMAC(s string) (hw HardwareAddr, err error)

解析硬件地址

### type Addr [Top](#net)

#### func InterfaceAddrs() ([]Addr, error)返回系统网络接口的地址列表

以slice的形式返回系统所有ip地址

```go
ipaddr,_:=net.InterfaceAddrs()
fmt.Printf("%+v",ipaddr)
fmt.Printf("%+v\n",ipaddr[2].Network())//获取网络名称
fmt.Printf("%+v\n",ipaddr[2].String())//获取ip地址
```

## ip地址操作 [Top](#net)

### type IPAddr [Top](#net)

```go
type IPAddr struct {
    IP   IP
    Zone string // IPv6范围寻址域
}
```

#### Network()(网络类型) String()

func (a *IPAddr) Network() string//返回地址的网络类型："ip"
func (a *IPAddr) String() string//返回ip字符串

#### 解析地址(域名)的ip信息 func ResolveIPAddr(network, address string) (*IPAddr, error)

ResolveIPAddr将addr作为一个格式为"host"或"ipv6-host%zone"的IP地址来解析。 函数会在参数net指定的网络类型上解析，net必须是"ip"、"ip4"或"ip6"

```go
ipaddr,_:=net.ResolveIPAddr("ip+net","10.0.2.15")
fmt.Printf("%+v\n",ipaddr.Network())// ip
fmt.Printf("%+v\n",ipaddr.String())

ipaddr,_:=net.ResolveIPAddr("ip6","google.com")
fmt.Printf("%+v\n",ipaddr.Network())
fmt.Printf("%+v\n",ipaddr.String())
```

### func ParseCIDR(s string) (IP, *IPNet, error)

本函数会返回IP地址和该IP所在的网络

将s作为一个CIDR（无类型域间路由）的IP地址和掩码字符串，如"192.168.100.1/24"或"2001:DB8::/48"，解析并返回IP地址和IP网络

```go
addr,netinfo,_:=net.ParseCIDR(ipOut)
fmt.Printf("%+v\n",addr) // ip addr
fmt.Printf("%+v\n",netinfo) // net
```

### func JoinHostPort SplitHostPort

func JoinHostPort(host, port string) string

将host和port合并为一个网络地址。一般格式为"host:port"；如果host含有冒号或百分号，格式为"[host]:port"

func SplitHostPort(hostport string) (host, port string, err error)
Ipv6的文字地址或者主机名必须用方括号括起来

```go
netPort:=net.JoinHostPort("8.8.8.8:80/%","8080")
ipA,netA,_:=net.SplitHostPort(netPort)
fmt.Printf("%+v\n",netPort)
fmt.Printf("%+v\n",ipA)
fmt.Printf("%+v\n",netA)
```

### type IP [Top](#net)

#### func (ip IP) String() string

#### marshal

func (ip IP) MarshalText() ([]byte, error) // 实现了encoding.TextMarshaler接口，返回值和String方法一样
func (ip *IP) UnmarshalText(text []byte) error// 实现了encoding.TextUnmarshaler接口

#### IPv4 LookupIP ParseIP

func IPv4(a, b, c, d byte) IP
func LookupIP(host string) ([]IP, error)
func ParseIP(s string) IP

#### func (ip IP) Equal(x IP) bool

如果ip和x代表同一个IP地址，Equal会返回真。代表同一地址的IPv4地址和IPv6地址也被认为是相等的


func (ip IP) IsGlobalUnicast() bool // ip是全局单播地址，则返回真
func (ip IP) IsInterfaceLocalMulticast() bool // ip是接口本地组播地址，则返回真
func (ip IP) IsLinkLocalMulticast() bool  // ip是链路本地组播地址，则返回真
func (ip IP) IsLinkLocalUnicast() bool //ip是链路本地单播地址，则返回真
func (ip IP) IsMulticast() bool // ip是组播地址，则返回真
func (ip IP) IsUnspecified() bool  // ip是未指定地址，则返回真
func (ip IP) To16() IP // 将一个IP地址转换为16字节表示
func (ip IP) To4() IP // 将一个IPv4地址转换为4字节表示


#### func (ip IP) IsLoopback() bool 回环

ip是环回地址，则返回真

#### func (ip IP) DefaultMask() IPMask

返回IP地址ip的默认子网掩码。只有IPv4有默认子网掩码

#### func (ip IP) Mask(mask IPMask) IP 

认为mask为ip的子网掩码，返回ip的网络地址部分的ip

### type IPMask [Top](#net)

func CIDRMask(ones, bits int) IPMask
返回一个IPMask类型值，该返回值总共有bits个字位，其中前ones个字位都是1，其余字位都是0
func IPv4Mask(a, b, c, d byte) IPMask

返回一个4字节格式的IPv4掩码
func (m IPMask) Size() (ones, bits int)
func (m IPMask) String() string // 返回m的十六进制格式

### type IPNet [Top](#net)

func (n *IPNet) String() string // 返回n的CIDR表示

#### Contains Network

func (n *IPNet) Contains(ip IP) bool // 报告该网络是否包含地址ip
func (n *IPNet) Network() string // 返回网络类型名："ip+net"，注意该类型名是不合法的


### 解析域名,DNS和ip [Top](#net)

func LookupAddr(addr string) (names []string, err error)//查询某个地址，返回映射到该地址的主机名序列
func LookupCNAME(host string) (cname string, err error)//查询name的规范DNS名(但该域名未必可以访问)
func LookupHost(host string) (addrs []string, err error)//查询主机的网络地址序列
func LookupPort(network, service string) (port int, err error)//查询指定网络和服务的(默认)端口
func LookupTXT(name string) ([]string, error)//返回指定主机的DNS TXT记录
func LookupIP(host string) ([]IP, error)//查询主机的ipv4和ipv6地址序列
func LookupMX(name string) ([]*MX, error)//返回指定主机的按Pref字段排好序的DNS MX记录
func LookupNS(name string) ([]*NS, error)//返回指定主机的DNS NS记录
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)

```go
host,_:=net.LookupHost("www.google.com")
domain,_ := net.ResolveIPAddr("ip6","www.google.com")
addr,_:=net.LookupAddr("8.8.8.8")
cname,_:=net.LookupCNAME("www.google.com")
port,_:=net.LookupPort("tcp","https|telnet")//查询端口
ltxt,_:=net.LookupTXT("google.com")//查找DNS信息
/*
[69.171.234.29 2404:6800:4005:80e::2004]
69.171.234.29
[google-public-dns-a.google.com.]
www.google.com.
443
[docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8= v=spf1 include:_spf.google.com ~all facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95]
66.220.149.32
*/
```
