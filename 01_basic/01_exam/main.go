package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("This is exam 1\n")
	const (
		Man    = 1
		Female = 2
	)
	for {
		second := time.Now().Unix()
		if second%2 == 0 {
			fmt.Println("Female")
		} else {
			fmt.Println("man")
		}
		time.Sleep(1000 * time.Millisecond)
	}

}
