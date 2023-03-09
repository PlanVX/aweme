package query

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserFind(t *testing.T) {
	assertions := assert.New(t)
	mock, gormDB, rdb, err := mockDB(t)
	assertions.NoError(err)
	model := NewUserModel(gormDB, rdb)
	const findOneUser = "SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1"
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
		// First() method will return gorm.ErrRecordNotFound if no record found
		assertions.Error(err)
		assertions.Equal(gorm.ErrRecordNotFound, err)
		assertions.Nil(user)
	})
	const findManyUser = "SELECT * FROM `users` WHERE id IN (?,?)"
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
		const findByUsername = "SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1"
		mock.ExpectQuery(findByUsername).
			WithArgs("test").
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "test"))
		user, err := model.FindByUsername(context.TODO(), "test")
		assertions.NoError(err)
		assertions.Equal("test", user.Username)
	})

}

func mockDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB, redis.UniversalClient, error) {
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))
	s := miniredis.RunT(t)
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{s.Addr()},
	})
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
