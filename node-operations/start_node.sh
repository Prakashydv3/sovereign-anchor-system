#!/bin/bash
# start_node.sh — Start the sovereign anchor node

set -e

NODE_DIR="${NODE_DIR:-./node-data}"
LOG_FILE="${LOG_DIR:-./logs}/node.log"

mkdir -p "$NODE_DIR" "$(dirname "$LOG_FILE")"

echo "[$(date -u +%Y-%m-%dT%H:%M:%SZ)] Starting sovereign anchor node..."

# Replace the binary path with your actual node binary
exec ./sovereign-node \
  --datadir "$NODE_DIR" \
  --log "$LOG_FILE" \
  "$@"
