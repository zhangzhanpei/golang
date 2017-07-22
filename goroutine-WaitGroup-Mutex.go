package main

import (
    "fmt"
    "runtime"
    "sync"
)

var sum = 0

func add(wg *sync.WaitGroup, mutex *sync.Mutex) {
    for i := 0; i < 100000; i++ {
        // mutex.Lock() //加锁
        sum = sum + 1 //如果没有对 sum + 1 操作加锁, 会导致多个 goroutine 操作 sum 时没取到最新值, 最后的结果小于 500000
        // mutex.Unlock() //释放锁
    }
    wg.Done() //调用 Done 方法使 WaitGroup 的 Counter 减一, 即通知 WaitGroup 当前 goroutine 执行完
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU()) //设置 goroutine 在多个 cpu 核心中运行

    wg := sync.WaitGroup{} //WaitGroup
    wg.Add(5) //设置 WaitGroup 的 Counter 为 5, 即需要等待完成的 goroutine 数量
    mutex := sync.Mutex{} //互斥锁

    for i := 0; i < 5; i++ { //起 5 个 goroutine
        go add(&wg, &mutex)
    }

    wg.Wait() //在这里阻塞, 直到 WaitGroup 的 Counter 为 0, 即等待所有 goroutine 执行完
    fmt.Println(sum)
}
