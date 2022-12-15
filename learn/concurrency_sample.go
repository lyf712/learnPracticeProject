package main

import (
	"fmt"
	"sync"
)

/**
Go语言中的并发程序主要是通过基于CSP（communicating sequential processes）的goroutine和channel来实现，
当然也支持使用传统的多线程共享内存的并发方式
*/

func main() {
	// 直接交由goroutine管理
	//go hello()
	//fmt.Println("i am main..")
	//time.Sleep(time.Second)
	sample1()
}

func hello(val int) {
	fmt.Printf("hello %d \n", val)
	wg.Done()
}

// java 中 CountLatch
var wg sync.WaitGroup

func sample1() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

// --channel..
