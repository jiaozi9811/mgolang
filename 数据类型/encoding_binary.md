# encoding/binary

[TOC]

此包实现了数字和字节序列之间的简单转换以及varint的编码和解码

包中有两个文件
>binary.go  用于数字和字节的转换
>varint.go  用于数字的编码和解码

varint函数使用可变长度编码对单个整数值进行编码和解码；较小的值需要较少的字节

## 字节序(Endianness)

就是字节的顺序，指的是多字节的数据在内存中的存放顺序

字节序分为:

- 大端序(Big Endian)--是指低地址端 存放 高位字节
- 小端序(Little Endian)--是指低地址端 存放 低位字节

>网络传输一般采用 Big Endian，也被称之为网络字节序，或网络序。当两台采用不同字节序的主机通信时，在发送数据之前都必须经过字节序的转换成为网络字节序后再进行传输

![大端序](https://upload.wikimedia.org/wikipedia/commons/thumb/5/54/Big-Endian.svg/280px-Big-Endian.svg.png "大端序")

![小端序](https://upload.wikimedia.org/wikipedia/commons/thumb/e/ed/Little-Endian.svg/280px-Little-Endian.svg.png "小端序")

提供了两个变量,用于定义大端和小端
>var BigEndian bigEndian
>var LittleEndian littleEndian

## 通过接口读写实现转换

func Read(r io.Reader, order ByteOrder, data interface{}) error
func Write(w io.Writer, order ByteOrder, data interface{}) error
func Size(v interface{}) int

## 通过varint转换

### 字节转换为数据

func Uvarint(buf []byte) (uint64, int)
func Varint(buf []byte) (int64, int)

### 数据转换为字节

func PutUvarint(buf []byte, x uint64) int
func PutVarint(buf []byte, x int64) int
func ReadVarint(r io.ByteReader) (int64, error)

```go
//binary.Read
var pi float64
b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
buf := bytes.NewReader(b)
err := binary.Read(buf, binary.LittleEndian, &pi)
if err != nil {
    fmt.Println("binary.Read failed:", err)
}
fmt.Print(pi)
```

```go
//binary.Write
buf := new(bytes.Buffer)
var pi float64 = math.Pi
err := binary.Write(buf, binary.LittleEndian, pi)
if err != nil {
    fmt.Println("binary.Write failed:", err)
}
fmt.Printf("% x", buf.Bytes())
```

```go
//PutUvarint
buf := make([]byte, binary.MaxVarintLen64)

for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
    n := binary.PutUvarint(buf, x)
    fmt.Printf("%x\n", buf[:n])
}
```

```go
//PutVarint
buf := make([]byte, binary.MaxVarintLen64)

for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
    n := binary.PutVarint(buf, x)
    fmt.Printf("%x\n", buf[:n])
}
```

```go
//Uvarint
inputs := [][]byte{
    []byte{0x01},
    []byte{0x02},
    []byte{0x7f},
    []byte{0x80, 0x01},
    []byte{0xff, 0x01},
    []byte{0x80, 0x02},
}
for _, b := range inputs {
    x, n := binary.Uvarint(b)
    if n != len(b) {
        fmt.Println("Uvarint did not consume all of in")
    }
    fmt.Println(x)
}
```