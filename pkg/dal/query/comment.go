package query

import (
	"context"

	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gorm"
)

// check if CommentModel implements dal.CommentModel
var _ dal.CommentModel = (*CommentModel)(nil)

// CommentModel is the implementation of dal.CommentModel
type CommentModel struct {
	queries comment
	db      *gorm.DB
}

// NewCommentModel creates a new comment model
func NewCommentModel(db *gorm.DB) *CommentModel {
	return &CommentModel{queries: Use(db).Comment, db: db}
}

// FindByVideoID finds comments by video id
func (c *CommentModel) FindByVideoID(ctx context.Context, videoID int64, limit, offset int) ([]*dal.Comment, error) {
	return c.queries.WithContext(ctx).FindByVideoID(videoID, limit, offset)
}

// Insert inserts a comment
func (c *CommentModel) Insert(ctx context.Context, comment *dal.Comment) error {
	return c.queries.WithContext(ctx).Create(comment)
}

// Delete deletes a comment by id and user id
func (c *CommentModel) Delete(ctx context.Context, id int64, uid int64) error {
	if r, err := c.queries.WithContext(ctx).DeleteByIDAndUserID(id, uid); err != nil {
		return err
	} else if r == 0 { // not found means no rows affected
		return gorm.ErrRecordNotFound
	}
	return nil // success
}
