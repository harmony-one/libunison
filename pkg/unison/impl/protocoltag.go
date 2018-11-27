package impl

import (
	"errors"
	"fmt"
)

// ProtocolTag is a binary tag that identifies a service provided by a node.
// Each network service randomly chooses and publish its own tag.
// Think of it as a protocol tag, akin to well-known port numbers.
type ProtocolTag [16]uint8

// MarshalUnison marshals the service tag into the given buffer.
func (tag *ProtocolTag) MarshalUnison(buf []uint8) (written int, err error) {
	copy(buf, tag[:])
	written = len(tag)
	if len(buf) < written {
		err = fmt.Errorf("insufficient marshal buffer space for service tag %+v", tag)
	}
	return
}

// UnmarshalUnison unmarshals the service tag out of the given buffer.
func (tag *ProtocolTag) UnmarshalUnison(buf []uint8) (consumed int, err error) {
	if len(buf) < len(tag) {
		err = errors.New("truncated service tag")
	} else {
		copy(tag[:], buf)
		consumed = len(tag)
	}
	return
}

