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

func TestActionComment2(t *testing.T) {
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

func TestListComment(t *testing.T) {
	resp, err := rpc.ListComment(ctx, &comment.CommentListRequest{
		Token:   TOKEN,
		VideoId: 7,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMCountVideoComment(t *testing.T) {
	resp, err := rpc.MCountComment(ctx, &comment.MCountVideoCommentRequest{
		VideoIdList: []int64{7},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
