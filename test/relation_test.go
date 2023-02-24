package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/relation"
)

func TestActionRelation1(t *testing.T) {
	doActionRelation1(t)
}

func TestActionRelation2(t *testing.T) {
	doActionRelation2(t)
}

func TestListFollowRelation(t *testing.T) {
	doListFollowRelation(t)
}

func TestListFollowerRelation(t *testing.T) {
	doListFollowerRelation(t)
}

func TestListFriendRelation(t *testing.T) {
	doListFriendRelation(t)
}

func TestMCheckFollowRelation(t *testing.T) {
	doMCheckFollowRelation(t)
}

func TestMCountRelation(t *testing.T) {
	doMCountRelation(t)
}

func BenchmarkActionRelation1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionRelation1(b)
		}
	})
}

func BenchmarkActionRelation2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionRelation2(b)
		}
	})
}

func BenchmarkListFollowRelation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doListFollowRelation(b)
		}
	})
}

func BenchmarkListFollowerRelation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doListFollowerRelation(b)
		}
	})
}

func BenchmarkListFriendRelation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doListFriendRelation(b)
		}
	})
}

func BenchmarkMCheckFollowRelation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMCheckFollowRelation(b)
		}
	})
}

func BenchmarkMCountRelation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMCountRelation(b)
		}
	})
}

func doActionRelation1(t assert.TestingT) {
	resp, err := rpc.ActionRelation(ctx, &relation.RelationActionRequest{
		Token:      TOKEN,
		ToUserId:   20,
		ActionType: 1,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doActionRelation2(t assert.TestingT) {
	resp, err := rpc.ActionRelation(ctx, &relation.RelationActionRequest{
		Token:      TOKEN,
		ToUserId:   20,
		ActionType: 2,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doListFollowRelation(t assert.TestingT) {
	resp, err := rpc.ListFollowRelation(ctx, &relation.RelationFollowListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doListFollowerRelation(t assert.TestingT) {
	resp, err := rpc.ListFollowerRelation(ctx, &relation.RelationFollowerListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doListFriendRelation(t assert.TestingT) {
	resp, err := rpc.ListFriendRelation(ctx, &relation.RelationFriendListRequest{
		UserId: 2,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMCheckFollowRelation(t assert.TestingT) {
	resp, err := rpc.MCheckFollowRelation(ctx, &relation.MCheckFollowRelationRequest{
		UserId:     2,
		UserIdList: []int64{3},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMCountRelation(t assert.TestingT) {
	resp, err := rpc.MCountRelation(ctx, &relation.MCountRelationRequest{
		UserIdList: []int64{2},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
