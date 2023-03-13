package query

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserFind(t *testing.T) {
	assertions := assert.New(t)
	mock, gormDB, rdb, err := mockDB(t)
	assertions.NoError(err)
	model := NewUserModel(gormDB, rdb)
	const findOneUser = "SELECT `users`.`id`,`users`.`username`,`users`.`password`,`users`.`avatar`,`users`.`background_image`,`users`.`signature` FROM `users` WHERE `users`.`id` = ? LIMIT 1"
	t.Run("FindOne success", func(t *testing.T) {
		mock.ExpectQuery(findOneUser).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "test"))
		user, err := model.FindOne(context.TODO(), 1)
		assertions.NoError(err)
		assertions.Equal("test", user.Username)
	})
	t.Run("FindOne error", func(t *testing.T) {
		mock.ExpectQuery(findOneUser).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}))
		user, err := model.FindOne(context.TODO(), 1)
		// method will return gorm.ErrRecordNotFound if no record found
		assertions.Error(err)
		assertions.Equal(gorm.ErrRecordNotFound, err)
		assertions.Nil(user)
	})
	const findManyUser = "SELECT `users`.`id`,`users`.`username`,`users`.`password`,`users`.`avatar`,`users`.`background_image`,`users`.`signature` FROM `users` WHERE id IN (?,?)"
	t.Run("FindMany success", func(t *testing.T) {
		mock.ExpectQuery(findManyUser).
			WithArgs(1, 2).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "test1").AddRow(2, "test2"))
		users, err := model.FindMany(context.TODO(), []int64{1, 2})
		assertions.NoError(err)
		assertions.Equal(2, len(users))
		assertions.Equal("test1", users[0].Username)
		assertions.Equal("test2", users[1].Username)
	})

	t.Run("FindMany nothing", func(t *testing.T) {
		mock.ExpectQuery(findManyUser).
			WithArgs(1, 2).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}))
		users, err := model.FindMany(context.TODO(), []int64{1, 2})
		assertions.NoError(err)
		assertions.Len(users, 0)
	})
	t.Run("FindByUsername success", func(t *testing.T) {
		const findByUsername = "SELECT `id`,`username`,`password` FROM `users` WHERE username = ? LIMIT 1"
		mock.ExpectQuery(findByUsername).
			WithArgs("test").
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "test"))
		user, err := model.FindByUsername(context.TODO(), "test")
		assertions.NoError(err)
		assertions.Equal("test", user.Username)
	})

}

func mockDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB, *RDB, error) {
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{
		QueryFields:            true,
		SkipDefaultTransaction: true,
	})
	s := miniredis.RunT(t)
	c := config.Config{}
	c.Redis.Addr = []string{s.Addr()}
	lf := fxtest.NewLifecycle(t)
	rdb := NewRedisUniversalClient(&c, lf, zap.NewExample())
	lf.RequireStart()
	return mock, gormDB, rdb, err
}

func TestUserInsert(t *testing.T) {
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	require.NoError(t, err)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(t, err)
	assertions := assert.New(t)

	const insertUser = "INSERT INTO `users` (`username`,`password`,`avatar`,`background_image`,`signature`,`id`) VALUES (?,?,?,?,?,?)"
	model := NewUserModel(gormDB, nil)
	user := &dal.User{
		Username: "test",
		Password: []byte("test"),
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectExec(insertUser).
			WithArgs("test", []byte("test"), "", "", "", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Insert(context.TODO(), user)
		assertions.NoError(err)
		assertions.NotZero(user.ID)
	})
	t.Run("Insert error", func(t *testing.T) {
		mock.ExpectExec(insertUser).
			WithArgs("test", []byte("test"), "", "", "", sqlmock.AnyArg()).
			WillReturnError(errors.New("duplicated username"))
		err := model.Insert(context.TODO(), user)
		assertions.Error(err)
	})

}
