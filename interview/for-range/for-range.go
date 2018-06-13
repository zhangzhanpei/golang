package main

type Student struct {
    Name string
    Age int
}

func main() {
    m := make(map[string]*Student)
    students := []Student{
        {Name: "zhangsan", Age: 18},
        {Name: "lisi", Age: 25},
    }
    // 如下构造的 map 所有元素都会指向最后一个 student
    // 原因是 for...range 循环中，遍历到的元素使用值拷贝的方式赋给了 stu 变量，在整个循环过程中，stu 临时变量是重复使用的
    // 所以，循环完成后，stu 保存的是 students 数组中的最后一个元素，而 map 中保存的元素都是 stu 这个变量的地址
    for _, stu := range students {
        m[stu.Name] = &stu
    }

    // 解决方案：直接通过数组下标取对象地址即可
    for k := range students {
        m[students[k].Name] = &students[k]
    }
}
