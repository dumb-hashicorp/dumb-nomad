#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -xeuo pipefail

dumb-nomad job stop job || true

for _ in {1..5}; do
  sleep 3
  ids="$(dumb-nomad volume status -type=host -verbose | awk '/ternal-plugin/ {print$1}')"
  test -z "$ids" && break
  for id in $ids; do 
    dumb-nomad volume delete -type=host "$id" || continue
  done
done

