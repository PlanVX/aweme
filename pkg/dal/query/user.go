package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gorm"
)

// check if UserModel implements UserModel interface
var _ dal.UserModel = (*CustomUserModel)(nil)

// CustomUserModel is the implementation of UserModel
type CustomUserModel struct {
	db *gorm.DB
	u  user
}

// NewUserModel returns a *CustomUserModel
func NewUserModel(db *gorm.DB) *CustomUserModel {

	return &CustomUserModel{
		db: db,
		u:  Use(db).User,
	}
}

// FindOne find one user by id
func (c *CustomUserModel) FindOne(ctx context.Context, id int64) (*dal.User, error) {
	return c.u.WithContext(ctx).FindOne(id)
}

// FindMany find many users by ids
func (c *CustomUserModel) FindMany(ctx context.Context, ids []int64) ([]*dal.User, error) {
	return c.u.WithContext(ctx).WithContext(ctx).FindMany(ids)
}

// FindByUsername find one user by username
func (c *CustomUserModel) FindByUsername(ctx context.Context, username string) (*dal.User, error) {
	return c.u.WithContext(ctx).FindByUsername(username)
}

// Insert insert a user
func (c *CustomUserModel) Insert(ctx context.Context, u *dal.User) error {
	return c.u.WithContext(ctx).Create(u)
}

// Update update a user
func (c *CustomUserModel) Update(context.Context, *dal.User) error {
	return nil
}
