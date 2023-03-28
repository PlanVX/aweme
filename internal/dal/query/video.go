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

package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

// check if VideoQuery implements VideoQuery interface
var _ dal.VideoQuery = (*VideoQuery)(nil)

// VideoQuery is the implementation of dal.VideoQuery
type VideoQuery struct {
	db       *gorm.DB
	uniqueID *UniqueID
	rdb      *RDB
}

// NewVideoQuery returns a *VideoQuery
func NewVideoQuery(db *gorm.DB, rdb *RDB) *VideoQuery {
	return &VideoQuery{
		db:       db,
		uniqueID: NewUniqueID(),
		rdb:      rdb,
	}
}

// FindOne find one video by id
func (c *VideoQuery) FindOne(ctx context.Context, id int64) (*dal.Video, error) {
	var v dal.Video
	err := c.db.WithContext(ctx).Take(&v, id).Error
	if err != nil {
		return nil, err
	}
	return c.FindOneStat(ctx, &v)
}

// FindMany find many videos by ids
func (c *VideoQuery) FindMany(ctx context.Context, ids []int64) ([]*dal.Video, error) {
	var videos []*dal.Video
	err := c.db.WithContext(ctx).
		Find(&videos, "id IN ?", ids).Error
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, videos)
}

// FindLatest find the latest videos
// timestamp is the millisecond timestamp
func (c *VideoQuery) FindLatest(ctx context.Context, timestamp int64, limit int64) ([]*dal.Video, error) {
	t := covertTime(timestamp)
	var videos []*dal.Video
	err := c.db.WithContext(ctx).
		Where("created_at < ?", t).
		Order("created_at desc").
		Limit(int(limit)).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, videos)
}

// FindByUserID find videos by user id
func (c *VideoQuery) FindByUserID(ctx context.Context, userID int64, limit, offset int) ([]*dal.Video, error) {
	var videos []*dal.Video
	err := c.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, videos)
}

// FindOneStat find one video stat by id
func (c *VideoQuery) FindOneStat(ctx context.Context, video *dal.Video) (*dal.Video, error) {
	err := c.rdb.HGetAll(ctx, GenRedisKey(TableVideo, video.ID)).Scan(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// FindManyStat find many video stats by ids
func (c *VideoQuery) FindManyStat(ctx context.Context, videos []*dal.Video) ([]*dal.Video, error) {
	cmder, err := c.rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, v := range videos {
			pipe.HGetAll(ctx, GenRedisKey(TableVideo, v.ID))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for i, cmd := range cmder {
		err := cmd.(*redis.MapStringStringCmd).Scan(videos[i])
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
}

// check if VideoCommand implements VideoCommand interface
var _ dal.VideoCommand = (*VideoCommand)(nil)

// VideoCommand is the implementation of dal.VideoQuery
type VideoCommand struct {
	db       *gorm.DB
	uniqueID *UniqueID
	rdb      *RDB
}

// NewVideoCommand returns a *VideoCommand
func NewVideoCommand(db *gorm.DB, rdb *RDB) *VideoCommand {
	return &VideoCommand{
		db:       db,
		uniqueID: NewUniqueID(),
		rdb:      rdb,
	}
}

// Insert insert a video
func (c *VideoCommand) Insert(ctx context.Context, video *dal.Video) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	video.ID = id
	err = c.db.WithContext(ctx).Create(video).Error
	if err != nil {
		return err
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableUser, video.UserID), CountVideo, 1)
	return err
}

// Update update a video
func (*VideoCommand) Update(context.Context, *dal.Video) error {
	return nil
}

// Delete a video by its id and correct user id
func (c *VideoCommand) Delete(ctx context.Context, id int64, uid int64) error {
	res := c.db.WithContext(ctx).
		Where("id = ?", id).
		Where("user_id = ?", uid).
		Delete(&dal.Video{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableUser, uid), CountVideo, -1)
	return nil
}

// covertTime converts timestamp to time.Time in milliseconds
func covertTime(timestamp int64) time.Time {
	var t time.Time
	if timestamp == 0 { // if timestamp is 0, use current time
		t = time.Now()
	} else {
		t = time.Unix(0, timestamp*int64(time.Millisecond))
	}
	return t
}
