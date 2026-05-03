#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1


dumb-nomad operator api "/v1/allocation/${DUMB_NOMAD_ALLOC_ID}" | jq '.NetworkStatus.Address | length'
