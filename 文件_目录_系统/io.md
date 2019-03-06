# io

[TOC]

io包提供了对I/O原语的基本接口。

本包的基本任务是包装这些原语已有的实现(（)如os包里的原语),使之成为共享的公共接口,这些公共接口抽象出了泛用的函数并附加了一些相关的原语的操作

输入和输出操作是使用原语实现的，这些原语将数据模拟成可读的或可写的字节流。
为此，Go 的 io 包提供了 io.Reader 和 io.Writer 接口，分别用于数据的输入和输出，如图：
![reader_writer](https://segmentfault.com/img/bVbdzja?w=1600&h=214)

## 一切皆文件

unix下有一切皆文件的思想，golang把这个思想贯彻到更远，因为本质上我们对文件的抽象就是一个可读可写的一个对象，也就是实现了io.Writer和io.Reader的对象我们都可以称为文件

## 接口

type Reader interface {
type Writer interface {
type ReadCloser interface {
type WriteCloser interface {
type ReadWriter interface {
type ReadWriteCloser interface {

type ReaderAt interface {
type WriterAt interface {

type ReaderFrom interface {
type WriterTo interface {

type Closer interface {
type Seeker interface {
type ReadSeeker interface {
type WriteSeeker interface {
type ReadWriteSeeker interface {

type ByteReader interface {
type ByteWriter interface {
type ByteScanner interface {
type RuneReader interface {
type RuneScanner interface {

### EOF变量

>var EOF = errors.New("EOF")
EOF当无法得到更多输入时，Read方法返回EOF。当函数一切正常的到达输入的结束时，就应返回EOF

### type Reader interface

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Read方法读取len(p)字节数据写入p。它返回写入的字节数和遇到的任何错误。即使Read方法返回值n < len(p)，本方法在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，Read按惯例会返回可以读取到的数据，而不是等待更多数据

Reader接口会在输入流的结尾返回非0的字节数,返回值err == EOF或err==nil.下一次Read调用必然返回(0, EOF)

```go
byte:=make([]byte,1024)
openFile:=os.Open("filepath")
//openFile:=os.Create("filepath")
//openFile:=strings.NewReader("strings")
//openFile:=bytes.NewBuffer([]byte("strings"))
for {
    n,err:=openFile.Read(byte)
    //每次以len(byte)的字节数读取openFile中的数据,暂存入byte,读取到的字节数为n.因为n的大小可以小于len(byte)
}
```

![reader](https://segmentfault.com/img/bVbdzru?w=1600&h=354)

[实现一个Read.go](./code/io_read.go)

通过io.Reader实现
[Read_v2.go](./code/Read_v2.go)

### type Writer interface

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Write方法len(p) 字节数据从p写入底层的数据流

它会返回写入的字节数(0 <= n <= len(p))和遇到的任何导致写入提取结束的错误。Write必须返回非nil的错误，如果它返回的 n < len(p)。Write不能修改切片p中的数据，即使临时修改也不行

![io_write](https://segmentfault.com/img/bVbdzWd?w=1600&h=358)

```go
func main() {
    proverbs := []string{
        "Channels orchestrate mutexes serialize",
        "Cgo is not Go",
        "Errors are values",
        "Don't panic",
    }
    var writer bytes.Buffer

    for _, p := range proverbs {
        n, err := writer.Write([]byte(p))
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if n != len(p) {
            fmt.Println("failed to write data")
            os.Exit(1)
        }
    }

    fmt.Println(writer.String())
}
```

[实现一个Writer.go](./code/Writer.go)

### type ByteReader interface

ReadByte读取输入中的单个字节并返回。如果没有字节可读取，返回错误

```GO
type ByteReader interface {
    ReadByte() (c byte, err error)
}
```

### type ByteScanner

```GO
type ByteScanner interface {
    ByteReader
    UnreadByte() error
}
```

ByteScanner接口在基本的ReadByte方法之外还添加了UnreadByte方法。

UnreadByte方法让下一次调用ReadByte时返回之前调用ReadByte时返回的同一个字节

### type RuneReader interface

ReadRune读取单个utf-8编码的字符，返回该字符和它的字节长度

```GO
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}
```

### type RuneScanner interface

RuneScanner接口在基本的ReadRune方法之外还添加了UnreadRune方法。

UnreadRune方法让下一次调用ReadRune时返回之前调用ReadRune时返回的同一个utf-8字符

```GO
type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```

### type Seeker interface

Seeker接口用于包装基本的移位方法

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

Seek方法设定下一次读写的位置：偏移量为offset，校准点由whence确定：0表示相对于文件起始；1表示相对于当前位置；2表示相对于文件结尾。Seek方法返回新的位置以及可能遇到的错误

### type ReaderAt interface

ReadAt从底层输入流的偏移量off位置读取len(p)字节数据写入p

```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```

### type WriterAt interface

WriteAt将p全部len(p)字节数据写入底层数据流的偏移量off位置

```go
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

### type ReaderFrom interface

```GO
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

ReadFrom方法从r读取数据直到EOF或者遇到错误。返回值n是读取的字节数，执行时遇到的错误（EOF除外）也会被返回

### type WriterTo interface

```GO
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

WriteTo方法将数据写入w直到没有数据可以写入或者遇到错误。返回值n是写入的字节数，执行时遇到的任何错误也会被返回

## 函数

func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
func ReadFull(r Reader, buf []byte) (n int, err error)
func WriteString(w Writer, s string) (n int, err error)

func Copy(dst Writer, src Reader) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

func MultiReader(readers ...Reader) Reader
func MultiWriter(writers ...Writer) Writer

### ReadFull

### 1

### 2

### 4

### 5

### 6

### 7

### 8