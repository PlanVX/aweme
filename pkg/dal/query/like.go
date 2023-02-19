package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gorm"
)

// check if LikeModel implements dal.LikeModel
var _ dal.LikeModel = (*LikeModel)(nil)

// LikeModel is the like model implementation for dal.LikeModel
type LikeModel struct {
	db      *gorm.DB
	queries like
}

// NewLikeModel creates a new like model
func NewLikeModel(db *gorm.DB) *LikeModel {
	l := Use(db).Like
	return &LikeModel{db: db, queries: l}
}

// Insert inserts a like
func (l *LikeModel) Insert(ctx context.Context, like *dal.Like) error {
	return l.queries.WithContext(ctx).Create(like)
}

// Delete deletes a like by video id and user id
func (l *LikeModel) Delete(ctx context.Context, vid, uid int64) error {
	if r, err := l.queries.WithContext(ctx).DeleteByVideoIDAndUserID(vid, uid); err != nil {
		return err
	} else if r == 0 { // not found means no rows affected
		return gorm.ErrRecordNotFound
	}
	return nil
}

// FindByVideoIDAndUserID finds a like by video id and user id
func (l *LikeModel) FindByVideoIDAndUserID(ctx context.Context, vid, uid int64) (*dal.Like, error) {
	return l.queries.WithContext(ctx).FindByVideoIDAndUserID(vid, uid)
}
