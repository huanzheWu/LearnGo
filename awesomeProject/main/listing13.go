package main

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go IncCounter(1)
	go IncCounter(2)
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func IncCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
