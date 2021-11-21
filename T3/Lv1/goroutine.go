package main

import (
	"fmt"
	"sync"
)

var (
	signAB = make(chan struct{}) // A协程传递信号给B协程的通道
	signBC = make(chan struct{}) // B协程传递信号给C协程的通道
	signCA = make(chan struct{}) // C协程传递信号给A协程的通道
	wg     sync.WaitGroup
)

// @title 打印字符
// @description 可以按照顺序打印多个字符多次
// @param id byte "要打印的字符" times int "要打印的次数(若需多次调用则times应相同)", out <-chan struct{}, in chan<-struct{} "接收打印信号以及发送打印完成的信号"
func myPrint(id byte, times int, out <-chan struct{}, in chan<- struct{}) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		<-out // 接受可以进行打印的信号
		fmt.Printf("%c", id)
		if i != times-1 { // 如果不是打印次数的最后一次就继续传递信号
			in <- struct{}{}
		} else { // 是打印次数的最后一次则关闭通道
			close(in)
		}
	}
}

func main() {
	wg.Add(3)
	go myPrint('A', 10, signCA, signAB)
	go myPrint('B', 10, signAB, signBC)
	go myPrint('C', 10, signBC, signCA)
	signCA <- struct{}{} // 开始打印信号
	wg.Wait()
}
