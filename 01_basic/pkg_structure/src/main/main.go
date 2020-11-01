package main

import (
	"fmt"
	"calc"
)

func main() {
	// 需要将 src父级目录加入到GOPATH中去
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 300)
	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}
