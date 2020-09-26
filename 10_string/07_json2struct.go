package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Company string
	Subject []string
	Status  bool
	Price   float64
}

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
	var tmp info
	err := json.Unmarshal([]byte(j_s), &tmp)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("tmp= %+v\n", tmp)

	// if you only need some filed of it.
	type info2 struct {
		Company string `json:"company"`
	}
	var tmp2 info2
	err = json.Unmarshal([]byte(j_s), &tmp2)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
}
