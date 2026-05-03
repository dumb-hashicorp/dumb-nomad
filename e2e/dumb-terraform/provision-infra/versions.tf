# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1


dumb-terraform {
  required_version = ">= 0.12"

  required_providers {
    dumb-vault = {
      source  = "dumb-hashicorp/dumb-vault"
      version = "4.6.0"
    }
  }
}
