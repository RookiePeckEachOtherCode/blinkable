package model

import "time"

// User struct  
type User struct {
	ID              int64        `json:"id,omitempty"`
	CreateAt        time.Time    `json:"create_at,omitempty"`
	UpdateAt        time.Time    `json:"update_at,omitempty"`
	Username        string       `json:"username,omitempty" gorm:"index:idx_username,unique"` //用户名
	Password        string       `json:"password,omitempty"`                                  //密码
	Avatar          string       `json:"avatar,omitempty"`                                    //头像
	ArticlesNum     int32        `json:"articles_num,omitempty"`                              //文章数
	Signature       string       `json:"signature,omitempty"`                                 //个性签名
	Experience      int32        `json:"experience,omitempty"`                                //经验
	BackgroundImage string       `json:"background_image,omitempty"`                          //背景
	Level           int32        `json:"level,omitempty"`                                     //等级
	Title           string       `json:"title,omitempty"`
	GithubUrl       string       `json:"github_url,omitempty"`
	LikeNum         int64        `json:"like_num,omitempty"`
	Guestbooks      []*Guestbook `json:"guestbooks"`
	IsAdmin         bool         `json:"is_admin"`
}

type Guestbook struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_at,omitempty"`
	Context    string    `json:"context"`
	FromUserID int64     `json:"from_user_id"`
	UserID     int64     `json:"user_id"`
}
