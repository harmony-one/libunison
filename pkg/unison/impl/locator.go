package impl

import (
	"github.com/harmony-one/libunison/pkg/unison"
	"net"
)

type ipLocator net.IPAddr

func NewIPLocator(addr net.IPAddr) unison.IPLocator {
	return &ipLocator{addr}
}

func (loc *ipLocator) IPAddr() net.IPAddr {
	return net.IPAddr(*loc)
}

type tcpLocator struct {
	ipLocator

	port uint16
}

func (loc *tcpLocator) Port() uint16 {
	return loc.port
}

type udpLocator struct {
	ipLocator

	port uint16
}

func (loc *udpLocator) Port() uint16 {
	return loc.port
}
