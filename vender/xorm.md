# xorm

[TOC]

https://gitea.com/xorm/manual-zh-CN

## 创建xorm engine引擎

```
import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
    var err error
    engine, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
}
```
or
```
import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
    var err error
    engine, err = xorm.NewEngine("sqlite3", "./test.db")
}
```

engine可以通过engine.Close来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭


## 创建xorm engine group引擎组
```
import (
    _ "github.com/lib/pq"
    "github.com/xormplus/xorm"
)

var eg *xorm.EngineGroup

func main() {
	conns := []string{
		"postgres://postgres:root@localhost:5432/test?sslmode=disable;", // 第一个默认是master
		"postgres://postgres:root@localhost:5432/test1?sslmode=disable;", // 第二个开始都是slave
		"postgres://postgres:root@localhost:5432/test2?sslmode=disable",
	}
    
    var err error
	eg, err = xorm.NewEngineGroup("postgres", conns)
}
```
