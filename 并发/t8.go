//select的通知退出机制
//关闭select中监听的某个通道，能够使select立即感知，然后进行处理

//例子：随机数生成器的通知退出机制，下游的消费者不需要随机数时，显式地通知生产者停止生产

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)

	go func() {
	Lable:
		for {
			select {
			//监听哦通道ch
			case ch <- rand.Int():
			//增加一路监听，对通知退出信号done进行监听
			case <-done:
				break Lable //跳出for循环
			}
		}
		//收到通知后关闭通道ch
		close(ch)
	}()

	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println("goroutine的数量：", runtime.NumGoroutine())

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//发送通知，告知生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("goroutine的数量：", runtime.NumGoroutine())

}
