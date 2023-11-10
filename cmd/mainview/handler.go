package main

import (
	"blinkable/kitex_gen/Mainview"
	"context"
)

// MainviewServiceImpl implements the last service interface defined in the IDL.
type MainviewServiceImpl struct{}

// LikeAction implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) LikeAction(ctx context.Context, req *mainview.LikeRequest) (resp *mainview.LikeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMainview implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) GetMainview(ctx context.Context, req *mainview.GetMainvewRequest) (resp *mainview.GetMainviewResponse, err error) {
	// TODO: Your code here...
	return
}

// AddGuestbook implements the MainviewServiceImpl interface.
func (s *MainviewServiceImpl) AddGuestbook(ctx context.Context, req *mainview.AddGuestbookRequest) (resp *mainview.AddGuestbookResponse, err error) {
	// TODO: Your code here...
	return
}
