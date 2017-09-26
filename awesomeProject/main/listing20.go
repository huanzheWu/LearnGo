package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var wg sync.WaitGroup

func player(name string, court chan int) {
	defer wg.Done()
	for {
		//如果通道关闭,那么选手胜利
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)

		//随机概率使某个选手Miss
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//关闭通道
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		//否则选手进行击球
		court <- ball
	}
}


func main() {
	rand.Seed(time.Now().Unix())
	court := make(chan int)
	//等待两个goroutine都执行完
	wg.Add(2)
	//选手1等待接球
	go player("candy", court)
	//选手2等待接球
	go player("luffic", court)
	//球进入球场（可以开始比赛了）
	court <- 1
	wg.Wait()
}
