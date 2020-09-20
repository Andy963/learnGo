package main

import "fmt"

// define interface, just define no implement, implement by other type (usually diy type)
type Humaner interface {
	// only method define
	say()
}

type Student struct {
	name string
	id   int
}

// Student implement say method
func (s *Student) say() {
	fmt.Println("student say: good good study.", s.name, s.id)
}

type Teacher struct {
	name  string
	grade int
}

// Teacher implement say method
func (s *Teacher) say() {
	fmt.Println("student say: good good study.", s.name, s.grade)
}

type myStr string

func (m *myStr) say() {
	fmt.Println("I'm a string", *m)
}
func main() {
	// define a interface type
	var i Humaner
	// if one type implement the method, this type variable(receiver) can use it
	s := &Student{"andy",1}
	i = s
	i.say()
	t := &Teacher{"wang",2}
	i =t
	t.say()
	var str1 myStr = "hello andy"
	i = &str1
	i.say()
}
//student say: good good study. andy 1
//student say: good good study. wang 2
//I'm a string hello andy
