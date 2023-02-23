package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/publish"
)

func TestActionPublish(t *testing.T) {
	filePath := "/home/illtamer/Code/golang/goland/github/tiktok-server/cmd/publish/test/bear.mp4"
	bytes, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	resp, err := rpc.ActionPublish(ctx, &publish.PublishActionRequest{
		Token: TOKEN,
		Data:  bytes,
		Title: "TestActionPublish测试视频",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestListPublish(t *testing.T) {
	resp, err := rpc.ListPublish(ctx, &publish.PublishListRequest{
		UserId: 3,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMGetVideos(t *testing.T) {
	resp, err := rpc.MGetVideos(ctx, &publish.VideosMGetRequest{
		UserId:      3,
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
