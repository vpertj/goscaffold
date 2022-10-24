package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"goscaffold/docs"
	"goscaffold/model/config"
	"runtime"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

var app *fiber.App

func init() {
	app = fiber.New()
	app.Use(cors.New())
	app.Use(compress.New())
	// 恐慌恢复中间件
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
			StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
				buf := make([]byte, 1024)
				buf = buf[:runtime.Stack(buf, false)]
				logrus.Error(fmt.Sprintf("panic: %v\n%s\n", e, buf))
			},
		},
	))
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))
	// 配置swagger
	// 修改 swagger 信息
	docs.SwaggerInfo.Title = "PMApi"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "工友帮智慧用工项目管理系统接口."
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if config.GetConfig().Server.Debug {
		// 注册swagger
		app.Get("/swagger/*", swagger.HandlerDefault)
	}
	// 设置接口地址
	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		//logrus.Debug("v1 😀")
		// 前置方法，/api/v1下面所有接口的前置执行方法
		return c.Next()
	})
	//拦截
	v1.Use(middlewareLog, middlewareAuth)
}

func GetV1Router() fiber.Router {
	return app.Group("/api").Group("/v1")
}

func New() *fiber.App {
	return app
}
