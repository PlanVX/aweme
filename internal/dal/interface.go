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

package dal

import "context"

//go generate for generating mock files for logic layer testing

// UserQuery is the interface for user model operations
//
//go:generate mockery --name=UserQuery --output=../logic --filename=mock_u_test.go --outpkg=logic
type UserQuery interface {
	FindOne(ctx context.Context, id int64) (*User, error)
	FindMany(ctx context.Context, ids []int64) ([]*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

// VideoQuery is the interface for video model operations
//
//go:generate mockery --name=VideoQuery --output=../logic --filename=mock_v_test.go --outpkg=logic
type VideoQuery interface {
	FindOne(ctx context.Context, id int64) (*Video, error)
	FindMany(ctx context.Context, ids []int64) ([]*Video, error)
	FindLatest(ctx context.Context, latestTime, limit int64) ([]*Video, error)
	FindByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Video, error)
}

// CommentQuery is the interface for comment model operations
//
//go:generate mockery --name=CommentQuery --output=../logic --filename=mock_c_test.go --outpkg=logic
type CommentQuery interface {
	FindByVideoID(ctx context.Context, videoID int64, limit, offset int) ([]*Comment, error)
}

// LikeQuery is the interface for like model operations
//
//go:generate mockery --name=LikeQuery --output=../logic --filename=mock_l_test.go --outpkg=logic
type LikeQuery interface {
	// FindByVideoIDAndUserID finds a like record by video id and user id
	FindByVideoIDAndUserID(ctx context.Context, vid, uid int64) (*Like, error)

	// FindVideoIDsByUserID finds video ids by user id
	FindVideoIDsByUserID(ctx context.Context, uid int64, limit, offset int) ([]int64, error)

	// FindWhetherLiked finds a like record by video id and user id
	// return a list of id that liked by userid
	FindWhetherLiked(ctx context.Context, userid int64, videoID []int64) ([]int64, error)
}

// RelationQuery is the interface for relation model operations
//
//go:generate mockery --name=RelationQuery --output=../logic --filename=mock_r_test.go --outpkg=logic
type RelationQuery interface {

	// FindWhetherFollowedList finds a relation record by userid and followTo
	// return a list of id that followed by userid
	FindWhetherFollowedList(ctx context.Context, userid int64, followTo []int64) ([]int64, error)

	// FindFollowerTo finds a relation record by userid
	// which means find the user who userid follows
	// return a list of user id
	FindFollowerTo(ctx context.Context, userid int64, limit, offset int) ([]int64, error)

	// FindFollowerFrom finds a relation record by followTo
	// which means find the user who followTo is followed by
	// return a list of user id
	FindFollowerFrom(ctx context.Context, followTo int64, limit, offset int) ([]int64, error)
}

// UserCommand is the interface for user model operations
//
//go:generate mockery --name=UserCommand --output=../logic --filename=mock_uc_test.go --outpkg=logic
type UserCommand interface {
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

// VideoCommand is the interface for video model operations
//
//go:generate mockery --name=VideoCommand --output=../logic --filename=mock_vc_test.go --outpkg=logic
type VideoCommand interface {
	Insert(ctx context.Context, video *Video) error
	Update(ctx context.Context, video *Video) error
	Delete(ctx context.Context, id int64, uid int64) error
}

// LikeCommand is the interface for like model operations
//
//go:generate mockery --name=LikeCommand --output=../logic --filename=mock_lc_test.go --outpkg=logic
type LikeCommand interface {
	//Insert inserts a like record
	Insert(ctx context.Context, like *Like) error
	//Delete deletes a like record by video id and user id
	Delete(ctx context.Context, vid, uid int64) error
}

// CommentCommand is the interface for comment model operations
//
//go:generate mockery --name=CommentCommand --output=../logic --filename=mock_cc_test.go --outpkg=logic
type CommentCommand interface {
	Insert(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id int64, uid int64, vid int64) error
}

// RelationCommand is the interface for relation model operations
//
//go:generate mockery --name=RelationCommand --output=../logic --filename=mock_rc_test.go --outpkg=logic
type RelationCommand interface {
	//Insert inserts a relation record
	Insert(ctx context.Context, rel *Relation) error
	//Delete deletes a relation record by userid and followTo
	Delete(ctx context.Context, userid, followTo int64) error
}
