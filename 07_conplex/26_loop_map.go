package main

import "fmt"

func main() {
	m := map[int]string{1: "go", 2: "python"}
	// range return: key,value
	for key, val := range m {
		fmt.Printf("%d ==>%s\n", key, val)
	}

	// : extract value, exists
	value, exists := m[1]
	if exists == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key not exists")
	}
}

//1 ==>go
//2 ==>python
//m[1] =  go
