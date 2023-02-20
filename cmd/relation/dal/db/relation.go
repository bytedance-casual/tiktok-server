package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FollowStatus bool

const (
	Normal   FollowStatus = true
	Canceled FollowStatus = false
)

type Follow struct {
	gorm.Model
	UserId         string       `json:"user_id"`
	FollowedUserId string       `json:"followed_user_id"`
	IsFollow       FollowStatus `json:"is_follow"`
}

type UserInfoResp struct {
	FollowCount   int64  `json:"follow_count"`                    // 关注总数
	FollowerCount int64  `json:"follower_count"`                  // 粉丝总数
	Id            int64  `json:"id"`                              // 用户id，和db.user不同
	IsFollow      bool   `json:"is_follow"`                       // true-已关注，false-未关注
	Name          string `json:"username" gorm:"column:username"` // 用户名称
}

func (f *Follow) TableName() string {
	return "follows"
}

func CreateFollow(ctx context.Context, actor string, toUser string) (err error) {
	err = DB.WithContext(ctx).Create(Follow{
		UserId:         actor,
		FollowedUserId: toUser,
		IsFollow:       Normal,
	}).Error
	return
}

func UpdateFollow(ctx context.Context, id uint, status bool) (err error) {
	err = DB.WithContext(ctx).Where("id", id).Update("is_follow", status).Error
	return
}

func QueryHasFollow(ctx context.Context, actor string, toUser string) (f *Follow, err error) {
	err = DB.WithContext(ctx).Where(map[string]string{
		"user_id":          actor,
		"followed_user_id": toUser,
	}).First(f).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// QueryFollowerList 查询粉丝列表 -连表查询 -返回粉丝信息
func QueryFollowerList(ctx context.Context, userID int64) ([]*UserInfoResp, error) {
	res := make([]*UserInfoResp, 0)
	if userID == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Table("users").Select("users.id, users.username, users.follow_count, users.follower_count, follows.is_follow").Joins("join follows on follows.user_id = users.id").Where("follows.followed_user_id = ? AND follows.is_follow = ?", userID, Normal).Find(&res).Error; err != nil {
		return nil, err
	}
	/*
		for k, _ := range res {
			fmt.Printf("%#v", res[k])
		}*/
	return res, nil
}

// QueryFollowList 查询关注列表--连表查询 -返回偶像信息
func QueryFollowList(ctx context.Context, userID int64) ([]*UserInfoResp, error) {
	res := make([]*UserInfoResp, 0)

	if userID == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Table("users").Select("users.id, users.username, users.follow_count, users.follower_count, follows.is_follow").Joins("join follows on follows.followed_user_id = users.id").Where("follows.user_id = ? AND follows.is_follow = ?", userID, Normal).Find(&res).Error; err != nil {
		return nil, err
	}
	/*
		for k, _ := range res {
			fmt.Printf("%#v", res[k])
		}
	*/
	return res, nil
}
