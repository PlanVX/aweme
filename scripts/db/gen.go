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

// Like defines the like model sql queries
type Like interface {
	// FindByVideoIDAndUserID
	//
	// select id, user_id, video_id, created_at
	// from likes
	// where video_id=@videoID and user_id=@userID
	FindByVideoIDAndUserID(videoID, userID int64) (*gen.T, error)
	// DeleteByVideoIDAndUserID delete by video id and user id
	//
	// delete from likes
	// where video_id=@videoID and user_id=@userID
	DeleteByVideoIDAndUserID(videoID, userID int64) (gen.RowsAffected, error)
	//FindVideoIDsByUserID
	//
	// select video_id
	// from likes
	// where user_id=@userID limit @limit offset @offset
	FindVideoIDsByUserID(userID int64, limit, offset int) ([]int64, error)
}

// Comment  defines the comment model sql queries
type Comment interface {
	// FindByVideoID
	//
	// select id, user_id, video_id, content, created_at
	// from comments
	// where video_id=@id
	// order by created_at desc
	// limit @limit offset @offset
	FindByVideoID(id int64, limit, offset int) ([]*gen.T, error)
	// DeleteByIDAndUserID delete comment by id and user id
	//
	// delete from comments
	// where id=@id and user_id=@userID
	DeleteByIDAndUserID(id, userID int64) (gen.RowsAffected, error)
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
	g.ApplyInterface(func(Like) {}, new(dal.Like))
	g.ApplyInterface(func(Comment) {}, new(dal.Comment))
	g.Execute()
}
