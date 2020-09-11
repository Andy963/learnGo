package main

import "fmt"

func main()  {
	var p *int
	fmt.Println("p=",p) // p 指向空
	// 不要操作没有合法指向的内存地址
	//*p = 666
}
