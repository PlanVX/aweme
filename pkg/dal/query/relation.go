package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// check if RelationModel implements RelationModel interface
var _ dal.RelationModel = (*RelationModel)(nil)

// RelationModel is the implementation of RelationModel
type RelationModel struct {
	rds redis.UniversalClient
	r   relation
}

// NewRelationModel returns a *RelationModel
func NewRelationModel(db relation, rds redis.UniversalClient) *RelationModel {
	return &RelationModel{rds: rds, r: db}
}

// Insert create a relation record
func (r *RelationModel) Insert(ctx context.Context, rel *dal.Relation) error {
	if err := r.r.Create(rel); err != nil {
		return err
	}
	r.rds.HIncrBy(ctx, GenRedisKey(TableUser, rel.UserID), CountFollow, 1)
	r.rds.HIncrBy(ctx, GenRedisKey(TableUser, rel.FollowTo), CountFans, 1)
	return nil
}

// Delete delete a relation record
func (r *RelationModel) Delete(ctx context.Context, userid, followTo int64) error {
	if r, err := r.r.WithContext(ctx).DeleteByUserIDAndFollowTo(userid, followTo); err != nil {
		return err
	} else if r == 0 {
		return gorm.ErrRecordNotFound
	}
	r.rds.HIncrBy(ctx, GenRedisKey(TableUser, userid), CountFollow, -1)
	r.rds.HIncrBy(ctx, GenRedisKey(TableUser, followTo), CountFans, -1)
	return nil
}

// FindWhetherFollowedList query whether the user follow the followTo ids
// return the followTo ids that the user follow
func (r *RelationModel) FindWhetherFollowedList(ctx context.Context, userid int64, followTo []int64) ([]int64, error) {
	var result []int64
	err := r.r.WithContext(ctx).
		Select(r.r.FollowTo).
		Where(r.r.UserID.Eq(userid), r.r.FollowTo.In(followTo...)).
		Scan(&result)
	return result, err
}

// FindFollowerTo query the user's followTo
// return the followTo ids that the user follow
func (r *RelationModel) FindFollowerTo(ctx context.Context, userid int64, limit, offset int) ([]int64, error) {
	return r.r.WithContext(ctx).FindFollowerTo(userid, limit, offset)
}

// FindFollowerFrom query the user's follower
// return the follower ids that the user is followed by
func (r *RelationModel) FindFollowerFrom(ctx context.Context, followTo int64, limit, offset int) ([]int64, error) {
	return r.r.WithContext(ctx).FindFollowerFrom(followTo, limit, offset)
}
