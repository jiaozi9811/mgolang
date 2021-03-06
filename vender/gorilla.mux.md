# gorilla/mux

[TOC]

https://github.com/gorilla/mux

https://www.gorillatoolkit.org/pkg/mux


https://www.jianshu.com/p/8ade70e51210

mux 名称来源于 HTTPrequest multiplexer,类似于官方包http.ServeMux

mux.Router将定义一个路由列表，其中每一个路由定义一个对应的请求url及处理方法

每个函数都有两个参数：一个http.ResponseWriter用于向请求者发送一个响应，一个指向http.Request的指针，该指针封装了底层HTTP请求的信息

## 使用
```
r:=mux.Router()                 初始化路由
r.HandleFunc("/",HomeHandler)   路由注册
func HomeHandler(w http.ResponseWriter,r *http.Request){
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w,"this is home")
}
```
### 带变量的url路由注册
```r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)```

### 组合使用
```r.HandleFunc("/products",ProductsHandler).Host("www.example.com").Methods("GET").Schemes("http")```

### 子路由
```
r:=mux.NewRouter()
s:=r.PathPrefix("/products").Subrouter()
s.handleFunc("/",ProductsHandler)
s.HandleFunc("/{key}/",ProductHandler)
```

### 定义路由别名
```r.HandleFunc("/articles/{category}/{id:[0-9]+}",ArticleHandler).Name("article")```

### 静态文件路由
```
flag.StringVar(&dir,"dir",".","the directory to serve files form.Defaults to the current dir")
flag.Parse()
r:=mux.NewRouter()
r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir(dir))))
```

### 生成url
```
url,err:=r.Get("router_name").URL("key1","val1","key2","val2")
```

```
r:=mux.NewRouter()
r.Host("{subdomain}.domain.com").
  Path("/articles/{category}/{id:[0-9]+}").
  Queries("filter", "{filter}").
  HandlerFunc(ArticleHandler).
  Name("article")
// url.String() will be "http://news.domain.com/articles/technology/42?filter=gorilla"
url,err:=r.Get("article").URL("subdomain","news",
                       "category","technology",
                       "id","42",
                       "filter","gorilla")
```

### walk方法
walk可以遍历访问所有已注册的路由
```
func main()  {
	r:=mux.NewRouter()
	r.HandleFunc("/",handler)
	r.HandleFunc("/products",handler).Methods(http.MethodGet)
	r.HandleFunc("/articles",handler).Methods(http.MethodGet)
	r.HandleFunc("articles/{id}",handler).Methods(http.MethodPut)
	r.HandleFunc("/authors",handler).Queries("surname","surname")

	err:=r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTempate,err:=route.GetPathTemplate()
		if err==nil{fmt.Println("ROUTE:",pathTempate)}

		pathRegexp,err:=route.GetPathRegexp()
		if err==nil{fmt.Println("Path regrxp:",pathRegexp)}

		queriexTemplates,err:=route.GetQueriesTemplates()
		if err==nil{fmt.Println("Queries templates:",strings.Join(queriexTemplates,","))}

		queriesRegrxps,err:=route.GetQueriesRegexp()
		if err==nil{fmt.Println("Queries regexps:",strings.Join(queriesRegrxps,","))}

		methods,err:=route.GetMethods()
		if err==nil{fmt.Println("Methods:",strings.Join(methods,","))}

		fmt.Println()
		return nil
	})

	if err!=nil{fmt.Println(err)}

	http.Handle("/",r)
}
```
### Middleware中间件
mux支持为路由添加中间件


```go
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func main()  {
	var dir string
	flag.StringVar(&dir,"dir",".","the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	//初始化router
	r:=mux.NewRouter()
	//静态文件路由
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir(dir))))
	//普通路由
	r.HandleFunc("/",HomeHandler)
	//指定host
	r.HandleFunc("/host",HostHandler).Host("localhost")
	//带变量的url路由
	r.HandleFunc("/user/{id}",GetUserHandler).Methods(http.MethodGet).Name("user")

	url,_:=r.Get("user").URL("id","5")
	fmt.Println("user url: ",url.String())

	// 遍历已注册的路由
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	r.Use(TestMiddleware)
	http.ListenAndServe(":3000", r)
}

func HomeHandler(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"this is home")
}

func HostHandler(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"the host is localhost")
}

func GetUserHandler(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"this is get user, and the user id is %v",vars["id"])
}
func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("middleware print: ",r.RequestURI)
		next.ServeHTTP(w,r)
	})
}
```

