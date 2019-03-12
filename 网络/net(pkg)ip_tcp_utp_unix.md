# net conn

[TOC]

net提供了对以下协议的连接支持
"ip","tcp","tcp4","tcp6","udp","udp4","udp6","unix"

每一种通信协议都使用xxConn结构体来表示，诸如IPConn、TCPConn等，这些结构体都实现了Conn接口，Conn接口实现了基本的读、写、关闭、获取远程和本地地址、设置timeout等功能

## func Pipe() (Conn, Conn)

Pipe创建一个内存中的同步、全双工网络连接。连接的两端都实现了Conn接口。一端的读取对应另一端的写入，直接将数据在两端之间作拷贝；没有内部缓冲

## type Conn

```go
type Conn interface{
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
    SetReadDeadline(t time.Time) error // 设定该连接的读操作deadline
    SetWriteDeadline(t time.Time) error // 设定该连接的写操作deadline
}
```

### func Dial(network, address string) (Conn, error)

在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
"tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"

对TCP和UDP网络，地址格式是host:port或[host]:port

对IP网络，network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名，地址必须是IP地址字面值。

`Dial("tcp", "google.com:http")`

### func DialTimeout(network, address string, timeout time.Duration) (Conn, error)

### func FileConn(f *os.File) (c Conn, err error)

## type ListenConfig

func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)

## type Listener

```go
type Listener interface {
    Addr() Addr// Addr返回该接口的网络地址
    Accept() (c Conn, err error)// Accept等待并返回下一个连接到该接口的连接
    Close() error//Close关闭该接口,并使任何阻塞的Accept操作都会不再阻塞并返回错误
}
```

func FileListener(f *os.File) (ln Listener, err error)
func Listen(network, address string) (Listener, error)

[listen and accept and read header](./code/net_listen_accept.go)

## type PacketConn 代表通用的面向数据包的网络连接

func FilePacketConn(f *os.File) (c PacketConn, err error)
func ListenPacket(network, address string) (PacketConn, error)

## type IPConn

func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
func (c *IPConn) Close() error
func (c *IPConn) File() (f *os.File, err error)
func (c *IPConn) LocalAddr() Addr
func (c *IPConn) Read(b []byte) (int, error)
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
func (c *IPConn) RemoteAddr() Addr
func (c *IPConn) SetDeadline(t time.Time) error
func (c *IPConn) SetReadBuffer(bytes int) error
func (c *IPConn) SetReadDeadline(t time.Time) error
func (c *IPConn) SetWriteBuffer(bytes int) error
func (c *IPConn) SetWriteDeadline(t time.Time) error
func (c *IPConn) SyscallConn() (syscall.RawConn, error)
func (c *IPConn) Write(b []byte) (int, error)
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)

## type TCPAddr

查找几个空闲的端口
[find_empty_port](./code/_find_empty_port.go)

```go
for len(ports)<5{
    addr,_:=net.Listen("tcp","localhost:0")
    defer addr.Close()

    _,sport,_:=net.SplitHostPort(addr.Addr().String())
    //port:=addr.Addr().(*net.TCPAddr).Port
    port,_:=strconv.Atoi(sport)
    ports=append(ports,port)
}
```

func ResolveTCPAddr(network, address string) (*TCPAddr, error)
func (a *TCPAddr) Network() string
func (a *TCPAddr) String() string

## type TCPConn

func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
func (c *TCPConn) Close() error
func (c *TCPConn) CloseRead() error
func (c *TCPConn) CloseWrite() error
func (c *TCPConn) File() (f *os.File, err error)
func (c *TCPConn) LocalAddr() Addr
func (c *TCPConn) Read(b []byte) (int, error)
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
func (c *TCPConn) RemoteAddr() Addr
func (c *TCPConn) SetDeadline(t time.Time) error
func (c *TCPConn) SetKeepAlive(keepalive bool) error
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
func (c *TCPConn) SetLinger(sec int) error
func (c *TCPConn) SetNoDelay(noDelay bool) error
func (c *TCPConn) SetReadBuffer(bytes int) error
func (c *TCPConn) SetReadDeadline(t time.Time) error
func (c *TCPConn) SetWriteBuffer(bytes int) error
func (c *TCPConn) SetWriteDeadline(t time.Time) error
func (c *TCPConn) SyscallConn() (syscall.RawConn, error)
func (c *TCPConn) Write(b []byte) (int, error)

## type TCPListener

func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
func (l *TCPListener) Accept() (Conn, error)
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
func (l *TCPListener) Addr() Addr
func (l *TCPListener) Close() error
func (l *TCPListener) File() (f *os.File, err error)
func (l *TCPListener) SetDeadline(t time.Time) error
func (l *TCPListener) SyscallConn() (syscall.RawConn, error)

## type UDPAddr

func ResolveUDPAddr(network, address string) (*UDPAddr, error)
func (a *UDPAddr) Network() string
func (a *UDPAddr) String() string

## type UDPConn

func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
func (c *UDPConn) Close() error
func (c *UDPConn) File() (f *os.File, err error)
func (c *UDPConn) LocalAddr() Addr
func (c *UDPConn) Read(b []byte) (int, error)
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
func (c *UDPConn) RemoteAddr() Addr
func (c *UDPConn) SetDeadline(t time.Time) error
func (c *UDPConn) SetReadBuffer(bytes int) error
func (c *UDPConn) SetReadDeadline(t time.Time) error
func (c *UDPConn) SetWriteBuffer(bytes int) error
func (c *UDPConn) SetWriteDeadline(t time.Time) error
func (c *UDPConn) SyscallConn() (syscall.RawConn, error)
func (c *UDPConn) Write(b []byte) (int, error)
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)

## type UnixAddr

func ResolveUnixAddr(network, address string) (*UnixAddr, error)
func (a *UnixAddr) Network() string
func (a *UnixAddr) String() string

## type UnixConn

func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
func (c *UnixConn) Close() error
func (c *UnixConn) CloseRead() error
func (c *UnixConn) CloseWrite() error
func (c *UnixConn) File() (f *os.File, err error)
func (c *UnixConn) LocalAddr() Addr
func (c *UnixConn) Read(b []byte) (int, error)
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
func (c *UnixConn) RemoteAddr() Addr
func (c *UnixConn) SetDeadline(t time.Time) error
func (c *UnixConn) SetReadBuffer(bytes int) error
func (c *UnixConn) SetReadDeadline(t time.Time) error
func (c *UnixConn) SetWriteBuffer(bytes int) error
func (c *UnixConn) SetWriteDeadline(t time.Time) error
func (c *UnixConn) SyscallConn() (syscall.RawConn, error)
func (c *UnixConn) Write(b []byte) (int, error)
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)
type UnixListener
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
func (l *UnixListener) Accept() (Conn, error)
func (l *UnixListener) AcceptUnix() (*UnixConn, error)
func (l *UnixListener) Addr() Addr
func (l *UnixListener) Close() error
func (l *UnixListener) File() (f *os.File, err error)
func (l *UnixListener) SetDeadline(t time.Time) error
func (l *UnixListener) SetUnlinkOnClose(unlink bool)
func (l *UnixListener) SyscallConn() (syscall.RawConn, error)