package main
import (
    "bytes"
    "fmt"
    "strings"
    "time"
)
func main() {
    var buffer bytes.Buffer
    s := time.Now()
    for i := 0; i < 100000; i++ {
        buffer.WriteString("test is here\n")
    }
    buffer.String()  // 拼接结果
    e := time.Now()
    fmt.Println("taked time is ", e.Sub(s).Seconds())
    s = time.Now()
    str := ""
    for i := 0; i < 100000; i++ {
        str += "test is here\n"
    }
    e = time.Now()
    fmt.Println("taked time is ", e.Sub(s).Seconds())
    s = time.Now()
    var sl []string
    for i := 0; i < 100000; i++ {
        sl = append(sl, "test is here\n")
    }
    strings.Join(sl, "")
    e = time.Now()
    fmt.Println("taked time is", e.Sub(s).Seconds())
}