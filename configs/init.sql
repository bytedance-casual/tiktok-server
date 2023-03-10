# 亮点：
# - 物理外键 -> 逻辑外键
# - 单独拆分评论内容
# - 热点操作如关注接口采用软删除
# - count 类数据动态计算
create table comments
(
    id         bigint auto_increment comment '逻辑主键'
        primary key,
    video_id   bigint   not null,
    user_id    bigint   not null comment '评论用户',
    updated_at datetime not null,
    deleted_at datetime null,
    created_at datetime null
);

create index comments_video_id_index
    on comments (video_id);

create table contents
(
    id         bigint auto_increment
        primary key,
    comment_id bigint   not null,
    content    text     not null,
    updated_at datetime not null,
    deleted_at datetime null,
    created_at datetime null
);

create table follows
(
    id               bigint auto_increment comment '逻辑主键'
        primary key,
    user_id          bigint   not null comment '关注者',
    followed_user_id bigint   not null comment '被关注者',
    is_follow        tinyint  null,
    updated_at       datetime not null,
    deleted_at       datetime null,
    created_at       datetime null,
    constraint follows_user_id_is_follow_followed_user_id_uindex
        unique (user_id, is_follow, followed_user_id)
);

create index follows_followed_user_id_is_follow_index
    on follows (followed_user_id, is_follow);

create index follows_user_id_is_follow_index
    on follows (user_id, is_follow);

create table likes
(
    id         bigint auto_increment comment '逻辑主键'
        primary key,
    video_id   bigint   not null comment '被点赞的视频',
    user_id    bigint   not null comment '点赞用户',
    updated_at datetime not null,
    deleted_at datetime null,
    created_at datetime null
);

create index likes_user_id_index
    on likes (user_id);

create index likes_user_id_video_id_index
    on likes (user_id, video_id);

create index likes_video_id_index
    on likes (video_id);

create table messages
(
    id           bigint auto_increment
        primary key,
    from_user_id bigint       not null,
    to_user_id   bigint       not null,
    content      varchar(255) not null,
    created_at   datetime     null,
    updated_at   datetime     null,
    deleted_at   datetime     null
);

create index messages_created_at_index
    on messages (created_at);

create index messages_from_user_id_to_user_id_created_at_index
    on messages (from_user_id, to_user_id, created_at);

create index messages_from_user_id_to_user_id_index
    on messages (from_user_id, to_user_id);

create table users
(
    id         bigint auto_increment
        primary key,
    username   varchar(45)  not null,
    password   varchar(255) not null,
    updated_at datetime     not null,
    deleted_at datetime     null,
    created_at datetime     null
);

create index users_username_index
    on users (username);

create table videos
(
    id         bigint auto_increment
        primary key,
    created_at datetime(3)  null,
    updated_at datetime(3)  null,
    deleted_at datetime(3)  null,
    author_id  bigint       not null,
    play_url   varchar(255) not null,
    cover_url  varchar(255) not null,
    title      longtext     null
);

create index idx_videos_deleted_at
    on videos (deleted_at);

create index videos_author_id_index
    on videos (author_id);

