package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// check if LikeModel implements dal.LikeModel
var _ dal.LikeModel = (*LikeModel)(nil)

// LikeModel is the implementation of dal.LikeModel
type LikeModel struct {
	db       *gorm.DB
	rdb      redis.UniversalClient
	uniqueID *UniqueID
}

// NewLikeModel creates a new comment like model
func NewLikeModel(db *gorm.DB, rdb redis.UniversalClient) *LikeModel {
	return &LikeModel{
		db:       db,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// FindByVideoIDAndUserID finds a like by video id and user id
func (c *LikeModel) FindByVideoIDAndUserID(ctx context.Context, vid, uid int64) (*dal.Like, error) {
	var like dal.Like
	err := c.db.
		WithContext(ctx).
		Where("video_id = ?", vid).
		Where("user_id = ?", uid).
		Take(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

// FindVideoIDsByUserID finds liked video ids by user id
func (c *LikeModel) FindVideoIDsByUserID(ctx context.Context, uid int64, limit, offset int) ([]int64, error) {
	var likes []int64
	err := c.db.WithContext(ctx).
		Model(&dal.Like{}).
		Table("likes").
		Where("user_id = ?", uid).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Pluck("video_id", &likes).Error
	if err != nil {
		return nil, err
	}
	return likes, nil
}

// FindWhetherLiked finds a like record by video ids and user id
// return a list of video id that liked by userid
func (c *LikeModel) FindWhetherLiked(ctx context.Context, userid int64, videoID []int64) ([]int64, error) {
	var likes []int64
	err := c.db.WithContext(ctx).
		Model(&dal.Like{}).
		Where("user_id = ?", userid).
		Where("video_id IN ?", videoID).
		Pluck("video_id", &likes).Error
	if err != nil {
		return nil, err
	}
	return likes, nil
}

// Insert inserts a like
func (c *LikeModel) Insert(ctx context.Context, like *dal.Like) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	like.ID = id
	err = c.db.WithContext(ctx).Create(like).Error
	if err != nil {
		return err
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, like.VideoID), CountLike, 1)
	c.rdb.HIncrBy(ctx, GenRedisKey(TableUser, like.UserID), CountLike, 1)
	return nil
}

// Delete deletes a like by video id and user id
func (c *LikeModel) Delete(ctx context.Context, vid, uid int64) error {
	res := c.db.WithContext(ctx).
		Where("video_id = ?", vid).
		Where("user_id = ?", uid).
		Delete(&dal.Like{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, vid), CountLike, -1)
	c.rdb.HIncrBy(ctx, GenRedisKey(TableUser, uid), CountLike, -1)
	return nil
}
