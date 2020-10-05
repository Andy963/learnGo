package main

import (
    "fmt"
    "time"
)

func main(){
    timer := time.NewTimer(2*time.Second)
    fmt.Println("now time is :",time.Now())

    // after 2 seconds send data to timer.C
    t := <-timer.C
    fmt.Println("t=",t)
}
