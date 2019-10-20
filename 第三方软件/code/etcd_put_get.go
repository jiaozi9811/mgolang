package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.0.0.19:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

    // Put
	resp,err:=client.Put(context.TODO(),"foo5","bar5")
	if err!=nil{
		fmt.Println(err)
    }
    fmt.Println(resp)

    // Get
	getpesp,err:=client.Get(context.TODO(),"foo5")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(getpesp.Kvs[0].Value))
	//for _,val:=range getpesp.Kvs{
	//	fmt.Println(string(val.Key))
	//}
}
