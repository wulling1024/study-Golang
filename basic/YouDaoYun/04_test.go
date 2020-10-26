package main

import (
	"fmt"
	"testing"
)
/**
这一块内容，主要对应的是basic/04内容
 */
func Test_main (t *testing.T){
	name := "OldName"
	user := user{Name: &name, age: 25}
	fmt.Println(user)
	f(user)
	fmt.Println(user)
}

type user struct {
	Name *string
	age  int
}

func f(user user) {
	fmt.Println(user)

	*(user.Name) = "NewName"
	user.age     = 26

	fmt.Println(user)
}