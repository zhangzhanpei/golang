package main

import "fmt"

type Employee struct {
    Age uint
}

func main() {
    // 实例化一个 struct
    e := Employee{18}
    fmt.Printf("%p \n", &e) // 0xc04204a080

    // 传递一个 struct
    modify(e)
    fmt.Println(e) // {18}
}

func modify(e Employee) {
    // 这里的变量地址已经改变了，说明 struct 是值传递，变量 e 是一个拷贝
    fmt.Printf("%p \n", &e) // 0xc04204a088
    // 因此，这里对 e 进行修改只是修改了副本
    e.Age = 25
}
