package main

import (
	"context"
	"tiktok-server/cmd/message/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// ChatMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ChatMessage(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.ToUserId <= 0 {
		resp = &message.MessageChatResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		errMsg := err.Error()
		resp = &message.MessageChatResponse{StatusCode: erren.AuthorizationFailedErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	messages, err := service.NewGetMessageService(ctx).GetMessage(req, claims.ID)
	if err != nil {
		errStr := err.Error()
		resp = &message.MessageChatResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &message.MessageChatResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, MessageList: messages}
	return resp, nil
}

// ActionMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ActionMessage(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.ToUserId <= 0 || len(req.Content) == 0 {
		resp = &message.MessageActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		errMsg := err.Error()
		resp = &message.MessageActionResponse{StatusCode: erren.AuthorizationFailedErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	if req.ActionType == 1 {
		err = service.NewCreateMessageService(ctx).CreateMessage(req, claims.ID)
	} else {
		err = erren.TypeNotSupportErr
	}
	if err != nil {
		errStr := err.Error()
		resp = &message.MessageActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &message.MessageActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, nil
}
