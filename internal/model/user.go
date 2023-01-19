package model

// User && UserRequest && UserResponse https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#mWFx8s
type User struct {
	Id            int64  `json:"id"`             // 用户 Id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // 已关注/未关注
}
type UserRequest struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
type UserResponse struct {
	Response
	User User `json:"user"`
}

// UserRegisterRequest && UserRegisterResponse  https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#1bQwYM
type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// UserLoginRequest & UserLoginResponse https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#9UOUMJ
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
