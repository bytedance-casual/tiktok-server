package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/favorite"
)

func TestActionFavorite1(t *testing.T) {
	doActionFavorite1(t)
}

func TestActionFavorite2(t *testing.T) {
	doActionFavorite2(t)
}

func TestListFavorite(t *testing.T) {
	doListFavorite(t)
}

func TestMCheckFavorite(t *testing.T) {
	doMCheckFavorite(t)
}

func TestMCountFavorite(t *testing.T) {
	doMCountFavorite(t)
}

func BenchmarkActionFavorite1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionFavorite1(b)
		}
	})
}

func BenchmarkActionFavorite2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionFavorite2(b)
		}
	})
}

func BenchmarkListFavorite(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doListFavorite(b)
		}
	})
}

func BenchmarkMCheckFavorite(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMCheckFavorite(b)
		}
	})
}

func BenchmarkMCountFavorite(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMCountFavorite(b)
		}
	})
}

func doActionFavorite1(t assert.TestingT) {
	resp, err := rpc.ActionFavorite(ctx, &favorite.FavoriteActionRequest{
		Token:      TOKEN,
		VideoId:    7,
		ActionType: 1,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doActionFavorite2(t assert.TestingT) {
	resp, err := rpc.ActionFavorite(ctx, &favorite.FavoriteActionRequest{
		Token:      TOKEN,
		VideoId:    7,
		ActionType: 2,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doListFavorite(t assert.TestingT) {
	resp, err := rpc.ListFavorite(ctx, &favorite.FavoriteListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMCheckFavorite(t assert.TestingT) {
	resp, err := rpc.MCheckFavorite(ctx, &favorite.MCheckFavoriteRequest{
		UserId:      2,
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMCountFavorite(t assert.TestingT) {
	resp, err := rpc.MCountFavorite(ctx, &favorite.MCountVideoFavoriteRequest{
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
