// 通道
// 通过无缓冲通道，使实现同步等待，取代定时器
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan struct{}) //创建通道不知道放什么东西就放结构体
	go func(chan struct{}) {
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		fmt.Println(sum)
		time.Sleep(5 * time.Second) //方便观察堵塞
		c <- struct{}{}             //写入通道，告知在堵塞的接收方
	}(c)

	fmt.Println("Goroutine的数目：", runtime.NumGoroutine())

	fmt.Println("开始等待,堵塞中")
	//读通道c，通过通道进行同步等待，此时堵塞
	<-c
	fmt.Println("通道接收到数据，堵塞结束")
	fmt.Println("Goroutine的数目：", runtime.NumGoroutine())

}
