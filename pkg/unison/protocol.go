package unison

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// ProtocolID is a protocol identity.
// Implementations shall ensure that identities are comparable,
// so that they can be used as map keys.
type ProtocolID interface {
	Marshaler
	Unmarshaler
}

// ProtocolTag is a binary tag that identifies a service provided by a node.
// Each network service randomly chooses and publish its own tag.
// Think of it as a protocol tag, akin to well-known port numbers.
type ProtocolTag [16]uint8

// MarshalUnisonSize returns the number of bytes required to marshall this
// protocol tag.
func (tag *ProtocolTag) MarshalUnisonSize() (marshalSize int) {
	return len(tag)
}

// MarshalUnison marshals the service tag into the given buffer.
func (tag *ProtocolTag) MarshalUnison(buf []uint8) (written int, err error) {
	if err = CheckMarshalSize(tag, buf); err != nil {
		return
	}
	copy(buf, tag[:])
	written = tag.MarshalUnisonSize()
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

// NewProtocolTag creates a new protocol tag from string or byte slice,
// by hashing it.
func NewProtocolTag(source interface{}) (tag *ProtocolTag) {
	tag = new(ProtocolTag)
	switch src := source.(type) {
	case string:
	case []byte:
		sum := sha256.Sum256([]byte(src))
		copy(tag[:], sum[:])
	default:
        panic(fmt.Sprintf("invalid protocol tag source %+v", src))
	}
	return
}

// NewProtocolID creates a new protocol ID from string or byte slice,
// by hashing it.
func NewProtocolID(source interface{}) (id ProtocolID) {
	return NewProtocolTag(source)
}

// Protocol is a service endpoint,
// bound to one protocol identified by a service ID.
// Protocol can be used for sending or receiving a message.
type Protocol interface {
	// ProtocolID returns the service ID.
	ProtocolID() (id ProtocolID)

	// LocalNodes returns the local nodes that this service uses to send and/or
	// receive messages.  The returned slice contains at least one entry.
	LocalNodes() (nodes []Node)

	// AddLocalNode adds a local node for which
}
