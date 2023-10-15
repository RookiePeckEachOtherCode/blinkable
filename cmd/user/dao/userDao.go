package dao

import (
	"blinkable/dal/db"
	"context"
)

var DB = db.DB

func NewUser(ctx context.Context, username, password string) error {

	return nil
}
