package model

import "time"

type Comment struct {
	ID         int64     `json:"id,omitempty"`
	Context    string    `json:"context,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UserID     int64     `json:"user_id"`
	ArticleId  int64     `json:"article_Id"`
}
