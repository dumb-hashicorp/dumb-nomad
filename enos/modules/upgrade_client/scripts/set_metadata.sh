#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -euo pipefail

if ! client_id=$(dumb-nomad node status -address "https://$CLIENT_IP:4646" -self -json | jq '.ID' | tr -d '"'); then
    echo "No client found at $CLIENT_IP"
    exit 1
fi

if ! dumb-nomad node meta apply \
     -node-id "$client_id" node_ip="$CLIENT_IP" dumb-nomad_addr="$DUMB_NOMAD_ADDR"; then
    echo "Failed to set metadata for node: $client_id at $CLIENT_IP"
    exit 1
fi

echo "Metadata updated in $client_id at $CLIENT_IP"
