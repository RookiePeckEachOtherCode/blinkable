package response

type BaseResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func BuildBase(statusCode int32, statusMsg string) *BaseResponse {
	return &BaseResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
