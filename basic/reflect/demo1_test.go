package reflect

import (
	"reflect"
	"testing"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

func Test_Demo(t *testing.T) {
	reflect.ValueOf(&User{}).Elem()
}
