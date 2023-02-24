package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/redis/go-redis/v9"
)

// check if UserModel implements UserModel interface
var _ dal.UserModel = (*CustomUserModel)(nil)

// CustomUserModel is the implementation of UserModel
type CustomUserModel struct {
	u        user
	rds      redis.UniversalClient
	uniqueID *UniqueID
}

// NewUserModel returns a *CustomUserModel
func NewUserModel(u user, rds redis.UniversalClient) *CustomUserModel {
	return &CustomUserModel{
		u:        u,
		rds:      rds,
		uniqueID: NewUniqueID(),
	}
}

// FindOne find one user by id
func (c *CustomUserModel) FindOne(ctx context.Context, id int64) (*dal.User, error) {
	u, err := c.u.WithContext(ctx).FindOne(id)
	if err != nil {
		return nil, err
	}
	stat, err := c.FindOneStat(ctx, u)
	return stat, err
}

// FindMany find many users by ids
func (c *CustomUserModel) FindMany(ctx context.Context, ids []int64) ([]*dal.User, error) {
	result, err := c.u.WithContext(ctx).WithContext(ctx).FindMany(ids)
	if err != nil {
		return nil, err
	}
	stat, err := c.FindManyStat(ctx, result)
	return stat, err
}

// FindByUsername find one user by username
func (c *CustomUserModel) FindByUsername(ctx context.Context, username string) (*dal.User, error) {
	return c.u.WithContext(ctx).FindByUsername(username)
}

// Insert insert a user
func (c *CustomUserModel) Insert(ctx context.Context, u *dal.User) error {
	uid, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	u.ID = uid
	return c.u.WithContext(ctx).Create(u)
}

// Update update a user
func (c *CustomUserModel) Update(context.Context, *dal.User) error {
	return nil
}

// FindOneStat find one user stat by id from redis
func (c *CustomUserModel) FindOneStat(ctx context.Context, user *dal.User) (*dal.User, error) {
	if err := c.rds.HGetAll(ctx, GenRedisKey(TableUser, user.ID)).Scan(user); err != nil {
		return nil, err
	}
	return user, nil
}

// FindManyStat find many user stats by ids from redis
func (c *CustomUserModel) FindManyStat(ctx context.Context, users []*dal.User) ([]*dal.User, error) {
	cmder, err := c.rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, u := range users {
			pipe.HGetAll(ctx, GenRedisKey(TableUser, u.ID))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for i, cmd := range cmder {
		err := cmd.(*redis.MapStringStringCmd).Scan(users[i])
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
