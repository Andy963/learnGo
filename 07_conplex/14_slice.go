package main

import "fmt"

func main()  {
	a := []int{1,2,3,4,5}
	s :=a[0:3:4]
	fmt.Println("s=",s)
	fmt.Println("len(s)=",len(s))
	fmt.Println("cap(s)=",cap(s))
}

//s= [1 2 3]
//len(s)= 3
//cap(s)= 5

