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

// the params is a interface type
// one method but diff performance
func whoSay(i Humaner) {
	i.say()
}
func main() {
	s := &Student{"andy", 1}
	t := &Teacher{"Wang", 2}
	var str1 myStr = "Im a string."
	whoSay(s)
	whoSay(t)
	whoSay(&str1)

	x := make([]Humaner,3)
	x[0] =s
	x[1] = t
	x[2] = &str1
	for _, i := range x{
		i.say()
	}
}
