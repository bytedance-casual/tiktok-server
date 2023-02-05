package service

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"os"
	"tiktok-server/cmd/publish/dal/db"
	"tiktok-server/cmd/publish/ffmpeg"
	"tiktok-server/cmd/publish/oss"
	"tiktok-server/internal/conf"
	"tiktok-server/kitex_gen/publish"
	"tiktok-server/kitex_gen/user"
)

type UploadVideoService struct {
	ctx context.Context
}

func NewUploadVideoService(ctx context.Context) *UploadVideoService {
	return &UploadVideoService{
		ctx: ctx,
	}
}

func (s *UploadVideoService) UploadVideo(req *publish.PublishActionRequest, user *user.User) error {
	uuidStr := uuid.NewV4().String()
	videoName := uuidStr + ".mp4"
	shotName := uuidStr + ".png"
	videoPath := conf.WebResourceFolder + videoName
	shotPath := conf.WebResourceFolder + shotName

	err := os.WriteFile(videoPath, req.Data, os.ModePerm)
	if err != nil {
		return err
	}
	err = ffmpeg.GetVideoShot(shotPath, videoPath)
	if err != nil {
		return err
	}

	err = oss.UploadFromPath(videoName, videoPath)
	if err != nil {
		return err
	}
	err = oss.UploadFromPath(shotName, shotPath)
	if err != nil {
		return err
	}
	go func() {
		_ = os.Remove(videoPath)
		_ = os.Remove(shotPath)
	}()

	publicURL := conf.Config.OssAliyun.PublicURL
	_, err = db.CreateVideo(&db.Video{
		AuthorId: user.Id,
		PlayUrl:  publicURL + videoName,
		CoverUrl: publicURL + shotName,
		Title:    req.Title,
	}, s.ctx)
	if err != nil {
		return err
	}
	return nil
}
