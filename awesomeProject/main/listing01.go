package main

import (
	"runtime"
	"sync"
	"fmt"
)

var (
	wg sync.WaitGroup
)

func main() {
	//分配一个逻辑处理器Ｐ给调度器使用
	runtime.GOMAXPROCS(1)
	//在这里,wg用于等待程序完成，计数器加2，表示要等待两个goroutine
	wg.Add(2)
	//声明1个匿名函数，并创建一个goroutine
	fmt.Printf("Begin Coroutines\n")
	go func() {
		//在函数退出时，wg计数器减1
		defer wg.Done()
		//打印3次小写字母表
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				if char=='k'{
					runtime.Gosched()
				}
				fmt.Printf("%c ", char)
			}
		}
	}()
	//声明1个匿名函数，并创建一个goroutine
	go func() {
		defer wg.Done()
		//打印大写字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				if char == 'K'{
					runtime.Gosched()
				}
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Printf("Waiting To Finish\n")
	//等待2个goroutine执行完毕
	wg.Wait()

}
