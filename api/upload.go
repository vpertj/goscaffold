package api

// import (
// 	"goscaffold/model/web"
// 	"goscaffold/routes"
// 	"goscaffold/service"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// )

// // 初始化路由
// func init() {
// 	upload := new(uploadApi)
// 	routes.GetV1Router().Post("/UploadFiles", upload.UploadFiles)
// }

// type uploadApi struct {
// }

// // UploadFiles 文件上传接口
// // @Summary     文件上传接口
// // @Description 文件上传接口
// // @Tags        Upload
// // @Accept      mp4
// // @Produce     json
// // @Param       file  formData file true "mp4"
// // @Success     200    {string} string
// // @Router      /UploadFiles [post]
// func (api *uploadApi) UploadFiles(c *fiber.Ctx) error {
// 	c.Request().Header.Set("content-Type", "video/mp4")
// 	// 获取上传文件
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return web.OutState(c, web.Fail, err.Error())
// 	}
// 	dst := "./upload/video/" + file.Filename
// 	// 上传文件至指定的完整文件路径
// 	c.SaveFile(file, dst)
// 	log.Println(1)
// 	ctx := c.Context()
// 	// 打开上传文件
// 	srcFile, err := file.Open()
// 	if err != nil {
// 		return web.OutState(c, web.Fail, err.Error())
// 	}
// 	log.Println(2)
// 	defer srcFile.Close()
// 	if err != nil {
// 		return web.OutState(c, web.Fail, err.Error())
// 	}

// 	data, err := new(service.UploadService).UploadFile(ctx, dst)
// 	if err != nil {
// 		return web.OutState(c, web.Fail, err.Error())
// 	}
// 	return web.OutResponse(c, web.Success, "ok", data)

// }
