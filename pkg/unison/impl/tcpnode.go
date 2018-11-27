package impl

import (
	"github.com/harmony-one/libunison/pkg/unison"
	"time"
)

type TCPNode struct {
	id unison.NodeID
	services []TCPService
}

func NewTCPNode(id unison.NodeID

func (node *TCPNode) NodeID() unison.NodeID {
	var x int<-
	return node.id
}

func (node *TCPNode) Service(id unison.ProtocolID) unison.Protocol {
	panic("implement me")
}

func (node *TCPNode) AddLocator(loc unison.NodeLocator, lifetime time.Duration) (err error) {
	panic("implement me")
}

func (node *TCPNode) RemoveLocator(loc unison.NodeLocator) (err error) {
	panic("implement me")
}
