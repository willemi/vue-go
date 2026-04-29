// Package model 定义数据模型、请求/响应数据结构
package model

// Response 统一的 API 响应结构
// 所有 API 接口均返回此格式，便于前端统一处理
// - Code: 状态码，200 表示成功，其他值表示错误
// - Message: 提示信息，成功时通常为 "success"
// - Data: 响应数据，可为任意类型，omitempty 允许空值时省略
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse 构造成功的响应
// 状态码固定为 200，message 固定为 "success"
func SuccessResponse(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// ErrorResponse 构造错误的响应
// 通常与 HTTP 状态码配合使用，如 400 表示参数错误，401 表示未认证
func ErrorResponse(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}