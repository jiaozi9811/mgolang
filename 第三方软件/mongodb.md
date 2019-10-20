# mongodb

[TOC]

## Mongo概念

SQL术语/概念|MongoDB术语/概念|解释/说明
-|-|-
database|database|数据库
table|collection|数据库表/集合
row|document|数据记录行/文档
column|field|数据字段/域
index|index|索引
table joins||表连接,MongoDB不支持
primary key|primary key|主键,MongoDB自动将_id字段设置为主键

### 文档(Document)
RDBMS|MongoDB
-|-
数据库|数据库
表格|集合
行|文档
列|字段
表联合|嵌入文档
主键|主键(MongoDB 提供了 key 为 _id)

## 元数据
在MongoDB数据库中名字空间 <dbname>.system.* 是包含多种系统信息的特殊集合(Collection)，如下:

集合命名空间|描述
-|-
dbname.system.namespaces|列出所有名字空间。
dbname.system.indexes|列出所有索引。
dbname.system.profile|包含数据库概要(profile)信息。
dbname.system.users|列出所有可访问数据库的用户。
dbname.local.sources|包含复制对端（slave）的服务器信息和状态。

## 数据类型
数据类型|描述
-|-
String|字符串。存储数据常用的数据类型。在 MongoDB 中，UTF-8 编码的字符串才是合法的。
Integer|整型数值。用于存储数值。根据你所采用的服务器，可分为 32 位或 64 位。
Boolean|布尔值。用于存储布尔值（真/假）。
Double|双精度浮点值。用于存储浮点值。
Min/Max keys|将一个值与 BSON（二进制的 JSON）元素的最低值和最高值相对比。
Array|用于将数组或列表或多个值存储为一个键。
Timestamp|时间戳。记录文档修改或添加的具体时间。
Object|用于内嵌文档。
Null|用于创建空值。
Symbol|  符号。该数据类型基本上等同于字符串类型，但不同的是，它一般用于采用特殊符号类型的语言。
Date|日期时间。用 UNIX 时间格式来存储当前日期或时间。你可以指定自己的日期时间：创建 Date|对象，传入年月日信息。
Object ID|对象 ID。用于创建文档的 ID。
Binary Data|二进制数据。用于存储二进制数据。
Code|代码类型。用于在文档中存储 JavaScript 代码。
Regular expression|正则表达式类型。用于存储正则表达

## mongod.conf

```
fork=true
bind_ip=0.0.0.0
port=27017
dbpath=/opt/mongo
logpath=/opt/mongo/log
directoryperdb=true
logappend=true
```

## 保留数据库
- admin： 从权限的角度来看，这是"root"数据库。要是将一个用户添加到这个数据库，这个用户自动继承所有数据库的权限。一些特定的服务器端命令也只能从这个数据库运行，比如列出所有的数据库或者关闭服务器。
- local: 这个数据永远不会被复制，可以用来存储限于本地单台服务器的任意集合
- config: 当Mongo用于分片设置时，config数据库在内部使用，用于保存分片的相关信息。

## 连接数据库
`mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]`

选项(options)

选项  |描述
-|-
replicaSet=name |验证replica set的名称。 Impliesconnect=replicaSet.
slaveOk=true\|false | true:在connect=direct模式下，驱动会连接第一台机器，即使这台服务器不是主。在connect=replicaSet模式下，驱动会发送所有的写请求到主并且把读取操作分布在其他从服务器。                     false: 在 connect=direct模式下，驱动会自动找寻主服务器. 在connect=replicaSet 模式下，驱动仅仅连接主服务器，并且所有的读写命令都连接到主服务器。
safe=true\|false |true: 在执行更新操作之后，驱动都会发送getLastError命令来确保更新成功。(还要参考 wtimeoutMS).false: 在每次更新之后，驱动不会发送getLastError来确保更新成功。
w=n |驱动添加 \{ w : n \} 到getLastError命令. 应用于safe=true。
wtimeoutMS=ms   |驱动添加 \{ wtimeout : ms \} 到 getlasterror 命令. 应用于 safe=true.
fsync=true\|false    |true: 驱动添加 \{ fsync : true \} 到 getlasterror 命令.应用于 safe=true. false: 驱动不会添加到getLastError命令中。
journal=true\|false  |如果设置为 true, 同步到 journal (在提交到数据库前写入到实体中). 应用于 safe=true
connectTimeoutMS=ms |可以打开连接的时间。
socketTimeoutMS=ms  |发送和接受sockets的时间。

