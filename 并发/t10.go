//每个请求一个routine
//计算100个自然数的和为例，将计算拆成多个task，每个task启动一个goroutine进行处理

//通道方向：使用通道作为参数，可以根据数据流的方向，只允许发送chan<-或者接收<-chan作为参数。

package main

import (
	"fmt"
	"sync"
)

// 工作任务
// 给每个任务都添加结果通道
type task struct {
	begin  int
	end    int
	result chan<- int //只能发送的通道
}

func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	//使用缓冲通道，将数据都存到缓存中，最后统一取出来
	t.result <- sum //向通道发送数据
}

func main() {
	//创建管道任务，数量10
	taskchan := make(chan task, 10)

	//创建结果管道，数量10
	resultchan := make(chan int, 10)

	//wait用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	//初始化task的goroutine，计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)

	//每个task启动一个goroutine进行处理
	go DistributeTask(taskchan, wait, resultchan)

	//结果汇总
	sum := ProcessResult(resultchan)
	fmt.Println(sum)
}

// 初始化任务
// 计算每个任务的begin和end
func InitTask(taskchan chan<- task, r chan int, p int) { //chan<-为只写通道，创建任务
	qu := p / 10  //每个任务的数据长度
	mod := p % 10 //余数是否为0
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk //向任务通道写入任务
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk //多出来的没有被整除的
	}

	//关闭任务通道，此时缓存仍在
	close(taskchan)
}

// 分配任务
func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, result chan int) { //只读通道，任务处理
	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait) //处理任务
	}
	wait.Wait()
	//等待任务完成后关闭结果通道，此时缓存仍在
	close(result)
}

func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// 取出缓存通道中的结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}
