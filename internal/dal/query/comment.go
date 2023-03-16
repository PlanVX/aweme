package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"gorm.io/gorm"
)

// check if CommentQuery implements dal.CommentQuery
var _ dal.CommentQuery = (*CommentQuery)(nil)

// CommentQuery is the implementation of dal.CommentQuery
type CommentQuery struct {
	db       *gorm.DB
	rdb      *RDB
	uniqueID *UniqueID
}

// NewCommentQuery is the constructor of CommentQuery
func NewCommentQuery(db *gorm.DB, rdb *RDB) *CommentQuery {
	return &CommentQuery{
		db:  db,
		rdb: rdb,
	}
}

// FindByVideoID finds comments by video id
func (c *CommentQuery) FindByVideoID(ctx context.Context, videoID int64, limit, offset int) ([]*dal.Comment, error) {
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

// check if CommentCommand implements dal.CommentCommand
var _ dal.CommentCommand = (*CommentCommand)(nil)

// CommentCommand is the implementation of dal.CommentCommand
type CommentCommand struct {
	db       *gorm.DB
	rdb      *RDB
	uniqueID *UniqueID
}

// NewCommentCommand is the constructor of CommentCommand
func NewCommentCommand(db *gorm.DB, rdb *RDB) *CommentCommand {
	return &CommentCommand{
		db:       db,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// Insert inserts a comment
func (c *CommentCommand) Insert(ctx context.Context, comment *dal.Comment) error {
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
func (c *CommentCommand) Delete(ctx context.Context, id int64, uid int64, vid int64) error {
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
