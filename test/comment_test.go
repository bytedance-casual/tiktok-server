package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/comment"
)

var commentId int64

func TestActionComment1(t *testing.T) {
	doActionComment1(t)
}

func TestActionComment2(t *testing.T) {
	doActionComment2(t)
}

func TestListComment(t *testing.T) {
	doListComment(t)
}

func TestMCountVideoComment(t *testing.T) {
	doMCountVideoComment(t)
}

func BenchmarkActionComment1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionComment1(b)
		}
	})
}

func BenchmarkActionComment2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionComment2(b)
		}
	})
}

func BenchmarkListComment(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doListComment(b)
		}
	})
}

func BenchmarkMCountVideoComment(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMCountVideoComment(b)
		}
	})
}

func doActionComment1(t assert.TestingT) {
	content := "TestComment评论内容"
	resp, err := rpc.ActionComment(ctx, &comment.CommentActionRequest{
		Token:       TOKEN,
		VideoId:     7,
		ActionType:  1,
		CommentText: &content,
		CommentId:   nil,
	})
	assert.NoError(t, err)
	commentId = resp.Comment.Id
	fmt.Printf("%v", resp)
}

func doActionComment2(t assert.TestingT) {
	resp, err := rpc.ActionComment(ctx, &comment.CommentActionRequest{
		Token:       TOKEN,
		VideoId:     7,
		ActionType:  2,
		CommentText: nil,
		CommentId:   &commentId,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doListComment(t assert.TestingT) {
	resp, err := rpc.ListComment(ctx, &comment.CommentListRequest{
		Token:   TOKEN,
		VideoId: 7,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMCountVideoComment(t assert.TestingT) {
	resp, err := rpc.MCountComment(ctx, &comment.MCountVideoCommentRequest{
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
