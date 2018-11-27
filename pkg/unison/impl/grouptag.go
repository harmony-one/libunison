package impl

import (
	"errors"
	"fmt"
)

// GroupTag is a binary tag that identifies a multicast group.
type GroupTag [16]uint8

// MarshalUnison marshals the group tag.
func (tag *GroupTag) MarshalUnison(buf []uint8) (written int, err error) {
	copy(buf, tag[:])
	written = len(tag)
	if len(buf) < written {
		err = fmt.Errorf("insufficient marshal buffer space for group tag %+v", tag)
	}
	return
}

// UnmarshalUnison unmarshals the group tag out of the given buf.
func (tag *GroupTag) UnmarshalUnison(buf []uint8) (consumed int, err error) {
	if len(buf) < len(tag) {
		err = errors.New("truncated group tag")
	} else {
		copy(tag[:], buf)
		consumed = len(tag)
	}
	return
}

