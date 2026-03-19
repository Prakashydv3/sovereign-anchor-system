# Artifact Determinism Proof

**Phase:** 2 — Deterministic Artifact Pipeline Hardening
**Status:** Complete

---

## Determinism Contract

> Same artifact input → same SHA-256 hash, always, on any machine, at any time.

### How it is enforced

| Mechanism | Detail |
|-----------|--------|
| Fixed struct field order | Go's `json.Marshal` serializes struct fields in declaration order — not alphabetical |
| No runtime state in hash | Hash depends only on artifact fields — no clock, no random, no counter |
| Canonical JSON | No pretty-printing, no trailing spaces, no locale-dependent formatting |
| SHA-256 | Deterministic, collision-resistant, platform-independent |

---

## Artifact Input Structure

```json
{
  "id":         "block-001",
  "state_root": "abc123def456",
  "height":     1,
  "timestamp":  1700000000
}
```

Field order in serialization is fixed: `id → state_root → height → timestamp`

---

## Reproducibility Log

| Run | Input | Hash Output |
|-----|-------|-------------|
| 1 | `{id:block-001, state_root:abc123def456, height:1, timestamp:1700000000}` | `e3b6f1...` (SHA-256) |
| 2 | same | identical |
| 3 | same | identical |
| field change: id=block-002 | different | different hash |
| field change: height=2 | different | different hash |

*Full hash values produced by `determinism_test.go` — run `go test ./artifact-tools/...`*

---

## Validation Rules (Structure Only)

| Field | Rule |
|-------|------|
| `id` | Non-empty string |
| `state_root` | Non-empty string |
| `height` | Integer > 0 |
| `timestamp` | Positive unix epoch |

No semantic validation — content meaning is never interpreted.

---

## Test Coverage

| Test | What it proves |
|------|---------------|
| `TestDeterminism` | Same input → same hash across 3 runs |
| `TestFieldSensitivity` | Any single field change → different hash |
| `TestValidation` | Invalid artifacts rejected before hashing |

---

**Conclusion:** Artifact hashing is deterministic, reproducible, and structurally validated.
