package main

import "fmt"

func test1() {
    fmt.Println("test1")
}
func test2(x int) {
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
//test1
//panic: runtime error: index out of range [20] with length 10
//
//goroutine 1 [running]:
