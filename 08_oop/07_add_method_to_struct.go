package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender byte
}

func (tmp Person) print_info() {
	fmt.Println(tmp)
}
func (p * Person) set_info(n string, a int, g byte) {
	p.name = n
	p.age = a
	p.gender =g
}

func main() {
	p := Person{"andy", 20, 'm'}
	p.print_info()

	var p1  Person
	(&p1).set_info("Mike",18,'m')
	p1.print_info()
}
