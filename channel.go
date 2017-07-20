//无缓冲channel, 放入数据后阻塞直到channel中的数据被取出
//有缓冲channel, 当channel满时试图继续插入数据到channel时阻塞
package main

import (
    "fmt"
    "time"
)

// func main() {
//     var c = make(chan bool) //无缓冲channel
//     go func() {
//         fmt.Println("in")
//         c <- true //这里会阻塞住, 直到其他地方接手了channel中的数据
//         fmt.Println("out")
//     }()
//     fmt.Println("start");
//     time.Sleep(time.Second * 3) //3秒后channel中的数据被取出, 然后goroutine继续执行, 打印出"out"
//     <- c
//     time.Sleep(time.Second * 30)
// }

func main() {
    var c = make(chan bool, 3) //有缓冲channel, 长度为3
    go func() {
        for i := 0; i < 10; i++ {
            c <- true
            fmt.Println(i) //这里先打印了0, 1, 2, 此时channel已满, goroutine阻塞住
        }
    }()
    fmt.Println("start");
    time.Sleep(time.Second * 3) //3秒后channel中的数据读出一个, 然后goroutine继续执行, 往channel中插入一个true, 打印出3
    <- c
    //主线程退出, goroutine强制结束
}
