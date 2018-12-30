package coopcast

import (
	"context"
	libraptorq "github.com/harmony-one/go-raptorq/pkg/raptorq"
	"net"
	"sync"
	"time"
)

const (
	metaReceived         byte          = 0
	pubKeySize           int           = 20
	stopBroadCastTime    time.Duration = 100 // unit is second
	cacheClearInterval   time.Duration = 250 // clear cache every xx seconds
	enforceClearInterval int64         = 300 // clear old cache eventually
	udpCacheSize         int           = 2 * 1024
	normalChunkSize      int           = 100 * 1200
	symbolSize           int           = 1200 // must be multiple of Al(=4) required by RFC6330

	hashSize  int     = 20  // sha1 hash size
	threshold float32 = 0.8 // threshold rate of number of neighors decode message successfully
)

// Peer represent identification information of a peer node
type Peer struct {
	IP      string
	TCPPort string
	UDPPort string
	PubKey  string
	Sid     int
}

// HashKey is the array of fixed size can be used as key in golang dictionary
type HashKey [hashSize]byte

// Node represents a node using coopcast to send and receive message
type Node struct {
	BroadCaster

	SelfPeer           Peer
	PeerList           []Peer
	AllPeers           []Peer
	InitialDelayTime   float64 // sender delay parameter
	MaxDelayTime       float64 // sender delay parameter
	ExpBase            float64 // sender delay parameter
	RelayTime          float64 // gossip delay parameter
	Hop                int
	SenderCache        map[HashKey]bool
	Cache              map[HashKey]*RaptorQImpl
	PeerDecodedCounter map[HashKey]map[int]int

	mux sync.Mutex
}

// RaptorQImpl represents raptorQ structure holding necessary information for encoding and decoding message
type RaptorQImpl struct {
	Encoder map[int]libraptorq.Encoder
	Decoder map[int]libraptorq.Decoder

	senderID        int
	rootHash        []byte
	numChunks       int
	chunkSize       int
	threshold       int
	receivedSymbols map[int]map[uint32]bool
	numDecoded      int
	initTime        int64 //instance initiate time
	successTime     int64 //success decode time, UnixNano time
	mux             sync.Mutex
	stats           map[int]float64 // for benchmark purpose
}

// BroadCaster interface define the broadcast interface for coopcast for both receiver and sender sides
type BroadCaster interface {
	BroadCast(msg []byte, pc net.PacketConn) (context.CancelFunc, *RaptorQImpl)
	StopBroadCast(cancel context.CancelFunc, raptorq *RaptorQImpl)
	ListeningOnBroadCast(pc net.PacketConn)
}
