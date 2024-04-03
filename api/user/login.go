package api

import (
	"goscaffold/model/web"
	"goscaffold/routes"
	service "goscaffold/service/user"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/vpertj/gotool"
)

// 初始化路由
func init() {
	user := new(userLogin)
	routes.GetV1Router().Post("/user/login", user.login) //用户登录

}

type userLogin struct {
}

// login 登录接口
// @Summary      用户登录接口
// @Description  用户登录接口
// @Tags          user 用户管理
// @Param	usertel		formData 	string		true	"手机号"
// @Success      200  {string} string
// @Router	/user/login [post]
func (api *userLogin) login(c *fiber.Ctx) error {

	id := gotool.ToInt(c.FormValue("usertel"))
	slog.Info("sssss")
	data, err := new(service.LoginService).Login(id)

	if err != nil {
		return web.OutState(c, web.Fail, err.Error())

	}
	return web.OutResponse(c, web.Success, "ok", data)

}
