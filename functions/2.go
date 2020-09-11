package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

//通过type 给函数类型取名，它的参数，返回值类型必须与对应的函数一致
type func_type func(int, int) int // 没有函数名，没有{}

func main() {
	var result int
	result = add(1, 2) // 传统调用方法
	fmt.Println("result=", result)

	// 声明一个变量 func_name, 类型是函数func_type类型
	var func_name func_type

	func_name = add
	result = func_name(10, 20) // 等价于： add(10,20)
	fmt.Println("func_name: result=", result)

}
