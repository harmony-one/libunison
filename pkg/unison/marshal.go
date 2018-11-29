package unison

import "fmt"

// BufferTooSmall signals that a given buffer is too small for marshaling the
// given object.
type BufferTooSmall struct {
	// Object to be marshaled
	Marshaler Marshaler

	// Given buffer
    Buf []uint8
}

func (e BufferTooSmall) Error() string {
	return fmt.Sprintf("buffer %v is too small (%v < %v) for %+v",
        e.Buf, len(e.Buf), e.Marshaler.MarshalUnisonSize(), e.Marshaler)
}

// CheckMarshalSize returns non-nil BufferTooSmall if and only if the
// length of the given buffer is smaller than required by the marshaler.
//
// Example pattern when we choose not to fill short buffer:
//
//     func (m *MyType) MarshalUnison(buf []uint8) (written int, err error) {
//         if err = CheckMarshalBufferSize(m, buf); err != nil {
//             return
//         }
//         ...  // Copy here
//         written = m.MarshalUnisonSize()
//         return
//     }
//
// Example pattern when we choose to fill the short buffer:
//
//     func (m *MyType) MarshalUnison(buf []uint8) (written int, err error) {
//         ...  // Fill the short buffer to the brim.
//         written = len(buf)
//         err = CheckMarshalBufferSize(m, buf)
//         return
//     }
func CheckMarshalSize(marshaler Marshaler, buf []byte) (err error) {
	if len(buf) < marshaler.MarshalUnisonSize() {
		err = &BufferTooSmall{marshaler, buf}
	}
	return
}

// Marshaler is an interface implemented by types that can marshal themselves
// into an on-the-wire octet stream in Unison protocols.
type Marshaler interface {
	// Return the number of bytes required to marshal this object.
    MarshalUnisonSize() (requiredSize int)

	// MarshalUnison marshals the object into an on-the-wire octet stream in
	// Unison protocols.
	//
	// It returns the number of bytes that were written into buf,
	// in [0..len(buf)] range.
	//
	// It returns a non-nil error to signal marshalling failure.
	//
	// If the buffer is too small, MarshalUnison returns a BufferTooSmall error.
	// MarshalUnison may still decide to fill it with a partially marshaled
	// object and return non-zero written count,
	// or it may decide not to fill the buffer and return zero written count.
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
