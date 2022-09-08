//缓冲通道

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//通道初始化，没有初始化的通道写入或读取数据会导致当前goroutine永久堵塞
	c := make(chan struct{})  //无缓冲通道
	ci := make(chan int, 100) //缓冲通道，容量100
	go func(chan struct{}, chan int) {
		for i := 0; i < 10; i++ {
			ci <- i //写通道ci，向缓冲区已满的通道写数据会导致goroutine堵塞
		}
		close(ci) //关闭通道ci
		time.Sleep(2 * time.Second)
		c <- struct{}{}

	}(c, ci)
	fmt.Println("Goroutine的数目：", runtime.NumGoroutine())
	//读通道c，通过通道同步等待
	fmt.Println("读通道，堵塞中")
	<-c

	//ci已经关闭，因为匿名函数启动的goroutine已经退出
	fmt.Println("Goroutine的数目：", runtime.NumGoroutine())

	//ci已经关闭，但是仍然能读取数据
	for v := range ci {
		fmt.Println(v)
	}
	fmt.Println(<-ci) //读取已经关闭的通道，不会堵塞，立即返回0值
	//可以使用comma,ok语法判断通道是否安全关闭
	value, ok := <-ci
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("通道已经关闭")
	}

}
