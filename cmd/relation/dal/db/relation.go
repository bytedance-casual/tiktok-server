package db

import (
	"context"
	"tiktok-server/kitex_gen/user"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"gorm.io/gorm"
)

type FollowStatus uint8

const (
	Nolmal   = true
	Canceled = false
)

type Follow struct {
	gorm.Model
	UserId         string `json:"user_id"`
	FollowedUserId string `json:"followed_user_id"`
	IsFollow       int    `json:"is_follow"`
}

/*
type UserInfoResp struct {
	FollowCount   int64  `json:"follow_count"`                    // 关注总数
	FollowerCount int64  `json:"follower_count"`                  // 粉丝总数
	Id            int64  `json:"id"`                              // 用户id，和db.user不同
	IsFollow      int    `json:"is_follow"`                       // true-已关注，false-未关注
	Name          string `json:"username" gorm:"column:username"` // 用户名称
}
*/

func (f *Follow) TableName() string {
	return "follows"
}

// 查询粉丝列表 -连表查询 -返回粉丝信息
func QueryFollowerList(ctx context.Context, userID int64) ([]*user.User, error) {
	res := make([]*user.User, 0)
	if userID == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Table("users").Select("users.id, users.username, users.follow_count, users.follower_count, follows.is_follow").Joins("join follows on follows.user_id = users.id").Where("follows.followed_user_id = ? AND follows.is_follow = ?", userID, Nolmal).Find(&res).Error; err != nil {
		return nil, err
	}
	/*
		for k, _ := range res {
			fmt.Printf("%#v", res[k])
		}*/
	return res, nil
}

// 查询关注列表--连表查询 -返回偶像信息
func QueryFollowList(ctx context.Context, userID int64) ([]*user.User, error) {
	res := make([]*user.User, 0)

	if userID == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Table("users").Select("users.id, users.username, users.follow_count, users.follower_count, follows.is_follow").Joins("join follows on follows.followed_user_id = users.id").Where("follows.user_id = ? AND follows.is_follow = ?", userID, Nolmal).Find(&res).Error; err != nil {
		return nil, err
	}
	/*
		for k, _ := range res {
			fmt.Printf("%#v", res[k])
		}
	*/
	return res, nil
}
