package main

import "fmt"
type Humaner interface {
    // only method define
    say()
}

type Person interface {
	Humaner // anonymous field, inherit Humaner
    read(book string)
}

type Student struct {
    name string
    id   int
}

// implement say
func (s *Student) say()  {
    fmt.Println("student say",s.name)
}

// implement read
func (s *Student) read(book string)  {
   fmt.Println(s.name,"read",book)
}

func main(){
    var i Person
    s := &Student{"andy",1}
    i = s
    i.say()
    i.read("gone with wind")
}
