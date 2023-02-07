namespace go message

struct MessageChatRequest {
    1:required string token   // 用户鉴权token
    2:required i64 to_user_id // 对方用户id
}

struct MessageChatResponse {
    1:required i32 status_code            // 状态码，0-成功，其他值-失败
    2:optional string status_msg          // 返回状态描述
    3:required list<Message> message_list // 消息列表
}

struct MessageActionRequest {
    1:required string token    // 用户鉴权token
    2:required i64 to_user_id  // 对方用户id
    3:required i32 action_type // 1-发送消息
    4:required string content  // 消息内容
}

struct MessageActionResponse {
    1:required i32 status_code            // 状态码，0-成功，其他值-失败
    2:optional string status_msg          // 返回状态描述
}

struct Message {
    1:required i64 id             // 消息id
    2:required i64 to_user_id     // 该消息接收者的id
    3:required i64 from_user_id   // 该消息发送者的id
    4:required string content     // 消息内容
    5:optional string create_time // 消息创建时间
}

service MessageService {
    // 聊天记录
    MessageChatResponse ChatMessage(1:required MessageChatRequest req)
    // 消息操作
    MessageActionResponse ActionMessage(1:required MessageActionRequest req)
}