# flag

[TOC]

[flag](./code/flag.go)
[flag demo](./code/flag_demo.go)

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

func Arg(i int) string
func Args() []string
func Visit(fn func(*Flag))
func VisitAll(fn func(*Flag))
func NArg() int
func NFlag() int
func Parse()
func Parsed() bool
func PrintDefaults()
func Set(name, value string) error

func Bool(name string, value bool, usage string) *bool
func BoolVar(p *bool, name string, value bool, usage string)
func Duration(name string, value time.Duration, usage string) *time.Duration
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
func Float64(name string, value float64, usage string) *float64
func Float64Var(p *float64, name string, value float64, usage string)
func Int(name string, value int, usage string) *int
func Int64(name string, value int64, usage string) *int64
func Int64Var(p *int64, name string, value int64, usage string)
func IntVar(p *int, name string, value int, usage string)
func String(name string, value string, usage string) *string
func StringVar(p *string, name string, value string, usage string)
func Uint(name string, value uint, usage string) *uint
func Uint64(name string, value uint64, usage string) *uint64
func Uint64Var(p *uint64, name string, value uint64, usage string)
func UintVar(p *uint, name string, value uint, usage string)
func UnquoteUsage(flag *Flag) (name string, usage string)
func Var(value Value, name string, usage string)

type Flag
func Lookup(name string) *Flag

type FlagSet
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
func (f *FlagSet) Arg(i int) string
func (f *FlagSet) Args() []string
func (f *FlagSet) Lookup(name string) *Flag
func (f *FlagSet) NArg() int
func (f *FlagSet) NFlag() int
func (f *FlagSet) Name() string
func (f *FlagSet) Output() io.Writer
func (f *FlagSet) Parse(arguments []string) error
func (f *FlagSet) Parsed() bool
func (f *FlagSet) PrintDefaults()
func (f *FlagSet) Set(name, value string) error
func (f *FlagSet) SetOutput(output io.Writer)
func (f *FlagSet) Visit(fn func(*Flag))
func (f *FlagSet) VisitAll(fn func(*Flag))

func (f *FlagSet) Bool(name string, value bool, usage string) *bool
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
func (f *FlagSet) ErrorHandling() ErrorHandling
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
func (f *FlagSet) Int(name string, value int, usage string) *int
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
func (f *FlagSet) String(name string, value string, usage string) *string
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
func (f *FlagSet) Var(value Value, name string, usage string)
