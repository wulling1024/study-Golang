package string

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {

	imgUrl := "https://img.ipcfun.com/uploads/ishoulu/pic/2013/05/9215193abd60b5ff099795216.jpg"
	//获取远端图片
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()

	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)

	imageBase64 := base64.StdEncoding.EncodeToString(data)
	fmt.Println("base64", imageBase64)
	fmt.Println(len(imageBase64))

	photo, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	fmt.Println(photo)
}
