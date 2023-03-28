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
		// 1. public APIs
		wrapPublic(NewRegister),
		wrapPublic(NewLogin),
		// 2. optional APIs
		wrapOptional(NewUserInfo),
		wrapOptional(NewFeed),
		wrapOptional(NewPublishList),
		wrapOptional(NewCommentList),
		wrapOptional(NewFavoriteList),
		// 3. private APIs
		wrapPrivate(NewUpload),
		wrapPrivate(NewFavoriteAction),
		wrapPrivate(NewCommentAction),
		wrapPrivate(NewRelationAction),
		wrapPrivate(NewRelationFollowList),
		wrapPrivate(NewRelationFollowerList),
		wrapPrivate(NewRelationFriendList),
		wrapPrivate(NewMessageChat),
		wrapPrivate(NewMessageAction),
	),
)
