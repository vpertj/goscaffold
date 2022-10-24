package web

import (
	"github.com/gofiber/fiber/v2"
)

// WebResponse 给web的相应
type WebResponse struct {
	State   int         `json:"state"`          // 当前相应是否正常完成，返回1代表成功，返回0代表（业务）不成功，返回-1代表（代码）出错
	Message string      `json:"message"`        // 不成功时必须存入相应原因
	Data    interface{} `json:"data,omitempty"` // 获取数据的接口需要在正常的情况下填充次字段
}

// ResponseStatus web相应状态
type ResponseStatus int

const (
	// Error 出错
	Error ResponseStatus = -1
	// Fail 失败
	Fail ResponseStatus = 0
	// Success 成功
	Success ResponseStatus = 1
	// Sign 未登录状态
	Sign ResponseStatus = -100
)

// OutResponse 输出web响应 ！！！输出内容后会关闭response！！！
// c *fiber.Ctx 控制器的请求
// status ResponseStatus 接口相应的状态
// message string 状态不成功时需要返回的信息
// data ...interface{} 传入要返回的数据和相应的操作数据
func OutResponse(c *fiber.Ctx, status ResponseStatus, message string, data ...interface{}) error {
	response := WebResponse{
		State:   int(status),
		Message: message,
	}
	if data != nil && len(data) > 0 {
		for _, d := range data {
			response.Data = d
		}
	}
	return c.JSON(response)
}

// OutData 输出成功状态的数据 ！！！输出内容后会关闭response！！！
// request *ghttp.Request 控制器的请求
// data ...interface{} 传入要返回的数据和相应的操作数据
func OutData(c *fiber.Ctx, data ...interface{}) error {
	return OutResponse(c, Success, "", data...)
}

// OutState 只输出状态和message  ！！！输出内容后会关闭response！！！
// request *ghttp.Request 控制器的请求
// status ResponseStatus 接口相应的状态
// message string 状态不成功时需要返回的信息
func OutState(c *fiber.Ctx, status ResponseStatus, message string) error {
	return OutResponse(c, status, message)
}
