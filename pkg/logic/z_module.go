// Package logic provides the business logic for the application
package logic

import (
	"go.uber.org/fx"
)

// Module is the business logic module
// It provides the business logic for the application
var Module = fx.Module("logic",
	fx.Provide(
		NewRegister,
		NewLogin,
		NewFeed,
		NewUpload,
		NewUserProfile,
		NewPublishList,
		NewJWTSigner,
	))
