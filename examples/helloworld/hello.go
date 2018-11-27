package helloworld

import (
	"github.com/harmony-one/libunison/pkg/unison"
)

func ucc(
	u unison.Unison,
	protocolID unison.ProtocolID, nodeID unison.NodeID,
	network string, addr string,
) (
	reply string, err error,
) {
	u.AddLocator(nodeID, unison.NewGoNetLocator(network, addr))
	proto := u.RegisterProtocol(protocolID)
	sess := proto.Connect(nodeID)
	err = sess.Send([]byte("Hello world"))
	if err != nil {
		return
	}
	msg, err := sess.Receive()
	if err != nil {
		return
	}
	reply = string(msg)
	return
}

func ucs(
	unison unison.Unison,
	protocolID unison.ProtocolID, nodeID unison.NodeID,
	network string, addr string,
) (
	reply string, err error,
) {
	unison.AddLocator(nodeID, unison.NewGoNetLocator(network, addr))
	proto := unison.RegisterProtocol(protocolID)
	sess := proto.Connect(nodeID)
	err = sess.Send([]byte("Hello world"))
	if err != nil {
		return
	}
	msg, err := sess.Receive()
	if err != nil {
		return
	}
	reply = string(msg)
	return
}

func mc(unison unison.Unison, protocolID unison.ProtocolID,
	groupID unison.GroupID) {
	proto := unison.RegisterProtocol(protocolID)
	group := proto.Join(groupID)
	group.Send([]byte("Hello world"))
	for {
		msg, sender, err := group.Receive()
	}
}
