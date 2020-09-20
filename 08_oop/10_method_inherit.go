package main

import "fmt"

type Person struct {
    name   string
    age    int
    gender byte
}

func (p *Person)PrintInfo()  {
    fmt.Println(p.name,p.age,p.gender)
}

type Student struct {
    Person // anonymous field
    id int
}

func main(){
    s := Student{Person{"andy",29,'m'},1}
    s.PrintInfo()
}
// s inherit PrintInfo from Person