# go-mgo

[TOC]

<http://labix.org/mgo>
<https://github.com/go-mgo/mgo>
<https://gopkg.in/mgo.v2>

import "gopkg.in/mgo.v2"

[mgo](./code/mgo.go)

## 连接数据库  
session, err :=mgo.Dial("mongodb://127.0.0.1:27017")  连接数据库
c := session.DB("test").C("people") 进入集合
err=c.Insert()  插入文档


## 查找文档
c.Find(nil).All(&result)
c.Find(bson.M{"name": "Ale"}).One(&result)

## 修改

func (c *Collection) Update(selector interface{}, change interface{}) error

## 删除

c.Remove(bson.M{"name": "Jimmy Kuu"})