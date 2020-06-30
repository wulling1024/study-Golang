package main

import "reflect"

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

func main() {
	reflect.ValueOf(&User{}).Elem()
}
