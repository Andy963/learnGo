package main

import "fmt"

func del_map(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "ruby", 2: "js"}
	fmt.Println("m=", m)
	del_map(m)
	fmt.Println("m=", m)
}
//m= map[1:ruby 2:js]
//m= map[2:js]