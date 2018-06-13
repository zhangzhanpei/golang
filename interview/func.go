func main() {
    // 这里打印的 i 是外部的变量，地址不变化，始终是同一个变量
    // i 在 for 循环中不断增加，当协程执行到打印 i 的时候，就打印此刻 i 的值
    // 所以，打印出来的 i 是不固定的，要看 for 循环和协程的执行速度
    // 注意，如果 runtime.GOMAXPROCS(1) 设置只使用一个处理器，那么 for 循环启动的协程没有时间片执行，会等 for 循环完毕后才能执行，此时打印出来的 i 都是 10
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("i: ", i)
        }()
    }

    // 这里的协程打印的 i 是值传递进来的，协程启动的时候就已经定了的
    // 所以，这里打印的就是 0 到 9
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("ii: ", i)
        }(i)
    }

    time.Sleep(5 * time.Second)
}
