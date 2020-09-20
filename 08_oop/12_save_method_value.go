package main

import "fmt"

type Person struct {
    name   string
    age    int
    gender byte
}
func (p Person) setInfo() {
    fmt.Println("setInfo",&p, p)
}

func (p *Person) setInfoPointer() {
    fmt.Println("setInfoPointer", p, *p)
}
func main(){
    p :=Person{"andy",20,'m'}
    setInfoVal := p.setInfo
    setInfoVal()
    setInfoPointer :=p.setInfoPointer
    setInfoPointer()
}
//setInfo &{andy 20 109} {andy 20 109}
//setInfoPointer &{andy 20 109} {andy 20 109}
