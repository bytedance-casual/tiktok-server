namespace go favorite
include "feed.thrift"

struct FavoriteActionRequest {
    1:required string token    // 用户鉴权token
    2:required i64 video_id    // 视频id
    3:required i32 action_type // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1:required i32 status_code
    2:optional string status_msg
}

struct FavoriteListRequest {
    1:required i64 user_id  // 用户id
    2:required string token // 用户鉴权token
}

struct FavoriteListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<feed.Video> video_list
}

service FavoriteService {
    // 赞操作
    FavoriteActionResponse ActionFavorite(1:required FavoriteActionRequest req)
    // 喜欢列表
    FavoriteListResponse ListFavorite(1:required FavoriteListRequest req)
}