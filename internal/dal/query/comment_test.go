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
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/internal/dal"
	"testing"
)

func TestCommentFind(t *testing.T) {
	assertions, mock, model := newMock(t, NewCommentQuery)
	const findCommentsByVideoID = "SELECT `comments`.`id`,`comments`.`content`,`comments`.`video_id`,`comments`.`user_id`,`comments`.`created_at` FROM `comments` WHERE video_id = ? LIMIT 2 OFFSET 10"
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

func TestCommentExec(t *testing.T) {
	assertions, mock, model := newMock(t, NewCommentCommand)
	const insertComment = "INSERT INTO `comments` (`content`,`video_id`,`user_id`,`created_at`,`id`) VALUES (?,?,?,?,?)"
	c := &dal.Comment{
		VideoID: 1,
		Content: "test",
		UserID:  1,
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectExec(insertComment).
			WithArgs(c.Content, c.VideoID, c.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Insert(context.TODO(), c)
		assertions.NoError(err)
		assertions.NotZero(c.ID)
		assertions.NotZero(c.CreatedAt)
	})
	t.Run("Insert failed", func(t *testing.T) {
		mock.ExpectExec(insertComment).
			WithArgs(c.Content, c.VideoID, c.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errors.New("failed"))
		err := model.Insert(context.TODO(), c)
		assertions.Error(err)
	})
	const deleteComment = "DELETE FROM `comments` WHERE id = ? AND user_id = ?"
	t.Run("Delete success", func(t *testing.T) {
		mock.ExpectExec(deleteComment).
			WithArgs(c.ID, c.UserID).
			WillReturnResult(
				sqlmock.NewResult(1, 1),
			)
		err := model.Delete(context.TODO(), c.ID, c.UserID, c.VideoID)
		assertions.NoError(err)
	})
	t.Run("Delete failed", func(t *testing.T) {
		mock.ExpectExec(deleteComment).
			WithArgs(c.ID, c.UserID).
			WillReturnError(errors.New("failed"))
		err := model.Delete(context.TODO(), c.ID, c.UserID, c.VideoID)
		assertions.Error(err)
	})
	t.Run("Delete nothing", func(t *testing.T) {
		mock.ExpectExec(deleteComment).WithArgs(c.ID, c.UserID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		err := model.Delete(context.TODO(), c.ID, c.UserID, c.VideoID)
		assertions.Error(err)
	})
}
