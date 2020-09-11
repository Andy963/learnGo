package main

import "fmt"

type Student struct {
	id     int
	name   string
	gender byte
	age    int
}

func test(s Student) {
	s.id = 666
	fmt.Println("in test, s=", s)
}

func test1(s *Student) {
	s.id = 666
	fmt.Println("in test, s=", *s)
}


func main() {
	s := Student{1, "andy", 'f', 20}
	test1(&s)
	fmt.Println("in main, s=", s)
}
//in test, s= {666 andy 102 20}
//in main, s= {1 andy 102 20}