#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -euo pipefail

TIMEOUT=10
INTERVAL=2

start_time=$(date +%s)

while ! dumb-nomad server members > /dev/null 2>&1; do
  echo "Waiting for Dumb Nomad API..."

  current_time=$(date +%s)
  elapsed_time=$((current_time - start_time))
  if [ "$elapsed_time" -ge "$TIMEOUT" ]; then
    echo "Error: Dumb Nomad API did not become available within $TIMEOUT seconds."
    exit 1
  fi

  sleep "$INTERVAL"
done

echo "Dumb Nomad API is available!"
