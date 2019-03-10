# strings

[TOC]

## Compare

func Compare(a, b string) int

>Compare比较字符串的速度比字符串内建的比较要快

## Equal

func EqualFold(s, t string) bool //判断两个utf-8编码字符串(将unicode大写、小写、标题三种格式字符视为相同)是否相同

fmt.Println(strings.EqualFold("Go", "go")) //true

## 大小写转换

func Title(s string) string//返回s中每个单词的首字母大写
func ToLower(s string) string//返回将所有字母都转为小写
func ToUpper(s string) string//返回将所有字母都转为大写

## Join操作

func Join(a []string, sep string) string//将一系列字符串连接为字符串,用sep分隔


## 重复打印 Repeat

func Repeat(s string, count int) string

fmt.Println("ba" + strings.Repeat("na", 2))//banana

## 替换 Replace

func Replace(s, old, new string, n int) string //用new替换old,一共替换n个,如果n< 0,则全部替换
func ReplaceAll(s, old, new string) string

## 去除字符串 Trim

func Trim(s string, cutset string) string//将s前后缀所有cutset的值去掉

func TrimSpace(s string) string//将s前后端所有空白(unicode.IsSpace指定)去掉
func TrimLeft(s string, cutset string) string
func TrimRight(s string, cutset string) string
func TrimPrefix(s, prefix string) string//去除s可能的前缀prefix
func TrimSuffix(s, suffix string) string//去除s可能的后缀suffix

## 是否包含子串 Contains

func Contains(s,substr string) bool       // 判断字符串s是否包含子串substr
func ContainsAny(s,char string) bool      // 判断字符串s是否包含字符串chars中的任一字符
func ContainsRune(s string,r rune) bool   // 判断字符串s是否包含utf-8码值r

```go
fmt.Println(strings.Contains("seafood", "foo")) //true
fmt.Println(strings.Contains("seafood", ""))    //true

fmt.Println(strings.ContainsAny("failure", "u & i"))//true

```

## 计数 Count

func Count(s, sep string) int     //子字符串sep出现次数
使用 Rabin-Karp 算法实现

```go
fmt.Println(strings.Count("cheese", "e")) //3
fmt.Println(strings.Count("five", ""))      //5
```

## 分割字符串 Fields Split

func Fields(s string) []string    //返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串
func FieldsFunc(s string, f func(rune) bool) []string

func Split(s, sep string) []string //去掉s中出现的sep的方式进行分割,返回生成的所有片段组成的切片
func SplitN(s, sep string, n int) []string
func SplitAfter(s, sep string) []string//从s中出现的sep后面切断的方式进行分割
func SplitAfterN(s, sep string, n int) []string

它们都调用了genSplit函数,如果sep为空,相当于分成一个个UTF-8字符
>Split(s, sep) 和 SplitN(s, sep, -1) 等价  
>SplitAfter(s, sep) 和 SplitAfterN(s, sep, -1) 等价

带N的方法可以通过最后一个参数n控制返回的结果中的slice中的元素个数
>当 n < 0 时，返回所有的子字符串
>当 n == 0 时，返回的结果是 nil
>当 n > 0 时，表示返回的slice中最多只有 n 个元素，其中，最后一个元素不会分割

```go
fmt.Println(strings.Fields("  foo bar  baz   "))//["foo" "bar" "baz"]

fmt.Println(strings.Split(" xyz ", ""))//[" " "x" "y" "z" " "]

fmt.Println(strings.SplitN("a,b,c", ",", 2))//["a" "b,c"]

fmt.Println(strings.SplitAfter("a,b,c", ","))//["a," "b," "c"]

fmt.Println(strings.SplitAfterN("a,b,c", ",", 2))//["a," "b,c"]
```


func HasPrefix(s, prefix string) bool     // s 中是否以 prefix 开始
func HasSuffix(s, suffix string) bool {   // s 中是否以 suffix 结尾
    return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

## 字符串索引index

func Index(s, sep string) int //在s中查找sep的第一次出现的位置,返回出现的索引值
func LastIndex(s, sep string) int //最后一次出现的位置

func IndexAny(s, chars string) int //chars中任一utf-8码值在s中第一次出现的位置
func LastIndexAny(s, chars string) int

func IndexFunc(s string, f func(rune) bool) int //查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
func LastIndexFunc(s string, f func(rune) bool) int

func IndexByte(s string, c byte) int//字符c在s中第一次出现的位置，不存在则返回-1

func IndexRune(s string, r rune) int //unicode码值r在s中第一次出现的位置

```go
fmt.Println(strings.Index("go gopher", "go"))//0
fmt.Println(strings.LastIndex("go gopher", "go"))//3

fmt.Println(strings.IndexAny("chicken", "aeiouy"))//2
```

## 前后缀判断

func HasPrefix(s, prefix string) bool//判断s是否有前缀字符串prefix
func HasSuffix(s, suffix string) bool//判断s是否有后缀字符串suffix

## type Reader

func NewReader(s string) *Reader
func (r *Reader) Len() int
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

### NewReader

### Read