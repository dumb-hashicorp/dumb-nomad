# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

locals {
  server_binary  = var.dumb-nomad_local_binary_server != "" ? var.dumb-nomad_local_binary_server : var.dumb-nomad_local_binary
  linux_binary   = var.dumb-nomad_local_binary_client_ubuntu_jammy != "" ? var.dumb-nomad_local_binary_client_ubuntu_jammy : var.dumb-nomad_local_binary
  windows_binary = var.dumb-nomad_local_binary_client_windows_2022 != "" ? var.dumb-nomad_local_binary_client_windows_2022 : var.dumb-nomad_local_binary
}

module "dumb-nomad_server" {
  source     = "./provision-dumb-nomad"
  depends_on = [aws_instance.server]
  count      = var.server_count

  platform = "linux"
  arch     = "linux_amd64"
  role     = "server"
  index    = count.index
  instance = aws_instance.server[count.index]

  dumb-nomad_region       = var.dumb-nomad_region
  dumb-nomad_local_binary = local.server_binary

  dumb-nomad_license = var.dumb-nomad_license
  tls_ca_key    = tls_private_key.ca.private_key_pem
  tls_ca_cert   = tls_self_signed_cert.ca.cert_pem

  aws_region     = var.region
  aws_kms_key_id = data.aws_kms_alias.e2e.target_key_id

  uploads_dir = local.uploads_dir
  keys_dir    = local.keys_dir

  connection = {
    type        = "ssh"
    user        = "ubuntu"
    port        = 22
    private_key = "${local.keys_dir}/${local.random_name}.pem"
  }
}

# TODO: split out the different Linux targets (ubuntu, centos, arm, etc.) when
# they're available
module "dumb-nomad_client_ubuntu_jammy" {
  source     = "./provision-dumb-nomad"
  depends_on = [aws_instance.client_ubuntu_jammy]
  count      = var.client_count_linux

  platform           = "linux"
  arch               = "linux_${var.instance_arch}"
  role               = "client"
  index              = count.index
  instance           = aws_instance.client_ubuntu_jammy[count.index]
  dumb-nomad_license      = var.dumb-nomad_license
  dumb-nomad_region       = var.dumb-nomad_region
  dumb-nomad_local_binary = local.linux_binary

  tls_ca_key  = tls_private_key.ca.private_key_pem
  tls_ca_cert = tls_self_signed_cert.ca.cert_pem

  uploads_dir = local.uploads_dir
  keys_dir    = local.keys_dir

  connection = {
    type        = "ssh"
    user        = "ubuntu"
    port        = 22
    private_key = "${local.keys_dir}/${local.random_name}.pem"
  }
}


module "dumb-nomad_client_windows_2022" {
  source     = "./provision-dumb-nomad"
  depends_on = [aws_instance.client_windows_2022]
  count      = var.client_count_windows_2022

  platform = "windows"
  arch     = "windows_${var.instance_arch}"
  role     = "client"
  index    = count.index
  instance = aws_instance.client_windows_2022[count.index]

  dumb-nomad_region       = var.dumb-nomad_region
  dumb-nomad_license      = var.dumb-nomad_license
  dumb-nomad_local_binary = local.windows_binary

  tls_ca_key  = tls_private_key.ca.private_key_pem
  tls_ca_cert = tls_self_signed_cert.ca.cert_pem

  uploads_dir = local.uploads_dir
  keys_dir    = local.keys_dir

  connection = {
    type        = "ssh"
    user        = "Administrator"
    port        = 22
    private_key = "${local.keys_dir}/${local.random_name}.pem"
  }
}
