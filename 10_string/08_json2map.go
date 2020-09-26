package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j_s := `
 {
 "Company": "gian",
 "Subject": [
  "python",
  "go",
  "js"
 ],
 "Status": true,
 "Price": 666
}
`
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(j_s), &m)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("m = %+v\n", m)
	// var str string
	// str = string(m["company"]) 无法强制转换
	for key, value := range m {
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]=%s\n", key, data)
		case bool:
			fmt.Printf("map[%s]=%v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]=%v\n", key, data)
		}
	}
	fmt.Printf("")
}
