package model

import "time"

type Article struct {
	ID         int64      `json:"id,omitempty"`
	CreateTime time.Time  `json:"create_at,omitempty"`
	Context    string     `json:"context"`
	CreaterID  int64      `json:"creater_id"`
	Comments   []*Comment `json:"comments"`
}
