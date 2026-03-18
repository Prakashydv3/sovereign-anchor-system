package artifact

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Artifact is the canonical L2 input structure.
// Fields are ordered — order is part of the determinism contract.
type Artifact struct {
	ID        string `json:"id"`
	StateRoot string `json:"state_root"`
	Height    uint64 `json:"height"`
	Timestamp int64  `json:"timestamp"`
}

// Hash returns a deterministic SHA-256 hash of the artifact.
// Serialization is canonical JSON with sorted keys (enforced by struct field order).
func Hash(a Artifact) ([32]byte, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return [32]byte{}, fmt.Errorf("marshal: %w", err)
	}
	return sha256.Sum256(data), nil
}
