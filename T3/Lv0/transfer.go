package main

import (
	"fmt"
	"sync"
)

var (
	myres  = make(map[int]int, 20)
	myChan = make(chan [2]int, 20) // 用来存放计算的阶乘值
	sign   = make(chan struct{})   // 标志是否储存结果完毕
	wg     sync.WaitGroup          //用来确定是否计算完毕
)

func factorial(n int) {
	defer wg.Done()
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myChan <- [2]int{n, res}
}

// 将计算的阶乘值从myChan中读出并储存
func store() {
	for v := range myChan {
		myres[v[0]] = v[1]
	}
	sign <- struct{}{}
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
		wg.Add(1)
	}
	go store()
	wg.Wait()
	close(myChan)
	<-sign
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}
