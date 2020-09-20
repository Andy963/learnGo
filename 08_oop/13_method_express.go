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
    p := Person{"andy",29, 'm'}

    fmt.Printf("%p %v\n", &p, p)
    // express f := p.SetInfoPointer, hide the receiver
    f :=(Person).setInfo // with out receiver, but the type should be same Person
    f(p)
    // pointer
    f1 := (*Person).setInfoPointer
    f1(&p)
}
//0xc00000c060 {andy 29 109}
//setInfo &{andy 29 109} {andy 29 109}
//setInfoPointer &{andy 29 109} {andy 29 109}

