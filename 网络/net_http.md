# net/http

tags : golang

---

func CanonicalHeaderKey(s string) string
func DetectContentType(data []byte) string // 确定数据的Content-Type

## func Error(w ResponseWriter, error string, code int)

使用指定的错误信息和状态码回复请求，将数据写入w。错误信息必须是明文

## func Handle(pattern string, handler Handler)

注册HTTP处理器handler和对应的模式pattern(注册到DefaultServeMux)

[Handle.go](./code/Handle.go)

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

册一个处理器函数handler和对应的模式pattern(注册到DefaultServeMux)

[HandleFunc.go](./code/HandelFunc.go)

## func ListenAndServe(addr string, handler Handler) error

监听TCP地址addr，并且会使用handler参数调用Serve函数处理接收到的连接。handler参数一般会设为nil，此时会使用DefaultServeMux

[ListenAndServe.go](./code/ListenAndServe.go)

## func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error

期望HTTPS连接
[ListenAndServeTLS.go](./code/ListenAndServeTLS.go)

## func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser

类似io.LimitReader,返回值的Read方法在读取的数据超过大小限制时会返回非EOF错误

## func NotFound(w ResponseWriter, r *Request)

返回一个简单的请求处理器，该处理器会对每个请求都回复"404 page not found"

## func ParseHTTPVersion(vers string) (major, minor int, ok bool)

解析HTTP版本字符串

## func ParseTime(text string) (t time.Time, err error)

用3种格式TimeFormat, time.RFC850和time.ANSIC尝试解析一个时间头的值(如Date: header)

## func ProxyFromEnvironment(req *Request) (*url.URL, error)

使用环境变量$HTTP_PROXY和$NO_PROXY(或$http_proxy和$no_proxy)的配置返回用于req的代理

## func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)

返回一个代理函数(用于Transport类型)

## func Redirect(w ResponseWriter, r *Request, url string, code int)

回复请求一个重定向地址urlStr和状态码code

## func Serve(l net.Listener, handler Handler) error

接手监听器l收到的每一个连接，并为每一个连接创建一个新的服务go程

## func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error

## func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)

使用提供的ReadSeeker的内容回复请求

## func ServeFile(w ResponseWriter, r *Request, name string)

回复请求name指定的文件或者目录的内容

### 加载图片或文件

```go
func foo(w http.ResponseWriter, r *http.Request) {
    fp := path.Join(".", "gopherbw.png")
    http.ServeFile(w, r, fp)
}
```

## func SetCookie(w ResponseWriter, cookie *Cookie)

在w的头域中添加Set-Cookie头

## func StatusText(code int) string

返回HTTP状态码code对应的文本，如220对应"OK"

## type client

### func (c *Client) CloseIdleConnections()

### func (c *Client) Do(req *Request) (*Response, error)

Do方法发送请求,返回HTTP回复.它会遵守客户端c设置的策略(如重定向、cookie、认证)

### func (c *Client) Get(url string) (resp *Response, err error)

向指定的URL发出一个GET请求

### func (c *Client) Head(url string) (resp *Response, err error)

向指定的URL发出一个HEAD请求

### func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)

向指定的URL发出一个POST请求

### func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)

向指定的URL发出一个POST请求，url.Values类型的data会被编码为请求的主体

```go
type Client struct {
    Transport RoundTripper // Transport指定执行独立,单次HTTP请求的机制.如果为nil，则使用DefaultTransport。
    CheckRedirect func(req *Request, via []*Request) error // CheckRedirect指定处理重定向的策略
    Jar CookieJar // Jar指定cookie管理器
    Timeout time.Duration// Timeout指定本类型的值执行请求的时间限制
}
```

### example

>func Get(url string) (resp *Response, err error)

```go
res, err := http.Get("http://www.google.com/robots.txt")
if err != nil {     log.Fatal(err) }
robots, err := ioutil.ReadAll(res.Body)
res.Body.Close()
if err != nil {    log.Fatal(err) }
fmt.Printf("%s", robots)
```

## type Dir

### func (d Dir) Open(name string) (File, error)

Dir使用限制到指定目录树的本地文件系统实现了http.FileSystem接口。空Dir被视为"."，即代表当前目录

