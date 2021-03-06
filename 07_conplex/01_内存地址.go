package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a=%d\n", a)
	fmt.Printf("a 的地址是:%v\n", &a)

	// 保存某个变量的地址，需要指针类型 *int保存int的地址， **int 保存 *int的地址
	// 定义一个变量p,类型为×int
	var p *int
	p = &a // p的值为a变量的内在地址
	fmt.Printf("p=%v, &=a=%v\n", p, &a)

	*p = 666 // *p 表示p变量所指向的内存，对它赋值，即修改了指向该地址的变量
	fmt.Printf("*p=%v, a=%v\n", *p, a)
}
