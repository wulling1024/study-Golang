package encoding

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

// 以Json格式，写信息到文件中
func Test_Encoder(t *testing.T) {
	// 内容信息
	info := []Website{
		{"Golang",
			"http://c.biancheng.net/golang/",
			[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"},
		},
		{"Java",
			"http://c.biancheng.net/java/",
			[]string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"},
		},
	}

	// 创建文件
	file, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	// 创建Json编码器
	jsonEncoder := json.NewEncoder(file)

	// 写入文件
	err = jsonEncoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func Test_Decoder(t *testing.T) {
	file, err := os.Open("./info.json")
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

func TestJSONFile(t *testing.T) {

}
