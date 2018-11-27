package unison

import (
	"github.com/harmony-one/libunison/pkg/unison/util"
)

// GroupID is a group identity.
// Implementations shall ensure that identities are comparable,
// so that they can be used as map keys.
type GroupID interface {
	util.Marshaler
	util.Unmarshaler
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

