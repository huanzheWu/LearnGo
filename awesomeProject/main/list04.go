package main

import (
	"sync"
	"runtime"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed prefix")
}

func main() {
	begTime := time.Nanosecond
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
	defer func(){
		endTime := time.Nanosecond
		fmt.Println("cost time :%d",endTime-begTime)
	}()
}
