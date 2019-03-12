# log

[TOC]

日志可以简单地分为

- 事件日志记录了发生在系统运行过程中的事件，可以用来审计操作、诊断问题等，对于理解复杂系统的运行非常关键
- 消息日志则被应用到如即时通信等软件里，用来记录来往的消息

## log 常量

在log包里首先定义了一些常量，它们是日志输出前缀的标识:

```go
 const (
    Ldate = 1 << iota  //当前时区的日期，格式是:2009/01/23
    Ltime          //当前时区的时间，格式是：01:23:23
    Lmicroseconds  //微秒解析
    Llongfile      //全文件名和行号
    Lshortfile     //当前文件名和行号, 会覆盖Llongfile
    LUTC           //使用UTC而非本地时区
    LstdFlags = Ldate | Ltime //标准Logger的默认值
 )
 ```

## 定制日志的抬头信息

[simple-logger](./code/simple_loggle.go)

```go
func init(){ 
    log.SetFlags(log.Ldate|log.Lshortfile) 
    //log.SetFlags(log.Ldate|log.Ltime |log.LUTC)

    // log.SetPrefix("【UserCenter】")
    // log.SetFlags(log.LstdFlags | log.Lshortfile |log.LUTC)
}
```

func Flags() int
func SetFlags(flag int)
func Prefix() string
func SetPrefix(prefix string)
func SetOutput(w io.Writer)
func Printf(format string, v ...interface{})
func Print(v ...interface{})
func Println(v ...interface{})
func Fatalf(format string, v ...interface{})
func Fatal(v ...interface{})
func Fatalln(v ...interface{})
func Panicf(format string, v ...interface{})
func Panic(v ...interface{})
func Panicln(v ...interface{})

## type Logger

func New(out io.Writer, prefix string, flag int) *Logger
func (l *Logger) Flags() int
func (l *Logger) SetFlags(flag int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)
func (l *Logger) Output(calldepth int, s string) error
func (l *Logger) Printf(format string, v ...interface{})
func (l *Logger) Print(v ...interface{})
func (l *Logger) Println(v ...interface{})
func (l *Logger) Fatalf(format string, v ...interface{})
func (l *Logger) Fatal(v ...interface{})
func (l *Logger) Fatalln(v ...interface{})
func (l *Logger) Panic(v ...interface{})
func (l *Logger) Panicf(format string, v ...interface{})
func (l *Logger) Panicln(v ...interface{})

## log级别

编号|等级||
-|-|-
0|Emergency|system is unusable
1|Alert|action must be taken immediately
2|Critical:|critical conditions
3|Error|error conditions
4|Warning|warning conditions
5|Notice|normal but significant condition
6|Informational|informational messages
7|Debug|debug-level messages