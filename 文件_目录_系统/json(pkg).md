# encoding/json

[TOC]

## 解析

>func Marshal(v interface{}) ([]byte, error)

## 反解析

>func Unmarshal(data []byte, v interface{}) error

## 结构体->json

```go
type User struct {

    UserName string 
    PassWord string 
}
    user := User{
        UserName: "tyming",
        PassWord: "1234567890",//这里逗号不能少
    }
    //json.Marshal() 编写为json格式
    str, _ := json.Marshal(user)
    fmt.Printf("%s\n",str)

    //有层次打印json,后两个参数分别为每一行的前缀,每一层的缩进
    str0, _ := json.MarshalIndent(user,""," ")
    fmt.Printf("%s\n",str0)
/*
{"UserName":"tyming","PassWord":"1234567890"}
{
    "UserName": "tyming",
    "PassWord": "1234567890"
}
*/
```

## json->结构体

//匹配的字段名忽略大小写
jsonStr :=`{
        "username":"tyming",
        "password":"1234567890"
    }`

    var u User
    //json.Unmarshal()解码json格式,参数1:json文本(比特序列),参数2:目标容器
    json.Unmarshal([]byte(jsonStr),&u)
    fmt.Println(u)
//{tyming 1234567890}

## 标签的使用

```go
type User struct {
    //`json:"-"` // 直接忽略字段
    //`json:"msg_name"` // 对应json字段名
    //`json:"msg_name,omitempty"` // 如果为空置则忽略字段
    UserName string `json:"username"`
    PassWord string `json:"password"`
}
func main(){

    user := User{
        UserName: "tyming",
        PassWord: "1234567890",//这里逗号不能少
    }
    //json.Marshal() 编写为json格式
    str, _ := json.Marshal(user)
    fmt.Printf("%s\n",str)

    //结构体加标签
    jsonStr :=`{
        "username":"tyming",
        "password":"1234567890"
    }`

    var u User
    //解码为结构体
    json.Unmarshal([]byte(jsonStr),&u)
    fmt.Println(u)
}
```