package main

import "fmt"

var global_var int

func main() {
	global_var = 100
	fmt.Println("global_var=", global_var)
}