### mongo命令连接
mongo参数  
    --port arg  
    --host arg  
    -u [ --username ] arg  
    -p [ --password ] arg  
    --nodb  

### MongoDB 连接命令格式
$ ./mongo
\> mongodb://admin:123456@localhost/  
mongodb://admin:123456@localhost/test  
mongodb://localhost  
mongodb://fred:foobar@localhost  
mongodb://fred:foobar@localhost/baz  
mongodb://localhost,localhost:27018,localhost:27019  
mongodb://host1,host2,host3/?connect=direct;slaveOk=true   直接连接第一个服务器，无论是replica set一部分或者主服务器或者从服务器  
mongodb://localhost/?safe=true  安全模式连接  
mongodb://host1,host2,host3/?safe=true;w=2;wtimeoutMS=2000  安全模式连接到replica set，并且等待至少两个复制服务器成功写入，超时时间设置为2秒   


## 常用命令

help  
db.help() 
db.mycoll.help()  集合help  
show dbs    查看数据库  
db          显示当前数据库对象  
use [db]    切换数据库  
show collections  
show tables  
db.serverStatus()  查看数据库服务器的状态  
use db && db.stats()  查询指定数据库统计信息  
db.getCollectionNames() 查询数据库集合列表  
db.shutdownServer()  终止服务器进程  

## 基本DDL和DML
### 创建数据库
`use DATABASE_NAME`     如果数据库不存在，则创建数据库，否则切换到指定数据库
### 插入集合
`db.db_name.insert({name:value})`

### 删除数据库
`db.dropDatabase()`     删除当前数据库，默认为 test
### 删除集合
`db.collection.drop()`

### 创建集合
`db.createCollection(name, options)`

`db.coll.insert({})`  在MongoDB中，不需要创建集合。当插入文档时，MongoDB会自动创建集合
options 可以是如下参数：

字段 | 类型 | 描述
-|-|-
capped  |布尔  |（可选）如果为 true，则创建固定集合。固定集合是指有着固定大小的集合，当达到最大值时，它会自动覆盖最早的文档。当该值为 true 时，必须指定 size 参数。
autoIndexId |布尔  |（可选）如为 true，自动在 _id 字段创建索引。默认为 false。
size    |数值  |（可选）为固定集合指定一个最大值（以字节计）。如果 capped 为 true，也需要指定该字段。
max |数值  |（可选）指定固定集合中包含文档的最大

### 删除集合
`db.collection_name.drop()`

### 插入文档
mongodb所有存储在集合中的数据都是BSON格式.BSON是类似JSON存储格式,是Binary JSON的简称  
使用 insert() 或 save() 方法向集合中插入文档  
`db.COLLECTION_NAME.insert(document)`

### 查看文档
`db.coll.find()[.pretty()]`  
`db.coll.find({})`

查询操作  
- 集合查询方法 find()  
- 查询内嵌文档  
- 查询操作符(内含 数组查询)  
** "$gt" 、"$gte"、 "$lt"、 "$lte"、"null查询"、"$all"、"$size"、"$in"、"$nin"、  
** "$and"、"$nor"、"$not"、"$or"、"$exists"、"$mod"、"$regex"、"$where"、"$slice"、"$elemMatch" 

`db.collection.find(query, projection)`  

在第二个参数中，指定键名且值为1或者true则是查询结果中显示的键；若值为0或者false，则为不显示键

