package query

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/internal/dal"
	"testing"
)

func TestLikeFind(t *testing.T) {
	assertions, mock, model := newMock(t, NewLikeQuery)
	const findByVideoIDAndUserID = "SELECT `likes`.`id`,`likes`.`video_id`,`likes`.`user_id`,`likes`.`created_at` FROM `likes` WHERE video_id = ? AND user_id = ? LIMIT 1"
	t.Run("FindByVideoIDAndUserID success", func(t *testing.T) {
		// video_id int64, user_id int64
		mock.ExpectQuery(findByVideoIDAndUserID).WithArgs(1, 1).
			WillReturnRows(mock.NewRows([]string{"id", "video_id", "user_id"}).AddRow(1, 1, 1))
		like, err := model.FindByVideoIDAndUserID(context.TODO(), 1, 1)
		assertions.NoError(err)
		assertions.Equal(int64(1), like.UserID)
		assertions.Equal(int64(1), like.VideoID)
	})
	t.Run("FindByVideoIDAndUserID fail", func(t *testing.T) {
		mock.ExpectQuery(findByVideoIDAndUserID).
			WithArgs(1, 1).
			WillReturnRows(mock.NewRows([]string{"id"}))
		like, err := model.FindByVideoIDAndUserID(context.TODO(), 1, 1)
		assertions.Error(err)
		assertions.Nil(like)
	})
	const findVideoIDsByUserID = "SELECT `video_id` FROM `likes` WHERE user_id = ? ORDER BY created_at desc LIMIT 2 OFFSET 1"
	t.Run("FindVideoIDsByUserID success", func(t *testing.T) {
		// user_id int64, limit int, offset int
		mock.ExpectQuery(findVideoIDsByUserID).
			WithArgs(int64(1)).
			WillReturnRows(mock.NewRows([]string{"video_id"}).AddRow(1).AddRow(2))
		like, err := model.FindVideoIDsByUserID(context.TODO(), 1, 2, 1)
		assertions.NoError(err)
		assertions.Equal(2, len(like))
		assertions.Equal(int64(1), like[0])
		assertions.Equal(int64(2), like[1])
	})
	t.Run("FindVideoIDsByUserID nothing", func(t *testing.T) {
		mock.ExpectQuery(findVideoIDsByUserID).
			WithArgs(int64(1)).
			WillReturnRows(mock.NewRows([]string{"video_id"}))
		like, err := model.FindVideoIDsByUserID(context.TODO(), 1, 2, 1)
		assertions.NoError(err)
		assertions.Equal(0, len(like))
	})
	const findWhetherLiked = "SELECT `video_id` FROM `likes` WHERE user_id = ? AND video_id IN (?,?)"
	t.Run("FindWhetherLiked success", func(t *testing.T) {
		// userid int64, videoID []int64
		mock.ExpectQuery(findWhetherLiked).
			WithArgs(int64(1), int64(1), int64(2)).
			WillReturnRows(mock.NewRows([]string{"video_id"}).AddRow(1))
		like, err := model.FindWhetherLiked(context.TODO(), 1, []int64{1, 2})
		assertions.NoError(err)
		assertions.Equal(1, len(like))
		assertions.Equal(int64(1), like[0])
	})
	t.Run("FindWhetherLiked nothing", func(t *testing.T) {
		mock.ExpectQuery(findWhetherLiked).
			WithArgs(2, 3, 5).
			WillReturnRows(mock.NewRows([]string{"video_id"}))
		like, err := model.FindWhetherLiked(context.TODO(), 2, []int64{3, 5})
		assertions.NoError(err)
		assertions.Empty(like)
	})
}

func TestLikeExec(t *testing.T) {
	assertions, mock, model := newMock(t, NewLikeCommand)
	const insert = "INSERT INTO `likes` (`video_id`,`user_id`,`created_at`,`id`) VALUES (?,?,?,?)"
	like := &dal.Like{
		VideoID: 1,
		UserID:  1,
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectExec(insert).
			WithArgs(like.VideoID, like.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Insert(context.TODO(), like)
		assertions.NoError(err)
		assertions.NotZero(like.ID)
		assertions.NotZero(like.CreatedAt)
	})
	t.Run("Insert fail", func(t *testing.T) {
		mock.ExpectExec(insert).
			WithArgs(like.VideoID, like.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errors.New("error"))
		err := model.Insert(context.TODO(), like)
		assertions.Error(err)
	})
	const deleteByVideoIDAndUserID = "DELETE FROM `likes` WHERE video_id = ? AND user_id = ?"
	t.Run("Delete success", func(t *testing.T) {
		mock.ExpectExec(deleteByVideoIDAndUserID).WithArgs(like.VideoID, like.UserID).
			WillReturnResult(
				sqlmock.NewResult(0, 1),
			)
		err := model.Delete(context.TODO(), like.VideoID, like.UserID)
		assertions.NoError(err)
	})
	t.Run("Delete fail", func(t *testing.T) {
		mock.ExpectExec(deleteByVideoIDAndUserID).
			WithArgs(like.VideoID, like.UserID).
			WillReturnError(errors.New("error"))
		err := model.Delete(context.TODO(), like.VideoID, like.UserID)
		assertions.Error(err)
	})
	t.Run("Delete nothing", func(t *testing.T) {
		mock.ExpectExec(deleteByVideoIDAndUserID).
			WithArgs(like.VideoID, like.UserID).
			WillReturnResult(sqlmock.NewResult(1, 0))
		err := model.Delete(context.TODO(), like.VideoID, like.UserID)
		assertions.Error(err)
	})
}
