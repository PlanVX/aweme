package query

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCVideoModel_covertTime(t *testing.T) {
	t.Run("pass 0 as timestamp", func(t *testing.T) {
		// for 0 timestamp, it will take current time as timestamp
		timestamp := covertTime(0)
		unixMilli := timestamp.UnixMilli()
		// so the timestamp should not be 0
		assert.NotEqual(t, int64(0), unixMilli)
	})
	t.Run("pass other timestamp", func(t *testing.T) {
		timestamp := covertTime(1234567890)
		assert.Equal(t, int64(1234567890), timestamp.UnixMilli())
	})
}

func TestVideoFind(t *testing.T) {
	assertions := assert.New(t)
	mock, db, rdb, err := mockDB(t)
	assertions.NoError(err)
	model := NewVideoModel(db, rdb)
	const findOne = "SELECT * FROM `videos` WHERE `videos`.`id` = ? ORDER BY `videos`.`id` LIMIT 1"
	t.Run("FindOne success", func(t *testing.T) {
		mock.ExpectQuery(findOne).
			WithArgs(1).
			WillReturnRows(mock.NewRows([]string{"id", "title"}).AddRow(1, "test"))
		video, err := model.FindOne(context.TODO(), 1)
		assertions.NoError(err)
		assertions.Equal("test", video.Title)
	})
	t.Run("FindOne error", func(t *testing.T) {
		mock.ExpectQuery(findOne).
			WithArgs(1).
			WillReturnRows(mock.NewRows([]string{"id", "title"}))
		video, err := model.FindOne(context.TODO(), 1)
		// First() method will return gorm.ErrRecordNotFound if no record found
		assertions.Error(err)
		assertions.Nil(video)
	})
	t.Run("FindMany success", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `videos` WHERE id IN (?,?)").
			WithArgs(1, 2).
			WillReturnRows(mock.NewRows([]string{"id", "title"}).AddRow(1, "test1").AddRow(2, "test2"))
		videos, err := model.FindMany(context.TODO(), []int64{1, 2})
		assertions.NoError(err)
		assertions.Equal(2, len(videos))
		assertions.Equal("test1", videos[0].Title)
		assertions.Equal("test2", videos[1].Title)
	})
	t.Run("FindMany nothing", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `videos` WHERE id IN (?,?)").
			WithArgs(1, 2).
			WillReturnRows(mock.NewRows([]string{"id", "title"}))
		videos, err := model.FindMany(context.TODO(), []int64{1, 2})
		assertions.NoError(err)
		assertions.Len(videos, 0)
	})
	ts := int64(1609459200000)
	//layout := "2006-01-02 15:04:05 -0700 MST"
	arg := time.Unix(0, ts*int64(time.Millisecond))
	t.Run("Find latest", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `videos` WHERE created_at < ? ORDER BY created_at desc LIMIT 2").
			WithArgs(arg).
			WillReturnRows(mock.NewRows([]string{"id", "title"}).AddRow(1, "test1").AddRow(2, "test2"))
		videos, err := model.FindLatest(context.TODO(), ts, 2)
		assertions.NoError(err)
		assertions.Len(videos, 2)
		assertions.Equal("test1", videos[0].Title)
		assertions.Equal("test2", videos[1].Title)
	})
	t.Run("Find latest nothing", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `videos` WHERE created_at < ? ORDER BY created_at desc LIMIT 2").
			WithArgs(arg).
			WillReturnRows(mock.NewRows([]string{"id", "title"}))
		videos, err := model.FindLatest(context.TODO(), ts, 2)
		assertions.NoError(err)
		assertions.Len(videos, 0)
	})
	const findByUserID = "SELECT * FROM `videos` WHERE user_id = ? ORDER BY created_at desc LIMIT 2 OFFSET 1"
	t.Run("FindByUserID success", func(t *testing.T) {
		// user_id int64, limit int, offset int
		mock.ExpectQuery(findByUserID).
			WithArgs(int64(1)).
			WillReturnRows(mock.NewRows([]string{"id", "title"}).AddRow(1, "test1").AddRow(2, "test2"))
		videos, err := model.FindByUserID(context.TODO(), 1, 2, 1)
		assertions.NoError(err)
		assertions.Len(videos, 2)
		assertions.Equal("test1", videos[0].Title)
		assertions.Equal("test2", videos[1].Title)
	})
	t.Run("FindByUserID nothing", func(t *testing.T) {
		mock.ExpectQuery(findByUserID).
			WithArgs(int64(1)).
			WillReturnRows(mock.NewRows([]string{"id", "title"}))
		videos, err := model.FindByUserID(context.TODO(), 1, 2, 1)
		assertions.NoError(err)
		assertions.Len(videos, 0)
	})
}

func TestVideoExec(t *testing.T) {
	assertions := assert.New(t)
	mock, db, rdb, err := mockDB(t)
	assertions.NoError(err)
	model := NewVideoModel(db, rdb)
	const insertVideo = "INSERT INTO `videos` (`user_id`,`video_url`,`cover_url`,`title`,`created_at`,`id`) VALUES (?,?,?,?,?,?)"
	v := &dal.Video{
		UserID:   1,
		VideoURL: "1",
		Title:    "1",
		CoverURL: "1",
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(insertVideo).
			WithArgs(1, "1", "1", "1", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := model.Insert(context.TODO(), v)
		assertions.NoError(err)
		assertions.NotZero(v.ID)
		assertions.NotZero(v.CreatedAt)
	})
	t.Run("Insert error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(insertVideo).
			WithArgs(1, "1", "1", "1", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errors.New("error"))
		mock.ExpectRollback()
		err := model.Insert(context.TODO(), v)
		assertions.Error(err)
	})
	t.Run("Delete success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `videos` WHERE id = ? AND user_id = ?").
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := model.Delete(context.TODO(), 1, 1)
		assertions.NoError(err)
	})
	t.Run("Delete error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `videos` WHERE id = ? AND user_id = ?").
			WithArgs(1, 1).
			WillReturnError(errors.New("error"))
		mock.ExpectRollback()
		err := model.Delete(context.TODO(), 1, 1)
		assertions.Error(err)
	})
	t.Run("Delete nothing", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `videos` WHERE id = ? AND user_id = ?").
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(1, 0))
		mock.ExpectCommit()
		err := model.Delete(context.TODO(), 1, 1)
		assertions.Error(err)
	})
}
