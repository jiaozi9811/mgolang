# context

[TOC]

[context](./code/context.go)

## type CancelFunc func()

## type Context interface

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{} // 当该context被取消的时候，该channel会被关闭
    Err() error
    Value(key interface{}) interface{}
}
```

- Done方法返回一个channel，这个channel对于以Context方式运行的函数而言，是一个取消信号。当这个channel关闭时，上面提到的这些函数应该终止手头的工作并立即返回。 之后，Err方法会返回一个错误，告知为什么Context被取消
- Deadline方法允许函数确定它们是否应该开始工作。如果剩下的时间太少，也许这些函数就不值得启动。代码中，我们也可以使用Deadline对象为 I/O 操作设置截止时间
- Value方法允许Context对象携带request作用域的数据

### func Background() Context

主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context，它不能被取消

```go
func Background() Context {
    return background
}

var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
)
```

### func TODO() Context

如果我们不知道该使用什么Context的时候，可以使用这个

```go
func TODO() Context {
    return todo
}
```

## func WithValue(parent Context, key, val interface{}) Context

它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到

## func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context

## func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消

## func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

这个表示是超时自动取消，

## 使用原则和技巧

- 不要把Context放在结构体中，要以参数的方式传递，parent Context一般为Background
- 应该要把Context作为第一个参数传递给入口请求和出口请求链路上的每一个函数，放在第一位，变量名建议都统一，如ctx。
- 给一个函数方法传递Context的时候，不要传递nil，否则在tarce追踪的时候，就会断了连-
- Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
- Context是线程安全的，可以放心的在多个goroutine中传递
- 可以把一个Context对象传递给任意个数的 gorotuine，对它执行 取消 操作时，所有 _ goroutine 都会接收到取消信号

## 常用方法

### 调用Context Done方法取消

```go
func Stream(ctx context.Context, out chan<- Value) error {
    for {
        v, err := DoSomething(ctx)

        if err != nil {
            return err
        }
        select {
        case <-ctx.Done():
            return ctx.Err()
        case out <- v:
        }
    }
}
```

### 通过 context.WithValue 来传值

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    valueCtx := context.WithValue(ctx, key, "add value")

    go watch(valueCtx)
    time.Sleep(10 * time.Second)
    cancel()
    time.Sleep(5 * time.Second)
}
func watch(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            //get value
            fmt.Println(ctx.Value(key), "is cancel")
            return
        default:
            //get value
            fmt.Println(ctx.Value(key), "int goroutine")
            time.Sleep(2 * time.Second)
        }
    }
}
```

### 超时取消 context.WithTimeout

```go
package main

import (
    "fmt"
    "sync"
    "time"
    "context"
)

var (
    wg sync.WaitGroup
)

func work(ctx context.Context) error {
    defer wg.Done()

    for i := 0; i < 10; i++ {
        select {
        case <- time.After(time.Second):
            fmt.Println("Doing some work ", i)

        // we received the signal of cancelation in this channel
        case <-ctx.Done():
            fmt.Println("Cancel the context ", i)
            return ctx.Err()
        }
    }
    return nil
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
    defer cancel()

    fmt.Println("Hey, I'm going to do some work")

    wg.Add(1)
    go work(ctx)
    wg.Wait()

    fmt.Println("Finished. I'm going home")
}
/*
Hey, I'm going to do some work
Doing some work  0
Doing some work  1
Doing some work  2
Cancel the context  3
Finished. I'm going home
*/
```

### 截止时间 取消 context.WithDeadline

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(4 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```