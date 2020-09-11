package main

import "fmt"

type Student struct {
	id     int
	name   string
	gender byte
	age    int
}

func main() {
	var p1 *Student = &Student{1, "andy", 'm', 29}
	fmt.Println("p1=", p1)
	p2 := &Student{name: "jack", age: 18}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("p2=", *p2)
}
