package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

// check if VideoModel implements VideoModel interface
var _ dal.VideoModel = (*CustomVideoModel)(nil)

// CustomVideoModel is the implementation of VideoModel
type CustomVideoModel struct {
	v        video                 // video is the gorm/gen generated query struct
	rdb      redis.UniversalClient // rdb is the redis client
	uniqueID *UniqueID             // uniqueID is the unique id generator
}

// NewVideoModel returns a *CustomVideoModel
func NewVideoModel(db video, rdb redis.UniversalClient) *CustomVideoModel {
	return &CustomVideoModel{
		v:        db,
		rdb:      rdb,
		uniqueID: NewUniqueID(),
	}
}

// FindOne find one video by id
func (c *CustomVideoModel) FindOne(ctx context.Context, id int64) (*dal.Video, error) {
	return c.v.WithContext(ctx).FindOne(id)
}

// FindLatest find the latest videos
// timestamp is the millisecond timestamp
func (c *CustomVideoModel) FindLatest(ctx context.Context, timestamp int64, limit int64) ([]*dal.Video, error) {
	var t time.Time
	if timestamp == 0 { // if timestamp is 0, use current time
		t = time.Now()
	} else {
		t = time.Unix(0, timestamp*int64(time.Millisecond))
	}
	result, err := c.v.WithContext(ctx).FindByTimestamp(t, limit)
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, result)
}

// FindMany find many videos by ids
func (c *CustomVideoModel) FindMany(ctx context.Context, ids []int64) ([]*dal.Video, error) {
	result, err := c.v.WithContext(ctx).FindMany(ids)
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, result)
}

// FindByUserID find videos by user id
func (c *CustomVideoModel) FindByUserID(ctx context.Context, userID int64, limit, offset int) ([]*dal.Video, error) {
	result, err := c.v.WithContext(ctx).FindByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.FindManyStat(ctx, result)
}

// Insert insert a video
func (c *CustomVideoModel) Insert(ctx context.Context, video *dal.Video) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	video.ID = id
	err = c.v.WithContext(ctx).Create(video)
	if err != nil {
		return err
	}
	c.rdb.HIncrBy(ctx, GenRedisKey(TableUser, video.UserID), CountVideo, 1)
	return nil
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

// FindOneStat find one video stat by id
func (c *CustomVideoModel) FindOneStat(ctx context.Context, video *dal.Video) (*dal.Video, error) {
	if err := c.rdb.HGetAll(ctx, GenRedisKey(TableVideo, video.ID)).Scan(video); err != nil {
		return nil, err
	}
	return video, nil
}

// FindManyStat find many video stats by ids
func (c *CustomVideoModel) FindManyStat(ctx context.Context, videos []*dal.Video) ([]*dal.Video, error) {
	cmder, err := c.rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, v := range videos {
			pipe.HGetAll(ctx, GenRedisKey(TableVideo, v.ID))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for i, cmd := range cmder {
		err := cmd.(*redis.MapStringStringCmd).Scan(videos[i])
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
}
