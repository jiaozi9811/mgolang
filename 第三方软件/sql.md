# sql

[TOC]

## database/sql

<http://go-database-sql.org/>

## 导入数据库引擎

mysql  github.com/go-sql-driver/mysql  
postgresql github.com/lib/pq  
sqlite3  github.com/mattn/go-sqlite3  

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
```

## 连接数据库

```go
func main() {
    db, err := sql.Open("mysql",
        "user:password@tcp(127.0.0.1:3306)/hello")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}
```

## 读取数据

```go
var (
    id int
    name string
)
rows, err := db.Query("select id, name from users where id = ?", 1)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
    err := rows.Scan(&id, &name)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(id, name)
}
err = rows.Err()
if err != nil {
    log.Fatal(err)
}
```

### Prepare查询

```go
stmt, err := db.Prepare("select id, name from users where id = ?")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()
rows, err := stmt.Query(1)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
    // ...
}
if err = rows.Err(); err != nil {
    log.Fatal(err)
}
```

### 单行查询

```go
var name string
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil {
    log.Fatal(err)
}
fmt.Println(name)
```

```go
stmt, err := db.Prepare("select name from users where id = ?")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()
var name string
err = stmt.QueryRow(1).Scan(&name)
if err != nil {
    log.Fatal(err)
}
fmt.Println(name)
```

## 更新数据

```go
stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
if err != nil {
    log.Fatal(err)
}
res, err := stmt.Exec("Dolly")
if err != nil {
    log.Fatal(err)
}
lastId, err := res.LastInsertId()
if err != nil {
    log.Fatal(err)
}
rowCnt, err := res.RowsAffected()
if err != nil {
    log.Fatal(err)
}
log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
```

```go

```

## 事务

db.Begin()  
Commit()  
Rollback()  

```go
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback()
stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close() // danger!
for i := 0; i < 10; i++ {
    _, err = stmt.Exec(i)
    if err != nil {
        log.Fatal(err)
    }
}
err = tx.Commit()
if err != nil {
    log.Fatal(err)
}
// stmt.Close() runs here!
```

## null 空值

```go
for rows.Next() {
    var s sql.NullString
    err := rows.Scan(&s)
    // check err
    if s.Valid {
       // use s.String
    } else {
       // NULL value
    }
}
```

## 不知表的内容

sql.RawBytes

```go
cols, err := rows.Columns() // Remember to check err afterwards
vals := make([]interface{}, len(cols))
for i, _ := range cols {
    vals[i] = new(sql.RawBytes)
}
for rows.Next() {
    err = rows.Scan(vals...)
    // Now you can check each element of vals for nil-ness,
    // and you can use type introspection and type assertions
    // to fetch the column into a typed variable.
}
```