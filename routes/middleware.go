package routes

import (
	"goscaffold/model/web"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	//添加排除列表
	AddAuthExclusionRole("/user/login")
	AddAuthExclusionRole("/user/register")
	AddAuthExclusionRole("/user/captcha")
	AddAuthExclusionRole("/user/sendsms")
	AddAuthExclusionRole("/user/resetpassword")

	// /api/v1/register/:type/      [:len(url)]
}

// _URI_PREFIX 路径前缀
const (
	_URI_PREFIX     = "/api/v1"
	_URI_PREFIX_LEN = len(_URI_PREFIX)
)

// _AuthExclusionList 排除列表
var _AuthExclusionList = map[string]int{}

// AddAuthExclusionRole 添加到排除列表
func AddAuthExclusionRole(path string) {
	_AuthExclusionList[strings.ToLower(path)] = len(path)
}

// middlewareLog 记录操作日志
func middlewareLog(c *fiber.Ctx) error {
	//uri := strings.ToLower(c.Path()[_URI_PREFIX_LEN:])
	//user, _ := new(u.UserService).GetLoginUserID(c)
	//body := string(c.Body())
	//if c.Method() == "GET" {
	//	body = c.OriginalURL()
	//}
	////添加到日志表 pmlogs
	//go new(service.LogsService).InsertLogs(int(user.Id), user.Name, uri, body)
	return c.Next()
}

// middlewareAuth
func middlewareAuth(c *fiber.Ctx) error {
	uri := strings.ToLower(c.Path()[_URI_PREFIX_LEN:])
	// 如果是排除目录则直接放行
	for path, pathLen := range _AuthExclusionList {
		if len(uri) < pathLen {
			continue
		}
		if uri[:pathLen] == path {
			logrus.Debug("😉 排除接口:", path)
			return c.Next()
		}
	}
	// 放行操作
	return c.Next()
}

// _intercept 拦截器
func _intercept(c *fiber.Ctx) error {
	return web.OutState(c, web.Sign, "用户未登录?")
}
