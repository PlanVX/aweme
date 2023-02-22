package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// check if LikeModel implements dal.LikeModel
var _ dal.LikeModel = (*LikeModel)(nil)

// LikeModel is the like model implementation for dal.LikeModel
type LikeModel struct {
	queries like
	rdb     redis.UniversalClient
}

// NewLikeModel creates a new like model
func NewLikeModel(db like, rdb redis.UniversalClient) *LikeModel {
	return &LikeModel{
		queries: db, rdb: rdb,
	}
}

// Insert inserts a like
func (l *LikeModel) Insert(ctx context.Context, like *dal.Like) error {
	if err := l.queries.WithContext(ctx).Create(like); err != nil {
		return err
	}
	l.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, like.VideoID), CountLike, 1)
	l.rdb.HIncrBy(ctx, GenRedisKey(TableUser, like.UserID), CountLike, 1)
	return nil
}

// Delete deletes a like by video id and user id
func (l *LikeModel) Delete(ctx context.Context, vid, uid int64) error {
	if r, err := l.queries.WithContext(ctx).DeleteByVideoIDAndUserID(vid, uid); err != nil {
		return err
	} else if r == 0 { // not found means no rows affected
		return gorm.ErrRecordNotFound
	}
	l.rdb.HIncrBy(ctx, GenRedisKey(TableVideo, vid), CountLike, -1)
	l.rdb.HIncrBy(ctx, GenRedisKey(TableUser, uid), CountLike, -1)
	return nil
}

// FindVideoIDsByUserID finds liked video ids by user id
func (l *LikeModel) FindVideoIDsByUserID(ctx context.Context, uid int64, limit, offset int) ([]int64, error) {
	return l.queries.WithContext(ctx).FindVideoIDsByUserID(uid, limit, offset)
}

// FindByVideoIDAndUserID finds a like by video id and user id
func (l *LikeModel) FindByVideoIDAndUserID(ctx context.Context, vid, uid int64) (*dal.Like, error) {
	return l.queries.WithContext(ctx).FindByVideoIDAndUserID(vid, uid)
}

// FindWhetherLiked finds a like record by video ids and user id
// return a list of video id that liked by userid
func (l *LikeModel) FindWhetherLiked(ctx context.Context, userid int64, videoID []int64) ([]int64, error) {
	var result []int64
	err := l.queries.WithContext(ctx).
		Select(l.queries.VideoID).
		Where(l.queries.UserID.Eq(userid), l.queries.VideoID.In(videoID...)).
		Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
