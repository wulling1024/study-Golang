package _for

import (
	"fmt"
	"testing"
)

func Test_range(t *testing.T){
	rates:=[]int32{1,2,3,4,5,6}
	for star, rate := range rates {
		if star+1 < 1 {
			panic("")
		}
		fmt.Println(star, rate)
	}
}
