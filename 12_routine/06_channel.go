package main

import (
    "fmt"
    "time"
)

func main(){
    ch := make(chan  string)
    defer fmt.Println("主协程结束")

    go func(){
        defer fmt.Println("子协程调用完毕")

        for i :=0; i<2;i++{
            fmt.Println("子协程i=",i)
            time.Sleep(time.Second)
        }
        ch <- "I am sub routine.I have job to do "
    }()
    // <-ch  means get mes from channel. <-ch "hello"  add msg to channel.
    str := <-ch // if ch is empty, it will blocked
    fmt.Println("str = ", str)
}
