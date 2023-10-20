package response

type User struct {
	BaseResponse
	UserId int32  `json:"user_id"`
	Token  string `json:"token"`
}
