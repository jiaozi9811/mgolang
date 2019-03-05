package main

import (
	"fmt"
	"time"
)

func main() {
	tab := []int{1, 3, 0, 5}

	ch := make(chan int)
	for _, value := range tab {
		go func(val int) {
			time.Sleep(time.Duration(val) * 10000000)
			fmt.Println(val)
			ch <- val
		}(value)
	}

	for _ = range tab {
		<-ch
	}
}

/*
睡眠排序通过为待排序的元素启动独立的任务，每个任务按照待排元素的key执行相对应的睡眠时间，然后及时的将序列中的元素收集到一起，达到排序的目的
*/