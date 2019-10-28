# negroni

[TOC]

```go get -u github.com/urfave/negroni```

```go
package main

import (
	"github.com/urfave/negroni"
	"io"
	"net/http"
)

func myHandler(rw http.ResponseWriter,r *http.Request)  {
	rw.Header().Set("CContent-Type","test/plain")
	io.WriteString(rw,"hello world")
}

func handler() http.Handler {
	return http.HandlerFunc(myHandler)
}

func main()  {
	n:=negroni.Classic()
  
  mux:=http.NewServeMux()
	mux.Handle("/",handler())
	mux.HandleFunc("/flysnow", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw,"blog:www.www.org\n")
		io.WriteString(rw,"wechat:www.org")
	})
  
	n.UseHandler(handler())
	n.Run(":1234")
}
```
negroni.Classic()返回一个negroni实例,通过这个实例可以添加一些中间件
negroni.Run()启动一个服务,Run方法和http.ListenAndServe是等价的
