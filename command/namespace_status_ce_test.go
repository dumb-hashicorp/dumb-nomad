// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

//go:build !ent
// +build !ent

package command

import "github.com/dumb-hashicorp/dumb-nomad/api"

func testQuotaSpec() *api.QuotaSpec {
	panic("not implemented - enterprise only")
}
