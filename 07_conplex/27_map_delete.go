package main

import "fmt"

func main() {
	m := map[int]string{1: "ruby", 2: "js"}
	fmt.Println("m=", m)
	delete(m, 1)
	fmt.Println("m=", m)
}

//m= map[1:ruby 2:js]
//m= map[2:js]