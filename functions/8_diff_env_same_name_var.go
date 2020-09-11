package main

import "fmt"

var a byte

func main() {
	var a int
	fmt.Printf("1:type of a is %T\n", a) // int
	{
		var a float32
		fmt.Printf("2:type of a is %T\n", a) // float 32
	}
	test()
}

func test() {
	fmt.Printf("3:type of a is %T\n", a) // byte
}