mongo sql| sql
-|-
`db.users.find()` |select * from users
`db.users.find({"age" : 27})` |select * from users where age = 27
`db.users.find({"username" : "joe", "age" : 27})`|select * from users where "username" = "joe" and age = 27
`db.users.find({}, {"username" : 1, "email" : 1})` |select username, email from users
`db.users.find({}, {"username" : 1, "_id" : 0})` // no case  // 即时加上了列筛选，_id也会返回；必须显式的阻止_id返回
`db.users.find({"age" : {"$gte" : 18, "$lte" : 30}})` |select * from users where age >=18 and age <= 30 // $lt(<) $lte(<=) $gt(>) $gte(>=)
`db.users.find({"username" : {"$ne" : "joe"}})` |select * from users where username <> "joe"
`db.users.find({"ticket_no" : {"$in" : [725, 542, 390]}})` |select * from users where ticket_no in (725, 542, 390)
`db.users.find({"ticket_no" : {"$nin" : [725, 542, 390]}})` |select * from users where ticket_no not in (725, 542, 390)
`db.users.find({"$or" : [{"ticket_no" : 725}, {"winner" : true}]})` |select * form users where ticket_no = 725 or winner = true
`db.users.find({"id_num" : {"$mod" : [5, 1]}})` |select * from users where (id_num mod 5) = 1
`db.users.find({"$not": {"age" : 27}})`| select * from users where not (age = 27)
`db.users.find({"username" : {"$in" : [null], "$exists" : true}})` |select * from users where username is null // 如果直接通过find({"username" : `null})进行查询，那么连带"没有username"的纪录一并筛选出来
`db.users.find({"name" : /joey?/i})` // 正则查询，value是符合PCRE的表达式

db.food.find({fruit : {$all : ["apple", "banana"]}}) // 对数组的查询, 字段fruit中，既包含"apple",又包含"banana"的纪录  
db.food.find({"fruit.2" : "peach"}) // 对数组的查询, 字段fruit中，第3个(从0开始)元素是peach的纪录  
db.food.find({"fruit" : {"$size" : 3}}) // 对数组的查询, 查询数组元素个数是3的记录，$size前面无法和其他的操作符复合使用  
db.users.findOne(criteria, {"comments" : {"$slice" : 10}}) // 对数组的查询，只返回数组comments中的前十条，还可以{"$slice" : -10}， {"$slice" : [23, 10]}; 分别返回最后10条，和中间10条  
db.people.find({"name.first" : "Joe", "name.last" : "Schmoe"})  // 嵌套查询  
db.blog.find({"comments" : {"$elemMatch" : {"author" : "joe", "score" : {"$gte" : 5}}}}) // 嵌套查询，仅当嵌套的元素是数组时使用,  
db.foo.find({"$where" : "this.x + this.y == 10"}) // 复杂的查询，$where当然是非常方便的，但效率低下。对于复杂查询，考虑的顺序应当是 正则 -> MapReduce -> $where  
db.foo.find({"$where" : "function() { return this.x + this.y == 10; }"}) // $where可以支持javascript函数作为查询条件  
db.foo.find().sort({"x" : 1}).limit(1).skip(10); // 返回第(10, 11]条，按"x"进行排序; 三个limit的顺序是任意的，应该尽量避免skip中使用large-number  

### 更新文档
MongoDB 使用 update() 和 save() 方法来更新集合中的文档
语法格式如下：
```
db.collection.update(
   <query>,
   <update>,
   {
     upsert: <boolean>,
     multi: <boolean>,
     writeConcern: <document>
   }
)
```
参数说明：  
query : update的查询条件，类似sql update查询内where后面的。
update : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
upsert : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
writeConcern :可选，抛出异常的级别。

`db.col.update({'title':'a'},{$set:{'title':'MongoDB'}},[,{multi:true}])`


#### save() 方法  
save() 方法通过传入的文档来替换已有文档。语法格式如下：
```
db.collection.save(
   <document>,
   {
     writeConcern: <document>
   }
)
```
参数说明：  
document : 文档数据。  
writeConcern :可选，抛出异常的级别





