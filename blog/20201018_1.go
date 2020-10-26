package main

import (
	"encoding/json"
	"fmt"
)

/*
	Website: https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651442371&idx=1&sn=e59aea7a9ccf5ac14fd825231d2850ae&chksm=80bb1231b7cc9b27fa9395e5d95aa2f0ac9afd66d16bf961bfac65691bf0cc16e5ddd8b0f1e5&scene=126&sessionid=1603024770&key=7f54b3443b6830331151116bd037df8130cea035b9470666785cea8a43faff2f33c6fbbb10640d3a7fb823ae71b3ad7ebbb2e0adc8d7ba9f7172f1e8100647371d5d79e41f248e95e1a9623a870ee4e841d0e6bc2ae6c24655ccec5d00c9e78e150bbe6d6ea8e94554b573fd3b480c100fe61f37b605d8f618364fa5cdfd3a2a&ascene=1&uin=MTIxMjU4MzMxNg%3D%3D&devicetype=Windows+10+x64&version=6300002f&lang=zh_CN&exportkey=AdlX9qkF8ggyVVAP9sfRzNA%3D&pass_ticket=%2FbJIhlbbmGNhP%2BGv3XmJe5DFj9hlmoxcrHEnG5oQeWOibjFb9jeK3jgfW2I5IpsH&wx_header=0
*/

type AutoGenerated struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Child []int  `json:"child"`
}

func main() {
	jsonStr1 := `{"age": 14,"name": "potter", "child":[1,2,3]}`
	a := AutoGenerated{}
	json.Unmarshal([]byte(jsonStr1), &a)
	aa := a.Child
	fmt.Println(aa)
	jsonStr2 := `{"age": 12,"name": "potter", "child":[3,4,5,7,8,9]}`
	json.Unmarshal([]byte(jsonStr2), &a)
	fmt.Println(aa)
}


