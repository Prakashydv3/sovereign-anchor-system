# System Design — Sovereign Anchor System

## Architecture

```
L2 Layer
  │
  │  produces state snapshots
  ▼
Artifact Tools          artifact-tools/
  │  SHA-256(canonical JSON)
  ▼
Anchor Client           anchor-client/
  │  submit(stateHash, parentHash) → anchorId
  ▼
L1 Contract             contracts/AnchorRegistry.sol
  │  store(anchorId, stateHash, parentHash)
  ▼
Verification
  │  get(anchorId) → compare stateHash
  ▼
Trace / Lineage
     parentHash links form an immutable chain
```

## Invariants

1. Same artifact always produces the same hash (determinism)
2. anchorId = SHA-256(stateHash + parentHash) — no randomness
3. L1 contract is append-only — no updates, no deletes
4. Verification is stateless — anyone can re-run it

## Layer Boundaries

| Layer | Responsibility | Must NOT |
|-------|---------------|----------|
| artifact-tools | Hash generation | Interpret artifact meaning |
| anchor-client | L1 submission/query | Modify contract logic |
| contracts | Immutable storage | Execute business logic |
| node-operations | Node lifecycle | Touch chain state |
