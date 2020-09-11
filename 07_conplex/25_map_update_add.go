package main

import "fmt"

func main() {
	m1 := map[int]string{1: "python", 2: "goland", 3: "web"}
	fmt.Println("old m1=", m1)
	// if key exists,update it
	m1[1] = "py"
	fmt.Println("new m1=", m1)

	// if not exists add it to map
	m1[4] = "database"
	fmt.Println("new m1=", m1)
}

//old m1= map[1:python 2:goland 3:web]
//new m1= map[1:py 2:goland 3:web]
//new m1= map[1:py 2:goland 3:web 4:database]
