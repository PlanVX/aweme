package dal

import "context"

//go generate for generating mock files for logic layer testing

// UserModel is the interface for user model operations
//
//go:generate mockery --name=UserModel --output=../logic --filename=mock_u_test.go --outpkg=logic
type UserModel interface {
	FindOne(ctx context.Context, id int64) (*User, error)
	FindMany(ctx context.Context, ids []int64) ([]*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

// VideoModel is the interface for video model operations
//
//go:generate mockery --name=VideoModel --output=../logic --filename=mock_v_test.go --outpkg=logic
type VideoModel interface {
	FindOne(ctx context.Context, id int64) (*Video, error)
	FindMany(ctx context.Context, ids []int64) ([]*Video, error)
	FindLatest(ctx context.Context, latestTime, limit int64) ([]*Video, error)
	FindByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Video, error)
	Insert(ctx context.Context, video *Video) error
	Update(ctx context.Context, video *Video) error
	Delete(ctx context.Context, id int64, uid int64) error
}

// CommentModel is the interface for comment model operations
//
//go:generate mockery --name=CommentModel --output=../logic --filename=mock_c_test.go --outpkg=logic
type CommentModel interface {
	FindByVideoID(ctx context.Context, videoID int64, limit, offset int) ([]*Comment, error)
	Insert(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id int64, uid int64) error
}

// LikeModel is the interface for like model operations
//
//go:generate mockery --name=LikeModel --output=../logic --filename=mock_l_test.go --outpkg=logic
type LikeModel interface {
	//Insert inserts a like record
	Insert(ctx context.Context, like *Like) error
	//Delete deletes a like record by video id and user id
	Delete(ctx context.Context, vid, uid int64) error

	// FindByVideoIDAndUserID finds a like record by video id and user id
	FindByVideoIDAndUserID(ctx context.Context, vid, uid int64) (*Like, error)
}
