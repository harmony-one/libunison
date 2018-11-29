package unison

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// GroupID is a group identity.
// Implementations shall ensure that identities are comparable,
// so that they can be used as map keys.
type GroupID interface {
	Marshaler
	Unmarshaler
}

// GroupTag is a binary tag that identifies a multicast group.
type GroupTag [16]uint8

// MarshalUnisonSize returns the number of bytes required to marshall this
// group tag.
func (tag *GroupTag) MarshalUnisonSize() (requiredSize int) {
	return len(tag)
}

// MarshalUnison marshals the group tag.
func (tag *GroupTag) MarshalUnison(buf []uint8) (written int, err error) {
	if err = CheckMarshalSize(tag, buf); err != nil {
		return
	}
	copy(buf, tag[:])
    written = tag.MarshalUnisonSize()
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

// NewGroupTag creates a new group tag from string or byte slice, by hashing it.
func NewGroupTag(source interface{}) (tag *GroupTag) {
	tag = new(GroupTag)
	switch src := source.(type) {
	case string:
		sum := sha256.Sum256([]byte(src))
		copy(tag[:], sum[:])
	case []byte:
		sum := sha256.Sum256(src)
		copy(tag[:], sum[:])
	default:
		panic(fmt.Sprintf("invalid group tag source %+v", src))
	}
	return
}

// NewGroupID creates a new group ID from string or byte slice, by hashing it.
func NewGroupID(source interface{}) (id GroupID) {
	return NewGroupTag(source)
}

// Group is a group of nodes, used for multicasting.
type Group interface {
	// GroupID returns the group ID.
	GroupID() GroupID

	// AddMember adds a member.
	AddMember(memberID NodeID)

	// RemoveMember removes a member.
	RemoveMember(memberID NodeID)

	// HasMember checks if the given identity is a member of the group.
	HasMember(memberID NodeID) bool

	// Members returns all identities in the group.
	// No specific ordering is imposed.
	Members() []NodeID

	// Protocol returns the group service instance for the given service ID.
	Service(id ProtocolID) Protocol
}
