package model

// Response 通用响应结构体
type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码 0成功，其他值失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}
