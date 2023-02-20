package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCommentList(t *testing.T) {
	var id int64 = 1
	ctx := context.WithValue(context.Background(), ContextKey, id)
	assertions := assert.New(t)
	mockComment := []*dal.Comment{{ID: 1, UserID: 1}, {ID: 2, UserID: 2}, {ID: 3, UserID: 3}}
	mockUser := []*dal.User{{ID: 1}, {ID: 2}, {ID: 3}}
	mockFollowTo := []int64{2}
	t.Run("test comment list query success", func(t *testing.T) {
		u := NewUserModel(t)
		c := NewCommentModel(t)
		r := NewRelationModel(t)
		l := NewCommentList(CommentListParam{
			UserModel:     u,
			CommentModel:  c,
			RelationModel: r,
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
