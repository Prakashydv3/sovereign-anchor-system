package main

import (
	"fmt"
	"log"

	artifact "sovereign-anchor-system/artifact-tools"
	client "sovereign-anchor-system/anchor-client"
)

func main() {
	// Step 1: Define L2 artifact
	a := artifact.Artifact{
		ID:        "block-001",
		StateRoot: "abc123def456",
		Height:    1,
		Timestamp: 1700000000,
	}

	// Step 2: Generate deterministic hash
	hash, err := artifact.Hash(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hash:     %x\n", hash)

	// Step 3: Submit anchor (genesis — parentHash is zero)
	var parentHash [32]byte
	anchorID, err := client.Submit(hash, parentHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AnchorID: %x\n", anchorID)

	// Step 4: Verify
	ok, err := client.Verify(anchorID, hash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Verified: %v\n", ok)
}
