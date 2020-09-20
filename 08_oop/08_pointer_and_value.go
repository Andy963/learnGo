package main

import "fmt"

type Person struct {
    name   string
    age    int
    gender byte
}

// the receiver is common, not pointer, will copy the value
func (p Person) set_info(n string, a int, g byte){
    p.name = n
    p.age = a
    p.gender = g
    fmt.Printf("set into &p=%p\n", &p)
    fmt.Println("in set info p= ",p)
}

// the receiver is pointer,just transfer the reference.
func (p *Person) set_info_by_pointer(n string, a int, g byte){
    p.name = n
    p.age = a
    p.gender = g
    fmt.Printf("set info p=%p\n", p)
    fmt.Println("in set info by pointer p=", *p)
}

func main(){
	p1 := Person{"andy",20,'m'}
    fmt.Printf("%p\n",&p1)

	// set value
	p1.set_info("andy",29,'m')
    fmt.Println("p1= ",p1)

    (&p1).set_info_by_pointer("jack",20,'m')
	fmt.Println("p1= ",p1)
}
//0xc00018c000
//set into &p=0xc00018c020
//in set info p=  {andy 29 109}
//p1=  {andy 20 109}
//set info p=0xc00018c000
//in set info by pointer p= {jack 20 109}
//p1=  {jack 20 109}
