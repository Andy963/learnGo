package main

import "fmt"

type Person struct{
	name string 
	gender byte
}

type Student struct{
	*Person // pointer type
	id int
	addr string
}

func main(){
	s1 := Student{
		&Person{"mike",'m'},
		666,
		"wuhan"}
	fmt.Println(s1.name,s1.gender,s1.id,s1.addr)
    var s2 Student
    s2.Person = new(Person) // 分配内存空间
    s2.name ="Mary"
    s2.gender = 'm'
    s2.id = 111
    s2.addr = "HuBei"
	fmt.Println(s2.name,s2.gender,s2.id,s2.addr)

}
