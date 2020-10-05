package main

import (
    "fmt"
    "time"
)

func main2()  {
   <-time.After(2*time.Second)
   fmt.Println("on time.")
}

func main1()  {
    time.Sleep(2*time.Second)
    fmt.Println("on time.")
}
func main(){
    timer := time.NewTimer(2*time.Second)
    <- timer.C
    fmt.Println("on time.")
}
