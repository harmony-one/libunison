package main

import (
	"flag"
	"github.com/harmony-one/libunison/internal/ida/coopcast"
	"github.com/harmony-one/libunison/internal/ida/manycast"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"time"
)

func initCoopCastNode(confignbr string, configallpeer string, t0 float64, t1 float64, t2 float64, base float64, hop int) *coopcast.Node {
	rand.Seed(time.Now().UTC().UnixNano())
	config1 := NewConfig()
	err := config1.ReadConfigFile(confignbr)
	if err != nil {
		log.Printf("unable to read config file %v", confignbr)
		return nil
	}
	selfPeer, peerList, _ := config1.GetPeerInfo()
	config2 := NewConfig()
	err = config2.ReadConfigFile(configallpeer)
	if err != nil {
		log.Printf("unable to read config file %v", configallpeer)
		return nil
	}

	_, _, allPeers := config2.GetPeerInfo()
	cache := make(map[coopcast.HashKey]*coopcast.RaptorQImpl)
	senderCache := make(map[coopcast.HashKey]bool)
	peerDecodedCounter := make(map[coopcast.HashKey]map[int]int)
	node := coopcast.Node{SelfPeer: selfPeer, PeerList: peerList, AllPeers: allPeers, Cache: cache, PeerDecodedCounter: peerDecodedCounter, SenderCache: senderCache, InitialDelayTime: t0, MaxDelayTime: t1, ExpBase: base, RelayTime: t2, Hop: hop}
	return &node
}

func initManyCastNode(confignbr string, configallpeer string) *manycast.Node {
	config1 := NewConfig()
	err := config1.ReadConfigFile(confignbr)
	if err != nil {
		log.Printf("unable to read config file %v", confignbr)
		return nil
	}
	selfPeer, peerList, _ := config1.GetPeerInfo()
	config2 := NewConfig()
	err = config2.ReadConfigFile(configallpeer)
	if err != nil {
		log.Printf("unable to read config file %v", configallpeer)
		return nil
	}
	_, _, allPeers := config2.GetPeerInfo()
	node := manycast.Node{SelfPeer: selfPeer, PeerList: peerList, AllPeers: allPeers}
	return &node
}

func main() {
	rand.Seed(time.Now().UnixNano())

	graphConfigFile := flag.String("graph_config", "graph0.txt", "file containing network structure")
	generateConfigFiles := flag.Bool("gen_config", false, "whether to generate config files from graph_config file")
	broadCast := flag.Bool("broadcast", false, "whether to broadcast a message")
	msgFile := flag.String("msg_file", "test.txt", "message file to broadcast")
	configFile := flag.String("nbr_config", "configs/config_0.txt", "config file contains neighbor peers")
	allPeerFile := flag.String("all_config", "configs/config_allpeers.txt", "config file contains all peer nodes info")
	mode := flag.String("mode", "coopcast", "choose broadcast testing mode, [coopcast|manycast]")
	t0 := flag.Float64("t0", 5, "initial delay time for symbol broadcasting")
	t1 := flag.Float64("t1", 50, "uppper bound delay time for symbol broadcasting")
	t2 := flag.Float64("t2", 7, "delay time for symbol relay")
	hop := flag.Int("hop", 1, "number of hops")
	base := flag.Float64("base", 1.05, "base of exponential increase of symbol broadcasting delay")
	flag.Parse()

	if *generateConfigFiles {
		GenerateConfigFromGraph(*graphConfigFile)
		return
	}

	switch *mode {
	case "coopcast":
		node := initCoopCastNode(*configFile, *allPeerFile, *t0, *t1, *t2, *base, *hop)
		if node == nil {
			log.Printf("unable to create node")
			return
		}
		uaddr := net.JoinHostPort("", node.SelfPeer.UDPPort)
		pc, err := net.ListenPacket("udp", uaddr)
		if err != nil {
			log.Printf("cannot listen on udp port")
			return
		}
		log.Printf("server start listening on udp port %s", node.SelfPeer.UDPPort)

		if *broadCast {
			go node.ListeningOnBroadCast(pc)
			filecontent, err := ioutil.ReadFile(*msgFile)
			if err != nil {
				log.Printf("cannot open file %s", *msgFile)
				return
			}
			log.Printf("file size is %v", len(filecontent))
			cancels, raptorq := node.BroadCast(filecontent, pc)
			node.StopBroadCast(cancels, raptorq)
		} else {
			node.ListeningOnBroadCast(pc)
		}
	case "manycast":
		node := initManyCastNode(*configFile, *allPeerFile)
		if node == nil {
			log.Printf("unable to create node")
			return
		}
		if *broadCast {
			filecontent, err := ioutil.ReadFile(*msgFile)
			if err != nil {
				log.Printf("cannot open file %s", *msgFile)
				return
			}
			log.Printf("file size is %v", len(filecontent))
			node.BroadCast(filecontent)
		} else {
			node.ListeningOnUniCast()
		}

	default:
		log.Printf("mode %v not supported", *mode)
	}
}
