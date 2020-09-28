package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("defer")
	runtime.Goexit()
	fmt.Println("test finish")
}

func main() {
	go func() {
		fmt.Println("start")

		test()
		fmt.Println("routine finish.")
	}()
	// forever loop, prevent main exit.
	for {

	}
}
