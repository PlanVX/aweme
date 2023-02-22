// Package query data access layer
package query

import (
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module is the module for dal.
// It provides the data access layer for the application.
// fx.Annotate is used to wrap the struct with an interface.
var Module = fx.Module("data access layer",
	fx.Provide(
		NewGormDB,
		newModel,
		NewRedisUniversalClient,
	))

// newModel returns the data access layer models.
func newModel(db *gorm.DB, rds redis.UniversalClient) (
	dal.UserModel,
	dal.CommentModel,
	dal.VideoModel,
	dal.RelationModel,
	dal.LikeModel) {
	use := Use(db)
	return NewUserModel(use.User, rds),
		NewCommentModel(use.Comment, rds),
		NewVideoModel(use.Video, rds),
		NewRelationModel(use.Relation, rds),
		NewLikeModel(use.Like, rds)
}
