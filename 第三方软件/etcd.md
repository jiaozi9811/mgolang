# etcd

[TOC]

<https://segmentfault.com/a/1190000008361945>

## 安装

https://github.com/coreos/etcd/releases/

安装 go get github.com/coreos/etcd/etcdctl 也可以直接下载etcd二进制 （包含etcd、etcdctl

版本 v3
https://github.com/etcd-io/etcd/tree/master/clientv3
<go get go.etcd.io/etcd/clientv3>


2379用于客户端通信，2380用于节点通信

## 环境变量

环境变量 export ETCDCTL_API=3

## 使用

存储 etcdctl put sensors "{aa:1, bb: 2}"
获取 etcdctl get sensors
监视 etcdctl watch sensors
获取所有值（或指定前缀 ）etcdctl get --prefix=true ""

etcd提供了leader选举、分布式时钟、分布式锁、持续监控（watch）和集群内各个成员的liveness监控等功能

etcd组成：
- HTTP Server： 用于处理用户发送的API请求以及其它etcd节点的同步与心跳信息请求。
- Store：用于处理etcd支持的各类功能的事务，包括数据索引、节点状态变更、监控与反馈、事件处理与执行等等，是etcd对用户提供的大多数API功能的具体实现。
- Raft：Raft强一致性算法的具体实现，是etcd的核心。
- WAL：Write Ahead Log（预写式日志），是etcd的数据存储方式。除了在内存中存有所有数据的状态以及节点的索引以外，etcd就通过WAL进行持久化存储。WAL中，所有的数据提交前都会事先记录日志。Snapshot是为了防止数据过多而进行的状态快照；Entry表示存储的具体日志内容。

## raft协议

## etcd概念词汇表

- Raft：etcd所采用的保证分布式系统强一致性的算法。
- Node：一个Raft状态机实例。
- Member： 一个etcd实例。它管理着一个Node，并且可以为客户端请求提供服务。
- Cluster：由多个Member构成可以协同工作的etcd集群。
- Peer：对同一个etcd集群中另外一个Member的称呼。
- Client： 向etcd集群发送HTTP请求的客户端。
- WAL：预写式日志，etcd用于持久化存储的日志格式。
- snapshot：etcd防止WAL文件过多而设置的快照，存储etcd数据状态。
- Proxy：etcd的一种模式，为etcd集群提供反向代理服务。
- Leader：Raft算法中通过竞选而产生的处理所有数据提交的节点。
- Follower：竞选失败的节点作为Raft中的从属节点，为算法提供强一致性保证。
- Candidate：当Follower超过一定时间接收不到Leader的心跳时转变为Candidate开始竞选。
- Term：某个节点成为Leader到下一次竞选时间，称为一个Term。
- Index：数据项编号。Raft中通过Term和Index来定位数据。


## lease租约

lease其实就是etcd支持申请定时器，比如：可以申请一个TTL=10秒的lease（租约），会返回给你一个lease ID标识定时器。你可以在set一个key的同时携带lease ID，那么就实现了一个自动过期的key。在etcd中，一个lease可以关联给任意多的Key，当lease过期后所有关联的key都将被自动删除

etcdctl lease grant 10

etcdctl put foo1 bar1 --lease=1234abcd
1234abcd是租约id 是指在创建租约10s时返回的id。这个ID然后可以被附加到密钥

## Compact数据整理

## 权限 user role

user 
可以为etcd创建多个用户并设置密码，子命令有：
    add 添加用户
    delete 删除用户
    get 取得用户详情
    list 列出所有用户
    passwd 修改用户密码
    grant-role 给用户分配角色
    revoke-role 给用户移除角色

role 
可以为etcd创建多个角色并设置权限，子命令有：
    add 添加角色
    delete 删除角色
    get 取得角色信息
    list 列出所有角色
    grant-permission 为角色设置某个key的权限
    revoke-permission 为角色移除某个key的权限

auth 
开启/关闭权限控制

```shell
root用户存在时才能开启权限控制
ming@ming:/tmp$ etcdctl auth enable
Error:  etcdserver: root user does not exist
ming@ming:/tmp$ etcdctl user add root
Password of root: 
Type password of root again for confirmation: 
User root created
ming@ming:/tmp$ etcdctl auth enable
Authentication Enabled
 
开启权限控制后需要用--user指定用户
ming@ming:/tmp$ etcdctl user list
Error:  etcdserver: user name not found
ming@ming:/tmp$ etcdctl user list --user=root
Password: 
root
ming@ming:/tmp$ etcdctl user get root --user=root
Password: 
User: root
Roles: root
 
添加用户，前两个密码是新用户的，后一个密码是root的
ming@ming:/tmp$ etcdctl user add mengyuan --user=root
Password of mengyuan: 
Type password of mengyuan again for confirmation: 
Password: 
User mengyuan created
 
使用新用户执行put命令，提示没有权限
ming@ming:/tmp$ etcdctl put key1 v1 --user=mengyuan
Password: 
Error:  etcdserver: permission denied
创建名为rw_key_的role，添加对字符串"key"做为前缀的key的读写权限，为mengyuan添加角色
ming@ming:/tmp$ etcdctl role add rw_key_ --user=root
Password: 
Role rw_key_ created
ming@ming:/tmp$ etcdctl --user=root role grant-permission rw_key_ readwrite key --prefix=true
Password: 
Role rw_key_ updated
ming@ming:/tmp$ etcdctl --user=root user grant-role mengyuan rw_key_
Password: 
Role rw_key_ is granted to user mengyuan
 
添加权限成功后执行put key1成功，执行put k1失败（因为上面只给前缀为"key"的key添加了权限）
ming@ming:/tmp$ etcdctl put key1 v1 --user=mengyuan
Password: 
OK
ming@ming:/tmp$ etcdctl put k1 v1 --user=mengyuan
Password: 
Error:  etcdserver: permission denied
 
执行user list命令失败，没有权限
ming@ming:/tmp$ etcdctl user list --user=mengyuan
Password: 
Error:  etcdserver: permission denied
为新用户添加root的角色后就能执行user list命令了，注意命令中第一个root是角色，第二个root是用户
ming@ming:/tmp$ etcdctl user grant-role mengyuan root --user=root
Password: 
Role root is granted to user mengyuan
ming@ming:/tmp$ etcdctl user list --user=mengyuan
Password: 
mengyuan
root
```


## 集群
etcd有三种集群化启动的配置方案：
    静态配置启动
    etcd自身服务发现
    通过DNS进行服务发现