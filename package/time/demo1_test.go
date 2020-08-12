package time

import (
	"fmt"
	"testing"
	"time"
)

func Test_Unix(t *testing.T){
	fmt.Println(time.Now().Unix() - 3600*24)
	//sum, _ := strconv.ParseFloat(strconv.FormatInt(100, 10), 64)
	fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 3, 0, 0, 0, time.Local).Unix())
}
