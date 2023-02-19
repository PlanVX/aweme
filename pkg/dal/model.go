package dal

import "time"

// User 用户表
type User struct {
	ID       int64  `gorm:"primary_key;auto_increment:false"` // 用户id
	Username string `gorm:"type:varchar(32);uniqueIndex"`     // 用户名
	Password []byte `gorm:"type:varchar(200);not null"`       // 密码
	Avatar   string `gorm:"type:varchar(200)"`                // 头像URL
}

// Video 视频表
type Video struct {
	ID        int64     `gorm:"primary_key;auto_increment:false" json:"id"`  // 视频id
	UserID    int64     `gorm:"type:bigint;not null"`                        // 用户id
	VideoURL  string    `gorm:"type:varchar(200);not null" json:"video_url"` // 视频URL
	CoverURL  string    `gorm:"type:varchar(200);not null" json:"cover_url"` // 封面URL
	Title     string    `gorm:"type:varchar(200);not null" json:"title"`     // 视频标题
	CreatedAt time.Time // 创建时间
}

// Comment 评论表
type Comment struct {
	ID        int64     `gorm:"primary_key;auto_increment:false" json:"id"` // 评论id
	Content   string    `gorm:"type:varchar(200);not null" json:"content"`  // 评论内容
	VideoID   int64     `gorm:"type:bigint;not null"`                       // 视频id
	UserID    int64     `gorm:"type:bigint;not null"`                       // 用户id
	CreatedAt time.Time // 创建时间
}

// Like 点赞表
type Like struct {
	ID        int64     `gorm:"primary_key;auto_increment:false" json:"id"` // 点赞id
	VideoID   int64     `gorm:"type:bigint;not null"`                       // 视频id
	UserID    int64     `gorm:"type:bigint;not null"`                       // 用户id
	CreatedAt time.Time // 创建时间
}

// Ralation 关系表
type Realtion struct {
	ID       int64     `gorm:"primary_key;auto_increment:false" json:"id"` // 点赞id
	Follower int64     `gorm:"type:bigint;not null"`                       // 关注者id
	Followed int64     `gorm:"type:bigint;not null"`                       // 被关注者id
	UpdateAt time.Time //关系创建时间
}
