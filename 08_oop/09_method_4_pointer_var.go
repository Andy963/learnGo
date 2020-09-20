package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender byte
}

func (p Person) setInfo() {
	fmt.Println("setInfo")
}

func (p *Person) setInfoPointer() {
	fmt.Println("setInfoPointer")
}

func main() {
	// struct var is a pinter,has a set of method
	p := Person{"andy", 20, 'm'}
	p.setInfo()        // func (p *Person) setInfoPointer()
	p.setInfoPointer() // change to (&p).setInfoPointer()
	(&p).setInfoPointer()

}

// p is common var if the func need a pointer, go will use &p,if not, use p
//p is a pointer if the func need a common val, go will use *p, if not use p
