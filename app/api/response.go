package api

// ApiResponse formate return info
type ApiResponse struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Next interface{} `json:"next"`
}
