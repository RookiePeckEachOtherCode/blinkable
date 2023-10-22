package response

type User struct {
	BaseResponse
	UserId int32  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfo struct {
	BaseResponse
	UserName        string `json:"user_name"`
	UserId          int32  `json:"user_id"`
	Avatar          string `json:"avatar"`
	ArticlesNum     int32  `json:"articles_num"`
	Level           int32  `json:"level"`
	Signature       string `json:"signature"`
	Experience      int32  `json:"experience"`
	BackgroundImage string `json:"background_image"`
}
