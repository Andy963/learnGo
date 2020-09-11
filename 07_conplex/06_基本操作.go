package main

import "fmt"

func main()  {
	var a [10]int
	var b [5]int
	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a),len(b))

	// 定义数组时，指定数组个数必须是常量
	// n :=5  var c [n]int 是错误的

	// 操作数组元素，从0开始，到len()-1，不对称元素， 这个数字叫下标
	// 下标可以是变量或者常量
	a[0] = 1
	i := 1
	a[i] = 2

	// 赋值

	for i :=0; i<len(a); i ++ {
		a[i] = i+1
	}

	// range a 返回第一个为下标，第二个为数据
	for i,data := range a {
		fmt.Printf("a[%d] = %d\n",i,data)
	}
}
