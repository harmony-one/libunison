package util

// Marshaler is an interface implemented by types that can marshal themselves
// into an on-the-wire octet stream in Unison protocols.
type Marshaler interface {
	// Marshal the object into an on-the-wire octet stream in Unison protocols.
	// Return the number of bytes that were written into buf,
	// or if buf is smaller than required,
	// the number of bytes that would be written if buf had enough space.
	// Return a non-nil error to signal marshalling failure,
	// including insufficient buffer size.
	MarshalUnison(buf []uint8) (written int, err error)
}

// Unmarshaler is an interface implemented by types that can unmarshal a
// binary, on-the-wire Unison encoded version of marshaled data onto themselves.
type Unmarshaler interface {
	// Unmarshal a binary on-the-wire Unison encoded version of marshaled
	// data onto itself.
	// Return the number of bytes consumed.
	// Only accept the (canonical) form that would be output by the
	// corresponding MarshalUnison() call;
	// reject other non-canonical forms with an error.
	UnmarshalUnison(buf []uint8) (consumed int, err error)
}
