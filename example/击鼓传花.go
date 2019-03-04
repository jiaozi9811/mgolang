/*
go版的击鼓传花（erlang和stackless的经典例子）
由n个节点组成一个环状网络，在上面传送共m个消息。
将每个消息（共m个），逐个发送给1号节点。
第1到n-1号节点在接收到消息后，都转发给下一号节点。
第n号节点每次收到消息后，不再继续转发。
当m个消息都从1号逐个到达第n号节点时，认为全部处理结束。
每次执行时设定n=300，m=10000
*/
package main

import (
	"flag"
	"fmt"
	//"os"
	"strconv"
)

var n, m int //定义两个全局变量

//3.定义节点(协程)
func node(i int, ch []chan int, result chan int) {
	for {
		msg := <-ch[i]
		//fmt.Println("node ", i," got msg ",msg)
		if i >= n-1 {
			//fmt.Println("msg ", msg," reached last node ",i)
			if msg >= m-1 {
				fmt.Println("final msg send back")
				result <- msg
			}
		} else {
			//fmt.Println("node ", i," pass msg  ",msg," to next node")
			ch[i+1] <- msg
		}
	}
}

func main() {
	//1.输入参数处理
	flag.Parse()
	args := flag.Args()
	if args != nil && len(args) > 0 {
		var err error
		n, err = strconv.Atoi(args[0])
		if err != nil {
			n = 300
		}
		m, err = strconv.Atoi(args[1])
		if err != nil {
			m = 10000
		}
	} else {
		n = 300
		m = 10000
	}
	fmt.Println("n=", n, "m=", m)

	//2.创建队列/channel
	//用于结束.go没有join之类的等待同步函数
	result := make(chan int)
	chs := make([]chan int, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		//5.启动节点(创建协程)
		go node(i, chs, result)
	}

	//4.初始化消息
	//go的channel写也是阻塞的
	for i := 0; i < m; i++ {
		//fmt.Println("put msg ", i," into  channel 0")
		chs[0] <- i
	}

	//等待结束
	<-result
}

//测试结果

/*
192.168.6.150
-bash-3.2$ go build ring.go
-bash-3.2$ time ./ring 3000 100
n= 3000 m= 100
final msg send back
real    0m0.356s
user    0m0.229s
sys     0m0.020s
-bash-3.2$ time go run ./ring.go 3000 100
n= 3000 m= 100
final msg send back
real    0m0.919s
user    0m0.631s
sys     0m0.079s
-bash-3.2$
-bash-3.2$ time ./ring 10000 100
n= 10000 m= 100
final msg send back
real    0m1.113s
user    0m0.774s
sys     0m0.038s
-bash-3.2$ time go run ./ring.go 10000 100
n= 10000 m= 100
final msg send back
real    0m1.747s
user    0m1.276s
sys     0m0.093s
-bash-3.2$ time ./ring 100 10000
n= 100 m= 10000
final msg send back
real    0m0.692s
user    0m0.493s
sys     0m0.004s
-bash-3.2$
-bash-3.2$ time go run ./ring.go 100 10000
n= 100 m= 10000
final msg send back
real    0m1.329s
user    0m0.903s
sys     0m0.074s
-bash-3.2$
-bash-3.2$ time go run ./ring.go 300 10000
n= 300 m= 10000
final msg send back
real    0m2.628s
user    0m1.924s
sys     0m0.073s
-bash-3.2$ time ./ring 300 10000
n= 300 m= 10000
final msg send back
real    0m1.922s
user    0m1.505s
sys     0m0.018s
*/
