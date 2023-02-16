package api

import "go.uber.org/fx"

/*
the group tags are used to distinguish the different types of APIs.
1. public APIs are exposed to the public, such as register/login APIs.
2. optional APIs are APIs which can be accessed with or without authentication.
3. private APIs are APIs which can only be accessed with authentication.
*/

// annotated with `group:"public"`
func wrapPublic[T any](t T) any { return fx.Annotate(t, fx.ResultTags(`group:"public"`)) }

// annotated with `group:"optional"`
func wrapOptional[T any](t T) any { return fx.Annotate(t, fx.ResultTags(`group:"optional"`)) }

// annotated with `group:"private"`
func wrapPrivate[T any](t T) any { return fx.Annotate(t, fx.ResultTags(`group:"private"`)) }

// Module is the module for api
// it provides all the APIs
var Module = fx.Module("api",
	fx.Provide(
		wrapPublic(NewRegister),
		wrapPublic(NewLogin),
		wrapOptional(NewFeed),
		wrapOptional(NewUserInfo),
		wrapPrivate(NewUpload),
	),
)
