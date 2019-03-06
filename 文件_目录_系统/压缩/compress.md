# archive/compress

[TOC]

## archive/tar

type Header

- func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
- func (h *Header) FileInfo() os.FileInfo
type Reader
- func NewReader(r io.Reader) *Reader
- func (tr *Reader) Next() (*Header, error)
- func (tr *Reader) Read(b []byte) (n int, err error)
type Writer
- func NewWriter(w io.Writer) *Writer
- func (tw *Writer) WriteHeader(hdr *Header) error
- func (tw *Writer) Write(b []byte) (n int, err error)
- func (tw *Writer) Flush() error
- func (tw *Writer) Close() error

>打完包后一定要用Close()关闭,因为 tar.Writer 使用了缓存，tw.Close()会将缓存中的数据写入到文件中，同时 tw.Close() 还会向 .tar 文件的最后写入结束信息，如果不关闭 tw 而直接退出程序，那么将导致 .tar 文件不完整

>存储在 .tar 文件中的每个文件都由两部分组成：文件信息和文件内容，所以向 .tar 文件中写入每个文件都要分两步：第一步写入文件信息，第二步写入文件数据。对于目录来说，由于没有内容可写，所以只需要写入目录信息即可

```go
type Header struct {
   Name       string    // 文件名称
   Mode       int64     // 文件的权限和模式位
   Uid        int       // 文件所有者的用户 ID
   Gid        int       // 文件所有者的组 ID
   Size       int64     // 文件的字节长度
   ModTime    time.Time // 文件的修改时间
   Typeflag   byte      // 文件的类型
   Linkname   string    // 链接文件的目标名称
   Uname      string    // 文件所有者的用户名
   Gname      string    // 文件所有者的组名
   Devmajor   int64     // 字符设备或块设备的主设备号
   Devminor   int64     // 字符设备或块设备的次设备号
   AccessTime time.Time // 文件的访问时间
   ChangeTime time.Time // 文件的状态更改时间
}
```

### 写入文件信息

- 首先将被打包文件的信息填入 tar.Header 结构体中
- 然后将结构体写入.tar文件中。这样完成第一步（写入文件信息）操作

**tar.FileInfoHeader**函数可以直接通过os.FileInfo创建tar.Header，并自动填写 tar.Header的大部分信息

```go
// srcFile 是要打包的文件的完整路径
fi, err := os.Stat(srcFile)
if err != nil {
    return err
}
// 根据 os.FileInfo 创建 tar.Header 结构体
hdr, err := tar.FileInfoHeader(fi, "")
if err != nil {
    return err
}
// 将 tar.Header 写入 .tar 文件中
err = tw.WriteHeader(hdr)
if err != nil {
    return err
}
```

这里的hdr是文件信息结构体，已经填写完毕。如果要填写的更详细，可以自己将hdr补充完整

下面通过tw.WriteHeader方法将hdr写入.tar文件中(tw是我们刚才创建的tar.Writer)：

### 写入文件数据

通过 tw.Write 方法写入

```go
// 打开要打包的文件准备读取
fr, err := os.Open(srcFile)
if err != nil {
    return err
}
defer fr.Close()
// 将文件数据写入 .tar 文件中，这里通过 io.Copy 函数实现数据的写入
_, err = io.Copy(tw, fr)
if err != nil {
    return err
}
```

### 读出文件数据

>从 .tar 文件中读出数据是通过 tar.Reader 完成的，所以首先要创建 tar.Reader，可以通过 tar.NewReader 方法来创建它，该方法要求提供一个 os.Reader 对象，以便从该对象中读出数据。可以先打开一个 .tar 文件，然后将该文件提供给 tar.NewReader 使用。

```go
// 打开要解包的文件，srcTar 是要解包的 .tar 文件的路径
fr, er := os.Open(srcTar)
if er != nil {
	return er
}
defer fr.Close()

// 创建 tar.Reader，准备执行解包操作
tr := tar.NewReader(fr)

//此时，我们就拥有了一个 tar.Reader 对象 tr，可以用 tr.Next() 来遍历包中的文件，然后将文件的数据保存到磁盘中：

// 遍历包中的文件
for hdr, er := tr.Next(); er != io.EOF; hdr, er = tr.Next() {
    if er != nil {
        return er
    }

    // 获取文件信息
    fi := hdr.FileInfo()

    // 创建空文件，准备写入解压后的数据
    fw, _ := os.Create(dstFullPath)
    if er != nil {
        return er
    }
    defer fw.Close()

    // 写入解压后的数据
    _, er = io.Copy(fw, tr)
    if er != nil {
        return er
    }
    // 设置文件权限
    os.Chmod(dstFullPath, fi.Mode().Perm())
}
```

[tar_untar.go](./tar_untar.go)

## compress/gzip [top](#archive/compress)

## archive/zip [top](#archive/compress)

## compress/bzip2 [top](#archive/compress)

## compress/flate [top](#archive/compress)

## compress/lzw [top](#archive/compress)

## compress/zlib [top](#archive/compress)

### compress tar.gz

```go

package main

import (
    // ...
    "compress/gzip" // 这里导入 compress/gzip 包
    // ...
)

func Tar(src string, dstTar string, failIfExist bool) (err error) {
    // ...
    fw, er := os.Create(dstTar)
    // ...
    gw := gzip.NewWriter(fw) // 这里添加一个 gzip.Writer
    // ...
    tw := tar.NewWriter(gw) // 这里传入 gw
    // ...
}

func UnTar(srcTar string, dstDir string) (err error) {
    // ...
    fr, er := os.Open(srcTar)
    // ...
    gr, er := gzip.NewReader(fr) // 这里添加一个 gzip.Reader
    // ...
    tr := tar.NewReader(gr) // 这里传入 gr
    // ...
}
```