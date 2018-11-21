package unison

import (
	"net"
	"time"
)

// Identity is a host identity.
type Identity interface {
	// TODO
}

// ServiceTag is a binary tag that identifies a service provided by a host.
// Each network service randomly chooses and publish its own tag.
type ServiceTag [16]byte

// Self is a local Unison interface.
type Self interface {
	// DialPeer connects to a remote peer.
	DialPeer(peerId Identity) (Peer, error)

	// DialService connects to a service on a remote peer.
	DialService(peerId Identity, tag ServiceTag) (Peer, Conn, error)

	// AddLocator manually adds an association between an identity and a
	// locator.  For use by discovery mechanisms.  Also usable for updating the
	// lifetime of the association.
	AddLocator(id Identity, addr net.IPAddr, lifetime time.Duration) error

	// RemoveLocator manually removes an association between an identity and a
	// locator.  For use by discovery mechanisms.
	RemoveLocator(id Identity, addr net.IPAddr) error
}

// Peer is a self-to-peer connection.
type Peer interface {
	// DialService dials a service and returns a connection to it.
	// service is the service key.  This always creates a new connection,
	// even if another connection to the same service is already open.
	DialService(tag ServiceTag) (Conn, error)

	// Close closes the peer association, and all connections over it, if any.
	Close()
}

// Conn is a service connection.
type Conn interface {
	// Send sends a message.
	Send(data []byte) error

	// Recv receives a message and places it into the given buffer.  On success,
	// it returns the number of octets written into the buffer and nil error.
	// Note: If this number is greater than len(buf), the message has been
	// truncated.  On failure, it returns 0 and a non-nil error.
	Receive(buf []byte) (uint, error)

	// Close closes the connection.
	Close()
}
