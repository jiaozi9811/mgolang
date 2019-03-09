# bufio

[TOC]

bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象

bytes.NewBuffer()

[bufio.go](./code/bufio.go)

https://www.cnblogs.com/golove/p/3282667.html
http://www.okyes.me/2016/05/30/go-bufio.html

## 数据类型及方法

func (b *Reader) Peek(n int) ([]byte, error) //返回缓冲区前n字节, 不移动读取指针
func (b *Reader) Buffered() int //缓冲区中缓冲的还没有读取的数据
func (b *Writer) Flush() error //刷新数据, 将缓冲区数据写入io writer
func (b *Writer) Available() int //写缓冲区可用的空间, 默认最大空间是4096
func (b *Reader) Reset(r io.Reader)//丢弃缓冲中的数据，清除任何错误

### type Reader

bufio.Reader 实现了如下接口：
- io.Reader
- io.WriterTo
- io.ByteScanner
- io.RuneScanner

func NewReader(rd io.Reader) *Reader //创建读缓冲区
func (b *Reader) Read(p []byte) (n int, err error) //读取数据到p中
func (b *Reader) ReadByte() (c byte, err error) //读取一个字节数据
func (b *Reader) ReadBytes(delim byte) (line []byte, err error) //读取delim之前的所有字节数据
func (b *Reader) ReadRune() (r rune, size int, err error) //读取一个utf-8编码的unicode码值
func (b *Reader) UnreadRune() error //设置最后一次读的Rune未读, 若最后一次不是ReadRune, 返回error
func (b *Reader) ReadSlice(delim byte) (line []byte, err error) //读取数据直到遇到delim
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error) //读取一行数据, 根据\n或者\r\n
func (b *Reader) ReadString(delim byte) (line string, err error) //读取delim之前的所有string数据
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)//WriteTo方法实现了io.WriterTo接口

### type Writer

bufio.Writer 实现了如下接口：

- io.Writer
- io.ReaderFrom
- io.ByteWriter

func NewWriter(w io.Writer) *Writer //创建写缓冲区
func (b *Writer) Reset(w io.Writer)//丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w
func (b *Writer) WriteString(s string) (int, error) //写入一个string
func (b *Writer) WriteByte(c byte) error //写入一个Byte
func (b *Writer) WriteRune(r rune) (size int, err error) //写入一个字符, 例如’你’或者’c’
func (b *Writer) Write(p []byte) (nn int, err error) //写入一个字节数组
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error) //ReadFrom 实现了 io.ReaderFrom

### type ReadWriter

func NewReadWriter(r *Reader, w *Writer) *ReadWriter//申请创建一个新的、将读写操作分派给r和w 的ReadWriter

## type Scanner

func NewScanner(r io.Reader) *Scanner
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Scan() bool
func (s *Scanner) Bytes() []byte
func (s *Scanner) Text() string
func (s *Scanner) Err() error

```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    fmt.Println(scanner.Text()) // Println will add back the final '\n'
}
if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
}
```

[bufio_scaner.go](./code/bufio_scaner.go)

```go
//统计字数
// An artificial input source.
const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
scanner := bufio.NewScanner(strings.NewReader(input))
// Set the split function for the scanning operation.
scanner.Split(bufio.ScanWords)
// Count the words.
count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading input:", err)
}
fmt.Printf("%d\n", count)
```

```go
func main() {
    input := "abcdefghijkl"
    scanner := bufio.NewScanner(strings.NewReader(input))
    split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
        return 0, nil, nil
    }
    scanner.Split(split)
    buf := make([]byte, 2)
    scanner.Buffer(buf, bufio.MaxScanTokenSize)
    for scanner.Scan() {
        fmt.Printf("%s\n", scanner.Text())
    }
}
```

## 函数

### bufio.flush()

bufio.flush() 会将缓存区内容写入文件，当所有写入完成后，因为缓存区会存储内容，所以需要手动flush()到文件

### bufio.Available()

bufio.Available() 为buf可用容量，等于len(buf) - n

## 示例

### 示例：Peek、Read、Discard、Buffered

```go
func main() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0)
	b := make([]byte, 10)

	fmt.Println(buf.Buffered()) // 0
	s, _ := buf.Peek(5)
	s[0], s[1], s[2] = 'a', 'b', 'c'
	fmt.Printf("%d   %q\n", buf.Buffered(), s) // 16   "abcDE"

	buf.Discard(1)

	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d   %q   %v\n", buf.Buffered(), b[:n], err)
	}
	// 5   "bcDEFGHIJK"   <nil>
	// 0   "LMNOP"   <nil>
	// 6   "QRSTUVWXYZ"   <nil>
	// 0   "123456"   <nil>
	// 0   "7890"   <nil>
	// 0   ""   EOF
}
```
------------------------------

### 示例：ReadLine

```go
func main() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")
	buf := bufio.NewReaderSize(sr, 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFGHIJKLMNOP"   true   <nil>
	// "QRSTUVWXYZ"   false   <nil>
	// "1234567890"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部有一个换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF
}
```

------------------------------

### 示例：ReadSlice

```go
func main() {
	// 尾部有换行标记
	buf := bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	// "ABCDEFG\n"   <nil>
	// ""   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	// "ABCDEFG"   EOF
}
```

### 示例：Available、Buffered、WriteString、Flush

```go
func main() {
	buf := bufio.NewWriterSize(os.Stdout, 0)
	fmt.Println(buf.Available(), buf.Buffered()) // 4096 0

	buf.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Println(buf.Available(), buf.Buffered()) // 4070 26

	// 缓存后统一输出，避免终端频繁刷新，影响速度
	buf.Flush() // ABCDEFGHIJKLMNOPQRSTUVWXYZ
}
```

### 示例：扫描

```go
func main() {
	// 逗号分隔的字符串，最后一项为空
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 定义匹配函数（查找逗号分隔的字符串）
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			// 告诉 Scanner 扫描结束。
			return 0, data, bufio.ErrFinalToken
		} else {
			// 告诉 Scanner 没找到匹配项，让 Scan 填充缓存后再次扫描。
			return 0, nil, nil
		}
	}
	// 指定匹配函数
	scanner.Split(onComma)
	// 开始扫描
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	// 检查是否因为遇到错误而结束
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
```
------------------------------

### 示例：带检查扫描

```go
func main() {
	const input = "1234 5678 1234567901234567890 90"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 自定义匹配函数
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 获取一个单词
		advance, token, err = bufio.ScanWords(data, atEOF)
		// 判断其能否转换为整数，如果不能则返回错误
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		// 这里包含了 return 0, nil, nil 的情况
		return
	}
	// 设置匹配函数
	scanner.Split(split)
	// 开始扫描
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
```