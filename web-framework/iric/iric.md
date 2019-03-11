# iric

<https://www.studyiris.com>

## install

go get -u github.com/kataras/iris

## hello world

```go
package main
import (
    "github.com/kataras/iris"
    //"github.com/kataras/iris/context"//<1.9需要单独引入
)
func main() {
    app := iris.New()
    app.Get("/", func(ctx context.Context){
        ctx.HTML("<h1>Hello World!</h1>")
    })
    app.Run(iris.Addr(":8080"))
}
```