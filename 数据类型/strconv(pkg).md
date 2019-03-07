# strconv

[TOC]

包实现了字符串与数字(整数、浮点数等)之间的互相转换

## strconv.Atoi

```go
v := "10"
if s, err := strconv.Atoi(v); err == nil {
    fmt.Printf("%T, %v", s, s)
}
//int, 10
```