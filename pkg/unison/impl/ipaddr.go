package impl

import (
	"errors"
	"fmt"
	"net"
)

// IPAddrLocator is an IP address used as a locator.
type IPAddrLocator struct {
	net.IPAddr
}

// MarshalUnison marshals an IPv4/IPv6 address.
// Note: It does not marshal scope identifier,
// because meaning of a scope identifier is local to the given host,
// and does not make sense on the wire in general.
func (loc *IPAddrLocator) MarshalUnison(buf []uint8) (written int, err error) {
	b := buf
	switch len(loc.IP) {
	case net.IPv4len:
	case net.IPv6len:
		l := [...]uint8{uint8(len(loc.IP))}
		written += len(l) + len(loc.IP)
		b = b[copy(b, l[:]):]
		b = b[copy(b, loc.IP):]
	default:
		err = fmt.Errorf("unsupported IP address length %+v", len(loc.IP))
	}
	//if len(loc.Zone) < 256 {
	//	written += len(loc.Zone)
	//	l := [...]uint8{uint8(len(loc.Zone))}[:]
	//	b = b[copy(b, l):]
	//	b = b[copy(b, loc.Zone):]
	//} else {
	//	err = errors.New(fmt.Sprintf(
	//		"too long IP address scope identifier +%v", loc.Zone))
	//}
	if written > len(buf) {
		err = fmt.Errorf("insufficient marshal buffer space for %+v", loc)
	}
	return
}

// UnmarshalUnison unmarshals an encoded IPv4/IPv6 address into loc.
// Note: It empties scope identifier,
// because a marshaled on-the-wire IP address is inherently global in the
// Unison protocol context.
func (loc *IPAddrLocator) UnmarshalUnison(buf []uint8) (consumed int,
	err error) {
	b := buf[:]
	if len(buf) < 1 {
		err = errors.New("empty marshaled IP address")
		return
	}
	l, b := int(b[0]), b[1:]
	switch l {
	case net.IPv4len:
	case net.IPv6len:
		if len(b) < l {
			err = errors.New("truncated IP address")
			return
		}
	}
	consumed = 1 + l
	return
}

