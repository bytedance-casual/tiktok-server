package main

import (
	"context"
	"tiktok-server/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// ActionPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ActionPublish(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// ListPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ListPublish(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
