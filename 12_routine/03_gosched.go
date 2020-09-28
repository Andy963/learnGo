package main

import (
    "fmt"
    "runtime"
)

func main(){
    go func(){
        for i := 0; i<5;i++{
            fmt.Println("go",i)
        }
    }()
    for i :=0;i<4;i++{
        runtime.Gosched()// 让出时间片，让别的程序先执行
        fmt.Println("main",i)
    }
}
