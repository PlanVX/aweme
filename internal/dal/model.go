/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dal

import (
	"time"
)

// User 用户表
type User struct {
	ID              int64  `gorm:"primary_key;auto_increment:false" redis:"-"` // 用户id
	Username        string `gorm:"type:varchar(32);uniqueIndex" redis:"-"`     // 用户名
	Password        []byte `gorm:"type:varchar(200);not null" redis:"-"`       // 密码
	Avatar          string `gorm:"type:varchar(200)" redis:"-"`                // 头像URL
	BackgroundImage string `gorm:"type:varchar(200)" redis:"-"`                // 背景图片URL
	Signature       string `gorm:"type:varchar(200)" redis:"-"`                // 个性签名
	// 不存入数据库
	VideoCount   int64 `gorm:"-" redis:"video_count"`    // 视频数量
	LikeCount    int64 `gorm:"-" redis:"like_count"`     // 点赞数量
	FansCount    int64 `gorm:"-" redis:"fans_count"`     // 粉丝数量
	FollowCount  int64 `gorm:"-" redis:"follow_count"`   // 关注数量
	BeLikedCount int64 `gorm:"-" redis:"be_liked_count"` // 被点赞数量
}

// Video 视频表
type Video struct {
	ID        int64     `gorm:"primary_key;auto_increment:false" redis:"-"` // 视频id
	UserID    int64     `gorm:"type:bigint;not null" redis:"-"`             // 用户id
	VideoURL  string    `gorm:"type:varchar(200);not null" redis:"-"`       // 视频URL
	CoverURL  string    `gorm:"type:varchar(200);not null" redis:"-"`       // 封面URL
	Title     string    `gorm:"type:varchar(200);not null" redis:"-"`       // 视频标题
	CreatedAt time.Time // 创建时间
	// 不存入数据库
	LikeCount    int64 `gorm:"-" redis:"like_count"`    // 点赞数量
	CommentCount int64 `gorm:"-" redis:"comment_count"` // 评论数量
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

// Relation 关系表
type Relation struct {
	ID        int64     `gorm:"primary_key;auto_increment:false" json:"id"` // 点赞id
	UserID    int64     `gorm:"type:bigint;not null"`                       // 关注者id
	FollowTo  int64     `gorm:"type:bigint;not null"`                       // 被关注者id
	CreatedAt time.Time //关系创建时间
}
