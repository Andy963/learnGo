package main

import "fmt"

func main() {
	a := 10
	str := "hello"

	//匿名函数，没有函数名，
	f1 := func() {
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}
	// 调用
	f1()

	// 给函数类型取别名
	type func_type func() // 函数无参，无返回值

	var f2 func_type
	f2 = f1
	f2()

	// 定义匿名函数，同时调用
	func() {
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}() //此括号代表调用函数

	// 带参数
	f3 := func(i, j int) {
		fmt.Println("i=", i)
		fmt.Println("j=", j)
	}
	f3(1, 2)

	// 定义匿名函数同时调用
	func(i, j int) {
		fmt.Println("i,j", i, j)
	}(2, 1)

	// 匿名函数有参有返回值
	rest := func(a, b int) (result int) {
		result = a + b
		return
	}(1, 3)
	fmt.Println("rest=", rest)
}
