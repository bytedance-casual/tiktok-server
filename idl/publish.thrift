namespace go publish

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
    3:required list<Video> video_list
}

struct Video {
    1:required i64 id
    2:required User author
    3:required string play_url
    4:required string cover_url
    5:required i64 favorite_count
    6:required i64 comment_count
    7:required bool is_favorite
    8:required string title
}

struct User {
    1:required i64 id
    2:required string name
    3:required i64 follow_count
    4:required i64 follower_count
    5:required bool is_follow
}

service PublishService {
    PublishActionResponse ActionPublish(1:required PublishActionRequest req)
    PublishListResponse ListPublish(1:required PublishListRequest req)
}