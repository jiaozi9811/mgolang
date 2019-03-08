# strconv

[TOC]

包实现了字符串与数字(整数、浮点数等)之间的互相转换

<https://www.cnblogs.com/golove/p/3262925.html>

## strconv包

### convert整形

func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (n uint64, err error)

func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string

// base表示进位制(2到36),如果base为0，则会从字符串前置判断,"0x"是16进制,"0"是8进制,否则是10进制
// bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64

```go
fmt.Println(strconv.ParseInt("FF", 16, 0))// 255
fmt.Println(strconv.ParseInt("0xFF", 16, 0))
// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
fmt.Println(strconv.ParseInt("0xFF", 0, 0))// 255
fmt.Println(strconv.ParseInt("9", 10, 4))
// 7 strconv.ParseInt: parsing "9": value out of range
```

### Atoi，Itoa

func Atoi(s string) (i int, err error)//Atoi是ParseInt(s, 10, 0)的简写
func Itoa(i int) string//Itoa是FormatInt(i, 10) 的简写

```go
fmt.Printf("%v\n",[]byte(strconv.Itoa(1)))  //[49]
fmt.Printf("%v",[]byte(string(1)))  //[1]
```

### convert布尔值

func FormatBool(b bool) string
func ParseBool(str string) (value bool, err error)
// 接受真值：1, t, T, TRUE, true, True
// 接受假值：0, f, F, FALSE, false, False

`fmt.Println(strconv.FormatBool(bool(false)))`
`fmt.Println(strconv.ParseBool("false"))`

### convert浮点数

func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func ParseFloat(s string, bitSize int) (f float64, err error)

// FormatFloat 将浮点数 f 转换为字符串形式
// f：要转换的浮点数
// fmt：格式标记（b、e、E、f、g、G）
// prec：精度（数字部分的长度，不包括指数部分）
// bitSize：指定浮点类型（32:float32、64:float64），结果会据此进行舍入。
//
// 格式标记：
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
//
// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
// 参考格式化输入输出中的旗标和精度说明

```go
s := "0.12345678901234567890"

f, err := strconv.ParseFloat(s, 32)
fmt.Println(f, err)                // 0.12345679104328156
fmt.Println(float32(f), err)       // 0.12345679

f, err = strconv.ParseFloat(s, 64)
fmt.Println(f, err)                // 0.12345678901234568
```

### IsPrint,IsGraphic,CanBackquote

func IsPrint(r rune) bool
// 判断一个字符是否是可打印的，和unicode.IsPrint一样。
// r必须是：字母（广义）、数字、标点、符号、ASCII空格

func IsGraphic(r rune) bool
// 判断 r 是否为 Unicode 定义的图形字符

```go
// 示例：获取不可打印字符和非图形字符
func main() {
    var rnp, rng, rpng, rgnp []rune
    const maxLen = 32
    for i := rune(0); i < utf8.MaxRune; i++ {
        if !strconv.IsPrint(i) { // 不可打印
            if len(rnp) < maxLen {
                rnp = append(rnp, i)
            }
            if strconv.IsGraphic(i) && len(rgnp) < maxLen { // 图形字符
                rgnp = append(rgnp, i)
            }
        }
        if !strconv.IsGraphic(i) { // 非图形字符
            if len(rng) < maxLen {
                rng = append(rng, i)
            }
            if strconv.IsPrint(i) && len(rpng) < maxLen { // 可打印
                rpng = append(rpng, i)
            }
        }
    }
    fmt.Printf("不可打印字符    ：%q\n", rnp)
    fmt.Printf("非图形字符      ：%q\n", rng)
    fmt.Printf("不可打印图形字符：%q\n", rgnp)
    fmt.Printf("可打印非图形字符：%q\n", rpng)
}

// 不可打印字符    ：['\x00' '\x01' '\x02' '\x03' '\x04' '\x05' '\x06' '\a' '\b' '\t' '\n' '\v' '\f' '\r' '\x0e' '\x0f' '\x10' '\x11' '\x12' '\x13' '\x14' '\x15' '\x16' '\x17' '\x18' '\x19' '\x1a' '\x1b' '\x1c' '\x1d' '\x1e' '\x1f']
// 非图形字符      ：['\x00' '\x01' '\x02' '\x03' '\x04' '\x05' '\x06' '\a' '\b' '\t' '\n' '\v' '\f' '\r' '\x0e' '\x0f' '\x10' '\x11' '\x12' '\x13' '\x14' '\x15' '\x16' '\x17' '\x18' '\x19' '\x1a' '\x1b' '\x1c' '\x1d' '\x1e' '\x1f']
// 不可打印图形字符：['\u00a0' '\u1680' '\u2000' '\u2001' '\u2002' '\u2003' '\u2004' '\u2005' '\u2006' '\u2007' '\u2008' '\u2009' '\u200a' '\u202f' '\u205f' '\u3000']
// 可打印非图形字符：[]
```

func CanBackquote(s string) bool
// 判断字符串是否可以不被修改的表示为一个单行的反引号字符串。
// 字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false

```go
    for i := rune(0); i < utf8.MaxRune; i++ {
        if !strconv.CanBackquote(string(i)) {
            fmt.Printf("%q, ", i)
        }
    }
// 结果如下：
// '\x00', '\x01', '\x02', '\x03', '\x04', '\x05', '\x06', '\a', '\b', '\n', '\v', '\f', '\r', '\x0e', '\x0f', '\x10', '\x11', '\x12', '\x13', '\x14', '\x15', '\x16', '\x17', '\x18', '\x19', '\x1a', '\x1b', '\x1c', '\x1d', '\x1e', '\x1f', '`', '\u007f', '\ufeff'
```

### Quote带双引号

func Quote(s string) string// s转换为双引号字符串
func QuoteToASCII(s string) string
func QuoteRune(r rune) string
func QuoteRuneToASCII(r rune) string
func Unquote(s string) (t string, err error)
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

```go
s := "Hello\tworld！\n"
fmt.Println(s)                         // Hello	world！（换行）
fmt.Println(strconv.Quote(s))          // "Hello\tworld！\n"
fmt.Println(strconv.QuoteToASCII(s))   // "Hello\t\u4e16\u754c\uff01\n"
fmt.Println(strconv.QuoteToGraphic(s)) // "Hello\tworld！\n"
```

### Append追加

// 将各种类型转换为字符串后追加到 dst 尾部
func AppendBool(dst []byte, b bool) []byte
func AppendInt(dst []byte, i int64, base int) []byte
func AppendUint(dst []byte, i uint64, base int) []byte
func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte

### AppendQuote

// 将各种类型转换为带引号字符串后追加到 dst 尾部
func AppendQuote(dst []byte, s string) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte
func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte