## etcd的go操作

## v3

版本 v3
https://github.com/etcd-io/etcd/tree/master/clientv3
<go get go.etcd.io/etcd/clientv3>

<https://yuerblog.cc/2017/12/12/etcd-v3-sdk-usage/>


[put and get](./code/etcd_put_get.go)

## 连接

```go
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
    }
    defer cli.Close()
```

Endpoints：etcd的多个节点服务地址
DialTimeout：创建client的首次连接超时，这里传了5秒，如果5秒都没有连接成功就会返回err；值得注意的是，一旦client创建成功，就不用再关心后续底层连接的状态了，client内部会重连


## type KV interface

type KV interface {
	Put(ctx context.Context, key, val string, opts ...OpOption) (*PutResponse, error)
	Get(ctx context.Context, key string, opts ...OpOption) (*GetResponse, error)
	Delete(ctx context.Context, key string, opts ...OpOption) (*DeleteResponse, error)
	Compact(ctx context.Context, rev int64, opts ...CompactOption) (*CompactResponse, error)
	Do(ctx context.Context, op Op) (OpResponse, error)
	Txn(ctx context.Context) Txn
}

### put

```go
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(resp)
```

## type Lease interface

type Lease interface {
	Grant(ctx context.Context, ttl int64) (*LeaseGrantResponse, error)
	Revoke(ctx context.Context, id LeaseID) (*LeaseRevokeResponse, error)
	TimeToLive(ctx context.Context, id LeaseID, opts ...LeaseOption) (*LeaseTimeToLiveResponse, error)
	Leases(ctx context.Context) (*LeaseLeasesResponse, error)
	KeepAlive(ctx context.Context, id LeaseID) (<-chan *LeaseKeepAliveResponse, error)
	KeepAliveOnce(ctx context.Context, id LeaseID) (*LeaseKeepAliveResponse, error)
	Close() error
}

- Grant：分配一个租约。
- Revoke：释放一个租约。
- TimeToLive：获取剩余TTL时间。
- Leases：列举所有etcd中的租约。
- KeepAlive：自动定时的续约某个租约。
- KeepAliveOnce：为某个租约续约一次。
- Close：貌似是关闭当前客户端建立的所有租约。


## Op

Op字面意思就是”操作”，Get和Put都属于Op，只是为了简化用户开发而开放的特殊API
func OpDelete(key string, opts …OpOption) Op
func OpGet(key string, opts …OpOption) Op
func OpPut(key, val string, opts …OpOption) Op
func OpTxn(cmps []Cmp, thenOps []Op, elseOps []Op) Op

## Txn事务
etcd中事务是原子执行的，只支持if … then … else …这种表达，能实现一些有意思的场景

type Txn interface {
	If(cs ...Cmp) Txn
	Then(ops ...Op) Txn
	Else(ops ...Op) Txn
	Commit() (*TxnResponse, error)
}

首先，需要开启一个事务，这是通过KV对象的方法实现的：

txn := kv.Txn(context.TODO())

txnResp, err := txn.If(clientv3.Compare(clientv3.Value("/hi"), "=", "hello")).
    Then(clientv3.OpGet("/hi")).
    Else(clientv3.OpGet("/test/", clientv3.WithPrefix())).
    Commit()


## type Auth

type Auth interface {
    AuthEnable(ctx context.Context) (*AuthEnableResponse, error)
    AuthDisable(ctx context.Context) (*AuthDisableResponse, error)
    UserAdd(ctx context.Context, name string, password string) (*AuthUserAddResponse, error)
    UserDelete(ctx context.Context, name string) (*AuthUserDeleteResponse, error)
    UserChangePassword(ctx context.Context, name string, password string) (*AuthUserChangePasswordResponse, error)
    UserGrantRole(ctx context.Context, user string, role string) (*AuthUserGrantRoleResponse, error)
    UserGet(ctx context.Context, name string) (*AuthUserGetResponse, error)
    UserList(ctx context.Context) (*AuthUserListResponse, error)
    UserRevokeRole(ctx context.Context, name string, role string) (*AuthUserRevokeRoleResponse, error)
    RoleAdd(ctx context.Context, name string) (*AuthRoleAddResponse, error)
    RoleGrantPermission(ctx context.Context, name string, key, rangeEnd string, permType PermissionType) (*AuthRoleGrantPermissionResponse, error)
    RoleGet(ctx context.Context, role string) (*AuthRoleGetResponse, error)
    RoleList(ctx context.Context) (*AuthRoleListResponse, error)
    RoleRevokePermission(ctx context.Context, role string, key, rangeEnd string) (*AuthRoleRevokePermissionResponse, error)
    RoleDelete(ctx context.Context, role string) (*AuthRoleDeleteResponse, error)
}

## type Cluster
type Cluster interface {
    MemberList(ctx context.Context) (*MemberListResponse, error)
    MemberAdd(ctx context.Context, peerAddrs []string) (*MemberAddResponse, error)
    MemberRemove(ctx context.Context, id uint64) (*MemberRemoveResponse, error)
    MemberUpdate(ctx context.Context, id uint64, peerAddrs []string) (*MemberUpdateResponse, error)
}

## type Watcher

type Watcher interface {
    Watch(ctx context.Context, key string, opts ...OpOption) WatchChan
    RequestProgress(ctx context.Context) error
    Close() error
}