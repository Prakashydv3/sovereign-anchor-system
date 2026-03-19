# Anchor Client Stability Report

**Phase:** 3 — Anchor Client Stabilization
**Status:** Complete

---

## Client Responsibilities

| Function | Input | Output | Error Condition |
|----------|-------|--------|-----------------|
| `Submit` | stateHash, parentHash | anchorId | zero stateHash / duplicate |
| `Query` | anchorId | AnchorRecord | not found |
| `Verify` | anchorId, expectedHash | bool | not found |

---

## AnchorID Generation

```
anchorId = SHA-256(stateHash ++ parentHash)
```

- No randomness
- No timestamp dependency
- No counter
- Same inputs → same anchorId, always

---

## Input Validation Rules

| Rule | Enforced In |
|------|------------|
| stateHash must not be zero | `Submit` |
| Duplicate anchorId rejected | `Submit` |
| anchorId must exist before query/verify | `Query`, `Verify` |

---

## Stability Guarantees

| Property | Status |
|----------|--------|
| Deterministic anchorId | ✅ |
| No semantic interpretation | ✅ |
| Duplicate prevention | ✅ |
| Consistent error returns | ✅ |
| No side effects on failed calls | ✅ |

---

## Test Coverage

| Test | What it validates |
|------|------------------|
| `TestSubmitRejectsZeroHash` | Zero stateHash blocked |
| `TestSubmitDeterministicAnchorID` | anchorId = SHA-256(stateHash+parentHash) |
| `TestSubmitDuplicateRejected` | Same anchor cannot be stored twice |
| `TestQueryMissingReturnsError` | Unknown anchorId returns error |
| `TestVerifyMatchAndMismatch` | Correct hash passes, wrong hash fails |

---

**Conclusion:** Anchor client is stable, deterministic, and rejects all invalid inputs before touching storage.
