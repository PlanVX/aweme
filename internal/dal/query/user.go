/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// check if UserQuery implements UserQuery interface
var _ dal.UserQuery = (*UserQuery)(nil)

// UserQuery is the implementation of dal.UserQuery
type UserQuery struct {
	db       *gorm.DB
	uniqueID *UniqueID
	rdb      *RDB
}

// NewUserQuery returns a *UserQuery
func NewUserQuery(db *gorm.DB, rdb *RDB) *UserQuery {
	return &UserQuery{
		db:       db,
		uniqueID: NewUniqueID(),
		rdb:      rdb,
	}
}

// FindOne find one user by id
func (c *UserQuery) FindOne(ctx context.Context, id int64) (*dal.User, error) {
	var u dal.User
	err := c.db.WithContext(ctx).Take(&u, id).Error
	if err != nil {
		return nil, err
	}
	return c.FindOneStat(ctx, &u)
}

// FindMany find many users by ids
// Even if there is no any user matched, it will return an empty slice
func (c *UserQuery) FindMany(ctx context.Context, ids []int64) ([]*dal.User, error) {
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
func (c *UserQuery) FindByUsername(ctx context.Context, username string) (*dal.User, error) {
	var u dal.User
	return &u, c.db.WithContext(ctx).Select("id", "username", "password").Take(&u, "username = ?", username).Error

}

// FindOneStat find one user stat by id from redis
func (c *UserQuery) FindOneStat(ctx context.Context, user *dal.User) (*dal.User, error) {
	err := c.rdb.HGetAll(ctx, GenRedisKey(TableUser, user.ID)).Scan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindManyStat find many user stats by ids from redis
func (c *UserQuery) FindManyStat(ctx context.Context, users []*dal.User) ([]*dal.User, error) {
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

// check if UserCommand implements dal.UserCommand interface
var _ dal.UserCommand = (*UserCommand)(nil)

// UserCommand is the implementation of dal.UserCommand
type UserCommand struct {
	db       *gorm.DB
	uniqueID *UniqueID
	rdb      *RDB
}

// NewUserCommand returns a *UserCommand
func NewUserCommand(db *gorm.DB, rdb *RDB) *UserCommand {
	return &UserCommand{
		db:       db,
		uniqueID: NewUniqueID(),
		rdb:      rdb,
	}
}

// Insert insert a user
func (c *UserCommand) Insert(ctx context.Context, u *dal.User) error {
	uid, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	u.ID = uid
	err = c.db.WithContext(ctx).Create(u).Error
	return err
}

// Update update a user
func (*UserCommand) Update(context.Context, *dal.User) error {
	return nil
}
