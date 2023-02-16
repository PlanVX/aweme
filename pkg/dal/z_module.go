// Package dal data access layer
package dal

import "go.uber.org/fx"

// Module is the module for dal.
// It provides the data access layer for the application.
// fx.Annotate is used to wrap the struct with an interface.
var Module = fx.Module("data access layer",
	fx.Provide(
		fx.Annotate(NewUserModel, fx.As(new(UserModel))),
		fx.Annotate(NewVideoModel, fx.As(new(VideoModel))),
	))
