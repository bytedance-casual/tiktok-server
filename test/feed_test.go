package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/feed"
)

// go test test2json -test.paniconexit0 -test.v
// go tool test2json -test.v -test.paniconexit0 -test.run ^\QTestFeed\E$

func TestFeed(t *testing.T) {
	resp, err := rpc.Feed(ctx, &feed.FeedRequest{
		LatestTime: nil,
		Token:      nil,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
