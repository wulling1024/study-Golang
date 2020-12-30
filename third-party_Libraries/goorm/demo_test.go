package goorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
	"time"
)
type StatusClear struct {
	Status int32 `gorm:"status"`
	Sum    int64 `gorm:"sum"`
}

func Test_Group(t *testing.T){
	db, err := gorm.Open("mysql","airpay_test_svr:yequnfeng1978!@#@(203.116.180.207:40080)/beepay_settlement_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("db connect fail")
	}
	defer db.Close()
	startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 3, 0, 0, 0, time.Local).Unix()
	fmt.Printf("startTime is %v \n", startTime)
	endTime := startTime - 3600*24
	fmt.Printf("endTime is %v \n", endTime)
	var records []StatusClear
	if err := db.Table("account_book_tab").Select("`status`, count(`status`) as sum").Where("pre_settle_time = ?", startTime).Group("`status`").Scan(&records).Error; err != nil {

	}
	fmt.Printf("len : %v \n", len(records))
	for _, record := range records{
		fmt.Printf("%v : %v \n", record.Status, record.Sum)
	}

}
