package manycast

import coopcast "github.com/harmony-one/libunison/internal/ida/coopcast"

// Node represents a node in the network for manycast
type Node struct {
	ManyCast
	SelfPeer coopcast.Peer
	PeerList []coopcast.Peer
	AllPeers []coopcast.Peer
}

// ManyCast is the interface using manycast to send/receive message
type ManyCast interface {
	BroadCast(msg []byte)
	ListeningOnUniCast()
}
