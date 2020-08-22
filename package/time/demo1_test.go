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
	//
	//nTime := time.Now()
	//yesTime := nTime.AddDate(0,0,-1)
	//fmt.Println(yesTime.Day())

	yesTime := time.Now().AddDate(0, 0, -1)
	startTime := time.Date(yesTime.Year(), yesTime.Month(), yesTime.Day(), 2, 0, 0, 0, time.Local).Unix()
	fmt.Println(startTime)

	fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 2, 0, 0, 0, time.Local).Unix())
}
