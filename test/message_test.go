package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/message"
)

func TestChatMessage(t *testing.T) {
	doChatMessage(t)
}

func TestActionMessage(t *testing.T) {
	doActionMessage(t)
}

func TestMGetLatestMessage(t *testing.T) {
	doMGetLatestMessage(t)
}

func BenchmarkChatMessage(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doChatMessage(b)
		}
	})
}

func BenchmarkActionMessage(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doActionMessage(b)
		}
	})
}

func BenchmarkMGetLatestMessage(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMGetLatestMessage(b)
		}
	})
}

func doChatMessage(t assert.TestingT) {
	resp, err := rpc.ChatMessage(ctx, &message.MessageChatRequest{
		Token:      TOKEN,
		ToUserId:   3,
		PreMsgTime: 0,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doActionMessage(t assert.TestingT) {
	resp, err := rpc.ActionMessage(ctx, &message.MessageActionRequest{
		Token:      TOKEN,
		ToUserId:   3,
		ActionType: 1,
		Content:    "TestActionMessage测试消息",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMGetLatestMessage(t assert.TestingT) {
	resp, err := rpc.MGetLatestMessage(ctx, &message.MGetLatestMessageRequest{
		UserId:       2,
		FriendIdList: []int64{3},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
