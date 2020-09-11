package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	*p = num
}

func getNum(s []int, num int) {
	s[0] = num / 1000       // 取千位
	s[1] = num % 1000 / 100 // 取百位
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func inputNum(randSlice []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Printf("请输入一个四位数：\n")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Printf("您输入的数不合格：")
		}
		//fmt.Println("num=", num)
		getNum(keySlice, num)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 { // 都猜对了
			break
		}
	}
}

func main() {
	var randNum int
	createNum(&randNum)
	fmt.Printf("randNum=%d\n", randNum)
	randSlice := make([]int, 4)
	getNum(randSlice, randNum)
	fmt.Printf("randSlice=%d\n", randSlice)
	inputNum(randSlice)
 fmt.Print("不知道百度双拼输入法怎么样")
}