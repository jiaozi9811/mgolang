# unicode/utf8

>utf8包实现了对utf-8文本的常用函数和常数的支持
>包括rune和utf-8编码byte序列之间互相翻译的函数

[统计字数](./code/wordcount.go)

```go
// 编码所需的基本数字
const (
    RuneError = '\uFFFD'     // 错误的 Rune 或 Unicode 代理字符
    RuneSelf  = 0x80         // ASCII 字符范围
    MaxRune   = '\U0010FFFF' // Unicode 码点的最大值
    UTFMax    = 4            // 一个字符编码的最大长度
)
```

## 统计字符个数

func RuneCount(p []byte) int
func RuneCountInString(s string) int

## 类型转换

func EncodeRune(p []byte, r rune) int   //将r转换为UTF-8编码
func DecodeRune(p []byte) (r rune, size int)    //解码p中的第一个字符，返回解码后的字符和p中被解码的字节数
func DecodeRuneInString(s string) (r rune, size int)
func DecodeLastRune(p []byte) (r rune, size int)    //解码p中的最后一个字符，返回解码后的字符，和p中被解码的字节数
func DecodeLastRuneInString(s string) (r rune, size int)

## 检查存在性

func FullRune(p []byte) bool    //检测p中第一个字符的UTF-8编码是否完整
func FullRuneInString(s string) bool

func Valid(p []byte) bool   //判断p是否为完整有效的UTF8编码序列
func ValidString(s string) bool
func ValidRune(r rune) bool //判断r能否被正确的转换为UTF8编码

func RuneLen(r rune) int        //需要多少字节来编码字符r
func RuneStart(b byte) bool     //判断b是否为UTF8字符的首字节编码

## code

```go
//https://www.cnblogs.com/golove/p/3271597.html
func main() {
    b := make([]byte, utf8.UTFMax)

    n := utf8.EncodeRune(b, '好')
    fmt.Printf("%v：%v\n", b, n) // [229 165 189 0]：3

    r, n := utf8.DecodeRune(b)
    fmt.Printf("%c：%v\n", r, n) // 好：3

    s := "大家好"
    for i := 0; i < len(s); {
        r, n = utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%c：%v   ", r, n) // 大：3   家：3   好：3
        i += n
    }
    fmt.Println()

    for i := len(s); i > 0; {
        r, n = utf8.DecodeLastRuneInString(s[:i])
        fmt.Printf("%c：%v   ", r, n) // 好：3   家：3   大：3
        i -= n
    }
    fmt.Println()

    b = []byte("好")
    fmt.Printf("%t, ", utf8.FullRune(b))     // true
    fmt.Printf("%t, ", utf8.FullRune(b[1:])) // true
    fmt.Printf("%t, ", utf8.FullRune(b[2:])) // true
    fmt.Printf("%t, ", utf8.FullRune(b[:2])) // false
    fmt.Printf("%t\n", utf8.FullRune(b[:1])) // false

    b = []byte("大家好")
    fmt.Println(utf8.RuneCount(b)) // 3

    fmt.Printf("%d, ", utf8.RuneLen('A'))          // 1
    fmt.Printf("%d, ", utf8.RuneLen('\u03A6'))     // 2
    fmt.Printf("%d, ", utf8.RuneLen('好'))          // 3
    fmt.Printf("%d, ", utf8.RuneLen('\U0010FFFF')) // 4
    fmt.Printf("%d\n", utf8.RuneLen(0x1FFFFFFF))   // -1

    fmt.Printf("%t, ", utf8.RuneStart("好"[0])) // true
    fmt.Printf("%t, ", utf8.RuneStart("好"[1])) // false
    fmt.Printf("%t\n", utf8.RuneStart("好"[2])) // false

    b = []byte("你好")
    fmt.Printf("%t, ", utf8.Valid(b))     // true
    fmt.Printf("%t, ", utf8.Valid(b[1:])) // false
    fmt.Printf("%t, ", utf8.Valid(b[2:])) // false
    fmt.Printf("%t, ", utf8.Valid(b[:2])) // false
    fmt.Printf("%t, ", utf8.Valid(b[:1])) // false
    fmt.Printf("%t\n", utf8.Valid(b[3:])) // true

    fmt.Printf("%t, ", utf8.ValidRune('好'))        // true
    fmt.Printf("%t, ", utf8.ValidRune(0))          // true
    fmt.Printf("%t, ", utf8.ValidRune(0xD800))     // false  代理区字符
    fmt.Printf("%t\n", utf8.ValidRune(0x10FFFFFF)) // false  超出范围
}
```