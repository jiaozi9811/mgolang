//https://segmentfault.com/a/1190000015591319
package main

import (
	"fmt"
	"strings"
	"io"
)
type alphaReader struct {
    // alphaReader 里组合了标准库的 io.Reader
    reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
    return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
    if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
        return r
    }
    return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
    // 这行代码调用的就是 io.Reader
    n, err := a.reader.Read(p)
    if err != nil {
        return n, err
    }
    buf := make([]byte, n)
    for i := 0; i < n; i++ {
        if char := alpha(p[i]); char != 0 {
            buf[i] = char
        }
    }

    copy(p, buf)
    return n, nil
}

func main() {
    //  使用实现了标准库 io.Reader 接口的 strings.Reader 作为实现
    reader := newAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
    p := make([]byte, 4)
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            break
        }
        fmt.Print(string(p[:n]))
    }
    fmt.Println()
}

/*
func main() {
    // file 也实现了 io.Reader
    file, err := os.Open("./alpha_reader3.go")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    
    // 任何实现了 io.Reader 的类型都可以传入 newAlphaReader
    // 至于具体如何读取文件，那是标准库已经实现了的，我们不用再做一遍，达到了重用的目的
    reader := newAlphaReader(file)
    p := make([]byte, 4)
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            break
        }
        fmt.Print(string(p[:n]))
    }
    fmt.Println()
}
*/