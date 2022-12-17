package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
Go语言中的并发程序主要是通过基于CSP（communicating sequential processes）的goroutine和channel来实现，
当然也支持使用传统的多线程共享内存的并发方式
*/

//func main() {
//	// 直接交由goroutine管理
//	//go hello()
//	//fmt.Println("i am main..")
//	//time.Sleep(time.Second)
//	sample1()
//}

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

func sample2() {

	var wg sync.WaitGroup

	ans := int64(0)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go newGoRoutine(wg, &ans)
	}
	wg.Wait()
}

func newGoRoutine(wg sync.WaitGroup, i *int64) {
	defer wg.Done()
	atomic.AddInt64(i, 1)
	return
}

// --channel..
// reference: https://www.jb51.net/article/228730.htm
func channelSample() {
	fmt.Println("this is a channel test!")
	// 进程→ 线程→协程
	// Go采用通道进行通信，而非共享内存--（对OS底层通信的封装选择）
	// 声明一个接受、发送 int类型的通道，nil
	var channel1 chan int
	fmt.Println(channel1)

	/*
		它是一个指针内存地址。通道变量默认是一个指针。多数情况下，当你想要和一个协程沟通的时候，
		你可以给函数或者方法 <传递一个通道> 作为参数。当从协程接收到通道参数后，你  不需要再对其进行<解引用> 就可以从通道接收或者发送数据。
	*/
	channel2 := make(chan int) //0xc000070300
	fmt.Println(channel2)

	//channel2 <- 2
	//fmt.Println(channel2) //fatal error: all goroutines are asleep - deadlock!
	// // 此时主线程将阻塞直到有协程接收这个数据. Go 的调度器开始调度协程接收通道 channel 的数据
	//    // 但是由于没有协程接受，没有协程是可被调度的。所有协程都进入休眠状态，即是主程序阻塞了

	// channel读写
	channel3 := make(chan string)
	go callChannelPrint(channel3)
	// 阻塞
	channel3 <- "channel3"
	close(channel3) //关闭通道

	fmt.Println("Stop!")

	channel01 := make(chan string)
	channel02 := make(chan string)
	go service1(channel01)
	go service1(channel02)
	select {
	case res := <-channel01:
		fmt.Println("Response form service 1", res)
	case res := <-channel02:
		fmt.Println("Response form service 2", res)
	}
}

func service1(channel chan string) {
	fmt.Printf("hello service1\n")
	// 处理服务1
	channel <- "service1 handle success!\n"
}

func service2(channel chan string) {
	fmt.Printf("hello service2\n")
	channel <- "service2 handle success!\n"
}

func callChannelPrint(channel chan string) {

	fmt.Println("Hello channel string " + <-channel)
}
