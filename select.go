package main

import "fmt"
import "time"

func main() {
    ch := make(chan bool)

    // 执行一个耗时任务
    go func () {
        fmt.Println("long time mission")
        time.Sleep(time.Duration(3) * time.Second)
        ch <- true
    }()

    // select 中只要有一个 case 能执行，则立即执行
    // 同一时刻有多个 case 可以执行，则随机执行一个
    // 如果 case 都不可执行，则执行 default
    select {
        // 耗时任务执行完毕
        case <- ch:
            fmt.Println("ok")
        // 5s 后超时
        case <- time.After(time.Duration(5) * time.Second):
            fmt.Println("time out")
        // 当其他 case 都不能执行时, 执行 default
        default:
            fmt.Println("default")
    }
}
