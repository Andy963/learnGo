package main

import "fmt"

func main() {
	// 数组[]里面的长度是一个固定的常量，数组不能修改长度
	a := [5]int{}
	fmt.Printf("len = %d, cap=%d\n", len(a), cap(a))

	// 切片[]里面为空，或者为...，切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))

	s = append(s, 11) // 给切片追加一个成员
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
//len = 5, cap=5
//len=0,cap=0
//len=1,cap=1
