namespace go feed
include "user.thrift"

struct FeedRequest {
    1:optional i64 latest_time
    2:optional string token
}

struct FeedResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<Video> video_list
    4:optional i64 next_time
}

struct Video {
    1:required i64 id
    2:required user.User author
    3:required string play_url
    4:required string cover_url
    5:required i64 favorite_count
    6:required i64 comment_count
    7:required bool is_favorite
    8:required string title
}

service FeedService {
    FeedResponse Feed(1:required FeedRequest req)
}