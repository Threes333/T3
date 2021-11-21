package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	prime   = make(chan int, 50000) // 存放以判断为质数的数
	primes  = make([]int, 0)        // 存放质数
	toJudge = make(chan int, 50000) // 待判断的数
	wg      sync.WaitGroup          // 用来确定是否判断完带判断的数
	wg2     sync.WaitGroup          // 用来确定是否存放完质数
)

// 将待判断的质数放入toJudge通道中
// n int "判断小于n的质数"
func producer(n int) {
	for i := 2; i <= n; i++ {
		toJudge <- i
	}
	close(toJudge)
}

//将已判断为质数的数存放起来
func storePrime() {
	defer wg2.Done()
	for v := range prime {
		primes = append(primes, v)
	}
}

//从toJudge中获取数字并判断是否为质数
func judgePrime() {
	defer wg.Done()
	for n := range toJudge {
		flag := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			prime <- n
		}
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	st := time.Now()
	go producer(50000)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		wg2.Add(1)
		go judgePrime()
		go storePrime()
	}
	wg.Wait()
	close(prime)
	wg2.Wait()
	fmt.Println(time.Now().Sub(st).Seconds())
}
