package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"gorm.io/gorm"
)

// check if LikeQuery implements dal.LikeQuery
var _ dal.LikeQuery = (*LikeQuery)(nil)

// LikeQuery is the implementation of dal.LikeQuery
type LikeQuery struct {
	db  *gorm.DB
	rdb *RDB
}

// NewLikeQuery creates a new comment like model
func NewLikeQuery(db *gorm.DB, rdb *RDB) *LikeQuery {
	return &LikeQuery{
		db:  db,
		rdb: rdb,
	}
}

// FindByVideoIDAndUserID finds a like by video id and user id
func (c *LikeQuery) FindByVideoIDAndUserID(ctx context.Context, vid, uid int64) (*dal.Like, error) {
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
func (c *LikeQuery) FindVideoIDsByUserID(ctx context.Context, uid int64, limit, offset int) ([]int64, error) {
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
func (c *LikeQuery) FindWhetherLiked(ctx context.Context, userid int64, videoID []int64) ([]int64, error) {
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

// check if LikeQuery implements dal.LikeCommand
var _ dal.LikeCommand = (*LikeCommand)(nil)

// LikeCommand is the implementation of dal.LikeCommand
type LikeCommand struct {
	db       *gorm.DB
	rdb      *RDB
	uniqueID *UniqueID
}

// NewLikeCommand creates a new comment like model
func NewLikeCommand(db *gorm.DB, rdb *RDB) *LikeCommand {
	return &LikeCommand{
		db:       db,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// Insert inserts a like
func (c *LikeCommand) Insert(ctx context.Context, like *dal.Like) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	like.ID = id
	err = c.db.WithContext(ctx).Create(like).Error
	if err != nil {
		return err
	}
	var fields = []HashField{
		{Key: GenRedisKey(TableVideo, like.VideoID), Field: CountLike},
		{Key: GenRedisKey(TableUser, like.UserID), Field: CountLike},
	}
	c.rdb.HKeyFieldsIncrBy(ctx, fields, 1)
	return nil
}

// Delete deletes a like by video id and user id
func (c *LikeCommand) Delete(ctx context.Context, vid, uid int64) error {
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
	var fields = []HashField{
		{Key: GenRedisKey(TableVideo, vid), Field: CountLike},
		{Key: GenRedisKey(TableUser, uid), Field: CountLike},
	}
	c.rdb.HKeyFieldsIncrBy(ctx, fields, -1)
	return nil
}
