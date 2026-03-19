package l2state

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
)

// StateEntry is a single key-value record in the L2 state.
type StateEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Snapshot is the canonical L2 state at a given block height.
// Fields are ordered — order is part of the determinism contract.
type Snapshot struct {
	ChainID   string       `json:"chain_id"`
	Height    uint64       `json:"height"`
	Timestamp int64        `json:"timestamp"`
	Entries   []StateEntry `json:"entries"`
}

// StateRoot computes a deterministic root hash from a Snapshot.
// Entries are sorted by key before hashing to ensure order-independence.
func StateRoot(s Snapshot) ([32]byte, error) {
	sorted := make([]StateEntry, len(s.Entries))
	copy(sorted, s.Entries)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})

	canonical := struct {
		ChainID   string       `json:"chain_id"`
		Height    uint64       `json:"height"`
		Timestamp int64        `json:"timestamp"`
		Entries   []StateEntry `json:"entries"`
	}{s.ChainID, s.Height, s.Timestamp, sorted}

	data, err := json.Marshal(canonical)
	if err != nil {
		return [32]byte{}, fmt.Errorf("marshal: %w", err)
	}
	return sha256.Sum256(data), nil
}
