package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"gorm.io/gorm"
)

// check if RelationModel implements RelationModel interface
var _ dal.RelationModel = (*RelationModel)(nil)

// RelationModel is the implementation of dal.RelationModel
type RelationModel struct {
	db       *gorm.DB
	rdb      *RDB
	uniqueID *UniqueID
}

// NewRelationModel creates a new comment relation model
func NewRelationModel(db *gorm.DB, rdb *RDB) *RelationModel {
	return &RelationModel{
		db:       db,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// FindWhetherFollowedList query whether the user follow the followTo ids
// return the followTo ids that the user follow
func (c *RelationModel) FindWhetherFollowedList(ctx context.Context, userid int64, followTo []int64) ([]int64, error) {
	var result []int64
	err := c.db.WithContext(ctx).
		Model(&dal.Relation{}).
		Where("follow_to IN ?", followTo).
		Where("user_id = ?", userid).
		Pluck("follow_to", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindFollowerTo query the user's followTo
// return the followTo ids that the user follow
func (c *RelationModel) FindFollowerTo(ctx context.Context, userid int64, limit, offset int) ([]int64, error) {
	var result []int64
	model := &dal.Relation{}
	err := c.db.WithContext(ctx).
		Model(model).
		Where("user_id = ?", userid).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Pluck("follow_to", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindFollowerFrom query the user's follower
// return the follower ids that the user is followed by
func (c *RelationModel) FindFollowerFrom(ctx context.Context, followTo int64, limit, offset int) ([]int64, error) {
	var result []int64
	err := c.db.WithContext(ctx).
		Model(&dal.Relation{}).
		Where("follow_to = ?", followTo).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Pluck("user_id", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Insert create a relation record
func (c *RelationModel) Insert(ctx context.Context, rel *dal.Relation) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	rel.ID = id
	err = c.db.WithContext(ctx).
		Create(rel).Error
	if err != nil {
		return err
	}
	keyFields := []HashField{
		{Key: GenRedisKey(TableUser, rel.UserID), Field: CountFollow},
		{Key: GenRedisKey(TableUser, rel.FollowTo), Field: CountFans},
	}
	c.rdb.HKeyFieldsIncrBy(ctx, keyFields, 1)
	return nil
}

// Delete a relation record by userid and followTo
func (c *RelationModel) Delete(ctx context.Context, userid, followTo int64) error {
	res := c.db.WithContext(ctx).
		Where("user_id = ?", userid).
		Where("follow_to = ?", followTo).
		Delete(&dal.Relation{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	keyFields := []HashField{
		{Key: GenRedisKey(TableUser, userid), Field: CountFollow},
		{Key: GenRedisKey(TableUser, followTo), Field: CountFans},
	}
	c.rdb.HKeyFieldsIncrBy(ctx, keyFields, -1)
	return nil
}
