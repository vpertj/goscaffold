# goscaffold
使用组件
fiber web框架
https://docs.gofiber.io/

https://learnku.com/docs/gofiber/2.x

swag:用户生成swagger文档
https://github.com/swaggo/swag 文档编写说明

go install github.com/swaggo/swag/cmd/swag@v1.8
go install github.com/swaggo/swag/cmd/swag@latest

air: 用于热重载运行
go install github.com/cosmtrek/air@latest

数据库访问组件 xorm db.GetOrm()
https://xorm.io/zh/

日志组件 logrus
https://github.com/sirupsen/logrus


数据库结构生成,依赖cgo需要安装gcc环境
go install xorm.io/reverse@latest
生成命令:
reverse -f mydbtables.yml


配置文件访问 config.GetConfig()
https://github.com/jinzhu/configor

lancet 工具库
https://github.com/duke-git/lancet

resty http请求库
https://github.com/go-resty/resty