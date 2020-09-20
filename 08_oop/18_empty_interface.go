package main

import "fmt"

func anyNumParam(arg ... interface{})  {
   // arg is an empty interface type
    // so this function can receive any number of parameters
}

func main(){
    // emtpy interface can store all type of data
    var i interface{} =1
    fmt.Println("i=",i)
    i = "abc"
    fmt.Println("i=",i)
}
