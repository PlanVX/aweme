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
