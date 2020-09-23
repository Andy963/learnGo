package main

import "fmt"

func test1() {
	fmt.Println("test1")
}
func test2() {
	panic("panic test, program will stop here.")
	fmt.Println("test2")
}
func test3() {
	fmt.Println("test3")
}

func main() {
	test1()
	test2()
	test3()
}
