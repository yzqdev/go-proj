package util

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Code    int         `json:"code"`
}

// JSON 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func JSON(c *fiber.Ctx, code int, msg string, data ...interface{}) error {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	var success = false
	if code == 200 {
		success = true
	}
	return c.JSON(ApiResponse{
		Data:    responseData,
		Msg:     msg,
		Success: success,
		Code:    code,
	})

}
