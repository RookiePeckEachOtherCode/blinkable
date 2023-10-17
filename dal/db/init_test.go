package db_test

import (
	"blinkable/dal/db"
	"testing"
)

func TestDBInit(t *testing.T) {
	if db.GetIsInit() == false {
		t.Error("db init failed")
	}
}
