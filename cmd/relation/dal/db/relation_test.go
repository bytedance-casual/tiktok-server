package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"tiktok-server/internal/conf"
)

type TestFollow struct {
	gorm.Model
	UserId         string `json:"user_id"`
	FollowedUserId string `json:"followed_user_id"`
	IsFollow       bool   `json:"is_follow"`
}

func (f *TestFollow) TableName() string {
	return "follows"
}

func testQueryFollowerList(ctx context.Context, userId int64) ([]int64, error) {
	followList := make([]*TestFollow, 0)
	if userId == 0 {
		return []int64{}, nil
	}

	if err := DB.WithContext(ctx).Where("followed_user_id = ? AND is_follow = ?", userId, Normal).Find(&followList).Error; err != nil {
		return nil, err
	}

	resp := make([]int64, len(followList))
	for i, follow := range followList {
		resp[i] = int64(follow.ID)
	}
	return resp, nil
}

func TestQueryFollowerList(t *testing.T) {
	conf.Init()
	Init()
	list, err := MCountFollower([]int64{2}, context.Background())
	assert.NoError(t, err)
	fmt.Printf("%v", list)
}
