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

// Package query data access layer
package query

import (
	"github.com/PlanVX/aweme/internal/dal"
	"go.uber.org/fx"
)

// Module is the module for dal.
// It provides the data access layer for the application.
// fx.Annotate is used to wrap the struct with an interface.
var Module = fx.Module("data access layer",
	fx.Provide(
		fx.Annotate(NewGormDB, fx.OnStop(gormClose)),
		fx.Annotate(NewRedisUniversalClient, fx.OnStop(closeRedis)),
		fx.Annotate(NewUserQuery, fx.As(new(dal.UserQuery))),
		fx.Annotate(NewUserCommand, fx.As(new(dal.UserCommand))),
		fx.Annotate(NewCommentQuery, fx.As(new(dal.CommentQuery))),
		fx.Annotate(NewCommentCommand, fx.As(new(dal.CommentCommand))),
		fx.Annotate(NewVideoQuery, fx.As(new(dal.VideoQuery))),
		fx.Annotate(NewVideoCommand, fx.As(new(dal.VideoCommand))),
		fx.Annotate(NewLikeQuery, fx.As(new(dal.LikeQuery))),
		fx.Annotate(NewLikeCommand, fx.As(new(dal.LikeCommand))),
		fx.Annotate(NewRelationQuery, fx.As(new(dal.RelationQuery))),
	),
	fx.Decorate(RedisOtel, withPlugins),
)
