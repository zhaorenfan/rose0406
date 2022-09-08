package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println(sum)
		time.Sleep(1 * time.Second)
	}()
	fmt.Println("goroutine的数量：", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
}
