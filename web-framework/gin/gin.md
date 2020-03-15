# gin

[TOC]

https://gin-gonic.com/

<https://github.com/gin-gonic/gin>

<https://github.com/skyhee/gin-doc-cn>

https://learnku.com/docs/gin-gonic/2019

## type Context 
```go
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter
	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string
	engine *Engine
	Keys map[string]interface{}
	Errors errorMsgs
	Accepted []string
	queryCache url.Values
	formCache url.Values
}
```

#http服务器

```go
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run()
```
```gin.H 是 map[string]interface{} 的一种快捷方式```

//生成具有转义的非ASCII字符的 ASCII-only JSON 
```c.AsciiJSON(http.StatusOK, map[string]interface{}{"lang": "GO语言", "tag": "<br>",}) ```

//按字面对特殊字符字符进行编码 
```c.PureJSON(200, gin.H{"html": "<b>Hello, world!</b>",})  ```

//防止json劫持.如果给定的结构是数组值，则默认预置 "while(1)," 到响应体 
```c.SecureJSON(http.StatusOK, []string{"lena","austin","foo"})```

//curl http://127.0.0.1:8080/JSONP?callback=x
```c.JSONP(http.StatusOK, gin.H{"foo": "bar"})```
