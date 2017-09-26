package main

import (
	"sync"
	"runtime"
	"fmt"
	"runtime/trace"
	"os"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()


	// Your program here
	wg.Add(2)
	go incCounter1(1)
	go incCounter1(2)
	wg.Wait()
	fmt.Printf("Final Counter:%d \n", counter)
}

func incCounter1(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
