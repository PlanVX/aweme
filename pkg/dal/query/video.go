package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gorm"
)

// check if VideoModel implements VideoModel interface
var _ dal.VideoModel = (*CustomVideoModel)(nil)

// CustomVideoModel is the implementation of VideoModel
type CustomVideoModel struct {
	v  video    // video is the gorm/gen generated query struct
	db *gorm.DB // db is the gorm db instance
}

// NewVideoModel returns a *CustomVideoModel
func NewVideoModel(db *gorm.DB) *CustomVideoModel {
	return &CustomVideoModel{
		db: db,
		v:  Use(db).Video,
	}
}

// FindLatest find latest videos
func (c CustomVideoModel) FindLatest(context.Context, int64, int64) ([]*dal.Video, error) {
	return []*dal.Video{}, nil
}

// FindOne find one video by id
func (c CustomVideoModel) FindOne(ctx context.Context, id int64) (*dal.Video, error) {
	return c.v.WithContext(ctx).FindOne(id)
}

// FindMany find many videos by ids
func (c CustomVideoModel) FindMany(ctx context.Context, ids []int64) ([]*dal.Video, error) {
	return c.v.WithContext(ctx).FindMany(ids)
}

// Insert insert a video
func (c CustomVideoModel) Insert(ctx context.Context, video *dal.Video) error {
	return c.v.WithContext(ctx).Create(video)
}

// Update update a video
func (c CustomVideoModel) Update(context.Context, *dal.Video) error {
	return nil
}

// Delete delete a video by id and user id
func (c CustomVideoModel) Delete(context.Context, int64, int64) error {
	return nil
}
