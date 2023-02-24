package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/feed"
)

func TestFeed(t *testing.T) {
	doFeed(t)
}

func BenchmarkFeed(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doFeed(b)
		}
	})
}

func doFeed(t assert.TestingT) {
	resp, err := rpc.Feed(ctx, &feed.FeedRequest{
		LatestTime: nil,
		Token:      nil,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
