package main

import "fmt"

type Person struct {
	name   string
	gender byte
	age    int
}

type mystr string // 自定义类型，给一个类型改名

type Student struct {
	Person  // 匿名字段，只有类型，没有名字，把Person中的字段都继承过来
	int     // 年级，匿名字段
	teacher mystr
	name    string // same field in Person
}

func main() {
	s := Student{Person{"mike",'m',18},1,"wang","zhou"}
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Person.name,s.gender,s.age,s.int,s.teacher,s.name)
}
