//WaitGroup

package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com/",
	"http://www.baidu.com/",
	"http://www.qq.com/",
}

func main() {

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			//defer wg.Done()
			defer wg.Add(-1)
			resp, err := http.Get(u)
			if err == nil {
				fmt.Println(resp.Status)
				//fmt.Println(resp.Header)
			}
		}(url)
	}
	fmt.Println("Goroutine的数目：", runtime.NumGoroutine()) //4
	wg.Wait()
	fmt.Println("Goroutine的数目：", runtime.NumGoroutine()) //9
}
