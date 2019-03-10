# bytes

bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似

该包定义了一些操作byte slice的便利操作。因为字符串可以表示为 []byte，因此，bytes 包定义的函数、方法等和 strings 包很类似

## type Reader

### 创建初始化byte(func NewReader)

func NewReader(b []byte) *Reader

// bytes.Reader 实现了如下接口：
// io.ReadSeeker
// io.ReaderAt
// io.WriterTo
// io.ByteScanner
// io.RuneScanner

### 读入byte(Read)

func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) ReadByte() (byte, error)
func (r *Reader) ReadRune() (ch rune, size int, err error)

### Len,Reset,Seek,Size

func (r *Reader) Len() int// 返回未读取部分的数据长度
func (r *Reader) Reset(b []byte)
func (r *Reader) Seek(offset int64, whence int) (int64, error)
func (r *Reader) Size() int64

func (r *Reader) UnreadByte() error
func (r *Reader) UnreadRune() error
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

## type Buffer

### 创建初始化NewBuffer，NewBufferString

func NewBuffer(buf []byte) *Buffer
func NewBufferString(s string) *

### 读取

func (b *Buffer) Read(p []byte) (n int, err error)
func (b *Buffer) ReadByte() (byte, error)
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
func (b *Buffer) ReadRune() (r rune, size int, err error)
func (b *Buffer) ReadString(delim byte) (line string, err error)

### Len,Reset

func (b *Buffer) Len() int
func (b *Buffer) Reset()

### 返回为读取部分

func (b *Buffer) Bytes() []byte
func (b *Buffer) String() string

### buffer容量及扩容Cap,Grow

func (b *Buffer) Cap() int
func (b *Buffer) Grow(n int)

### Next

func (b *Buffer) Next(n int) []byte

### 丢弃缓冲中的byte

func (b *Buffer) Truncate(n int)//丢弃缓冲中除前n字节数据外的其它数据,如果n小于零或者大于缓冲容量将panic

### Unread

func (b *Buffer) UnreadByte() error
func (b *Buffer) UnreadRune() error

### Write

func (b *Buffer) Write(p []byte) (n int, err error)
func (b *Buffer) WriteByte(c byte) error
func (b *Buffer) WriteRune(r rune) (n int, err error)
func (b *Buffer) WriteString(s string) (n int, err error)
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

## bytes函数

### 转换为rune

func Runes(s []byte) []rune//转换为 []rune 类型返回

### 比较

func Compare(a, b []byte) int
func Equal(a, b []byte) bool
func EqualFold(s, t []byte) bool

### 转换为Map类型

func Map(mapping func(r rune) rune, s []byte) []byte//将 s 中的字符替换为 mapping(r) 的返回值

### Count统计

func Count(s, sep []byte) int

### 包含Contains

func Contains(b, subslice []byte) bool

### 前后缀

func HasPrefix(s, prefix []byte) bool
func HasSuffix(s, suffix []byte) bool

### 索引Index

func Index(s, sep []byte) int
func IndexByte(s []byte, c byte) int
func IndexRune(s []byte, r rune) int
func IndexAny(s []byte, chars string) int
func IndexFunc(s []byte, f func(r rune) bool) int
func LastIndex(s, sep []byte) int
func LastIndexAny(s []byte, chars string) int
func LastIndexFunc(s []byte, f func(r rune) bool) int

### 大小写Title

func Title(s []byte) []byte
func ToLower(s []byte) []byte
func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte
func ToUpper(s []byte) []byte
func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte
func ToTitle(s []byte) []byte
func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte

### 重复和替换 Repeat,Replace

func Repeat(b []byte, count int) []byte
func Replace(s, old, new []byte, n int) []byte

### 删除指定字节Trim

func Trim(s []byte, cutset string) []byte
func TrimSpace(s []byte) []byte
func TrimFunc(s []byte, f func(r rune) bool) []byte
func TrimLeft(s []byte, cutset string) []byte
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
func TrimPrefix(s, prefix []byte) []byte
func TrimRight(s []byte, cutset string) []byte
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
func TrimSuffix(s, suffix []byte) []byte

### 拆分Split,Fields

func Fields(s []byte) [][]byte
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
func Split(s, sep []byte) [][]byte
func SplitN(s, sep []byte, n int) [][]byte
func SplitAfter(s, sep []byte) [][]byte
func SplitAfterN(s, sep []byte, n int) [][]byte

### 组合 Join

func Join(s [][]byte, sep []byte) []byte