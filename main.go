package main

import (
	"fmt"
	"goscaffold/model/config"

	_ "github.com/duke-git/lancet/v2/strutil"
	_ "github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	// 这个必须在第一行
	_ "goscaffold/boot"
	// 🚧注意引入顺序🚧
	// 1.先引入核心路由代码
	"goscaffold/routes"
	// 2.引入接口代码
	_ "goscaffold/api"
	_ "goscaffold/api/user"

	// 3.数据库初始化
	_ "goscaffold/model/db"
)

func main() {
	logrus.Info("接口开始启动")
	app := routes.New()
	logrus.Info(app.Listen(fmt.Sprintf(":%v", config.GetConfig().Server.Port)))

}
