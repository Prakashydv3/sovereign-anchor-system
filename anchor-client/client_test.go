package client_test

import (
	"crypto/sha256"
	"testing"

	client "sovereign-anchor-system/anchor-client"
)

func makeHash(s string) [32]byte { return sha256.Sum256([]byte(s)) }

// TestSubmitRejectsZeroHash ensures zero stateHash is rejected.
func TestSubmitRejectsZeroHash(t *testing.T) {
	_, err := client.Submit([32]byte{}, [32]byte{})
	if err == nil {
		t.Fatal("expected error for zero stateHash")
	}
}

// TestSubmitDeterministicAnchorID proves anchorId is deterministic.
func TestSubmitDeterministicAnchorID(t *testing.T) {
	sh := makeHash("state-A")
	ph := makeHash("parent-A")

	id1, err := client.Submit(sh, ph)
	if err != nil {
		t.Fatal(err)
	}

	// Recompute expected anchorId independently
	combined := append(sh[:], ph[:]...)
	expected := sha256.Sum256(combined)

	if id1 != expected {
		t.Fatalf("anchorId mismatch: got %x want %x", id1, expected)
	}
}

// TestSubmitDuplicateRejected ensures duplicate anchors are rejected.
func TestSubmitDuplicateRejected(t *testing.T) {
	sh := makeHash("state-B")
	ph := makeHash("parent-B")

	_, err := client.Submit(sh, ph)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Submit(sh, ph)
	if err == nil {
		t.Fatal("expected error on duplicate submit")
	}
}

// TestQueryMissingReturnsError ensures querying unknown anchorId returns error.
func TestQueryMissingReturnsError(t *testing.T) {
	_, err := client.Query(makeHash("nonexistent"))
	if err == nil {
		t.Fatal("expected error for missing anchor")
	}
}

// TestVerifyMatchAndMismatch covers both verify outcomes.
func TestVerifyMatchAndMismatch(t *testing.T) {
	sh := makeHash("state-C")
	ph := makeHash("parent-C")

	id, err := client.Submit(sh, ph)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := client.Verify(id, sh)
	if err != nil || !ok {
		t.Fatal("expected verify to pass")
	}

	ok, err = client.Verify(id, makeHash("wrong"))
	if err != nil || ok {
		t.Fatal("expected verify to fail on wrong hash")
	}
}