## type Cookie

```go
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string
    // MaxAge=0表示未设置Max-Age属性
    // MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
    // MaxAge>0表示存在Max-Age属性，单位是秒
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // 未解析的“属性-值”对的原始文本
}
```

## func (c *Cookie) String() string 返回该cookie的序列化结果

返回该cookie的序列化结果

## type Handler interface { ServeHTTP(ResponseWriter, *Request) }

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

### func FileServer(root FileSystem) Handler

返回一个简单的请求处理器，该处理器会对每个请求都回复"404 page not found"
`log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))`
`http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))`

### func NotFoundHandler() Handler

返回一个简单的请求处理器，该处理器会对每个请求都回复"404 page not found"。

```go
mux := http.NewServeMux()
mux.Handle("/resources", http.NotFoundHandler())
mux.Handle("/resources/people/", newPeopleHandler())
log.Fatal(http.ListenAndServe(":8080", mux))
```

### func RedirectHandler(url string, code int) Handler

返回一个请求处理器，该处理器会对每个请求都使用状态码code重定向到网址url

```go
mux := http.NewServeMux()

rh := http.RedirectHandler("http://example.org", 307)
mux.Handle("/foo", rh)

log.Println("Listening...")
http.ListenAndServe(":3000", mux)
```

### func StripPrefix(prefix string, h Handler) Handler

返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理

### func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

r返回一个采用指定时间限制的请求处理器

## type HandlerFunc func(ResponseWriter, *Request)

HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Handler接口

### func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

ServeHTTP方法会调用f(w, r)

## type Header 代表HTTP头域的键值对

type Header map[string][]string

### Add Del Get Set Write

func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
//WriteSubset以有线格式将头域写入w

### 带header的client

```go
client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://example.com")

req, err := http.NewRequest("GET", "http://example.com", nil)
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
```

## type Request

```go
type Request struct {
    Method string   // 指定HTTP方法(GET,POST,PUT等).对客户端,""代表GET
    URL *url.URL    // 在服务端表示被请求的URI,在客户端表示要访问的URL,在客户端,URL的Host字段指定了要连接的服务器
    Header Header   // Header字段用来表示HTTP请求的头域
    Body io.ReadCloser   // Body是请求的主体
    GetBody func() (io.ReadCloser, error)
    ContentLength int64    // 记录相关内容的长度
    Response *Response
    Close bool   // 在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接
    // 在服务端，Host指定URL会在其上寻找资源的主机。
    // Host的格式可以是"host:port"。
    // 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
    // 如过该字段为""，Request.Write方法会使用URL字段的Host。
    Host string
    RemoteAddr string   // 允许HTTP服务器和其他软件记录该请求的来源地址,一般用于日志;客户端会忽略
    RequestURI string  // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
    Cancel <-chan struct{}
    Form url.Values   // 是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据
    PostForm url.Values   // 解析好的POST或PUT的表单数据

    Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
    TransferEncoding []string  // 按从最外到最里的顺序列出传输编码
    // MultipartForm是解析好的多部件表单，包括上传的文件。
    // 本字段只有在调用ParseMultipartForm后才有效
    MultipartForm *multipart.Form
    Trailer Header    // 指定了会在请求主体之后发送的额外的头域
    TLS *tls.ConnectionState  // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
    ctx context.Context
}
```

### func NewRequest(method, url string, body io.Reader) (*Request, error)
### func ReadRequest(b *bufio.Reader) (*Request, error)
### func (r *Request) AddCookie(c *Cookie)
### func (r *Request) BasicAuth() (username, password string, ok bool)
### func (r *Request) Context() context.Context
### func (r *Request) Cookie(name string) (*Cookie, error)
### func (r *Request) Cookies() []*Cookie
### func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
### func (r *Request) FormValue(key string) string
### func (r *Request) MultipartReader() (*multipart.Reader, error)
### func (r *Request) ParseForm() error
### func (r *Request) ParseMultipartForm(maxMemory int64) error
### func (r *Request) PostFormValue(key string) string
### func (r *Request) ProtoAtLeast(major, minor int) bool
### func (r *Request) Referer() string
### func (r *Request) SetBasicAuth(username, password string)
### func (r *Request) UserAgent() string
### func (r *Request) WithContext(ctx context.Context) *Request
### func (r *Request) Write(w io.Writer) error
### func (r *Request) WriteProxy(w io.Writer) error


