package strings

import (
	"fmt"
	"strings"
	"testing"
)

func Test_fmt(t *testing.T) {
	fmt.Println(fmt.Sprintf("123%d", 4))
	fmt.Println(fmt.Sprintf("123%s", "4"))
}

func Test_Merge(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"a", "b", "1"}
	c := append(a, b...)
	fmt.Println(c)

}

func Test_Change(t *testing.T) {
	// \u0040
	fmt.Println(strings.Replace("jiayi.shao@shopee.com", "@", "\\u0040", 1))
}

func Test_Interception(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5, 6}
	sli = sli[0:3:3]
	fmt.Println(sli)
}
