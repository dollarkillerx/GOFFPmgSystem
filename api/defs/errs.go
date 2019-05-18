package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{400,Err{"Request body is not correct","001"}}
	ErrorNotAuthUser = ErroResponse{401,Err{"User authentication failed","002"}}
	ErrorDBError = ErroResponse{500,Err{"DB ops failed","003"}}
	ErrorInternalFaults = ErroResponse{500,Err{"Internal service error","004"}}
)