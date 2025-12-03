package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

func main() {
  lockAdd1000()
  fmt.Println("===============================")
  atomicAdd1000()
}

// 题目 ：
// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func lockAdd1000() {
  var counter int
  var mx sync.Mutex
  var wg sync.WaitGroup

  for i := 0; i < 10; i++ {
    wg.Add(1)

    go func() {
      defer wg.Done()
      for i := 0; i < 1000; i++ {
        mx.Lock()
        counter++
        mx.Unlock()
      }
    }()
  }
  wg.Wait()

  fmt.Println("final counter:", counter)

}

// 题目 ：
// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func atomicAdd1000() {
  var counter int32
  var wg sync.WaitGroup

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for i := 0; i < 1000; i++ {
        atomic.AddInt32(&counter, 1)
      }
    }()

    wg.Wait()
  }

  fmt.Println("final counter:", counter)
}
