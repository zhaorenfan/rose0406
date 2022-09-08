//select多路复用，用于多路监听多个通道
//当监听的通道没有状态是可读或可写的，select是阻塞的
//只要监听的通道中有一个状态是可读或可写的，则select就不会阻塞，
//而是进入进入处理就绪通道的分支流程
//如果监听的通道有多个可读或可写的状态，则select随机选取一个处理

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func(chan int, chan int) {
		for {
			select { //先阻塞，等待通道状态
			//监听两个通道，ch和ch2，准备就绪等待读通道
			//向通道ch写入的数据是随机的，0或1
			case ch <- 0:
				fmt.Println("向通道ch写入0")
			case ch <- 1:
				fmt.Println("向通道ch写入1")
			//向通道ch写入的数据是随机的，2或3
			case ch2 <- 2:
			case ch2 <- 3:
			}
		}
	}(ch, ch2)

	time.Sleep(2 * time.Second)
	//读取通道
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch2)
	}
}
