package query

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentFind(t *testing.T) {
	assertions, mock, model := commentTest(t)
	const findCommentsByVideoID = "SELECT * FROM `comments` WHERE video_id = ? LIMIT 2 OFFSET 10"
	t.Run("FindByVideoID success", func(t *testing.T) {
		// video_id int64,limit int,offset int
		mock.ExpectQuery(findCommentsByVideoID).
			WithArgs(int64(1)).
			WillReturnRows(mock.NewRows([]string{"id", "video_id"}).AddRow(1, 1).AddRow(2, 1))
		comment, err := model.FindByVideoID(context.TODO(), 1, 2, 10)
		assertions.NoError(err)
		assertions.Equal(2, len(comment))
		assertions.Equal(int64(1), comment[0].VideoID)
		assertions.Equal(int64(1), comment[1].VideoID)
	})
	t.Run("FindByVideoID nothing", func(t *testing.T) {
		mock.ExpectQuery(findCommentsByVideoID).
			WithArgs(8).
			WillReturnRows(mock.NewRows([]string{"id"}))
		comment, err := model.FindByVideoID(context.TODO(), 8, 2, 10)
		assertions.NoError(err)
		assertions.Len(comment, 0)
	})
}

func commentTest(t *testing.T) (*assert.Assertions, sqlmock.Sqlmock, *CommentModel) {
	assertions := assert.New(t)
	mock, db, rdb, err := mockDB(t)
	assertions.NoError(err)
	model := NewCommentModel(db, rdb)
	return assertions, mock, model
}

func TestCommentExec(t *testing.T) {
	assertions, mock, model := commentTest(t)
	const insertComment = "INSERT INTO `comments` (`content`,`video_id`,`user_id`,`created_at`,`id`) VALUES (?,?,?,?,?)"
	c := &dal.Comment{
		VideoID: 1,
		Content: "test",
		UserID:  1,
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(insertComment).
			WithArgs(c.Content, c.VideoID, c.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := model.Insert(context.TODO(), c)
		assertions.NoError(err)
		assertions.NotZero(c.ID)
		assertions.NotZero(c.CreatedAt)
	})
	t.Run("Insert failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(insertComment).
			WithArgs(c.Content, c.VideoID, c.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		err := model.Insert(context.TODO(), c)
		assertions.Error(err)
	})
	const deleteComment = "DELETE FROM `comments` WHERE id = ? AND user_id = ?"
	t.Run("Delete success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(deleteComment).
			WithArgs(c.ID, c.UserID).
			WillReturnResult(
				sqlmock.NewResult(1, 1),
			)
		mock.ExpectCommit()
		err := model.Delete(context.TODO(), c.ID, c.UserID, c.VideoID)
		assertions.NoError(err)
	})
	t.Run("Delete failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(deleteComment).
			WithArgs(c.ID, c.UserID).
			WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		err := model.Delete(context.TODO(), c.ID, c.UserID, c.VideoID)
		assertions.Error(err)
	})
	t.Run("Delete nothing", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(deleteComment).WithArgs(c.ID, c.UserID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()
		err := model.Delete(context.TODO(), c.ID, c.UserID, c.VideoID)
		assertions.Error(err)
	})
}
