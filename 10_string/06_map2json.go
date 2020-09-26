package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["company"] = "giant"
	m["subjects"] = []string{"python", "go", "js"}
	m["status"] = true
	m["price"] = 666.6
	result, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("result = ",string(result))
}
