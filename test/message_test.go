package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/message"
)

func TestChatMessage(t *testing.T) {
	resp, err := rpc.ChatMessage(ctx, &message.MessageChatRequest{
		Token:      TOKEN,
		ToUserId:   3,
		PreMsgTime: 0,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestActionMessage(t *testing.T) {
	resp, err := rpc.ActionMessage(ctx, &message.MessageActionRequest{
		Token:      TOKEN,
		ToUserId:   3,
		ActionType: 1,
		Content:    "TestActionMessage测试消息",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMGetLatestMessage(t *testing.T) {
	resp, err := rpc.MGetLatestMessage(ctx, &message.MGetLatestMessageRequest{
		UserId:       2,
		FriendIdList: []int64{3},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
