package toml

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg     *tomlConfig
	once    sync.Once
	cfgLock sync.Mutex
)

func Config() *tomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("config.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

// 热更新配置信息
func config2() *tomlConfig {
	once.Do(ReloadConfig)
	cfgLock.Lock()
	defer cfgLock.Unlock()
	return cfg
}

func ReloadConfig() {
	filePath, err := filepath.Abs("config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse toml file once. filePath: %s\n", filePath)
	config := new(tomlConfig)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}
	cfgLock.Lock()
	defer cfgLock.Unlock()
	cfg = config
}
