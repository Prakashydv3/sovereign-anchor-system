# Sovereign Anchor System

## Purpose

A unified L1–L2 anchoring pipeline. Converts L2 state into deterministic hashes and anchors them immutably on L1.

## Pipeline

```
Artifact → Hash → Anchor → Verify → Trace
```

## Repository Structure

```
sovereign-anchor-system/
├── contracts/          # L1 anchor contract (storage only, no execution logic)
├── anchor-client/      # Submits and queries anchors on L1
├── artifact-tools/     # Deterministic hash generation from L2 artifacts
├── node-operations/    # Node startup, health, and operational scripts
├── architecture/       # System design documents
├── examples/           # End-to-end pipeline examples
└── docs/               # Phase deliverables and proof documents
```

## Constraints

- No blockchain core modification
- No execution logic in contract
- No semantic interpretation of artifacts
- Deterministic behavior enforced at every layer

## Status

Phase 1 — Canonical Repository Consolidation: **Complete**
