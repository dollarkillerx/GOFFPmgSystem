package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}


type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400,Error:Err{Error:"Request body is not correct",ErrorCode:"001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC:401,Error:Err{Error:"User authentication failed.",ErrorCode:"002"}}
)