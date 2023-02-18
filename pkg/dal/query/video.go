package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"gorm.io/gorm"
	"time"
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

// FindOne find one video by id
func (c *CustomVideoModel) FindOne(ctx context.Context, id int64) (*dal.Video, error) {
	return c.v.WithContext(ctx).FindOne(id)
}

// FindLatest find the latest videos
// timestamp is the millisecond timestamp
func (c *CustomVideoModel) FindLatest(ctx context.Context, timestamp int64, limit int64) ([]*dal.Video, error) {
	t := time.Unix(0, timestamp*int64(time.Millisecond)) // convert millisecond to time.Time
	return c.v.WithContext(ctx).FindByTimestamp(t, limit)
}

// FindMany find many videos by ids
func (c *CustomVideoModel) FindMany(ctx context.Context, ids []int64) ([]*dal.Video, error) {
	return c.v.WithContext(ctx).FindMany(ids)
}

// FindByUserID find videos by user id
func (c *CustomVideoModel) FindByUserID(ctx context.Context, userID int64, limit, offset int) ([]*dal.Video, error) {
	return c.v.WithContext(ctx).FindByUserID(userID, limit, offset)
}

// Insert insert a video
func (c *CustomVideoModel) Insert(ctx context.Context, video *dal.Video) error {
	return c.v.WithContext(ctx).Create(video)
}

// Update update a video
func (c *CustomVideoModel) Update(context.Context, *dal.Video) error {
	return nil
}

// Delete a video by its id and correct user id
func (c *CustomVideoModel) Delete(ctx context.Context, id, userID int64) error {
	resultInfo, err := c.v.WithContext(ctx).Delete(&dal.Video{ID: id, UserID: userID})
	if err != nil {
		return err
	} else if resultInfo.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
