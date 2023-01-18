package model

// PublishActionRequest && PublishActionResponse https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#Fjyk8n
type PublishActionRequest struct {
	Token string `json:"token"`
	Data  []byte `json:"data"`
	Title string `json:"title"`
}
type PublishActionResponse struct {
	Response
}

// PublishListRequest && PublishListResponse https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#Rq5tgJ
type PublishListRequest struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
type PublishListResponse struct {
	Response
	VideoList Video `json:"video_list"`
}
