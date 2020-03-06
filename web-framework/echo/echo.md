# echo

<http://go-echo.org/>

## 安装

go get github.com/labstack/echo/...

##  Hello, World

```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```

## Context
Context 是一个接口,含有请求和相应的引用，路径，路径参数，数据，注册的业务处理方法和读取请求和输出响应的API
```
type Context interface {
	Request() *http.Request
	SetRequest(r *http.Request)
	SetResponse(r *Response)
	Response() *Response
	IsTLS() bool
	IsWebSocket() bool
	Scheme() string
	RealIP() string
	Path() string
	SetPath(p string)
	Param(name string) string
	ParamNames() []string
	SetParamNames(names ...string)
	ParamValues() []string
	SetParamValues(values ...string)
	QueryParam(name string) string
	QueryParams() url.Values
	QueryString() string
	FormValue(name string) string
	FormParams() (url.Values, error)
	FormFile(name string) (*multipart.FileHeader, error)
	MultipartForm() (*multipart.Form, error)
	Cookie(name string) (*http.Cookie, error)
	SetCookie(cookie *http.Cookie)
	Cookies() []*http.Cookie
	Get(key string) interface{}
	Set(key string, val interface{})
	Bind(i interface{}) error
	Validate(i interface{}) error
	Render(code int, name string, data interface{}) error
	HTML(code int, html string) error
	HTMLBlob(code int, b []byte) error
	String(code int, s string) error
	JSON(code int, i interface{}) error
	JSONPretty(code int, i interface{}, indent string) error
	JSONBlob(code int, b []byte) error
	JSONP(code int, callback string, i interface{}) error
	JSONPBlob(code int, callback string, b []byte) error
	XML(code int, i interface{}) error
	XMLPretty(code int, i interface{}, indent string) error
	XMLBlob(code int, b []byte) error
	Blob(code int, contentType string, b []byte) error
	Stream(code int, contentType string, r io.Reader) error
	File(file string) error
	Attachment(file string, name string) error
	Inline(file string, name string) error
	NoContent(code int) error
	Redirect(code int, url string) error
	Error(err error)
	Handler() HandlerFunc
	SetHandler(h HandlerFunc)
	Logger() Logger
	SetLogger(l Logger)
	Echo() *Echo
	Reset(r *http.Request, w http.ResponseWriter)
}
```


## Cookie
```
c.SetCookie 	 //创建cookie
c.Cookie     	//读取cookie
cCookies	//所有cookie
```
### 创建cookie
```
func writeCookie(c echo.Context) error {
    cookie := new(http.Cookie)
    cookie.Name = "username"
    cookie.Value = "jon"
    cookie.Expires = time.Now().Add(24 * time.Hour)
    c.SetCookie(cookie)
    return c.String(http.StatusOK, "write a cookie")
}
使用 new(http.Cookie) 创建Cookie。
cookie 的属性值会被赋值给 http.Cookie 的可导出属性。
最后，使用 c.SetCookie(cookies) 来给响应添加 Set-Cookie 头。

	c.SetCookie(&http.Cookie{
		Name:       FlashName,
		Value:      url.QueryEscape(flashValue),
	})
```

### 读取cookie
```
cookie,err:=c.Cookie(FlashName)
fmt.Println(cookie.Name)
fmt.Println(cookie.Value)
```

## bind,render,alidate捆绑,渲染,验证
c.Bind
c.Render
c.Alidate

## 响应
c.String	//发送带状态码的纯文本响应
c.HTML		//发送一个带状态码的简单 html 响应
