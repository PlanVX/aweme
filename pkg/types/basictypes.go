package types

// User 用户信息
type User struct {
	ID            int64  `json:"id"`             // 用户id
	Username      string `json:"name"`           // 用户名称
	Avatar        string `json:"avatar"`         // 用户头像 URL
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

// Video 视频信息
type Video struct {
	ID            int64  `json:"id"`             // 视频唯一标识
	Author        *User  `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// Comment 评论
type Comment struct {
	ID         int64  `json:"id"`          // 评论id
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	User       *User  `json:"user"`        // 评论用户信息
}

// Response 基础响应
type Response struct {
	Code int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	Msg  string `json:"status_msg"`  // 返回状态描述
}

// Message 消息
type Message struct {
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 该消息接收者的id
	FromUserID int64  `json:"from_user_id"` // 该消息发送者的id
	Content    string `json:"content"`      // 消息内容
	CreateTime string `json:"create_time"`  // 消息创建时间
}
