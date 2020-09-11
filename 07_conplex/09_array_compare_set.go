package main

import "fmt"

func main() {
	a := [3]int{1, 2}
	b := [3]int{1, 2}
	c := [3]int{1, 2, 3}

	fmt.Println(a == b)
	fmt.Println(a == c)

	var d [3]int
	d = c
	fmt.Println("d=", d)
}
