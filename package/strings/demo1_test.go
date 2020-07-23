package strings

import (
	"fmt"
	"testing"
	"time"
)

func Test_Sprintf(t *testing.T) {
	fmt.Println(fmt.Sprintf("%s_%s_%s", "AirPay_VersionNotes", "product", "2020"))
	start, _ := time.Parse("2006-01-02", "2020-03-30")
	now ,_ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	//month, _ := strconv.ParseFloat(beego.AppConfig.String("date"),64)
	fmt.Println(time.Since(start).Hours()/24)
	d := now.Sub(start)
	if d.Hours()/24 > 30*4 {
		fmt.Println(d.Hours()/24)
		fmt.Println("true")
	}else {
		fmt.Println(d.Hours()/24)
		fmt.Println("false")
	}
}
