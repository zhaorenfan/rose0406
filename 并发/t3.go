package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前的GOMAXPROCS值
	fmt.Println("Go可以并发的数目：", runtime.GOMAXPROCS(0))

	//设置GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)

	//获取当前的GOMAXPROCS值
	fmt.Println("Go可以并发的数目：", runtime.GOMAXPROCS(0))

}
