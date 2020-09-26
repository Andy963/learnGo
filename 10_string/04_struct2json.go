package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量名首字母必须大写,小写字母开头的将会被忽略
type info struct {
	Company string
	Subject []string
	Status  bool
	Price   float64
}

func main() {
	s := info{"gian", []string{"python", "go", "js"}, true, 666}
	//j_s, err := json.Marshal(s)
	j_s, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println("j_s=", string(j_s))
}
