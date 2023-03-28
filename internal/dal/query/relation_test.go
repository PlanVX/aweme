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
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PlanVX/aweme/internal/dal"
	"testing"
)

func TestRelationFind(t *testing.T) {
	assertions, mock, model := newMock(t, NewRelationQuery)
	const FindWhetherFollowedList = "SELECT `follow_to` FROM `relations` WHERE follow_to IN (?,?) AND user_id = ?"
	t.Run("FindWhetherFollowed success", func(t *testing.T) {
		// userid int64, videoID []int64
		mock.ExpectQuery(FindWhetherFollowedList).
			WithArgs(int64(1), int64(2), int64(1)).
			WillReturnRows(mock.NewRows([]string{"followed_id"}).AddRow(1))
		like, err := model.FindWhetherFollowedList(context.TODO(), 1, []int64{1, 2})
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1}, like)
	})
	const FindFollowerTo = "SELECT `follow_to` FROM `relations` WHERE user_id = ? ORDER BY created_at desc LIMIT 3 OFFSET 1"
	t.Run("FindFollowerTo success", func(t *testing.T) {
		mock.ExpectQuery(FindFollowerTo).
			WithArgs(int64(1)).
			WillReturnRows(mock.NewRows([]string{"follow_to"}).AddRow(1).AddRow(2).AddRow(3))
		like, err := model.FindFollowerTo(context.TODO(), 1, 3, 1)
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1, 2, 3}, like)
	})
	const FindFollowerFrom = "SELECT `user_id` FROM `relations` WHERE follow_to = ? ORDER BY created_at desc LIMIT 1 OFFSET 1"
	t.Run("FindFollowerFrom success", func(t *testing.T) {
		mock.ExpectQuery(FindFollowerFrom).
			WithArgs(int64(5)).
			WillReturnRows(mock.NewRows([]string{"user_id"}).AddRow(1))
		like, err := model.FindFollowerFrom(context.TODO(), 5, 1, 1)
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1}, like)
	})
}

func TestRelationExec(t *testing.T) {
	assertions, mock, model := newMock(t, NewRelationCommand)
	const InsertRelation = "INSERT INTO `relations` (`user_id`,`follow_to`,`created_at`,`id`) VALUES (?,?,?,?)"
	rel := &dal.Relation{
		UserID:   1,
		FollowTo: 1,
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectExec(InsertRelation).
			WithArgs(rel.UserID, rel.FollowTo, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Insert(context.TODO(), rel)
		assertions.NoError(err)
		assertions.NotZero(rel.CreatedAt)
		assertions.NotZero(rel.ID)
	})
	const DeleteRelation = "DELETE FROM `relations` WHERE user_id = ? AND follow_to = ?"
	t.Run("Delete success", func(t *testing.T) {
		mock.ExpectExec(DeleteRelation).
			WithArgs(rel.UserID, rel.FollowTo).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Delete(context.TODO(), rel.UserID, rel.FollowTo)
		assertions.NoError(err)
	})
}
