package api

const (
	Success         = 0    // 成功
	ErrParamInvalid = 1001 // 参数错误
)

// ErrMessage. err message
var ErrMessage = map[int]string{ // err message
	Success:         "请求成功",   // msg
	ErrParamInvalid: "请求参数错误", // param invalid
}
