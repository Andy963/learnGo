package main

import "fmt"

func main() {
	type Student struct {
		id     int
		name   string
		gender byte
		age    int
	}
	// you should init all field
	var s1 Student = Student{1, "andy", 'm', 20}
	fmt.Println("s1=", s1)

	// init few field, other field will init to 0 automatically
	s2 := Student{id: 1, age: 19}
	fmt.Println("s2=", s2)
}
//s1= {1 andy 109 20}
//s2= {1  0 19}