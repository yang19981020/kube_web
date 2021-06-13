package response

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(code int, message string, data ...interface{}) (j *JsonResponse) {
	j = new(JsonResponse)
	j.Code = code
	j.Data = data
	j.Message = message
	return
}
