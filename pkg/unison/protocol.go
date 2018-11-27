package unison

import (
	"fmt"
	"github.com/harmony-one/libunison/pkg/unison/util"
)

// ProtocolID is a protocol identity.
// Implementations shall ensure that identities are comparable,
// so that they can be used as map keys.
type ProtocolID interface {
	util.Marshaler
	util.Unmarshaler
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

type HelloWorldProtocol struct {
}

func (hwp *HelloWorldProtocol) HandleReceived(msg []uint8, node Sender) {
	fmt.Printf("received hello message %+v", msg)
	hwp.Send()
}
