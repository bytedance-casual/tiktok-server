namespace go comment
include "user.thrift"

struct CommentActionRequest {
    1:required string token
    2:required i64 video_id
    3:required i32 action_type     // 1-发布评论，2-删除评论
    4:optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5:optional i64 comment_id      // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
    1:required i32 status_code   // 状态码，0-成功，其他值-失败
    2:optional string status_msg // 返回状态描述
    3:optional Comment comment   // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest {
    1:required string token
    2:required i64 video_id
}

struct CommentListResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<Comment> comment_list // 评论列表
}

struct MCountVideoCommentRequest {
    1:required list<i64> video_id_list
}

struct MCountVideoCommentResponse {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<i64> count_list
}

struct Comment {
    1:required i64 id             // 视频评论id
    2:required user.User user          // 评论用户信息
    3:required string content     // 评论内容
    4:required string create_date // 评论发布日期，格式 mm-dd
}

service CommentService {
    // 评论操作
    CommentActionResponse ActionComment(1:required CommentActionRequest req)
    // 视频评论列表
    CommentListResponse ListComment(1:required CommentListRequest req)
    // protect
    // 批量查询评论数
    MCountVideoCommentResponse MCountVideoComment(1:required MCountVideoCommentRequest req)
}