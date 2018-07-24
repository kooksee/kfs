package sp2p

import (
	"fmt"
	"encoding/hex"
	"strings"
)

// BytesID converts a byte slice to a NodeID
func BytesID(b []byte) (Hash, error) {
	var id Hash
	if len(b) != len(id) {
		return id, fmt.Errorf("wrong length, want %d bytes", len(id))
	}
	copy(id[:], b)
	return id, nil
}

// MustBytesID converts a byte slice to a NodeID.
// It panics if the byte slice is not a valid NodeID.
func MustBytesID(b []byte) Hash {
	id, err := BytesID(b)
	if err != nil {
		panic(Errs("check node id error", err.Error()))
	}
	return id
}

// HexID converts a hex string to a NodeID.
// The string may be prefixed with 0x.
func HexID(in string) (Hash, error) {
	var id Hash
	b, err := hex.DecodeString(strings.TrimPrefix(in, "0x"))
	if err != nil {
		return id, err
	} else if len(b) != len(id) {
		return id, fmt.Errorf("wrong length, want %d hex chars", len(id)*2)
	}
	copy(id[:], b)
	return id, nil
}

// MustHexID converts a hex string to a NodeID.
// It panics if the string is not a valid NodeID.
func MustHexID(in string) Hash {
	id, err := HexID(in)
	if err != nil {
		panic(Errs("check nodeid error", err.Error()))
	}
	return id
}
