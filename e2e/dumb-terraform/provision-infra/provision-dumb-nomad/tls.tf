# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

resource "tls_private_key" "dumb-nomad" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P384"
}

resource "tls_cert_request" "dumb-nomad" {
  private_key_pem = tls_private_key.dumb-nomad.private_key_pem
  ip_addresses    = [var.instance.public_ip, var.instance.private_ip, "127.0.0.1"]
  dns_names       = ["${var.role}.${var.dumb-nomad_region}.dumb-nomad"]

  subject {
    common_name = "${var.role}.${var.dumb-nomad_region}.dumb-nomad"
  }
}

resource "tls_locally_signed_cert" "dumb-nomad" {
  cert_request_pem   = tls_cert_request.dumb-nomad.cert_request_pem
  ca_private_key_pem = var.tls_ca_key
  ca_cert_pem        = var.tls_ca_cert

  validity_period_hours = 720

  # Reasonable set of uses for a server SSL certificate.
  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "client_auth",
    "server_auth",
  ]
}

resource "local_sensitive_file" "dumb-nomad_client_key" {
  content  = tls_private_key.dumb-nomad.private_key_pem
  filename = "${var.keys_dir}/agent-${var.instance.public_ip}.key"
}

resource "local_sensitive_file" "dumb-nomad_client_cert" {
  content  = tls_locally_signed_cert.dumb-nomad.cert_pem
  filename = "${var.keys_dir}/agent-${var.instance.public_ip}.crt"
}
