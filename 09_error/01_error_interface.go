package main

import (
	"errors"
	"fmt"
)

func main() {
	error1 := errors.New("this is normal error.")
	fmt.Println("error1", error1)
}

//func main(){
//    var err1  error = fmt.Errorf("%s","this is an error.")
//    fmt.Println("err1",err1)
//    err2 := fmt.Errorf("%s","this is a normal err2")
//    fmt.Println("err2",err2)
//}
