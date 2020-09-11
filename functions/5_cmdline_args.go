package main

import "fmt"
import "os"

func main() {

	arg_list := os.Args
	for i := 0; i < len(arg_list); i++ {
		fmt.Printf("args_list[%d]=%s\n", i, arg_list[i])
	}
}
