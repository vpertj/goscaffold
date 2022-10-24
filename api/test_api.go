package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"goscaffold/model"
	"goscaffold/model/db"
	"goscaffold/model/web"
	"goscaffold/routes"
	"goscaffold/service"
)

//初始化路由
func init() {
	test := new(testApi)
	routes.GetV1Router().Get("/test", test.test)
	routes.GetV1Router().Post("/test1", test.test1)
}

type testApi struct {
}

// test 一个测试接口
// @Summary     测试接口
// @Description 用来测试的接口
// @Tags        demo
// @Success     200 {string} string
// @Router      /test [get]
func (api *testApi) test(c *fiber.Ctx) error {
	data, err := new(service.TestService).Test()
	if err != nil {
		return web.OutState(c, web.Fail, err.Error())
	}
	return web.OutResponse(c, web.Success, "ok", data)

}

//func (*testApi) testAdmin(c *fiber.Ctx) error {
//	ts := new(service.TestService)
//	ts.Test()
//	return c.SendString("你瞅啥?")
//}
//
//func (*testApi) testUser(c *fiber.Ctx) error {
//	ts := new(service.TestService)
//	ts.Test()
//	return c.SendString("你瞅啥?ds")
//}

// test1 一个测试接口
// @Summary     测试接口1
// @Description 用来测试的接口1
// @Tags        demo
// @Param       name             formData string true   "你的名字"
// @Param       usertoken header                 string  false "用户Token"
// @Success     200              {object} model.TestData
// @Router      /test1 [post]
func (*testApi) test1(c *fiber.Ctx) error {
	orm := db.GetOrm()
	count, err := orm.Count(new(db.Test))
	if err != nil {
		logrus.Error(err)
	}
	return c.JSON(&model.TestData{
		Name:  c.FormValue("name", "无名人士"),
		Count: int(count),
	})
}
