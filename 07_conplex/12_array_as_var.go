package main

import "fmt"

// 数组作为函数参数，是值传递
// 实参数组的每个元素给形参数组拷贝一份
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify a =", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify(a)
	fmt.Println("main: a= ", a)
}
