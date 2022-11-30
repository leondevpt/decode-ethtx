package handler

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func RespOK(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
}

func BuildResponse(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
