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
c.JSON		//发送一个带状态码的 json 对象
c.JSONPretty	//打印出的json数据带有缩紧(可以使用空格和tab)，更为好看
c.JSONBlob	//直接发送一个已经转换好的 json 对象
c.XML		//转换 golang 对象为 xml 数据发送响应
c.XMLPretty	
c.File		//发送一个文件为内容的响应
c.Attachment	//发送文件的方法类似，只是它会多提供一个名称
c.Blob		//发送任意类型的数据。需要提供 content type
c.Stream	//发送任意数据流响应。需要提供 content type，io.Reader 和状态码
c.NoContent	//发送空内容

c.Redirect	//提供一个 url 用于重定向

## 路由 

路径匹配顺序 
- Static (固定路径)
- Param (参数路径)
- Match any (匹配所有)

### 组路由 

### 路由列表 
Echo#Routes() []*Route会根据定义的顺序列出所有已经注册的路由。每一个路由包含 http 方法，路径和对应的处理器
```
data, err := json.MarshalIndent(e.Routes(), "", "  ")
if err != nil {
    return err
}
ioutil.WriteFile("routes.json", data, 0644)
```

## 静态文件 

用法1 
```
e := echo.New()
e.Static("/static", "assets")
```

用法2 
```
e := echo.New()
e.Static("/", "assets")
```

## Echo#File() 
```Echo#File(path, file string)```  
使用 url 路径注册一个新的路由去访问某个静态文件 

用法 1  
将public/index.html作为主页  
```e.File("/", "public/index.html")```  

用法 2  
给images/favicon.ico一个静态路径  
```e.File("/favicon.ico", "images/favicon.ico")``` 

## 模板

### 模板渲染 
Context#Render(code int, name string, data interface{}) error用于渲染一个模板，然后发送一个 text/html 的状态响应。可以使用任何模板引擎，只要赋值给Echo.Renderer 

## 中间件 
Recover中间件从 panic 链中的任意位置恢复程序， 打印堆栈的错误信息，并将错误集中交给HTTPErrorHandler处理 
e.User(middleware.Recover()) 

Logger 中间件记录了每一个请求的信息
e.Use(middleware.Logger()) 

HTTPS 重定向中间件将 http 请求重定向到 https,例如，http://laily.net将被重定向到https://laily.net。 
e.Pre(middleware.HTTPSRedirect()) 

HTTPS WWW 重定向将 http 请求重定向到带 www 的https 请求 。例如，http://laily.net将被重定向到https://www.laily.net。
e.Pre(middleware.HTTPSWWWRedirect()) 

HTTPS NonWWW 将 http 请求重定向到不带 www 的 https 请求。例如，http://www.laily.net将被重定向到https://laily.net 
e.Pre(middleware.HTTPSNonWWWRedirect())

将不带 www 的请求重定向到带 www 的请求。例如，http://laily.net重定向到http://www.laily.net 
e.Pre(middleware.WWWRedirect()) 

将带 www 的请求重定向到不带 www 的请求。例如，http://www.laily.net重定向到http://laily.net 
e.Pre(middleware.NonWWWRedirect()) 

## HTTP2/HTTPS

生成自验证证书,命令会生一个cert.pem和key.pem文件
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost

e.Logger.Fatal(e.StartTLS(":443","cert.pem","key.pem"))