更多实例
只更新第一条记录：  
`db.col.update( { "count" : { $gt : 1 } } , { $set : { "test2" : "OK"} } );`  
全部更新：  
`db.col.update( { "count" : { $gt : 3 } } , { $set : { "test2" : "OK"} },false,true );`  
只添加第一条：  
`db.col.update( { "count" : { $gt : 4 } } , { $set : { "test5" : "OK"} },true,false );`  
全部添加进去:  
`db.col.update( { "count" : { $gt : 5 } } , { $set : { "test5" : "OK"} },true,true );`  
全部更新：  
`db.col.update( { "count" : { $gt : 15 } } , { $inc : { "count" : 1} },false,true );`  
只更新第一条记录：  
`db.col.update( { "count" : { $gt : 10 } } , { $inc : { "count" : 1} },false,false );`  

### 删除文档
```
db.collection.remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
```
参数说明：
query :（可选）删除的文档的条件。
justOne : （可选）如果设为 true 或 1，则只删除一个文档，如果不设置该参数，或使用默认值 false，则删除所有匹配条件的文档。
writeConcern :（可选）抛出异常的级别。

删除所有数据
db.col.remove({})

现在官方推荐使用 deleteOne() 和 deleteMany() 方法
如删除集合下全部文档：
`db.inventory.deleteMany({})`
删除 status 等于 A 的全部文档：
`db.inventory.deleteMany({ status : "A" })`
删除 status 等于 D 的一个文档：
`db.inventory.deleteOne( { status: "D" } )`

### 查询文档  
语法   
`db.collection.find(query, projection)`  
* query ：可选，使用查询操作符指定查询条件  
* projection ：可选，使用投影操作符指定返回的键。查询时返回文档中所有键值， 只需省略该参数即可（默认省略）

`db.col.find().pretty()`

#### MongoDB 与 RDBMS Where 语句比较  
操作 | 格式|  范例 | RDBMS中的类似语句
-|-|-|-
等于 | `{<key>:<value>}` |`db.col.find({"by":"菜鸟教程"}).pretty()` |where by = '菜鸟教程'
小于 | `{<key>:{$lt:<value>}}`|   `db.col.find({"likes":{$lt:50}}).pretty()`   | where likes < 50
小于或等于|  ` {<key>:{$lte:<value>}}`| ` db.col.find({"likes":{$lte:50}}).pretty()` |  where likes <= 50
大于 | `{<key>:{$gt:<value>}}` |  `db.col.find({"likes":{$gt:50}}).pretty()`   | where likes > 50
大于或等于 |  `{<key>:{$gte:<value>}}` | `db.col.find({"likes":{$gte:50}}).pretty()`  | where likes >= 50
不等于| `{<key>:{$ne:<value>}}`  | `db.col.find({"likes":{$ne:50}}).pretty()`    |where likes != 50

#### AND 条件  
`db.col.find({key1:value1, key2:value2}).pretty()`

#### OR 条件  
```
db.col.find(
   {
      $or: [
         {key1: value1}, {key2:value2}
      ]
   }
).pretty()
```
#### Limit() 方法
读取指定数量的数据记录  
`db.COLLECTION_NAME.find().limit(NUMBER)`

#### Skip() 方法  
跳过指定数量的数据  
`db.COLLECTION_NAME.find().limit(NUMBER).skip(NUMBER)`

#### 排序 sort() 方法  
sort() 方法可以通过参数指定排序的字段，并使用 1 和 -1 来指定排序的方式，其中 1 为升序排列，而 -1 是用于降序排列  
`db.COLLECTION_NAME.find().sort({KEY:1})`

## 索引

### createIndex() 方法  
语法  
`db.collection.createIndex(keys, options)`  
 Key值为你要创建的索引字段，1为指定按升序创建索引，如果你想按降序来创建索引指定为-1即可

options参数列表  

