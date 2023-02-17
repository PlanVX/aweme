package main

import (
	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gen"
)

// FindByUsername defines the find by username interface
type FindByUsername interface {
	// FindByUsername
	//
	// where("username=@name")
	FindByUsername(name string) (*gen.T, error)
}

// FindByID find by id or id list
type FindByID interface {
	// FindOne find one item by id
	//
	// where("id=@id")
	FindOne(id int64) (*gen.T, error)

	// FindMany find many items by ids
	//
	// where("id in @ids")
	FindMany(ids []int64) ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.ApplyBasic(new(dal.User), new(dal.Video))
	g.ApplyInterface(func(FindByID) {}, new(dal.User), new(dal.Video))
	g.ApplyInterface(func(FindByUsername) {}, new(dal.User))
	g.Execute()
}
