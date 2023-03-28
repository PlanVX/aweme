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

package query

import (
	"errors"
	"github.com/sony/sonyflake"
)

// UniqueID provides unique id generation
type UniqueID struct {
	sf *sonyflake.Sonyflake
}

// NewUniqueID creates a new unique id generator
func NewUniqueID() *UniqueID {
	// the default machine id is the lower 16 bits of the private IP address
	// for large scale deployment,
	// it is recommended to provide a custom machine id
	return &UniqueID{sf: sonyflake.NewSonyflake(sonyflake.Settings{})}
}

// NextID generates a new unique id
// according to the generator's algorithm
func (u *UniqueID) NextID() (int64, error) {
	if u.sf == nil {
		return 0, errors.New("unique id generator is not initialized")
	}
	uid, err := u.sf.NextID()
	if err != nil {
		return 0, err
	} else {
		// the 1st bit is now unused,
		// so it is safe to convert to int64
		return int64(uid), nil
	}
}
