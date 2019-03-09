# flag

[TOC]

flag 提供用来解析命令行参数的包,使得开发命令行工具更为简单
    .Usage
    .Type(参数名,默认值,使用提示)
    .Type(指针,参数名,默认值,使用提示)
    .Args 返回命令行的参数
    .Parse 执行解析
    
<http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/>

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


flag
定义参数
1）flag.Xxx()，其中 Xxx 可以是 Int、String 等；返回一个相应类型的指针，如：
var ip = flag.Int("flagname", 1234, "help message for flagname")

2）flag.XxxVar()，将 flag 绑定到一个变量上，如：
var flagvar int
flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")

自定义参数



通过调用 flag.Parse() 进行解析
从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。
命令行 flag 的语法有如下三种形式：
    -flag // 只支持bool类型
    -flag=x
    -flag x // 只支持非bool类型