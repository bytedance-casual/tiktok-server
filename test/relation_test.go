package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/relation"
)

func TestActionRelation1(t *testing.T) {
	resp, err := rpc.ActionRelation(ctx, &relation.RelationActionRequest{
		Token:      TOKEN,
		ToUserId:   20,
		ActionType: 1,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestActionRelation2(t *testing.T) {
	resp, err := rpc.ActionRelation(ctx, &relation.RelationActionRequest{
		Token:      TOKEN,
		ToUserId:   20,
		ActionType: 2,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestListFollowRelation(t *testing.T) {
	resp, err := rpc.ListFollowRelation(ctx, &relation.RelationFollowListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestListFollowerRelation(t *testing.T) {
	resp, err := rpc.ListFollowerRelation(ctx, &relation.RelationFollowerListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestListFriendRelation(t *testing.T) {
	resp, err := rpc.ListFriendRelation(ctx, &relation.RelationFriendListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMCheckFollowRelation(t *testing.T) {
	resp, err := rpc.MCheckFollowRelation(ctx, &relation.MCheckFollowRelationRequest{
		UserId:     2,
		UserIdList: []int64{3},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMCountRelation(t *testing.T) {
	resp, err := rpc.MCountRelation(ctx, &relation.MCountRelationRequest{
		UserIdList: []int64{2},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
