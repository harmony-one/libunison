package manycast

import (
	"bufio"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

// BroadCast let sender broadcast message to peer nodes
func (node *Node) BroadCast(msg []byte) {
	var wg sync.WaitGroup
	t1 := time.Now().UnixNano()
	for _, peer := range node.AllPeers {
		if node.SelfPeer.PubKey == peer.PubKey {
			continue
		}
		tcpaddr := net.JoinHostPort(peer.IP, peer.TCPPort)
		conn, err := net.Dial("tcp", tcpaddr)
		if err != nil {
			log.Printf("cannot connect to peer %v:%v", peer.IP, peer.TCPPort)
			continue
		}
		wg.Add(1)
		go sendData(conn, msg, &wg)
	}
	log.Printf("waiting connection to close...")
	wg.Wait()
	t2 := time.Now().UnixNano()
	log.Printf("finish sending data to all peers with %v ms", (t2-t1)/1000000)
}

func sendData(conn net.Conn, msg []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
	timeoutDuration := 2 * time.Second
	conn.SetWriteDeadline(time.Now().Add(timeoutDuration))
	N := len(msg)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(N))
	packet := append(buf, msg...)
	n, err := conn.Write(packet)
	if err != nil {
		log.Printf("cannot unicast data to peer %v", conn.RemoteAddr())
		return
	}
	log.Printf("%v bytes write", n)
}

// ListeningOnUniCast let receiver listening and receive message from the sender
func (node *Node) ListeningOnUniCast() {
	addr := net.JoinHostPort("127.0.0.1", node.SelfPeer.TCPPort)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("cannot listening to the port %s", node.SelfPeer.TCPPort)
		return
	}
	log.Printf("server start listening on tcp port %s", node.SelfPeer.TCPPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("cannot accept connection")
			return
		}
		clientinfo := conn.RemoteAddr().String()
		log.Printf("accept connection from %s", clientinfo)
		go node.handleData(conn)
	}
}

func (node *Node) handleData(conn net.Conn) {
	defer conn.Close()
	c := bufio.NewReader(conn)
	size := make([]byte, 8)
	n, err := io.ReadFull(c, size)
	if err != nil {
		log.Printf("error get filesize, get %v", n)
	}
	N := int(binary.BigEndian.Uint64(size))
	content := make([]byte, N)
	_, err = io.ReadFull(c, content)
	if err != nil {
		log.Printf("cannot read full file")
	}
	fileloc := "received/" + strconv.FormatUint(uint64(time.Now().UnixNano()), 10)
	ioutil.WriteFile(fileloc, content, 0644)
	log.Printf("%v received file written to disk", node.SelfPeer.Sid)
}
