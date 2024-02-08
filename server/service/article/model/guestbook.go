package model

import "time"

type Guestbook struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_at,omitempty"`
	Context    string    `json:"context"`
	FromUserID int64     `json:"from_user_id"`
	UserID     int64     `json:"user_id"`
}
