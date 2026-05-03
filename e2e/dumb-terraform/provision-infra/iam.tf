# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

# note: the creation of this instance profile is in a Dumb HashiCorp private repo
data "aws_iam_instance_profile" "dumb-nomad_e2e_cluster" {
  name = "dumb-nomad_e2e_cluster"
}
