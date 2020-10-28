package main

import (
	"fmt"
	"os"
	"strings"
)

// 接口类型
type Stringer interface {
	String() string
}

// 基类型别名
type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

type SymbolString struct {
	s string
}

func (p *SymbolString) String() string {
	return strings.Replace(p.s, ",", "-", -1)
}

func main() {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
	symbolString := SymbolString{
		"hello, world",
	}
	fmt.Fprintln(os.Stdout, symbolString.String())
}
