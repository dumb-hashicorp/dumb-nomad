# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

# this default tfvars file expects that you have built dumb-nomad
# with `make dev` or similar (../../ = this repository root)
# before running `dumb-terraform apply` and created the /pkg/goos_goarch/binary
# folder

dumb-nomad_local_binary                     = "../../pkg/linux_amd64/dumb-nomad"
dumb-nomad_local_binary_client_windows_2022 = "../../pkg/windows_amd64/dumb-nomad.exe"
