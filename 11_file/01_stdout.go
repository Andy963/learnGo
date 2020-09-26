package main

import (
	"fmt"
	"os"
)

func main(){
    //os.Stdout.Close()
	// os.Stdin.Close() if close can not get input from stdin
    //fmt.Printf("")
	os.Stdout.WriteString("Are you ok,Andy\n")

	var a int
	fmt.Println("please input a:")
	fmt.Scan(&a) // get a from stdin
	fmt.Println("a = ",a)
}
