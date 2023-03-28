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

package logic

import (
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCommentList(t *testing.T) {
	var id int64 = 1
	ctx := ContextWithOwner(id)
	assertions := assert.New(t)
	mockComment := []*dal.Comment{{ID: 1, UserID: 1}, {ID: 2, UserID: 2}, {ID: 3, UserID: 3}}
	mockUser := []*dal.User{{ID: 1}, {ID: 2}, {ID: 3}}
	mockFollowTo := []int64{2}
	t.Run("test comment list query success", func(t *testing.T) {
		u := NewUserQuery(t)
		c := NewCommentQuery(t)
		r := NewRelationQuery(t)
		l := NewCommentList(CommentListParam{
			UserQuery:     u,
			CommentQuery:  c,
			RelationQuery: r,
		})
		assertions.NotNil(l)
		c.On("FindByVideoID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockComment, nil)
		u.On("FindMany", mock.Anything, mock.Anything).Return(mockUser, nil)
		r.On("FindWhetherFollowedList", mock.Anything, mock.Anything, mock.Anything).Return(mockFollowTo, nil)
		list, err := l.CommentList(ctx, &types.CommentListReq{VideoID: 1})
		assertions.NoError(err)
		assertions.NotNil(list)
		assertions.Equal(3, len(list.CommentList))
		assertions.Equal(false, list.CommentList[0].User.IsFollow)
		assertions.Equal(true, list.CommentList[1].User.IsFollow)
		assertions.Equal(false, list.CommentList[2].User.IsFollow)
	})
}
