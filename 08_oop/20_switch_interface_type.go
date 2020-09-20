package main

import "fmt"

type Student struct {
    name string
    id   int
}

func main() {
    i := make([]interface{}, 3)
    i[0] = 1
    i[1] = "hello world"
    i[2] = Student{"andy", 1}

    for index, data := range i{
        switch value :=data.(type) {
        case int:
            fmt.Printf("i[%d] type is int, data is %d\n", index, value)
        case string:
            fmt.Printf("i[%d] type is string, data is %s\n", index, value)
        case Student:
            fmt.Printf("i[%d] type is student, Student name is %s\n", index, value.name)
        }
    }
}
