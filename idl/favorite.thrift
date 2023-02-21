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

struct MCheckFavoriteRequest {
    1:required i64 user_id
    2:required list<i64> video_id_list
}

struct MCheckFavoriteResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<bool> exist_list
}

struct MCountVideoFavoriteRequest {
    1:required list<i64> video_id_list
}

struct MCountVideoFavoriteResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<i64> count_list
}

service FavoriteService {
    // 赞操作
    FavoriteActionResponse ActionFavorite(1:required FavoriteActionRequest req)
    // 喜欢列表
    FavoriteListResponse ListFavorite(1:required FavoriteListRequest req)
    // protect
    // 批查询是否点赞
    MCheckFavoriteResponse MCheckFavorite(1:required MCheckFavoriteRequest req)
    // 查询视频点赞数
    MCountVideoFavoriteResponse MCountFavorite(1:required MCountVideoFavoriteRequest req)
}