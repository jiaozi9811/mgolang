# 获取

## 获取硬件信息

```go
	ips,_:=net.Interfaces()//获取硬件的网卡信息
		//fmt.Printf("%#v",ips)
	for k,v:=range ips{
		fmt.Println(k,v)
	}
	netIp,_:=net.InterfaceAddrs()//获取硬件的ip信息
    fmt.Println(netIp)
```