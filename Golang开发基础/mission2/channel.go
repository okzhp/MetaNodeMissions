package main

import (
	"fmt"
	"sync"
)

func main() {
	channelWithoutBuffer()
	fmt.Println("==============================")
	channelWithBuffer()
}

// 题目 ：
// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func channelWithoutBuffer() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(ch chan<- int) {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			fmt.Println("向ch发送数据:", i)
			ch <- i
		}
		wg.Done()
	}(ch)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println("从ch接收到数据:", i)
		}
		wg.Done()
	}(ch)

	wg.Wait()

}

// 题目 ：
// 实现一个带有缓冲的通道，
// 生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
func channelWithBuffer() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(ch chan<- int) {
		defer close(ch)
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			fmt.Println("生产者协程发送:", i)
			ch <- i
		}
	}(ch)

	go func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println("消费者协程接收到:", i)
		}
	}(ch)

	wg.Wait()
}
