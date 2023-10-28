package minio_test

import (
	"blinkable/pkg/minio"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	err := minio.CreateBucket("test")
	if err != nil {
		t.Error(err)
	}
}
