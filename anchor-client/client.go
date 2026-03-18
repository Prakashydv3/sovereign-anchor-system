package client

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// AnchorRecord mirrors what is stored on L1.
type AnchorRecord struct {
	AnchorID   [32]byte
	StateHash  [32]byte
	ParentHash [32]byte
	Timestamp  uint64
}

// store is the in-memory anchor store (replaced by L1 call in production).
var store = map[[32]byte]AnchorRecord{}

// Submit stores an anchor. Returns anchorId.
// anchorId = SHA-256(stateHash + parentHash) — deterministic, no randomness.
func Submit(stateHash [32]byte, parentHash [32]byte) ([32]byte, error) {
	if stateHash == ([32]byte{}) {
		return [32]byte{}, errors.New("stateHash must not be zero")
	}

	combined := append(stateHash[:], parentHash[:]...)
	anchorID := sha256.Sum256(combined)

	if _, exists := store[anchorID]; exists {
		return anchorID, fmt.Errorf("anchor %x already exists", anchorID)
	}

	store[anchorID] = AnchorRecord{
		AnchorID:   anchorID,
		StateHash:  stateHash,
		ParentHash: parentHash,
	}
	return anchorID, nil
}

// Query retrieves an anchor by its ID.
func Query(anchorID [32]byte) (AnchorRecord, error) {
	r, ok := store[anchorID]
	if !ok {
		return AnchorRecord{}, fmt.Errorf("anchor %x not found", anchorID)
	}
	return r, nil
}

// Verify checks that the stored stateHash matches the expected hash.
func Verify(anchorID [32]byte, expectedHash [32]byte) (bool, error) {
	r, err := Query(anchorID)
	if err != nil {
		return false, err
	}
	return r.StateHash == expectedHash, nil
}
