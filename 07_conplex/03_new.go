package main

import "fmt"

func main() {
	a := 10
	var p *int

	p = &a // p指向一个合法内存

	//p 是*int, 指向int顾炎武
	p = new(int)

	*p = 666
	fmt.Println("*p=", *p)
}
