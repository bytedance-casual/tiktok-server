package main

import (
	"context"
	"tiktok-server/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// ChatMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ChatMessage(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// ActionMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ActionMessage(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...
	return
}
