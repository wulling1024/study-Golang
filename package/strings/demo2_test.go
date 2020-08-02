package strings

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Name struct{
	Name *string `json:"name"`
	Age  int	`json:"age"`
}
// 结构体定义
type robot struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func Test_PointerOfString(t *testing.T){
	str := "{\"name\":\"0xc00005c5c0\", \"age\":23}"
	var info Name
	if err := json.Unmarshal([]byte(str), &info); err != nil {
		fmt.Println("解析失败")
	}
	fmt.Println(info)
	fmt.Println(time.Now().Unix())
}


// 解析到结构体数组
func Test_parse_array(t *testing.T) {
	fmt.Println("解析json字符串为结构体数组")
	str := "[{\"name\":\"name1\",\"amount\":100},{\"name\":\"name2\",\"amount\":200},{\"name\":\"name3\",\"amount\":300},{\"name\":\"name4\",\"amount\":400}]"
	all := []robot{}
	err := json.Unmarshal([]byte(str), &all)
	if err != nil {
		fmt.Printf("err=%v", err)
	}
	for _, one := range all {
		fmt.Printf("name=%v, amount=%v\n", one.Name, one.Amount)
	}
}
// 解析到结构体指针的数组
func Test_parse_pointer_array(t *testing.T) {
	fmt.Println("解析json字符串为结构体指针的数组")
	str := "[{\"name\":\"name1\",\"amount\":100},{\"name\":\"name2\",\"amount\":200},{\"name\":\"name3\",\"amount\":300},{\"name\":\"name4\",\"amount\":400}]"
	all := []*robot{}
	err := json.Unmarshal([]byte(str), &all)
	if err != nil {
		fmt.Printf("err=%v", err)
	}
	for _, one := range all {
		fmt.Printf("name=%v, amount=%v\n", one.Name, one.Amount)
	}
}
