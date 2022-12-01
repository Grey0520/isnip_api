package main

import (
    "fmt"

    "github.com/Grey0520/isnip_api/settings"
)
func main() {
    // 从`conf/conf.yaml`加载配置信息
    if err := settings.Init(); err != nil {
        fmt.Printf("load config failed, err:%v\n", err)
        return
    }
}