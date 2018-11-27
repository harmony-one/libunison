package unison

import (
	"github.com/harmony-one/libunison/pkg/unison/util"
	"time"
)

// NodeID is a node identity.
// Implementations shall ensure that identities are comparable,
// so that they can be used as map keys.
type NodeID interface {
	util.Marshaler
	util.Unmarshaler
}

// NodeLocator is a node locator (that is, for now, an IP address).
// Implementations shall ensure that locators are comparable,
// so that they can be used as map keys.
type NodeLocator interface {
	util.Marshaler
	util.Unmarshaler
}

// Node is a node interface, common to both local and remote.
type Node interface {
	// ID returns the node ID.
	NodeID() NodeID

	// Protocol returns the node service instance for the given service ID.
	Service(id ProtocolID) Protocol

	// AddLocator manually adds an association between an identity and a
	// locator.
	// For use by discovery mechanisms.
	// Also usable for updating the lifetime of the association.
	AddLocator(loc NodeLocator, lifetime time.Duration) (err error)

	// RemoveLocator manually removes an association between an identity and
	// a locator.
	// For use by discovery mechanisms.
	RemoveLocator(loc NodeLocator) (err error)
}

