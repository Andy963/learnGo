package main

import "fmt"

func main() {

	source := []int{1, 2, 3, 7, 9, 10, 7}
	destination := []int{6, 6, 6, 6, 6, 6}
	copy(destination, source)
	fmt.Println(destination)

}

//[1 2 3 6 6 6]
