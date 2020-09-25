package main

import (
    "fmt"
    "strconv"
)

func main(){
    slice :=make([]byte,0,1024)
    slice = strconv.AppendBool(slice,true)
    slice = strconv.AppendInt(slice,1234,10)
    slice = strconv.AppendQuote(slice,"Andy")
    fmt.Println("slice=", string(slice))

    var str string
    str = strconv.FormatBool(false)
    //'f' format, -1精度， 64 float64
    str = strconv.FormatFloat(3.14,'f',-1,64)
    fmt.Println("str=",str)
    // str convert to bool
    var flag bool
    var err error
    flag,err = strconv.ParseBool("true")
    if err == nil{
        fmt.Println("flag = ",flag)
    }else{
        fmt.Println("err = ",err)
    }
    // str convert to int
    a, _ := strconv.Atoi("567")
    fmt.Println("a = ",a)
}

//slice= true1234"Andy"
//str= 3.14
//flag =  true
//a =  567
