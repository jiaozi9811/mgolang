//https://segmentfault.com/a/1190000015591319
package main

import (
	"io"
	"fmt"
)

type alphaReader struct {
    src string// 资源
    cur int// 当前读取到的位置 
}

// 创建一个实例
func newAlphaReader(src string) *alphaReader {
    return &alphaReader{src: src}
}

// 过滤函数
func alpha(r byte) byte {
    if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
        return r
    }
    return 0
}

// Read 方法
func (a *alphaReader) Read(p []byte) (int, error) {
    // 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
    if a.cur >= len(a.src) {
        return 0, io.EOF
    }

    // x 是剩余未读取的长度
    x := len(a.src) - a.cur
    n, bound := 0, 0
    if x >= len(p) {        
        bound = len(p) //剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
    } else if x < len(p) {        
        bound = x// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
    }

    buf := make([]byte, bound)
    for n < bound {
        // 每次读取一个字节，执行过滤函数
        if char := alpha(a.src[a.cur]); char != 0 {
            buf[n] = char
        }
        n++
        a.cur++
    }    
    copy(p, buf)// 将处理后得到的 buf 内容复制到 p 中
    return n, nil
}

func main() {
    reader := newAlphaReader("Hello! It's 9am, where is the sun?")
    p := make([]byte, 4)
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            break
        }
        fmt.Print(string(p[:n])+`|`)
    }
    fmt.Println()
}