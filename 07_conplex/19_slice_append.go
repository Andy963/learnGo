package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Printf("s1=%d, len=%d, cap=%d\n", s1, len(s1), cap(s1)) //s1=[], len=0, cap=0

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	fmt.Printf("s1=%d, len=%d, cap=%d\n", s1, len(s1), cap(s1)) //s1=[], len=0, cap=0
}