Parameter  | Type  |  Description
-|-|-
background | Boolean |建索引过程会阻塞其它数据库操作，background可指定以后台方式创建索引，即增加 "background" 可选参数。 "background" 默认值为false。
unique | Boolean |建立的索引是否唯一。指定为true创建唯一索引。默认值为false.
name  |  string | 索引的名称。如果未指定，MongoDB的通过连接索引的字段名和排序顺序生成一个索引名称。
dropDups  |  Boolean |3.0+版本已废弃。在建立唯一索引时是否删除重复记录,指定 true 创建唯一索引。默认值为 false.
sparse | Boolean| 对文档中不存在的字段数据不启用索引；这个参数需要特别注意，如果设置为true的话，在索引字段中不会查询出不包含对应字段的文档.。默认值为 false.
expireAfterSeconds | integer| 指定一个以秒为单位的数值，完成 TTL设定，设定集合的生存时间。
v |  index version|  索引的版本号。默认的索引版本取决于mongod创建索引时运行的版本。
weights| document   | 索引权重值，数值在 1 到 99,999 之间，表示该索引相对于其他索引字段的得分权重。
default_language  |  string | 对于文本索引，该参数决定了停用词及词干和词器的规则的列表。 默认为英语
language_override  | string  |对于文本索引，该参数指定了包含在文档中的字段名，语言覆盖默认的language，默认值为 language.

### getIndexes()

### dropIndex()

### reIndex()

## 聚合  
聚合(aggregate)主要用于处理数据(诸如统计平均值,求和等)，并返回计算后的数据结果。有点类似sql语句中的 count(*)

`db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)`

聚合的表达式:

表达式| 描述 | 实例
-|-|-
$sum  |  计算总和。  | `db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$sum : "$likes"}}}])`
$avg  |  计算平均值 |  `db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$avg : "$likes"}}}])`
$min  |  获取集合中所有文档对应值得最小值。 |  `db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$min : "$likes"}}}])`
$max   | 获取集合中所有文档对应值得最大值。 |  `db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$max : "$likes"}}}])`
$push |  在结果文档中插入值到一个数组中。  |  `db.mycol.aggregate([{$group : {_id : "$by_user", url : {$push: "$url"}}}])`
$addToSet |  在结果文档中插入值到一个数组中，但不创建副本。| `db.mycol.aggregate([{$group : {_id : "$by_user", url : {$addToSet : "$url"}}}])`
$first | 根据资源文档的排序获取第一个文档数据。| `db.mycol.aggregate([{$group : {_id : "$by_user", first_url : {$first : "$url"}}}])`
$last |  根据资源文档的排序获取最后一个文档数据 |`db.mycol.aggregate([{$group : {_id : "$by_user", last_url : {$last : "$url"}}}])`

## 管道

$project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。  
$match：用于过滤数据，只输出符合条件的文档。$match使用MongoDB的标准查询操作。  
$limit：用来限制MongoDB聚合管道返回的文档数。  
$skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。  
$unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。  
$group：将集合中的文档分组，可用于统计结果。  
$sort：将输入文档排序后输出。  
$geoNear：输出接近某一地理位置的有序文档。  


## 复制(副本集)

### 副本集设置

`mongod --port "PORT" --dbpath "YOUR_DB_DATA_PATH" --replSet "REPLICA_SET_INSTANCE_NAME"`

在Mongo客户端使用命令rs.initiate()来启动一个新的副本集  
我们可以使用rs.conf()来查看副本集的配置  
查看副本集状态使用 rs.status() 命令  

### 添加副本集

`rs.add(HOST_NAME:PORT)`


## 分片

![分片集群][1]
[1]: http://www.runoob.com/wp-content/uploads/2013/12/sharding.png

分片主要组件：  
**Shard**:用于存储实际的数据块，实际生产环境中一个shard server角色可由几台机器组个一个replica set承担，防止主机单点故障  
**Config Server**:mongod实例，存储了整个 ClusterMetadata，其中包括 chunk信息  
**Query Routers**:前端路由，客户端由此接入，且让整个集群看上去像单一数据库，前端应用可以透明使用  


## 备份恢复 mongodump mongorestore

`mongodump -h dbhost -d dbname -o dbdirectory`

`mongorestore -h <hostname><:port> -d dbname <path>`


## 监控 mongostat mongotop  
提供了mongostat 和 mongotop 两个命令来监控MongoDB的运行情况

### mongostat

### mongotop

## Rockmongo(MongoDB管理工具)

## 用户安全管理

### 开启安全认证  
使用--auth选项启动mongod进程即可启用认证模式
设置auth=true

### 添加用户

` db.createUser({user: "admin",pwd: "1234",roles: [ "readWrite", "dbAdmin" ]})`




