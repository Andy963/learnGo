package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("a=,b=", *a, *b)
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println("main:a=,b=", a, b)
}
