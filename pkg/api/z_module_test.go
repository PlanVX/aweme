package api

import (
	"github.com/PlanVX/aweme/pkg/logic"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

type Option struct {
	fx.In
	Pubs []*Api `group:"public"`
	Opt  []*Api `group:"optional"`
	Pri  []*Api `group:"private"`
}

type Dep struct {
	fx.Out
	Register    *logic.Register
	Login       *logic.Login
	Feed        *logic.Feed
	Upload      *logic.Upload
	UserProfile *logic.UserProfile
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
