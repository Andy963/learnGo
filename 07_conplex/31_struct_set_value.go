package main

import "fmt"


type Student struct {
    id     int
    name   string
    gender byte
    age    int
}

func main(){
    var s Student
    s.id = 1
    s.name = "andy"
    s.gender='m'
    s.age = 29
    fmt.Println("s=",s)
}
