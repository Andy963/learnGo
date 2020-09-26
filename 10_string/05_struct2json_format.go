package main

import (
    "encoding/json"
    "fmt"
)

// 成员变量名首字母必须大写,小写字母开头的将会被忽略
// 如果要输出时的字段是小写，则 `json: "new name"`
type info struct {
    Company string `json:"company"`
    Subject []string `json:"-"` // 此字段不会输出
    Status  bool `json:",string"` // 此字段不会显示true,而是字符串"true"
    Price   float64 `json:",string"`
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
