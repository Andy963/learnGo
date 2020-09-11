package main

import "fmt"

type Student struct {
	id     int
	name   string
	gender byte
	age    int
}

func main() {
	var s1 Student
	var p1 *Student
	// use pointer after have a valid point
	p1 = &s1
	// use p1.id is same as (*p1).id,but can only use dot
	p1.id = 1
	(*p1).name = "andy"
	p1.gender = 'm'
	p1.age = 29
	fmt.Println("p1=", p1)

	p2 := new(Student)
	p2.id = 2
	p2.name = "jack"
	p2.gender = 'm'
	p2.age = 18
	fmt.Println("p2=", p2)

}
//p1= &{1 andy 109 29}
//p2= &{2 jack 109 18}
