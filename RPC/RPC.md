# RPC

[TOC]

RPC(remote procedure call 远程过程调用)
    分布式系统中不同节点间的通信方式

go语言的rpc包在net/rpc下

RPC 框架大致有两种不同的侧重方向，一种偏重于服务治理，另一种偏重于跨语言调用
服务治理型的 RPC 框架有Alibab Dubbo、Motan 等
跨语言调用型的 RPC 框架有 Thrift、gRPC、Hessian、Finagle 等