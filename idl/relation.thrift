namespace go relation
include "user.thrift"

struct RelationActionRequest {
    1:required string token    // 用户鉴权token
    2:required i64 to_user_id  // 对方用户id
    3:required i32 action_type // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1:required i32 status_code
    2:optional string status_msg
}

struct RelationFollowListRequest {
    1:required i64 user_id  // 用户id
    2:required string token // 用户鉴权token
}

struct RelationFollowListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<user.User> user_list // 用户信息列表
}

struct RelationFollowerListRequest {
    1:required i64 user_id  // 用户id
    2:required string token // 用户鉴权token
}

struct RelationFollowerListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<user.User> user_list // 用户信息列表
}

struct RelationFriendListRequest {
    1:required i64 user_id  // 用户id
    2:required string token // 用户鉴权token
}

struct RelationFriendListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<FriendUser> user_list // 用户列表
}

struct MCheckFollowRelationRequest {
    1:required i64 user_id            // 用户id
    2:required list<i64> user_id_list // 目标用户id列表
}

struct MCheckFollowRelationResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<bool> check_list
}

struct MCountRelationRequest {
    1:required list<i64> user_id_list
}

struct MCountRelationResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<i64> follow_count_list
    4:required list<i64> follower_count_list
}

struct FriendUser {
    1:optional string message // 和该好友的最新聊天消息
    2:required i64 msgType    // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
    3:required i64 id
    4:required string name
    5:required i64 follow_count
    6:required i64 follower_count
    7:required bool is_follow
}

service RelationService {
    // 关系操作
    RelationActionResponse ActionRelation(1:required RelationActionRequest req)
    // 用户关注列表
    RelationFollowListResponse ListFollowRelation(1:required RelationFollowListRequest req)
    // 用户粉丝列表
    RelationFollowerListResponse ListFollowerRelation(1:required RelationFollowerListRequest req)
    // 用户好友列表
    RelationFriendListResponse ListFriendRelation(1:required RelationFriendListRequest req)
    // protect
    // 批查询是否关注
    MCheckFollowRelationResponse MCheckFollowRelation(1:required MCheckFollowRelationRequest req)
    // 批获取统计关系
    MCountRelationResponse MCountRelation(1:required MCountRelationRequest req)
}