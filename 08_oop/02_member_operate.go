package main

import "fmt"

type Person struct {
	name   string
	gender byte
	age    int
}

type Student struct {
	Person  // 匿名字段，只有类型，没有名字，把Person中的字段都继承过来
	grade   int
	teacher string
}

func main() {
	var s1 = Student{Person{"andy", 'm', 29}, 1, "Mr wang"}
	fmt.Println("s1=", s1)

	s1.name = "Jack"
	s1.gender = 'm'
	s1.age = 20
	fmt.Println("s1=",s1)
	// 局部赋值
	s1.Person = Person{"Mary",'f',18}
	fmt.Println("s1=",s1)
}
