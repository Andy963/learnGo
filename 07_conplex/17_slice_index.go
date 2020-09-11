package main

import "fmt"

func main()  {
	array := []int{0,1,2,3,4,5,6,7,8,9}
	//[low:high:max]取下标从low开始，len=high-low个元素，cap = max-low
	s1 := array[:] // [0:len(array):len(array)]
	fmt.Println("s1=",s1)
	fmt.Printf("len = %d, cap = %d\n",len(s1),cap(s1))

	s2 := array[2:] // 从2开始到结必
	fmt.Println("s2=",s2)
	fmt.Printf("len = %d, cap = %d\n",len(s2),cap(s2))

	s3 := array[:5] // 从0开始，到5,容量为10
	fmt.Println("s3=",s3)
	fmt.Printf("len = %d, cap = %d\n",len(s3),cap(s3))

	s4 := array[2:3] // 从2开始，到3, 容量为8
	fmt.Println("s4=",s4)
	fmt.Printf("len = %d, cap = %d\n",len(s4),cap(s4))


}

//s1= [0 1 2 3 4 5 6 7 8 9]
//len = 10, cap = 10
//s2= [2 3 4 5 6 7 8 9]
//len = 8, cap = 8
//s3= [0 1 2 3 4]
//len = 5, cap = 10
//s4= [2]
//len = 1, cap = 8