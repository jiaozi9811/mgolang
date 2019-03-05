# flag

[TOC]

flag 提供用来解析命令行参数的包,使得开发命令行工具更为简单
    .Usage
    .Type(参数名,默认值,使用提示)
    .Type(指针,参数名,默认值,使用提示)
    .Args 返回命令行的参数
    .Parse 执行解析

## 判断参数是否存在

```go
flag.Parse()
args := flag.Args()
var n, m int
var err error
if args == nil || len(args) < 1 {
    n, m = 300, 10000
} else {
    n, err = strconv.Atoi(args[0])
    if err != nil {
        n = 300
    }
    if len(args) < 2 {
        m = 10000
    } else {
        m, err = strconv.Atoi(args[1])
        if err != nil {
            m = 10000
        }
    }
}
fmt.Println(n, m)
```