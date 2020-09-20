package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender byte
}

func (p *Person) PrintInfo() {
	fmt.Println(p.name, p.age, p.gender)
}

type Student struct {
	Person // anonymous field
	id     int
}

func (p *Student) PrintInfo() {
	fmt.Println("student:=", p)
}

func main() {
	s := Student{Person{"andy", 29, 'm'}, 1}
	s.PrintInfo()
	// use the method which is closer
	//if you want to use the inherit method
	// you should declare it, like this:
	s.Person.PrintInfo()
}
