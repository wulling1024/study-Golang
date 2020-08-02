package toml

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T){
	//fmt.Println(Config().infoDate[1].appleInfo.age)
	fmt.Println(Config().infoDate[0].appleInfo.age)
	//fmt.Println(Config().infoDate[1].ruanjianInfo.name)
	fmt.Println(Config().infoDate[0].ruanjianInfo.name)
}