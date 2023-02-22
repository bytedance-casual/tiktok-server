package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"tiktok-server/internal/utils"
)

const (
	Normal   bool = true
	Canceled bool = false
)

type Follow struct {
	gorm.Model
	UserId         string `json:"user_id"`
	FollowedUserId string `json:"followed_user_id"`
	IsFollow       bool   `json:"is_follow"`
}

func (f *Follow) TableName() string {
	return "follows"
}

func CreateFollow(ctx context.Context, actor string, toUser string) (err error) {
	err = DB.WithContext(ctx).Create(&Follow{
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
	}).First(&f).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// QueryFollowerList 查询粉丝列表
func QueryFollowerList(ctx context.Context, userId int64) ([]int64, error) {
	followList := make([]*Follow, 0)
	if userId == 0 {
		return []int64{}, nil
	}

	if err := DB.WithContext(ctx).Where("followed_user_id = ? AND is_follow = ?", userId, Normal).Find(&followList).Error; err != nil {
		return nil, err
	}

	resp := make([]int64, len(followList))
	for i, follow := range followList {
		atoi, err := strconv.Atoi(follow.UserId)
		if err != nil {
			return nil, err
		}
		resp[i] = int64(atoi)
	}
	return resp, nil
}

// QueryFollowList 查询关注列表
func QueryFollowList(ctx context.Context, userId int64) ([]int64, error) {
	followList := make([]*Follow, 0)
	if userId == 0 {
		return []int64{}, nil
	}

	if err := DB.WithContext(ctx).Where("user_id = ? AND is_follow = ?", userId, Normal).Find(&followList).Error; err != nil {
		return nil, err
	}

	resp := make([]int64, len(followList))
	for i, follow := range followList {
		temp, err := strconv.Atoi(follow.FollowedUserId)
		if err != nil {
			return nil, err
		}
		resp[i] = int64(temp)
	}
	return resp, nil
}

func MCheckFollow(userId int64, userIdList []int64, ctx context.Context) ([]bool, error) {
	followList := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND is_follow = ? AND followed_user_id IN ?", userId, Normal, userIdList).Find(&followList).Error; err != nil {
		return nil, err
	}

	set := utils.NewSet[string]()
	for _, follow := range followList {
		set.Add(follow.FollowedUserId)
	}

	boolList := make([]bool, len(userIdList))
	for i, id := range userIdList {
		boolList[i] = set.Contains(strconv.FormatInt(id, 10))
	}
	return boolList, nil
}

// MCountFollow 统计用户关注数
func MCountFollow(userIdList []int64, ctx context.Context) ([]int64, error) {
	countList := make([]int64, len(userIdList))
	for i, userId := range userIdList {
		if err := DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ? AND is_follow = ?", userId, Normal).Count(&countList[i]).Error; err != nil {
			return nil, err
		}
	}
	return countList, nil
}

// MCountFollower 统计用户粉丝数
func MCountFollower(userIdList []int64, ctx context.Context) ([]int64, error) {
	countList := make([]int64, len(userIdList))
	for i, userId := range userIdList {
		if err := DB.WithContext(ctx).Model(&Follow{}).Where("followed_user_id = ? AND is_follow = ?", userId, Normal).Count(&countList[i]).Error; err != nil {
			return nil, err
		}
	}
	return countList, nil
}
