package main

import "fmt"

type Person struct {
	name   string
	gender byte
	age    int
}

type Student struct {
	Person  // 匿名字段，只有类型，没有名字，把Person中的字段都继承过来
	grade   int
	teacher string
}

func main() {
	// 顺序初始化, s1后面的student可以省略不写
	var s1 Student = Student{Person{"andy", 'm', 29}, 1, "Mr wang"}
	fmt.Println("s1=", s1)
	// 自动推导类型
	s2 := Student{Person{"andy", 'm', 29}, 1, "Mr wang"}
	fmt.Println("s2=", s2)
	// +v会显示详细字段信息
	fmt.Printf("s2= %+v\n", s2)
	//部分赋值
	s3 := Student{Person: Person{name: "mike"}, grade: 2}
	fmt.Println("s3=", s3)
}
