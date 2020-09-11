package main

import "fmt"

func main() {
	// 定义在 {}里的变量就是局部变量
	// 只能在{}里面用，作用域就是当前{}
	if flag := 0; flag == 0 {
		fmt.Println("flag=", flag)
	}
	// flag = 1
	// flag 的作用域仅仅为if语句，出了它就失效了
}
