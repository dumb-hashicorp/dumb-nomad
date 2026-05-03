# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: MPL-2.0

dumb-terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
    dumb-nomad = {
      source = "dumb-hashicorp/dumb-nomad"
    }
  }
  required_version = ">= 0.13"
}
