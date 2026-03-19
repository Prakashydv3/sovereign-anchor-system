package l2state_test

import (
	"testing"

	l2state "sovereign-anchor-system/artifact-tools"
)

func baseSnapshot() l2state.Snapshot {
	return l2state.Snapshot{
		ChainID:   "sovereign-l2",
		Height:    10,
		Timestamp: 1700000000,
		Entries: []l2state.StateEntry{
			{Key: "account-A", Value: "100"},
			{Key: "account-B", Value: "200"},
		},
	}
}

// TestStateRootDeterminism proves same snapshot → same root.
func TestStateRootDeterminism(t *testing.T) {
	s := baseSnapshot()
	r1, _ := l2state.StateRoot(s)
	r2, _ := l2state.StateRoot(s)
	if r1 != r2 {
		t.Fatal("state root is not deterministic")
	}
}

// TestStateRootEntryOrderIndependent proves entry order does not affect root.
func TestStateRootEntryOrderIndependent(t *testing.T) {
	s1 := baseSnapshot()
	s2 := baseSnapshot()
	s2.Entries = []l2state.StateEntry{
		{Key: "account-B", Value: "200"},
		{Key: "account-A", Value: "100"},
	}

	r1, _ := l2state.StateRoot(s1)
	r2, _ := l2state.StateRoot(s2)
	if r1 != r2 {
		t.Fatal("entry order should not affect state root")
	}
}

// TestStateRootChangesOnMutation proves any state change changes the root.
func TestStateRootChangesOnMutation(t *testing.T) {
	base := baseSnapshot()
	mutated := baseSnapshot()
	mutated.Entries[0].Value = "999"

	r1, _ := l2state.StateRoot(base)
	r2, _ := l2state.StateRoot(mutated)
	if r1 == r2 {
		t.Fatal("mutated state should produce different root")
	}
}
