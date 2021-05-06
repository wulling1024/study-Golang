package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

type tomlConfig struct {
	Title string
	Owner ownerInfo
	DB database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org string `toml:"organization"`
	Bio string
	DOB time.Time
}

type database struct {
	Server string
	Ports []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data [][]interface{}
	Hosts []string
}

func main() {
	var config tomlConfig
	// 参数 fpath 的地址是项目的根目录开始的
	if _, err := toml.DecodeFile("third-party_Libraries/toml/example.toml", &config); err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}
