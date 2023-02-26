package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// check if CommentModel implements dal.CommentModel
var _ dal.CommentModel = (*CommentModel)(nil)

// CommentModel is the implementation of dal.CommentModel
type CommentModel struct {
	db       *gorm.DB
	rdb      redis.UniversalClient
	uniqueID *UniqueID
}

// NewCommentModel is the constructor of CommentModel
func NewCommentModel(db *gorm.DB, rdb redis.UniversalClient) *CommentModel {
	return &CommentModel{
		db:       db,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// FindByVideoID finds comments by video id
func (c *CommentModel) FindByVideoID(ctx context.Context, videoID int64, limit, offset int) ([]*dal.Comment, error) {
	var comments []*dal.Comment
	err := c.db.WithContext(ctx).
		Where("video_id = ?", videoID).
		Limit(limit).
		Offset(offset).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Insert inserts a comment
func (c *CommentModel) Insert(ctx context.Context, comment *dal.Comment) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	comment.ID = id
	err = c.db.WithContext(ctx).Create(comment).Error
	if err != nil {
		return err
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, comment.VideoID), CountComment, 1)
	return nil
}

// Delete deletes a comment by id and user id
func (c *CommentModel) Delete(ctx context.Context, id int64, uid int64, vid int64) error {
	res := c.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, uid).
		Delete(&dal.Comment{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, vid), CountComment, -1)
	return nil
}
