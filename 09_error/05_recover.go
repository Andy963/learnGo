package main

import "fmt"

func test1() {
    fmt.Println("test1")
}
func test2(x int) {
	// recover 可以打印出panic的信息
	// 当没有崩溃时，recover返回 nil
	// recover can only use after defer
	// recover 返回之后，再次调用返回空值
	defer func(){
	    if err := recover(); err !=nil{
	        fmt.Println(err) // 这里再写recover 返回nil
        }
    }()
    var a [10]int
    a[x] = 111 //当x超出边界时，产生一个panic
}
func test3() {
    fmt.Println("test3")
}

func main() {
    test1()
    test2(20)
    test3()
}