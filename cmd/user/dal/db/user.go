package db

import (
	"context"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	//Deleted        gorm.DeletedAt
}
type UserInfoResp struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id，和db.user不同
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Username      string `json:"name"`           // 用户名称，需要改成name字段来返回，与user结构体不同，
}

func (u *User) TableName() string {
	return "users"
}

// MGetUsers multiple get list of user info by userid
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) (int64, error) {

	result := DB.WithContext(ctx).Create(users)
	return int64(users[0].ID), result.Error
}

// QueryUser query user info by userName
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryIsFollow   querying if followed someone by followedUserId. userid is parsed from your token.
func QueryIsFollow(ctx context.Context, userid, followedUserId int64) (bool, error) {
	res := make([]*User, 0)
	if result := DB.WithContext(ctx).Where("user_id = ? AND followed_user_id = ?", userid).Find(&res); result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
