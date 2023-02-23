package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/favorite"
)

func TestActionFavorite1(t *testing.T) {
	resp, err := rpc.ActionFavorite(ctx, &favorite.FavoriteActionRequest{
		Token:      TOKEN,
		VideoId:    7,
		ActionType: 1,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestActionFavorite2(t *testing.T) {
	resp, err := rpc.ActionFavorite(ctx, &favorite.FavoriteActionRequest{
		Token:      TOKEN,
		VideoId:    7,
		ActionType: 2,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestListFavorite(t *testing.T) {
	resp, err := rpc.ListFavorite(ctx, &favorite.FavoriteListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMCheckFavorite(t *testing.T) {
	resp, err := rpc.MCheckFavorite(ctx, &favorite.MCheckFavoriteRequest{
		UserId:      2,
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMCountFavorite(t *testing.T) {
	resp, err := rpc.MCountFavorite(ctx, &favorite.MCountVideoFavoriteRequest{
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
