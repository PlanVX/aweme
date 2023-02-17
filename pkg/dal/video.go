package dal

import "context"

// check if VideoModel implements VideoModel interface
var _ VideoModel = (*CustomVideoModel)(nil)

// CustomVideoModel is the implementation of VideoModel
type CustomVideoModel struct {
}

// NewVideoModel returns a *CustomVideoModel
func NewVideoModel() *CustomVideoModel {
	return &CustomVideoModel{}
}

// FindLatest find latest videos
func (c CustomVideoModel) FindLatest(context.Context, int64, int64) ([]*Video, error) {
	return []*Video{}, nil
}

// FindOne find one video by id
func (c CustomVideoModel) FindOne(context.Context, int64) (*Video, error) {
	return new(Video), nil
}

// FindMany find many videos by ids
func (c CustomVideoModel) FindMany(context.Context, []int64) ([]*Video, error) {
	return nil, nil
}

// Insert insert a video
func (c CustomVideoModel) Insert(context.Context, *Video) error {
	return nil
}

// Update update a video
func (c CustomVideoModel) Update(context.Context, *Video) error {
	return nil
}

// Delete delete a video by id and user id
func (c CustomVideoModel) Delete(context.Context, int64, int64) error {
	return nil
}
