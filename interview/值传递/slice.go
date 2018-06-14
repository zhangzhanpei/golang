package main

import "fmt"

func main() {
    // 定义一个 slice，sli 变量保存了切片的地址
    sli := []int{1, 2}
    fmt.Printf("%p \n", &sli) // 0xc0420023e0

    // 传递一个 slice
    modify(sli)
    fmt.Println(sli) // 这里依然打印 [1 2]
}

func modify(sli []int) {
    // 值传递，传进来的是 sli 变量的拷贝，这里打印的是这个拷贝的变量地址
    fmt.Printf("%p \n", &sli) // 0xc042002420
    // 这里将变量拷贝赋值为 nil，只是修改了这个拷贝，不影响变量保存的指针指向的底层数组
    sli = nil
    // 这里可以通过下标修改底层数组元素的值，但无法修改底层数组的长度和容量，如果要修改只能传指针
    // sli[0] = 2
}
