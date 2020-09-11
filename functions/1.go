package main

import "fmt"

func my_func(args ...int) {

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d\n", i, args[i])
	}

	for i, data := range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
}

func my_func1(a int, args ...int) {
	// 固定参数一定要传值，不定参数可以不传，且不定参数放在固定参数之后
}
func my_func2() int { // 只有一个返回值，这里不用括号
	return 666
}

func my_func3() (result int) { //给返回值取个名字
	result = 666 // 给返回值赋值
	return
}
func my_func4() (int, int, int) {
	return 1, 2, 3
}

func my_func5() (a, b, c int) { // a int, b int, c int
	a, b, c = 1, 2, 3
	return
}

func min_max(a, b int) (min, max int) {
	if a > b {
		min, max = b, a
	} else {
		min, max = a, b
	}
	return
}
func main() {
	min, max := min_max(10, 20)
	fmt.Printf("min=%d,max=%d\n", min, max)

}
