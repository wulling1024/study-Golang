package time

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Unix(t *testing.T){
	fmt.Println(time.Now().Unix() - 3600*24)
	sum, _ := strconv.ParseFloat(strconv.FormatInt(100, 10), 64)
	fmt.Println(sum)
}
