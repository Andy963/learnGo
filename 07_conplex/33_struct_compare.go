package main

import "fmt"

type Student struct {
	id     int
	name   string
	gender byte
	age    int
}

func main() {
	s1 := Student{1, "andy", 'm', 19}
	s2 := Student{1, "andy", 'm', 19}
	s3 := Student{1, "andy", 'f', 19}
	fmt.Println("s1==s2", s1 == s2)
	fmt.Println("s1==s3", s1 == s3)
	var temp Student
	temp = s3
	fmt.Println("temp=", temp)
}

//s1==s2 true
//s1==s3 false
//temp= {1 andy 102 19}
