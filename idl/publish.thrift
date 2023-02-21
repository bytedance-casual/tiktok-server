namespace go publish
include "feed.thrift"

struct PublishActionRequest {
    1:required string token
    2:required binary data
    3:required string title
}

struct PublishActionResponse {
    1:required i32 status_code
    2:optional string status_msg
}

struct PublishListRequest {
    1:required i64 user_id
    2:required string token
}

struct PublishListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<feed.Video> video_list
}

struct PublishVideoActionRequest {
    1:required i64 video_id
    2:required i32 action_type // 1-更新视频喜欢人数, 2-更新视频评论人数
    3:required bool increase   // true +1; false -1
}

struct PublishVideoActionResponse {
    1:required i32 status_code
    2:optional string status_msg
}

struct VideosMGetRequest {
    1:required i64 user_id             // 执行查询操作的用户 id
    2:required list<i64> video_id_list // 需要查询的视频 id 列表
}

struct VideosMGetResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<feed.Video> videos
}

service PublishService {
    // 发布视频
    PublishActionResponse ActionPublish(1:required PublishActionRequest req)
    // 查询自身发布视频
    PublishListResponse ListPublish(1:required PublishListRequest req)
    // protect
    // 更新视频数据
    PublishVideoActionResponse VideoActionPublish(1:required PublishVideoActionRequest req)
    // 批获取最简视频信息
    VideosMGetResponse MGetVideos(1:required VideosMGetRequest req)
}