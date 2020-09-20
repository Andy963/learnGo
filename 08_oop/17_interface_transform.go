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
    var iPro Person
    iPro = &Student{"andy",1}
    var i Humaner
    // iPro = i  this will throw err
    i = iPro  // 超集转化为子集
    i.say()
}
