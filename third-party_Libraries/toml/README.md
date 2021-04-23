具体实现的功能：

    1、读取TOML配置文件的信息
    2、热更新配置文件（也就是在不关停应用程序的情况下，实现对配置信息的更新）
```    
├── README.md       # README.md的介绍说明
├── config.go       # 单例化配置信息的初始化，并实现`POSIX信号`的热更新
├── config.toml     # 配置信息
├── toml_config.go  # 配置信息与结构体的映射
└── toml_test.go    # 测试程序
```
关于**POSIX信号**做个简单的说明：

![](https://cdn.jsdelivr.net/gh/wulling1024/Picture-bed@master/img20200620200724144621.png)