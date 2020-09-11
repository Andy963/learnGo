package main

import "fmt"

func main()  {
	a := []int{0,1,2,3,4,5,6,7}

	// new slice
	s1 :=a[2:5] // start from a[2],catch 3 ele [2,3,4], s1 cap=6
	s1[1] = 666
	fmt.Println("s1=",s1)

	// new slice
	s2 :=s1[2:6] // [4,5,6,7]
	fmt.Println("s2=",s2)
}
