package main

import "fmt"
// common function
func add(a,b int) int{
    return a +b
}
// oop method
type long int

func (tmp long)  Add(other long) long{
    return tmp + other
}
func main(){
    // use function
    var result int
    result = add(1,2)
    fmt.Println("result=",result)
    // define a long var as a receiver
    var r long = 1
    // use oop method
    rest := r.Add(2)
    fmt.Println("rest=",rest)
}
