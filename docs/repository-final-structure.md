# Repository Final Structure

**Phase:** 1 — Canonical Repository Consolidation
**Status:** Complete

---

## Canonical Layout

```
sovereign-anchor-system/
│
├── contracts/
│   └── AnchorRegistry.sol          # L1 anchor storage contract
│
├── anchor-client/
│   ├── client.go                   # Anchor submission and query logic
│   └── config.go                   # RPC endpoint and contract address config
│
├── artifact-tools/
│   ├── hash_generator.go           # Deterministic SHA-256 hash from artifact
│   └── artifact_schema.go          # Artifact input structure definition
│
├── node-operations/
│   ├── start_node.sh               # Node startup script
│   └── health_check.go             # Node liveness probe
│
├── architecture/
│   └── system-design.md            # L1–L2 anchoring architecture overview
│
├── examples/
│   └── e2e_anchor_example.go       # Full pipeline: artifact → hash → anchor → verify
│
├── docs/
│   ├── repository-final-structure.md       ← this file
│   ├── artifact-determinism-proof.md       (Phase 2)
│   ├── anchor-client-stability-report.md   (Phase 3)
│   ├── anchor-verification-proof.md        (Phase 4)
│   ├── l2-state-standard.md                (Phase 5)
│   ├── l2-l1-anchor-flow-proof.md          (Phase 6)
│   ├── lineage-continuity-report.md        (Phase 7)
│   ├── replay-consistency-proof.md         (Phase 8)
│   ├── anchoring-observability-map.md      (Phase 9)
│   └── mainnet-readiness-report.md         (Phase 10)
│
├── .gitignore
└── README.md
```

---

## Consolidation Rules

| Rule | Detail |
|------|--------|
| Single source of truth | All components live in this repository |
| No binaries committed | `.exe`, `.bin`, `.out` excluded via `.gitignore` |
| No node data committed | `node-data/`, `chaindata/` excluded |
| No temp files | `tmp/`, `*.log`, `*.bak` excluded |
| One pipeline direction | Artifact → Hash → Anchor → Verify → Trace |

---

## Component Ownership

| Folder | Responsibility |
|--------|---------------|
| `contracts/` | L1 immutable anchor storage |
| `anchor-client/` | L1 interaction (submit, query) |
| `artifact-tools/` | Deterministic hash generation |
| `node-operations/` | Node lifecycle management |
| `architecture/` | System design and decisions |
| `examples/` | Runnable pipeline demonstrations |
| `docs/` | Phase deliverables and proofs |

---

## Removed / Excluded

- All compiled binaries
- All node runtime data
- All temporary build artifacts
- All IDE-specific files

---

**Signed off:** Phase 1 complete. Repository is the single canonical source.
