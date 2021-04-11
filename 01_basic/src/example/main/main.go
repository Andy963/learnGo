package main

import (
	packageAdd "../add"
	"fmt"
)

func main() {
	fmt.Println("Name=", packageAdd.Name)
	fmt.Println("Age=", packageAdd.Age)
}
