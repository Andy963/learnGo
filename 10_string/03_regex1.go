package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aaa acc tac"
	pattern := regexp.MustCompile(`a.c`)
	if pattern == nil {
		fmt.Println("pattern err")
		return
	}

	result := pattern.FindAllStringSubmatch(buf, -1)
	fmt.Println("result= ", result)
	// result=  [[abc] [azc] [a7c] [acc]] 返回的是二维数组
}
