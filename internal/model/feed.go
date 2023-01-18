package model

// FeedRequest && FeedResponse && Video https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#gYvQW0
type FeedRequest struct {
	LatestTime int64  `json:"latest_time,omitempty"`
	Token      string `json:"token,omitempty"`
}
type FeedResponse struct {
	Response
	VideoList Video `json:"video_list"`
	NextTime  int64 `json:"next_time,omitempty"`
}
type Video struct {
	Id            int64  `json:"id"`             // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频评论总数
	IsFavorite    bool   `json:"is_favorite"`    // 已点赞/未点赞
	Title         string `json:"title"`          // 视频标题
}
