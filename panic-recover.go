package main

import "fmt"

func main() {
    defer func() {
        if err := recover(); err != nil { //recover必须是defer调用
            fmt.Println(err)
        }
    }()

    go func() { //如果发生panic的goroutine和声明recover的goroutine不是同一个, 此时recover无法捕获错误
        panic(404) //这个panic无法被main goroutine中的recover捕获
    }()

    panic(500) //这里的panic才可以被捕获
}
