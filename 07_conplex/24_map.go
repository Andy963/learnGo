package main

import "fmt"

func main() {
	var m1 map[int]string     // int is the type of key, string is the type of value
	fmt.Printf("m1=%v\n", m1) // only have len method, no cap method

	// create by make
	m2 := make(map[int]string)
	fmt.Printf("m2=%v\n", m2)
	fmt.Println("len=", len(m2))
	// if you set the len of a map, but no key-value in it, the len(m)=0
	// you can enlarge the map by add new key-value,even if set a len
	m3 := make(map[int]string, 2) // len =2
	m3[1] = "go"
	m3[2] = "py"
	m3[3] = "c"
	fmt.Println("m3=, len=", m3, len(m3))

	// initial while create
	m4 := map[int]string{1: "c++"}
	fmt.Println("m4=", m4)
}

//m1=map[]
//m2=map[]
//len= 0
//m3=, len= map[1:go 2:py 3:c] 3
//m4= map[1:c++]
