# io/ioutil

[TOC]

## func NopCloser(r io.Reader) io.ReadCloser

## func ReadAll(r io.Reader) ([]byte, error)

//l从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误

## func ReadFile(filename string) ([]byte, error)

// 从filename指定的文件中读取数据并返回文件的内容

## func ReadDir(dirname string) ([]os.FileInfo, error)

返回dirname指定的目录的目录信息的有序列表

```go
files, err := ioutil.ReadDir(".")
if err != nil {
    log.Fatal(err)
}

for _, file := range files {
    fmt.Println(file.Name())
}
```

## func TempDir(dir, prefix string) (name string, err error)

在dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径

```go
content := []byte("temporary file's content")
dir, err := ioutil.TempDir("", "example")
if err != nil {
    log.Fatal(err)
}

defer os.RemoveAll(dir) // clean up

tmpfn := filepath.Join(dir, "tmpfile")
if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
    log.Fatal(err)
}
```

## func TempFile(dir, prefix string) (f *os.File, err error)

在dir目录下创建一个新的、使用prefix为前缀的临时文件，以读写模式打开该文件并返回os.File指针

```go
content := []byte("temporary file's content")
tmpfile, err := ioutil.TempFile("", "example.*.txt")
if err != nil {
    log.Fatal(err)
}

defer os.Remove(tmpfile.Name()) // clean up

if _, err := tmpfile.Write(content); err != nil {
    tmpfile.Close()
    log.Fatal(err)
}
if err := tmpfile.Close(); err != nil {
    log.Fatal(err)
}
```

## func WriteFile(filename string, data []byte, perm os.FileMode) error

函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件