# RPC

[TOC]

rpcx
<https://legacy.gitbook.com/book/smallnest/go-rpc-programming-guide/details>

RPC(remote procedure call 远程过程调用)
    该协议允许运行于一台计算机的程序调用另一台计算机的子程序，而程序员无需额外地为这个交互作用编程。如果涉及的软件采用面向对象编程，那么远程过程调用亦可称作远程调用或远程方法调用

    net/rpc库使用encoding/gob进行编解码，支持tcp或http数据传输方式

    net/rpc/jsonrpc库实现RPC方法，JSON RPC采用JSON进行数据编解码，因而支持跨语言调用。但目前的jsonrpc库是基于tcp协议实现的，暂时不支持使用http进行数据传输

go语言的rpc包在net/rpc下

RPC 框架大致有两种不同的侧重方向，一种偏重于服务治理，另一种偏重于跨语言调用
服务治理型的 RPC 框架有Alibab Dubbo、Motan 等
跨语言调用型的 RPC 框架有 Thrift、gRPC、Hessian、Finagle 等

## simple example

[server.go](./code/simple_server.go)
[client.go](./code/simple_client.go)

func Accept(lis net.Listener)
func HandleHTTP()
func Register(rcvr interface{}) error
func RegisterName(name string, rcvr interface{}) error
func ServeCodec(codec ServerCodec)
func ServeConn(conn io.ReadWriteCloser)
func ServeRequest(codec ServerCodec) error

## type Call

type Call struct {
    ServiceMethod string      // 调用的服务和方法的名称
    Args          interface{} // 函数的参数（下层为结构体指针）
    Reply         interface{} // 函数的回复（下层为结构体指针）
    Error         error       // 在调用结束后，保管错误的状态
    Done          chan *Call  // 对其的接收操作会阻塞，直到远程调用结束
}

## type ClientCodec

type ClientCodec interface {
    // 本方法必须能安全的被多个go程同时使用
    WriteRequest(*Request, interface{}) error
    ReadResponseHeader(*Response) error
    ReadResponseBody(interface{}) error
    Close() error
}

## type ServerCodec

type ServerCodec interface {
    ReadRequestHeader(*Request) error
    ReadRequestBody(interface{}) error
    // 本方法必须能安全的被多个go程同时使用
    WriteResponse(*Response, interface{}) error
    Close() error
}

## type Request

type Request struct {
    ServiceMethod string // 格式："Service.Method"
    Seq           uint64 // 由客户端选择的序列号
    // 内含隐藏或非导出字段
}

## type Response

type Response struct {
    ServiceMethod string // 对应请求的同一字段
    Seq           uint64 // 对应请求的同一字段
    Error         string // 可能的错误
    // 内含隐藏或非导出字段
}

## type Client

func Dial(network, address string) (*Client, error)
func DialHTTP(network, address string) (*Client, error)
func DialHTTPPath(network, address, path string) (*Client, error)
func NewClient(conn io.ReadWriteCloser) *Client
func NewClientWithCodec(codec ClientCodec) *Client
func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
func (client *Client) Close() error
func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call

## type Server

func NewServer() *Server
func (server *Server) Accept(lis net.Listener)
func (server *Server) HandleHTTP(rpcPath, debugPath string)
func (server *Server) Register(rcvr interface{}) error
func (server *Server) RegisterName(name string, rcvr interface{}) error
func (server *Server) ServeCodec(codec ServerCodec)
func (server *Server) ServeConn(conn io.ReadWriteCloser)
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)
func (server *Server) ServeRequest(codec ServerCodec) error