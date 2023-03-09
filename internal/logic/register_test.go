package logic

import (
	"context"
	"errors"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewRegister(t *testing.T) {
	assertions := assert.New(t)
	text := "a quick brown fox jumps over the lazy dog"
	userReq := types.UserReq{Username: text, Password: text}
	t.Run("success", func(t *testing.T) {
		register := mockRegister(t, nil)
		resp, err := register.Register(context.TODO(), &userReq)
		assertions.NoError(err)
		assertions.NotNil(resp)
		assertions.NotEmpty(resp.Token)
	})
	t.Run("fail on Insert", func(t *testing.T) {
		register := mockRegister(t, errors.New("error"))
		resp, err := register.Register(context.TODO(), &userReq)
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

// mockRegister is a helper function to mock the UserModel and return a Register
// err is the error that will be returned by the UserModel.Insert method
func mockRegister(t *testing.T, err error) *Register {
	m := NewUserModel(t)
	register := NewRegister(RegisterParam{UserModel: m, J: mockJwt()})
	m.On("Insert", mock.Anything, mock.Anything).Return(err)
	return register
}
