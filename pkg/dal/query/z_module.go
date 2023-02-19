// Package query data access layer
package query

import (
	"github.com/PlanVX/aweme/pkg/dal"
	"go.uber.org/fx"
)

// Module is the module for dal.
// It provides the data access layer for the application.
// fx.Annotate is used to wrap the struct with an interface.
var Module = fx.Module("data access layer",
	fx.Provide(
		NewGormDB,
		fx.Annotate(NewUserModel, fx.As(new(dal.UserModel))),
		fx.Annotate(NewVideoModel, fx.As(new(dal.VideoModel))),
		fx.Annotate(NewLikeModel, fx.As(new(dal.LikeModel))),
		fx.Annotate(NewCommentModel, fx.As(new(dal.CommentModel))),
	))
