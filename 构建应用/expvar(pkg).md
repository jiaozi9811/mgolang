# expvar

[TOC]

expvar包提供了公共变量的标准接口，如服务的操作计数器。本包通过HTTP在/debug/vars位置以JSON格式导出了这些变量。

对这些公共变量的读写操作都是原子级的

<http://blog.studygolang.com/2017/06/expvar-in-action/>

[triv.go](./example/triv.go)
//src/net/http/triv.go