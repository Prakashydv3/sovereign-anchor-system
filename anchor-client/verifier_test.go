package client_test

import (
	"crypto/sha256"
	"testing"

	client "sovereign-anchor-system/anchor-client"
)

func TestFullVerifyAllPass(t *testing.T) {
	sh := sha256.Sum256([]byte("state-verify-1"))
	ph := sha256.Sum256([]byte("parent-verify-1"))

	id, err := client.Submit(sh, ph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := client.FullVerify(id, sh, ph)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Valid {
		t.Fatalf("expected valid: %+v", result)
	}
}

func TestFullVerifyHashMismatch(t *testing.T) {
	sh := sha256.Sum256([]byte("state-verify-2"))
	ph := sha256.Sum256([]byte("parent-verify-2"))

	id, _ := client.Submit(sh, ph)

	result, err := client.FullVerify(id, sha256.Sum256([]byte("wrong")), ph)
	if err != nil {
		t.Fatal(err)
	}
	if result.Valid || result.HashMatch {
		t.Fatal("expected hash mismatch to invalidate result")
	}
}

func TestFullVerifyParentMismatch(t *testing.T) {
	sh := sha256.Sum256([]byte("state-verify-3"))
	ph := sha256.Sum256([]byte("parent-verify-3"))

	id, _ := client.Submit(sh, ph)

	result, err := client.FullVerify(id, sh, sha256.Sum256([]byte("wrong-parent")))
	if err != nil {
		t.Fatal(err)
	}
	if result.Valid || result.ParentLinked {
		t.Fatal("expected parent mismatch to invalidate result")
	}
}

func TestFullVerifyNotFound(t *testing.T) {
	_, err := client.FullVerify(sha256.Sum256([]byte("ghost")), [32]byte{}, [32]byte{})
	if err == nil {
		t.Fatal("expected error for non-existent anchor")
	}
}
