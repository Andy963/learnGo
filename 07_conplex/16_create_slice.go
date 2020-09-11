package main

import "fmt"

func main()  {
	// 自动推导，同时初始化
	s1 := []int{1,2,3,4}
	fmt.Println("s1=",s1)

	// 借助make函数，格式： make(切片类型，长度，容量）
	s2 := make([]int,5,10)
	fmt.Printf("len=%d, cap=%d\n",len(s2),cap(s2))

	// if no cap, the cap will same as len
	s3 := make([]int, 5)
	fmt.Printf("len=%d, cap=%d\n",len(s3),cap(s3))
}

//s1= [1 2 3 4]
//len=5, cap=10
//len=5, cap=5
