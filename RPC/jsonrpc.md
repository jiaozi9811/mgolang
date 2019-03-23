# /net/prc/jsonrpc

[TOC]

[jsonrpc-server.go](./code/jsonrpc_server.go)
[jsonrpc-client.go](./code/jsonrpc_client.go)

func Dial(network, address string) (*rpc.Client, error)
func NewClient(conn io.ReadWriteCloser) *rpc.Client
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec
func ServeConn(conn io.ReadWriteCloser)