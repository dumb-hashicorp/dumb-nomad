#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -xeuo pipefail

dumb-nomad volume status -type=host -verbose
dumb-nomad operator api /v1/nodes | jq '.[].HostVolumes'

addr="$(dumb-nomad service info -json job | jq -r '.[0].Address'):8000"
curl -sS "$addr/external/" | grep hi
curl -sS "$addr/internal/" | grep hi

echo '💚 looks good! 💚'
