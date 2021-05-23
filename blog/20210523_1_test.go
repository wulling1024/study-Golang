package blog

import (
	"fmt"
	"sync"
	"testing"
)

/**
URL :
*/

func Test20210523_1(t *testing.T) {
	var m sync.Map

	m.Store("address", map[string]string{"province": "广东省", "city": "深圳市"})
	load, _ := m.Load("address")

	fmt.Println(load["address"]["province"])
}

func Test20210523_1_1(t *testing.T) {
	var m sync.Map
	m.Store("address", map[string]string{"province": "江苏", "city": "南京"})
	v, _ := m.Load("address")
	fmt.Println(v["province"])
}
