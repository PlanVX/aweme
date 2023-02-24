package query

import (
	"context"
	"github.com/redis/go-redis/v9"

	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gorm"
)

// check if CommentModel implements dal.CommentModel
var _ dal.CommentModel = (*CommentModel)(nil)

// CommentModel is the implementation of dal.CommentModel
type CommentModel struct {
	queries  comment
	rdb      redis.UniversalClient
	uniqueID *UniqueID
}

// NewCommentModel creates a new comment model
func NewCommentModel(c comment, rdb redis.UniversalClient) *CommentModel {
	return &CommentModel{
		queries:  c,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// FindByVideoID finds comments by video id
func (c *CommentModel) FindByVideoID(ctx context.Context, videoID int64, limit, offset int) ([]*dal.Comment, error) {
	return c.queries.WithContext(ctx).FindByVideoID(videoID, limit, offset)
}

// Insert inserts a comment
func (c *CommentModel) Insert(ctx context.Context, comment *dal.Comment) error {
	uid, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	comment.ID = uid
	err = c.queries.WithContext(ctx).Create(comment)
	if err != nil {
		return err
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, comment.VideoID), CountComment, 1)
	return nil
}

// Delete deletes a comment by id and user id
func (c *CommentModel) Delete(ctx context.Context, id int64, uid int64, vid int64) error {
	if r, err := c.queries.WithContext(ctx).DeleteByIDAndUserID(id, uid); err != nil {
		return err
	} else if r == 0 { // not found means no rows affected
		return gorm.ErrRecordNotFound
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, vid), CountComment, -1)
	return nil // success
}
