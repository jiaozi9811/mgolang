# flag

[TOC]

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