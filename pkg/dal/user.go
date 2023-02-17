package dal

import "context"

// check if UserModel implements UserModel interface
var _ UserModel = (*CustomUserModel)(nil)

// CustomUserModel is the implementation of UserModel
type CustomUserModel struct {
}

// NewUserModel returns a *CustomUserModel
func NewUserModel() *CustomUserModel {
	return &CustomUserModel{}
}

// FindOne find one user by id
func (c *CustomUserModel) FindOne(context.Context, int64) (*User, error) {
	return new(User), nil
}

// FindMany find many users by ids
func (c *CustomUserModel) FindMany(context.Context, []int64) ([]*User, error) {
	return nil, nil
}

// FindByUsername find one user by username
func (c *CustomUserModel) FindByUsername(context.Context, string) (*User, error) {
	return new(User), nil
}

// Insert insert a user
func (c *CustomUserModel) Insert(context.Context, *User) error {
	return nil
}

// Update update a user
func (c *CustomUserModel) Update(context.Context, *User) error {
	return nil
}
