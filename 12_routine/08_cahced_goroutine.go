package main

import (
    "fmt"
    "time"
)

func main(){

    ch := make(chan int, 3) // set capacity to 0  means no cache

    fmt.Printf("len(ch)= %d, cap(ch) =%d\n", len(ch), cap(ch))

    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
            fmt.Printf("sub go routine. index[%d], len(ch)= %d, cap(ch) =%d\n",i, len(ch), cap(ch))
        }
    }()

    time.Sleep(2 * time.Second)
    for i := 0; i < 3; i++ {
        num := <-ch
        fmt.Println("num = ", num)
    }
}
//len(ch)= 0, cap(ch) =3
//sub go routine. index[0], len(ch)= 1, cap(ch) =3
//sub go routine. index[1], len(ch)= 2, cap(ch) =3
//sub go routine. index[2], len(ch)= 3, cap(ch) =3
//num =  0
//num =  1
//num =  2
