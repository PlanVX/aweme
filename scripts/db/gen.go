package main

import (
	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gen"
	"time"
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

// Video find by user id
type Video interface {
	// FindByUserID
	//
	// select id, user_id, title, created_at, video_url, cover_url
	// from videos
	// where user_id=@id
	// order by created_at desc
	// limit @limit offset @offset
	FindByUserID(id int64, limit, offset int) ([]*gen.T, error)

	// FindByTimestamp
	//
	//  SELECT id, user_id, title, created_at, video_url, cover_url
	//  FROM `videos`
	//  WHERE `videos`.`created_at` > @timestamp
	//  ORDER BY `videos`.`created_at` desc
	//  LIMIT @limit
	FindByTimestamp(timestamp time.Time, limit int64) ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.ApplyBasic(new(dal.User), new(dal.Video))
	g.ApplyInterface(func(FindByID) {}, new(dal.User), new(dal.Video))
	g.ApplyInterface(func(FindByUsername) {}, new(dal.User))
	g.ApplyInterface(func(Video) {}, new(dal.Video))
	g.Execute()
}
