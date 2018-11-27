package unison

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Locator traffic type
const (
	LocatorTrafficTypeSignalingAndUserData = 0
	LocatorTrafficTypeSignalingOnly        = 1
	LocatorTrafficTypeUserDataOnly         = 2
)

// Locator type
const (
	LocatorTypeIP        = 0
	LocatorTypeIPESP     = 1
	LocatorTypeTransport = 2
)

// Locator is a host locator, as defined in RFC 8046, section 4.
type Locator interface {
	// TrafficType returns the traffic type for which the locator is valid.
	TrafficType() (trafficType uint8)

	// SetTrafficType sets the traffic type field.
	SetTrafficType(trafficType uint8)

	// LocatorType returns the locator type constant as defined in RFC
	// 8046 section 4.2.
	LocatorType() (locatorType uint8)

	// Lifetime returns the locator lifetime in seconds.
	Lifetime() (lifetime uint32)

	// SetLifetime sets the locator lifetime in seconds.
	SetLifetime(lifetime uint32)

	// Preferred returns whether this locator is preferred,
	// as defined in RFC 8046, section 3.3.3.
	Preferred() (preferred bool)

	// SetPreferred sets whether this locator is preferred.
	SetPreferred(preferred bool)

	// EncodeLocator encodes the locator into the given buffer if not nil,
	// and returns its length.
	EncodeLocator(buf []byte) (length uint8)

	// DecodeLocator decodes the locator into this object.
	DecodeLocator(buf []byte) (err error)
}

// LocatorBase is the base implementation of locator,
// common to all locator types.
// Concrete implementations may embed this,
// but should still implement LocatorType() and Encode-/DecodeLocator().
// They should also expose type-specific field getters/setters.
type LocatorBase struct {
	trafficType uint8
	preferred   bool
	lifetime    uint32
}

// TrafficType returns the traffic type for which the locator is valid.
func (loc *LocatorBase) TrafficType() (trafficType uint8) {
	return loc.trafficType
}

// SetTrafficType sets the traffic type field.
func (loc *LocatorBase) SetTrafficType(trafficType uint8) {
	loc.trafficType = trafficType
}

// Lifetime returns the locator lifetime in seconds.
func (loc *LocatorBase) Lifetime() (lifetime uint32) {
	return loc.lifetime
}

// SetLifetime sets the locator lifetime in seconds.
func (loc *LocatorBase) SetLifetime(lifetime uint32) {
	loc.lifetime = lifetime
}

// Preferred returns whether this locator is preferred,
// as defined in RFC 8046, section 3.3.3.
func (loc *LocatorBase) Preferred() (preferred bool) {
	return loc.preferred
}

// SetPreferred sets whether this locator is preferred.
func (loc *LocatorBase) SetPreferred(preferred bool) {
	loc.preferred = preferred
}

// IPLocator is an IP-address-only locator.  Used (mainly) for HIP signaling.
type IPLocator struct {
	LocatorBase

	ip net.IP
}

// LocatorType returns the locator type constant as defined in RFC
// 8046 section 4.2.
func (loc *IPLocator) LocatorType() (locatorType uint8) {
	return LocatorTypeIP
}

// EncodeLocator encodes the locator into the given buffer if not nil,
// and returns its length.
func (loc *IPLocator) EncodeLocator(buf []byte) (length uint8) {
	if buf != nil {
		copy(buf, loc.ip.To16())
	}
	length = 16
	return
}

// DecodeLocator decodes the locator into this object.
func (loc *IPLocator) DecodeLocator(buf []byte) (err error) {
	ip := net.IP(buf)
	switch len(ip) {
	case net.IPv4len:
		loc.ip = ip
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			loc.ip = ipv4
		} else {
			loc.ip = ip
		}
	default:
		err = fmt.Errorf("invalid encoded IP address %+v", buf)
	}
	return
}

// IP returns the IP address of this locator.
func (loc *IPLocator) IP() (ip net.IP) {
	return loc.ip
}

// SetIP sets the IP address of this locator.
func (loc *IPLocator) SetIP(ip net.IP) {
	loc.ip = ip
}

// IPESPLocator is an IP-address-only locator.  Used (mainly) for HIP signaling.
type IPESPLocator struct {
	LocatorBase

	spi uint32
	ip  net.IP
}

// LocatorType returns the locator type constant as defined in RFC
// 8046 section 4.2.
func (loc *IPESPLocator) LocatorType() (locatorType uint8) {
	return LocatorTypeIPESP
}

// EncodeLocator encodes the locator into the given buffer if not nil,
// and returns its length.
func (loc *IPESPLocator) EncodeLocator(buf []byte) (length uint8) {
	if buf != nil {
		var spiBuf [4]byte
		binary.BigEndian.PutUint32(spiBuf[:], loc.spi)
		copy(buf, spiBuf[:])
		copy(buf[4:], loc.ip.To16())
	}
	length = 20
	return
}

// DecodeLocator decodes the locator into this object.
func (loc *IPESPLocator) DecodeLocator(buf []byte) (err error) {
	if len(buf) < 4 {
		err = fmt.Errorf("truncated IP-ESP locator %+v", buf)
		return
	}
	spi := binary.BigEndian.Uint32(buf)
	ip := net.IP(buf[4:])
	switch len(ip) {
	case net.IPv4len:
		loc.ip = ip
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			ip = ipv4
		}
	default:
		err = fmt.Errorf("invalid encoded IP address %+v", buf)
		return
	}
	loc.spi = spi
	loc.ip = ip
	return
}

