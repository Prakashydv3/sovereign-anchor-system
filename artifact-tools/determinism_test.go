package artifact_test

import (
	"fmt"
	"testing"

	artifact "sovereign-anchor-system/artifact-tools"
)

// TestDeterminism proves same input always produces same hash.
func TestDeterminism(t *testing.T) {
	a := artifact.Artifact{
		ID:        "block-001",
		StateRoot: "abc123def456",
		Height:    1,
		Timestamp: 1700000000,
	}

	h1, _ := artifact.Hash(a)
	h2, _ := artifact.Hash(a)
	h3, _ := artifact.Hash(a)

	if h1 != h2 || h2 != h3 {
		t.Fatalf("non-deterministic: %x %x %x", h1, h2, h3)
	}
	fmt.Printf("PASS determinism: %x\n", h1)
}

// TestFieldSensitivity proves any field change produces a different hash.
func TestFieldSensitivity(t *testing.T) {
	base := artifact.Artifact{ID: "block-001", StateRoot: "abc123", Height: 1, Timestamp: 1700000000}
	variants := []artifact.Artifact{
		{ID: "block-002", StateRoot: "abc123", Height: 1, Timestamp: 1700000000},
		{ID: "block-001", StateRoot: "abc124", Height: 1, Timestamp: 1700000000},
		{ID: "block-001", StateRoot: "abc123", Height: 2, Timestamp: 1700000000},
		{ID: "block-001", StateRoot: "abc123", Height: 1, Timestamp: 1700000001},
	}

	baseHash, _ := artifact.Hash(base)
	for _, v := range variants {
		h, _ := artifact.Hash(v)
		if h == baseHash {
			t.Fatalf("field change did not change hash: %+v", v)
		}
	}
	fmt.Println("PASS field sensitivity")
}

// TestValidation proves invalid artifacts are rejected before hashing.
func TestValidation(t *testing.T) {
	cases := []artifact.Artifact{
		{ID: "", StateRoot: "abc", Height: 1, Timestamp: 1700000000},
		{ID: "x", StateRoot: "", Height: 1, Timestamp: 1700000000},
		{ID: "x", StateRoot: "abc", Height: 0, Timestamp: 1700000000},
		{ID: "x", StateRoot: "abc", Height: 1, Timestamp: 0},
	}
	for _, c := range cases {
		if err := artifact.Validate(c); err == nil {
			t.Fatalf("expected validation error for: %+v", c)
		}
	}
	fmt.Println("PASS validation")
}
