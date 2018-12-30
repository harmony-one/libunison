package main

import (
	"flag"
	"github.com/harmony-one/libunison/internal/ida/coopcast"
	"github.com/harmony-one/libunison/internal/ida/manycast"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// InitCoopCastNode creates a coopcast node
func InitCoopCastNode(confignbr string, configallpeer string, t0 float64, t1 float64, t2 float64, base float64, hop int) *coopcast.Node {
	rand.Seed(time.Now().UTC().UnixNano())
	config1 := NewConfig()
	err := config1.ReadConfigFile(confignbr)
	if err != nil {
		log.Printf("unable to read config file %v", confignbr)
	}
	selfPeer, peerList, _ := config1.GetPeerInfo()
	config2 := NewConfig()
	err = config2.ReadConfigFile(configallpeer)
	if err != nil {
		log.Printf("unable to read config file %v", configallpeer)
	}

	_, _, allPeers := config2.GetPeerInfo()
	Cache := make(map[coopcast.HashKey]*coopcast.RaptorQImpl)
	SenderCache := make(map[coopcast.HashKey]bool)
	PeerDecodedCounter := make(map[coopcast.HashKey]map[int]int)
	node := coopcast.Node{SelfPeer: selfPeer, PeerList: peerList, AllPeers: allPeers, Cache: Cache, PeerDecodedCounter: PeerDecodedCounter, SenderCache: SenderCache, InitialDelayTime: t0, MaxDelayTime: t1, ExpBase: base, RelayTime: t2, Hop: hop}
	return &node
}

// InitManyCastNode creates a manycast node
func InitManyCastNode(confignbr string, configallpeer string) *manycast.Node {
	config1 := NewConfig()
	err := config1.ReadConfigFile(confignbr)
	if err != nil {
		log.Printf("unable to read config file %v", confignbr)
	}
	selfPeer, peerList, _ := config1.GetPeerInfo()
	config2 := NewConfig()
	err = config2.ReadConfigFile(configallpeer)
	if err != nil {
		log.Printf("unable to read config file %v", configallpeer)
	}
	_, _, allPeers := config2.GetPeerInfo()
	node := manycast.Node{SelfPeer: selfPeer, PeerList: peerList, AllPeers: allPeers}
	return &node
}

func main() {
	graphConfigFile := flag.String("graph_config", "graph0.txt", "file containing network structure")
	generateConfigFiles := flag.Bool("gen_config", false, "whether to generate config files from graph_config file")
	broadCast := flag.Bool("broadcast", false, "whether to broadcast a message")
	msgFile := flag.String("msg_file", "test.txt", "message file to broadcast")
	configFile := flag.String("nbr_config", "configs/config_0.txt", "config file contains neighbor peers")
	allPeerFile := flag.String("all_config", "configs/config_allpeers.txt", "config file contains all peer nodes info")
	mode := flag.String("mode", "coopcast", "choose broadcast testing mode, [coopcast|manycast]")
	t0 := flag.String("t0", "5", "initial delay time for symbol broadcasting")
	t1 := flag.String("t1", "50", "uppper bound delay time for symbol broadcasting")
	t2 := flag.String("t2", "7", "delay time for symbol relay")
	hop := flag.String("hop", "1", "number of hops")
	base := flag.String("base", "1.05", "base of exponential increase of symbol broadcasting delay")
	flag.Parse()

	if *generateConfigFiles {
		GenerateConfigFromGraph(*graphConfigFile)
		return
	}

	switch *mode {
	case "coopcast":
		var ta, tb, tc, b float64
		var h int
		var err error
		if ta, err = strconv.ParseFloat(*t0, 64); err != nil {
			log.Printf("unable to parse t0 %v with error %v", t0, err)
			return
		}
		if tb, err = strconv.ParseFloat(*t1, 64); err != nil {
			log.Printf("unable to parse t1 %v with error %v", t1, err)
			return
		}
		if tc, err = strconv.ParseFloat(*t2, 64); err != nil {
			log.Printf("unable to parse t2 %v with error %v", t2, err)
			return
		}
		if b, err = strconv.ParseFloat(*base, 64); err != nil {
			log.Printf("unable to parse base %v with error %v", base, err)
			return
		}
		if h, err = strconv.Atoi(*hop); err != nil {
			log.Printf("unable to parse hop %v with error %v", hop, err)
			return
		}
		node := InitCoopCastNode(*configFile, *allPeerFile, ta, tb, tc, b, h)
		uaddr := net.JoinHostPort("", node.SelfPeer.UDPPort)
		pc, err := net.ListenPacket("udp", uaddr)
		if err != nil {
			log.Printf("cannot connect to udp port")
			return
		}
		log.Printf("server start listening on udp port %s", node.SelfPeer.UDPPort)

		if *broadCast {
			go node.ListeningOnBroadCast(pc)
			filecontent, err := ioutil.ReadFile(*msgFile)
			log.Printf("file size is %v", len(filecontent))
			if err != nil {
				log.Printf("cannot open file %s", *msgFile)
				return
			}
			cancels, raptorq := node.BroadCast(filecontent, pc)
			node.StopBroadCast(cancels, raptorq)
		} else {
			node.ListeningOnBroadCast(pc)
		}
	case "manycast":
		node := InitManyCastNode(*configFile, *allPeerFile)
		if *broadCast {
			filecontent, err := ioutil.ReadFile(*msgFile)
			log.Printf("file size is %v", len(filecontent))
			if err != nil {
				log.Printf("cannot open file %s", *msgFile)
				return
			}
			node.BroadCast(filecontent)
		} else {
			node.ListeningOnUniCast()
		}

	default:
		log.Printf("mode %v not supported", *mode)
	}
}