## get方法
```
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	hostname string
	port int
)

func init()  {
	flag.StringVar(&hostname,"hostname","0.0.0.0","The hostname or IP on which the REST server will listen")
	flag.IntVar(&port,"port",8080,"The port on which the TEST server will listen")
}
func main()  {
	flag.Parse()
	address:=fmt.Sprintf("%s:%d",hostname,port)
	log.Println("REST service listening on",address)

	r:=mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/service/get",MyGetHandler).Methods(http.MethodGet)

	err:=http.ListenAndServe(address,r)
	if err!=nil{log.Fatalln("ListenAndServe err:",err)}
	log.Println("server end")
}

func MyGetHandler(w http.ResponseWriter,r *http.Request)  {
	vals:=r.URL.Query()
	param,_:=vals["servicename"]

	var res=map[string]string{"result":"succ","name":param[0]}
	response,_:=json.Marshal(res)
	w.Header().Set("Content-Type:","application/json")
	w.Write(response)
}

//$ curl -X GET http://localhost:8080/api/service/get?servicename=myname
//{"name":"myname","result":"succ"}
```


## POST方法
```
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	hostname string
	port int
)

func init() {
	flag.StringVar(&hostname, "hostname", "0.0.0.0", "The hostname or IP on which the REST server will listen")
	flag.IntVar(&port, "port", 8080, "The port on which the TEST server will listen")
}
func main()  {
	flag.Parse()
	address:=fmt.Sprintf("%s:%d",hostname,port)
	log.Println("REST service listening on",address)

	r:=mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/service/{servicename}/post",MyPostHandler).Methods(http.MethodPost)

	err:=http.ListenAndServe(address,r)
	if err!=nil{log.Fatalln("ListenAndServe err:",err)}
	log.Println("server end")
}
func MyPostHandler(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	servicename:=vars["servicename"]

	var req map[string]interface{}
	body,_:=ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&req)
	servicetype:=req["servicetype"].(string)

	res:=map[string]string{"result":"succ","name":servicename,"type":servicetype}
	respose,_:=json.Marshal(res)
	w.Header().Set("Content-Type","application/json")
	w.Write(respose)
}
// curl -X POST -d '{"servicetype":"mytype"}' http://localhost:8080/api/service/myservice/post
// {"name":"myservice","result":"succ","type":"mytype"}
```

## http返回值
```
func MyGetHandler(w http.ResponseWriter,r *http.Request)  {
	var res map[string]string=make(map[string]string)
	var status = http.StatusOK

	vals:=r.URL.Query()
	param,ok:=vals["name"]
	fmt.Println(param)
	if (!ok) {
		res["result"]="fail"
		res["error"]="required parameter name is missing"
		status=http.StatusBadRequest
	}else{
	res["result"]="succ"
	res["name"]=param[0]
	status=http.StatusOK
	}

	response,_:=json.Marshal(res)
	w.WriteHeader(status)
	w.Header().Set("Content-type","application/json")
	w.Write(response)
}
// curl -v -X GET http://localhost:8080/api/service/get?name=myname
// curl -v -X GET http://localhost:8080/api/service/get

```

## 返回error
```
func MyGetHandler(w http.ResponseWriter,r *http.Request)  {
	var res map[string]string=make(map[string]string)
	var status = http.StatusOK

	vals:=r.URL.Query()
	param,ok:=vals["name"]
	fmt.Println(param)
	if (!ok) {
		res["result"]="fail"
		res["error"]="required parameter name is missing"
		status=http.StatusBadRequest
	}else{
	res["result"]="succ"
	res["name"]=param[0]
	status=http.StatusOK
	}

	response,_:=json.Marshal(res)
	w.WriteHeader(status)
	w.Header().Set("Content-type","application/json")
	w.Write(response)
}
```
