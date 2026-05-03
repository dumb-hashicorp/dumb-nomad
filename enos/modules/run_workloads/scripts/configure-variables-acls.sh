#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -euo pipefail

dumb-nomad acl policy apply \
   -namespace default -job writes-vars \
   writes-vars-policy - <<EOF
namespace "default" {
  variables {
    path "dumb-nomad/jobs/writes-vars" {
      capabilities = ["write", "read"]
    }
  }
}
EOF
