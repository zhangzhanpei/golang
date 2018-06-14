package main

import "fmt"

func main() {
    // 定义一个 map，这里的 m 实际上是一个指针类型
    m := make(map[string]int)
    m["语文"] = 85
    m["数学"] = 90
    fmt.Printf("%p \n", &m) // 0xc042004018

    // 传递一个 map
    modify(m)
    fmt.Println(m) // map[语文:80 数学:90]
}

func modify(m map[string]int) {
    // 值传递，传进来的是指针的拷贝，指向原来的 map
    fmt.Printf("%p \n", &m) // 0xc042004028
    // 因此，这里修改 map 直接修改了原先的 map
     m["语文"] = 80
}
