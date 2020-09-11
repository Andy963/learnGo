package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [5]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Println("rand array n= ", a)
	fmt.Println("bubble sort:")

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("after sort:")
	fmt.Println("sorted a = ", a)
}
