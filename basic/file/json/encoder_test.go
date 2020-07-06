package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test_encoder(t *testing.T) {
	file, err := os.Open("./basic/file/json/info.json")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	var info []Website

	// Json解码器
	newDecoder := json.NewDecoder(file)
	err = newDecoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
