# errors

[TOC]

## error类型

error类型来返回函数执行过程中遇到的错误，如果返回的 error 值为 nil，则表示未遇到错误，否则 error 会返回一个字符串，用于说明遇到了什么错误

其实error是一个接口,定义如下：

```go
type error interface {
    Error() string
}
```

## New

func New(text string) error

```go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
    fmt.Print(err)
}
```

## fmt.Errorf

```go
const name, id = "bimmler", 17
err := fmt.Errorf("user %q (id %d) not found", name, id)
if err != nil {
    fmt.Print(err)
}
```