package client

import "fmt"

// VerificationResult holds the outcome of a full anchor verification.
type VerificationResult struct {
	AnchorExists  bool
	HashMatch     bool
	ParentLinked  bool
	Valid         bool
}

// FullVerify checks anchor existence, hash match, and parent linkage.
func FullVerify(anchorID [32]byte, expectedHash [32]byte, expectedParent [32]byte) (VerificationResult, error) {
	r := VerificationResult{}

	record, err := Query(anchorID)
	if err != nil {
		return r, fmt.Errorf("anchor not found: %w", err)
	}
	r.AnchorExists = true
	r.HashMatch = record.StateHash == expectedHash
	r.ParentLinked = record.ParentHash == expectedParent
	r.Valid = r.AnchorExists && r.HashMatch && r.ParentLinked

	return r, nil
}
