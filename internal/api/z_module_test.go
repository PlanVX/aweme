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

import (
	"github.com/PlanVX/aweme/internal/logic"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

type Option struct {
	fx.In
	Pubs []*API `group:"public"`
	Opt  []*API `group:"optional"`
	Pri  []*API `group:"private"`
}

type Dep struct {
	fx.Out
	Register    *logic.Register
	Login       *logic.Login
	Feed        *logic.Feed
	Upload      *logic.Upload
	UserProfile *logic.UserProfile
	PublishList *logic.PublishList
	Like        *logic.Like
	LikeList    *logic.LikeList
	Comment     *logic.CommentAction
	CommentList *logic.CommentList
}

func TestGroupedValues(t *testing.T) {
	fxtest.New(t, Module, fx.Provide(func() Dep { return Dep{} }),
		fx.Invoke(func(option Option) {
			assert.NotEmpty(t, option.Pubs)
			assert.NotEmpty(t, option.Opt)
			assert.NotEmpty(t, option.Pri)
		}),
	)
}
