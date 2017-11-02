package main

import "fmt"
import "time"

func main() {
    ch := make(chan bool)

    //执行一个耗时任务
    go func () {
        fmt.Println("long time mission")
        time.Sleep(time.Duration(3) * time.Second)
        ch <- true
    }()

    //select在进入时会选择一个可执行的case来执行, 如果都不可执行则阻塞
    select {
        //耗时任务执行完毕
        case <- ch:
            fmt.Println("ok")
        //5s后超时
        case <- time.After(time.Duration(5) * time.Second):
            fmt.Println("time out")
        //当其他case都不能执行时, 执行default
        default:
            fmt.Println("default")
    }
}
