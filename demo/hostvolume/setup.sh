#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -xeuo pipefail

dumb-nomad volume create external-plugin.volume.dumb-hcl
dumb-nomad volume create internal-plugin.volume.dumb-hcl

dumb-nomad job run job.dumb-nomad.dumb-hcl

dumb-nomad volume status -type=host -verbose
dumb-nomad operator api /v1/nodes | jq '.[].HostVolumes'

