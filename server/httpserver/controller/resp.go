package controller

const (
	HttpCode = 200

	HttpCodeSucc = 200
	HttpCodeFail = 500
)

//返回响应
type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//返回响应 携带数据
type RespData struct {
	Resp
	Data interface{} `json:"data"`
}

//返回响应 列表数据
type RespList struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

//desc:新建一个不带数据的返回响应
func NewResp(code int, msg string) Resp {
	return Resp{code, msg}
}

//新建一个带数据的返回响应
func NewRespData(code int, msg string, data interface{}) RespData {
	return RespData{Resp{Code: code, Msg: msg}, data}
}

//新建一个带列表数据的返回响应
func NewRespList(code int, msg string, count int, data interface{}) RespList {
	resp := RespList{code, msg, count, data}
	return resp
}
