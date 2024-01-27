package model

import "time"

type Guestbook struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_at,omitempty"`
	Context    string    `json:"context"`
	FromuserID int64     `json:"fromuser_id"`
	UserID     int64     `json:"user_id"`
}
