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

service PublishService {
    PublishActionResponse ActionPublish(1:required PublishActionRequest req)
    PublishListResponse ListPublish(1:required PublishListRequest req)
}