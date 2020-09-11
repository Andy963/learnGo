package main

import "fmt"

func zero(x int) {
	result := 10 / x
	fmt.Println("result=", result)
}

func main1() {

	defer fmt.Println("bbbbbbbbb")
	defer zero(0)
	defer fmt.Println("aaaaaaaa")
}

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}(a, b)
	a = -10
	b = -20
	fmt.Printf("a=%d,b=%d\n", a, b)
}
