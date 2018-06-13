func deferCall() {
    // defer 函数延迟到 return 之前执行，多个 defer 之间按后进先出的顺序执行
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()
    // 这里触发异常，deferCall 方法停止执行并 return
    panic("触发异常")

    // 所以 deferCall 方法先倒序执行 defer 函数，再执行 panic，输出如下：
    // 打印后
    // 打印中
    // 打印前
    // panic: 触发异常
}
