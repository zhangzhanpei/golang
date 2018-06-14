package main

import "fmt"

func main() {
    // 定义一个数组
    arr := [2]int{1, 2}
    fmt.Printf("%p \n", &arr) // 0xc04204a080

    // 传递一个数组
    modify(arr)
    fmt.Println(arr) // 这里依然打印 [1 2]
}

func modify(arr [2]int) {
    // 值传递，传进来的是数组的拷贝
    fmt.Printf("%p \n", &arr) // 0xc04204a0b0
    // 这里通过下标修改数组元素的值，但只是修改了数组拷贝，不影响原数组
    // 因此，传递一个大的数组时最好传指针，否则数组拷贝很耗内存
     arr[0] = 2
}
