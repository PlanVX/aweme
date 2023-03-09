package logic

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	// Like is the like logic layer struct
	Like struct {
		likeModel dal.LikeModel
	}
	// LikeParam is the param for NewLike
	LikeParam struct {
		fx.In
		LikeModel dal.LikeModel
	}
)

// NewLike returns a new Like logic
func NewLike(param LikeParam) *Like {
	return &Like{
		likeModel: param.LikeModel,
	}
}

// Like is the like logic
// handle the like action
func (l *Like) Like(c context.Context, req *types.FavoriteActionReq) (*types.FavoriteActionResp, error) {
	owner, _ := c.Value(ContextKey).(int64) // get the owner from context
	switch req.ActionType {
	case int32(1): // means add like for a video

		like := &dal.Like{
			VideoID: req.VideoID,
			UserID:  owner,
		}

		err := l.likeModel.Insert(c, like)
		if err != nil {
			return nil, err
		}
	case int32(2): // means remove like for a video

		err := l.likeModel.Delete(c, req.VideoID, owner)
		if err != nil {
			return nil, err
		}
	default:
		return nil, echo.ErrBadRequest
	}
	return &types.FavoriteActionResp{}, nil
}
