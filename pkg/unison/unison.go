package unison

// Sender is a message sender.
type Sender interface {
	// Send a message to the destination.
	Send(msg []byte, destination Node) (err error)
}

// Receiver is a message receiver.
type Receiver interface {
	// Receive a message into the given buffer.
	//
	// Return the message length in bytes,
	// the source node from which the message came,
	// and a nil error on success.
	//
	// Note that msgLen > len(msg) means the received message is truncated.
	// On failure, return a non-nil error (msgLen and sender is undefined).
	Receive(msg []byte) (msgLen int, source Node, err error)
}

// Unison is a top-level Unison interface.
type Unison interface {
	// AddLocator associates the given locator with the given node ID.
	AddLocator(id NodeID, locator NodeLocator)
}