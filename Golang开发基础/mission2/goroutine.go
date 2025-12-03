package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine
func main() {
	//测试打印奇偶数
	printNum()

	//测试任务调度
	tasks := []Task{
		func() { time.Sleep(300 * time.Millisecond) },
		func() { time.Sleep(200 * time.Millisecond) },
		func() { time.Sleep(500 * time.Millisecond) },
	}
	durations := dispatcher(tasks)
	for i, duration := range durations {
		fmt.Printf("task %d 执行时间(ms) : %d \n", i, duration)
	}
}

// 题目 ：
// 编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printNum() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数", i)
		}
	}()

	wg.Wait()
}

// 题目 ：
// 设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
type Task func()

// 任务调度器，使用多个协程并发执行多个任务并记录每个任务的执行时间
func dispatcher(tasks []Task) []time.Duration {
	var wg sync.WaitGroup
	durations := make([]time.Duration, len(tasks))

	for i, task := range tasks {
		wg.Add(1)

		go func(i int, t Task) {
			defer wg.Done()

			start := time.Now()
			t()
			durations[i] = time.Since(start)
		}(i, task)
	}
	wg.Wait()
	return durations
}