## type Response

```go
type Response struct {
	Header Header   //保管头域的键值对
    Body io.ReadCloser  //代表回复的主体
	ContentLength int64 //记录相关内容的长度
	Close bool  // 记录头域是否指定应在读取完主体后关闭连接。（即Connection头）
	Request *Request    // Request是用来获取此回复的请求
	
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	TransferEncoding []string   //按从最外到最里的顺序列出传输编码
  	Uncompressed bool  //是否压缩
	Trailer Header  // 保存和头域相同格式的trailer键值对，和Header字段相同类型
	TLS *tls.ConnectionState    // 包含接收到该回复的TLS连接的信息
```

func Get(url string) (resp *Response, err error)
func Head(url string) (resp *Response, err error)
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
func PostForm(url string, data url.Values) (resp *Response, err error)
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
func (r *Response) Cookies() []*Cookie
func (r *Response) Location() (*url.URL, error)
func (r *Response) ProtoAtLeast(major, minor int) bool
func (r *Response) Write(w io.Writer) error

## type ResponseWriter

```go
mux := http.NewServeMux()
mux.HandleFunc("/sendstrailers", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Trailer", "AtEnd1, AtEnd2")
    w.Header().Add("Trailer", "AtEnd3")

    w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
    w.WriteHeader(http.StatusOK)

    w.Header().Set("AtEnd1", "value 1")
    io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
    w.Header().Set("AtEnd2", "value 2")
    w.Header().Set("AtEnd3", "value 3") // These will appear as trailers.
})
```


## 服务端 server

### type Server struct

```go
type Server struct {
	Addr    string  
	Handler Handler // handler to invoke, http.DefaultServeMux if nil
	ReadTimeout time.Duration// 请求的读取操作在超时前的最大持续时间
	WriteTimeout time.Duration // 回复的写入操作在超时前的最大持续时间
 	MaxHeaderBytes int   // 请求的头域最大长度，如为0则用DefaultMaxHeaderBytes
	mu         sync.Mutex

	TLSConfig *tls.Config// 可选的TLS配置，用于ListenAndServeTLS方法
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	ConnState func(net.Conn, ConnState)
 	ErrorLog *log.Logger   // ErrorLog指定一个可选的日志记录器

	ReadHeaderTimeout time.Duration
	IdleTimeout time.Duration

	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	doneChan   chan struct{}
	onShutdown []func()

	disableKeepAlives int32     // accessed atomically.
	inShutdown        int32     // accessed atomically (non-zero means we're in Shutdown)
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used


}
```

func (srv *Server) Close() error
func (srv *Server) ListenAndServe() error
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
func (srv *Server) RegisterOnShutdown(f func())
func (srv *Server) Serve(l net.Listener) error
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
func (srv *Server) SetKeepAlivesEnabled(v bool)
func (srv *Server) Shutdown(ctx context.Context) error

### type ServeMux

ServeMux类型是HTTP请求的多路转接器。它会将每一个接收的请求的URL与一个注册模式的列表进行匹配，并调用和URL最匹配的模式的处理器

```go
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}
```
func NewServeMux() *ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

## index.html

```go
func main() {
    http.HandleFunc("/hello", func(res http.ResponseWriter, request *http.Request) {
        res.Write([]byte("hello"))
    })
    http.HandleFunc("/index.html",indexM)
    http.ListenAndServe(":8080", nil)
}
func indexM(res http.ResponseWriter,req *http.Request){
    fmt.Fprintf(res,"index.html")
}
//http.ListenAndServe     监听端口
//http.HandleFunc          定义路由
```

## 简单的web服务器

```go
import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
    http.HandleFunc("/", sayhelloName)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

## type Transport

func (t *Transport) CancelRequest(req *Request)
func (t *Transport) CloseIdleConnections()
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
func (t *Transport) RoundTrip(req *Request) (*Response, error)