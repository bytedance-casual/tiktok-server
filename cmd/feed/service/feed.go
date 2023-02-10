package service

import (
	"context"
	"tiktok-server/cmd/feed/dal/db"
	"tiktok-server/cmd/feed/rpc"
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

	userIdList := make([]int64, len(videos))
	for i, video := range videos {
		userIdList[i] = video.AuthorId
	}

	resp, err := rpc.MGetUsers(s.ctx, &user.UsersMGetRequest{UserId: userId, UserIdList: userIdList})
	if err != nil {
		return nil, 0, err
	}

	users := resp.Users
	fVideos := make([]*feed.Video, len(videos))
	for i, video := range videos {
		fVideos[i] = &feed.Video{
			Id:            int64(video.ID),
			Author:        users[i],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
		}
	}

	nextTime := videos[len(videos)-1].UpdatedAt.UnixMilli()
	return fVideos, nextTime, nil
}
