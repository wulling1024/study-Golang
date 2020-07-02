package escapeAnalysis

import (
	"fmt"
	"testing"
)

type data struct {
	name string
}

func patent1()data{
	p := data{"keke"}
	fmt.Printf("patent1: %p \n", &p)
	return p
}

func patent2() *data {
	p := data{"jame"}
	fmt.Printf("patent2: %p \n", &p)
	return &p
}
func Test_Demo(t *testing.T){
	p1 := patent1()
	fmt.Printf("main: %p \n", &p1)
	p2 := patent2()
	fmt.Printf("main: %p \n", p2)
}


