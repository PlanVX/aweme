package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// check if UserModel implements UserModel interface
var _ dal.UserModel = (*UserModel)(nil)

// UserModel is the implementation of dal.UserModel
type UserModel struct {
	db       *gorm.DB
	uniqueID *UniqueID
	rdb      redis.UniversalClient
}

// NewUserModel returns a *UserModel
func NewUserModel(db *gorm.DB, rdb redis.UniversalClient) *UserModel {
	return &UserModel{
		db:       db,
		uniqueID: NewUniqueID(),
		rdb:      rdb,
	}
}

// FindOne find one user by id
func (c *UserModel) FindOne(ctx context.Context, id int64) (*dal.User, error) {
	var u dal.User
	err := c.db.WithContext(ctx).Take(&u, id).Error
	if err != nil {
		return nil, err
	}
	return c.FindOneStat(ctx, &u)
}

// FindMany find many users by ids
// Even if there is no any user matched, it will return an empty slice
func (c *UserModel) FindMany(ctx context.Context, ids []int64) ([]*dal.User, error) {
	var users []*dal.User
	err := c.db.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, users)
}

// FindByUsername find one user by username
func (c *UserModel) FindByUsername(ctx context.Context, username string) (*dal.User, error) {
	var u dal.User
	return &u, c.db.WithContext(ctx).Select("id", "username", "password").Take(&u, "username = ?", username).Error

}

// Insert insert a user
func (c *UserModel) Insert(ctx context.Context, u *dal.User) error {
	uid, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	u.ID = uid
	err = c.db.WithContext(ctx).Create(u).Error
	return err
}

// Update update a user
func (*UserModel) Update(context.Context, *dal.User) error {
	return nil
}

// FindOneStat find one user stat by id from redis
func (c *UserModel) FindOneStat(ctx context.Context, user *dal.User) (*dal.User, error) {
	err := c.rdb.HGetAll(ctx, GenRedisKey(TableUser, user.ID)).Scan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindManyStat find many user stats by ids from redis
func (c *UserModel) FindManyStat(ctx context.Context, users []*dal.User) ([]*dal.User, error) {
	cmder, err := c.rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
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
