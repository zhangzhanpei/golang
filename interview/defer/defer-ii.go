/**
 * return xx 并不是一个原子操作，是先给返回值赋值，再 return，共两步操作
 * defer 函数插入到 return 之前执行，所以 return xx 共三步：赋值 + defer + return
 * 所以，defer 可以对返回值进行修改
 * 技巧：先忽略 defer 函数，先做返回值的赋值，再看 defer 函数有没有修改返回值
 */

func deferCallA() (ret int) {
    defer func() {
        ret++
    }()
    // 三步
    // ret = 5
    // 执行 defer 后 ret = 6
    // return
    return 5
}

func deferCallB() (ret int) {
    tmp := 5
    defer func() {
        tmp = tmp + 5
    }()
    // 三步
    // ret = tmp 即 ret = 5
    // 执行 defer 函数修改了 tmp，ret 仍然为 5
    // return
    return tmp
}

func deferCallC() (ret int) {
    defer func(ret int) {
        ret = ret + 5
    }(ret)
    // 三步
    // ret = 0
    // 执行 defer 函数修改的是传入的 ret，ret 是值传递所以修改的是 ret 的一个副本，原来的 ret 仍然为 0
    // return
    return 0
}
