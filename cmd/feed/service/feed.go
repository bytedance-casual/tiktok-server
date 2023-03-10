package service

import (
	"context"
	"tiktok-server/cmd/feed/dal/db"
	"tiktok-server/cmd/feed/rpc"
	"tiktok-server/internal/utils"
	"tiktok-server/kitex_gen/comment"
	"tiktok-server/kitex_gen/favorite"
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/user"
	"time"
)

type GetFeedService struct {
	ctx context.Context
}

func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{
		ctx: ctx,
	}
}

func (s *GetFeedService) GetFeedInfo(req *feed.FeedRequest, userId int64) ([]*feed.Video, int64, error) {
	videoList := make([]*feed.Video, 0)
	videos, err := db.MGetVideosByTime(s.ctx, userId, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}

	if len(videos) == 0 {
		nextTime := time.Now().UnixMilli()
		return videoList, nextTime, nil
	}

	videoIdList := make([]int64, len(videos))
	userIdList := make([]int64, len(videos))
	favoriteList := make([]bool, len(videos))
	for i, video := range videos {
		userIdList[i] = video.AuthorId
		videoIdList[i] = int64(video.ID)
	}

	// TODO 整体事务
	resp1, err := rpc.MGetUsers(s.ctx, &user.UsersMGetRequest{UserId: userId, UserIdList: userIdList})
	if err != nil {
		return nil, 0, err
	}
	if userId != 0 {
		resp2, err := rpc.MCheckFavorite(s.ctx, &favorite.MCheckFavoriteRequest{UserId: userId, VideoIdList: videoIdList})
		if err != nil {
			return nil, 0, err
		}
		favoriteList = resp2.ExistList
	}
	resp3, err := rpc.MCountFavorite(s.ctx, &favorite.MCountVideoFavoriteRequest{VideoIdList: videoIdList})
	if err != nil {
		return nil, 0, err
	}
	resp4, err := rpc.MCountComment(s.ctx, &comment.MCountVideoCommentRequest{VideoIdList: videoIdList})
	if err != nil {
		return nil, 0, err
	}

	userMap := resp1.Users

	favoriteCountList := resp3.CountList
	commentCountList := resp4.CountList
	for i, video := range videos {
		videoList = append(videoList, &feed.Video{
			Id:            int64(video.ID),
			Author:        userMap[userIdList[i]],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			IsFavorite:    utils.Ternary[bool](userId != 0, favoriteList[i], false),
			FavoriteCount: favoriteCountList[i],
			CommentCount:  commentCountList[i],
			Title:         video.Title,
		})
	}

	nextTime := videos[len(videos)-1].UpdatedAt.UnixMilli()
	return videoList, nextTime, nil
}
