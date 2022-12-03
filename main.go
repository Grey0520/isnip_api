package main

import (
	"fmt"

	"github.com/Grey0520/isnip_api/controller"
	"github.com/Grey0520/isnip_api/dao/mysql"
	"github.com/Grey0520/isnip_api/logger"
	"github.com/Grey0520/isnip_api/routers"
	"github.com/Grey0520/isnip_api/settings"
	"github.com/Grey0520/isnip_api/utils/snowflake"
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
	// mysql的连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// 翻译
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator Trans failed,err:%v\n", err)
		return
	}
	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
