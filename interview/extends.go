type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    // 通过 t.People.ShowA() 调用的 ShowA() 方法，此时真正的调用者是 People，并不会把调用者替换成 Teacher，毕竟 People 并不知道被谁组合
    // 所以，这里的 ShowB() 方法调的是 People 类型下的 ShowB()
    p.ShowB()
}

func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    // 组合了 People 类型
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    // 这里调 ShowA() 方法时，Teacher 类型并没有 ShowA() 方法，因此调的是 People 类型的 ShowA() 方法，即等同于 t.People.ShowA()
    t.ShowA()
}
