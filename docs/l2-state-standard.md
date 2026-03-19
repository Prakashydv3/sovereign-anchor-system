# L2 State Standard

**Phase:** 5 — L2 State Definition Formalization
**Status:** Complete

---

## What L2 Produces for Anchoring

L2 produces a **Snapshot** at each block height. The Snapshot is the only input to the anchoring pipeline.

---

## Snapshot Format

```json
{
  "chain_id":  "sovereign-l2",
  "height":    10,
  "timestamp": 1700000000,
  "entries": [
    { "key": "account-A", "value": "100" },
    { "key": "account-B", "value": "200" }
  ]
}
```

| Field | Type | Rule |
|-------|------|------|
| `chain_id` | string | Non-empty, identifies the L2 chain |
| `height` | uint64 | Monotonically increasing block height |
| `timestamp` | int64 | Positive unix epoch |
| `entries` | array | Key-value state records |

---

## State Root Generation

```
StateRoot = SHA-256(canonical_json(sort_by_key(entries)))
```

Steps:
1. Sort `entries` by `key` (ascending, lexicographic)
2. Serialize full snapshot as canonical JSON
3. Apply SHA-256

Sorting ensures entry insertion order does not affect the root.

---

## Determinism Guarantees

| Property | How enforced |
|----------|-------------|
| Entry order independence | Entries sorted by key before hashing |
| No runtime state | Hash depends only on snapshot fields |
| Canonical serialization | Fixed struct field order in JSON marshal |
| Platform independence | SHA-256, standard JSON — no locale/platform variance |

---

## No Hidden Assumptions

- L2 does not interpret entry values — they are opaque strings
- L2 does not validate business logic — structure only
- L2 does not know about L1 — it only produces a Snapshot

---

## Test Coverage

| Test | What it proves |
|------|---------------|
| `TestStateRootDeterminism` | Same snapshot → same root |
| `TestStateRootEntryOrderIndependent` | Entry order does not affect root |
| `TestStateRootChangesOnMutation` | Any value change → different root |

---

**Conclusion:** L2 state format is unambiguous, deterministic, and produces a canonical root hash suitable for L1 anchoring.
