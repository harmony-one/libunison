package manycast

import coopcast "github.com/harmony-one/libunison/internal/ida/coopcast"

type Node struct {
	UniCast
	SelfPeer coopcast.Peer
	PeerList []coopcast.Peer
	AllPeers []coopcast.Peer
}

// IDA broadcast using RaptorQ interface
type UniCast interface {
	BroadCast(msg []byte)
	ListeningOnUniCast()
}
