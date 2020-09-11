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
    name string // same field in Person
}

func main(){
    var s  Student
    s.grade =1
    s.name = "andy" // 就近原则,{Person:{name: gender:0 age:0} grade:1 teacher: name:andy}
    fmt.Printf("%+v\n",s)
    //if you want to set name field in Person, it should be like this:
    s.Person.name = "mike"
    fmt.Printf("%+v\n",s)
    //{Person:{name:mike gender:0 age:0} grade:1 teacher: name:andy}
}