// SPI returns the security parameter index (SPI) of this locator.
func (loc *IPESPLocator) SPI() (spi uint32) {
	return loc.spi
}

// SetSPI sets the security parameter index (SPI) of this locator.
func (loc *IPESPLocator) SetSPI(spi uint32) {
	loc.spi = spi
}

// IP returns the IP address of this locator.
func (loc *IPESPLocator) IP() (ip net.IP) {
	return loc.ip
}

// SetIP sets the IP address of this locator.
func (loc *IPESPLocator) SetIP(ip net.IP) {
	loc.ip = ip
}

// Transport locator kind, as defined in Table 2 in RFC 5770, section 5.7.
const (
	// Host (interface) address.
	TransportLocatorKindHost = 0

	// Address observed by STUN server.
	TransportLocatorKindServerReflexive = 1

	// Address observed by the peer.
	TransportLocatorKindPeerReflexive = 2

	// Address provided by TURN (RFC 5766) or HIP relay server (RFC 5770).
	TransportLocatorKindRelayed = 3
)

// TransportLocator is a transport (TCP/UDP) locator, used for NAT traversal.
type TransportLocator struct {
	LocatorBase

	port     uint16
	protocol uint8
	kind     uint8
	priority uint32
	spi      uint32
	ip       net.IP
}

// LocatorType returns the locator type constant as defined in RFC
// 8046 section 4.2.
func (loc *TransportLocator) LocatorType() (locatorType uint8) {
	return LocatorTypeTransport
}

// EncodeLocator encodes the locator into the given buffer if not nil,
// and returns its length.
func (loc *TransportLocator) EncodeLocator(buf []byte) (length uint8) {
	if buf != nil {
		var portBuf [2]byte
		var priorityBuf, spiBuf [4]byte
		binary.BigEndian.PutUint16(portBuf[:], loc.port)
		binary.BigEndian.PutUint32(priorityBuf[:], loc.priority)
		binary.BigEndian.PutUint32(spiBuf[:], loc.spi)
		copy(buf, portBuf[:])
		buf[2] = loc.protocol
		buf[3] = loc.kind
		copy(buf[4:], priorityBuf[:])
		copy(buf[8:], spiBuf[:])
		copy(buf[12:], loc.ip.To16())
	}
	length = 28
	return
}

// DecodeLocator decodes the locator into this object.
func (loc *TransportLocator) DecodeLocator(buf []byte) (err error) {
	if len(buf) < 12 {
		err = fmt.Errorf("truncated IP-ESP locator %+v", buf)
		return
	}
	port := binary.BigEndian
	spi := binary.BigEndian.Uint32(buf)
	buf = buf[4:]
	ip := net.IP(buf)
	switch len(ip) {
	case net.IPv4len:
		loc.ip = ip
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			ip = ipv4
		}
	default:
		err = fmt.Errorf("invalid encoded IP address %+v", buf)
		return
	}
	loc.spi = spi
	loc.ip = ip
	return
}

// Port returns the transport (TCP/UDP) port number of this locator.
func (loc *TransportLocator) Port() (port uint16) {
	return loc.port
}

// SetPort sets the transport (TCP/UDP) port number of this locator.
func (loc *TransportLocator) SetPort(port uint16) {
	loc.port = port
}

// Protocol returns the transport (TCP/UDP) protocol number of this locator.
func (loc *TransportLocator) Protocol() (protocol uint8) {
	return loc.protocol
}

// SetProtocol sets the transport (TCP/UDP) protocol number of this locator.
func (loc *TransportLocator) SetProtocol(protocol uint8) {
	loc.protocol = protocol
}

// Kind returns the ICE candidate kind of this locator.
// See TransportLocatorKind* constants.
func (loc *TransportLocator) Kind() (kind uint8) {
	return loc.kind
}

// SetKind sets the ICE candidate kind of this locator.
// See TransportLocatorKind* constants.
func (loc *TransportLocator) SetKind(kind uint8) {
	loc.kind = kind
}

// Priority returns the ICE priority of this locator.
func (loc *TransportLocator) Priority() (priority uint32) {
	return loc.priority
}

// SetPriority sets the ICE priority of this locator.
func (loc *TransportLocator) SetPriority(priority uint32) {
	loc.priority = priority
}

// SPI returns the security parameter index (SPI) of this locator.
func (loc *TransportLocator) SPI() (spi uint32) {
	return loc.spi
}

// SetSPI sets the security parameter index (SPI) of this locator.
func (loc *TransportLocator) SetSPI(spi uint32) {
	loc.spi = spi
}

// IP returns the IP address of this locator.
func (loc *TransportLocator) IP() (ip net.IP) {
	return loc.ip
}

// SetIP sets the IP address of this locator.
func (loc *TransportLocator) SetIP(ip net.IP) {
	loc.ip = ip
}

func decodeIP(buf []byte) (ip net.IP, err error) {
	clone = net.IP(buf)
	switch len(ip) {
	case net.IPv4len:
		loc.ip = ip
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			loc.ip = ipv4
		} else {
			loc.ip = ip
		}
	default:
		err = fmt.Errorf("invalid encoded IP address %+v", buf)
	}
	return
}
