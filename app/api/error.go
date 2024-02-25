package api

const (
	Success         = 0    // 成功
	ErrParamInvalid = 1001 // 参数错误
)

// ErrMessage. err message
var ErrMessage = map[int]string{
	Success:         "请求成功",
	ErrParamInvalid: "请求参数错误",
}
