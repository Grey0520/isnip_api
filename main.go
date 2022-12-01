package main

import (
    "fmt"

    "github.com/Grey0520/isnip_api/settings"
    "github.com/Grey0520/isnip_api/logger"
)
func main() {
    // 从`conf/conf.yaml`加载配置信息
    if err := settings.Init(); err != nil {
        fmt.Printf("load config failed, err:%v\n", err)
        return
    }
    // 自定义logger
    if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
        fmt.Printf("init logger failed, err:%v\n", err)
        return
    }
}