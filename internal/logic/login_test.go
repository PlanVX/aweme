package logic

import (
	"context"
	"errors"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewLogin(t *testing.T) {
	assertions := assert.New(t)
	text := []byte("a quick brown fox jumps over the lazy dog")
	bytearray, err := bcrypt.GenerateFromPassword(text, bcrypt.DefaultCost)
	assertions.NoError(err)
	arguments := &dal.User{Username: string(text), Password: bytearray}
	t.Run("success", func(t *testing.T) {
		login := mockLogin(t, arguments, nil)
		resp, err := login.Login(context.TODO(), &types.UserReq{Username: string(text), Password: string(text)})
		assertions.NoError(err)
		assertions.NotNil(resp)
		assertions.NotEmpty(resp.Token)
	})
	t.Run("fail on FindByUsername", func(t *testing.T) {
		login := mockLogin(t, nil, errors.New("error"))
		resp, err := login.Login(context.TODO(), &types.UserReq{Username: string(text), Password: string(text)})
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("fail on bcrypt verify", func(t *testing.T) {
		login := mockLogin(t, arguments, nil)
		resp, err := login.Login(context.TODO(), &types.UserReq{Username: string(text), Password: "wrong password"})
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

// mockLogin is a helper function to mock the login logic
// u is the user that will be returned by mocked UserModel.FindByUsername method
// err is the error that will be returned by mocked UserModel.FindByUsername method
func mockLogin(t *testing.T, u *dal.User, err error) *Login {
	m := NewUserModel(t)
	login := NewLogin(LoginParam{UserModel: m, J: mockJwt()})
	m.On("FindByUsername", mock.Anything, mock.Anything).Return(u, err)
	return login
}

// mockJwt is a helper function to mock the jwt signer
func mockJwt() *JWTSigner {
	return NewJWTSigner(&config.Config{JWT: struct {
		Secret    string   `yaml:"secret"`
		TTL       int64    `yaml:"ttl"`
		Whitelist []string `yaml:"whitelist"`
	}{Secret: "1234", TTL: 1234, Whitelist: nil}})
}
