package api

// ApiResponse formate return info
type ApiResponse struct {
	Data interface{} `json:"data"` // data info
	Msg  string      `json:"msg"`  // success or err msg
	Code int         `json:"code"` // code
	Next interface{} `json:"next"` // next info
}

// InterfaceResponse formate interface info
type InterfaceResponse struct {
	Data interface{} `json:"data"` // data info
	Msg  string      `json:"msg"`  // success or err msg
	Code int         `json:"code"` // code
	Next interface{} `json:"next"` // next info
}

// InterfaceResponse formate interface info
type InterfaceResponseInfo struct {
	Data interface{} `json:"data"` // data info
	Msg  string      `json:"msg"`  // success or err msg
	Code int         `json:"code"` // code
	Next interface{} `json:"next"` // next info
}
