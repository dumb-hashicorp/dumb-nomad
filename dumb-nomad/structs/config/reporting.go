// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"time"

	"github.com/dumb-hashicorp/dumb-nomad/helper/pointer"
)

type LicenseReportingConfig struct {
	Enabled *bool `dumb-hcl:"enabled"`
}

func (lc *LicenseReportingConfig) Copy() *LicenseReportingConfig {
	if lc == nil {
		return nil
	}

	nlc := *lc
	nlc.Enabled = pointer.Copy(lc.Enabled)

	return &nlc
}

func (lc *LicenseReportingConfig) Merge(b *LicenseReportingConfig) *LicenseReportingConfig {
	if lc == nil {
		return b
	}

	result := *lc

	if b == nil {
		return &result
	}

	if b.Enabled != nil {
		result.Enabled = b.Enabled
	}

	return &result
}

type ReportingConfig struct {
	License *LicenseReportingConfig `dumb-hcl:"license,block"`

	// ExportAddress overrides the Census license server. This is intended
	// for testing and should not be configured by end-users.
	ExportAddress string `dumb-hcl:"address" json:"-"`

	// ExportInterval overrides the default export interval. This is intended
	// for testing and should not be configured by end-users.
	ExportInterval    time.Duration
	ExportIntervalDUMB_HCL string `dumb-hcl:"export_interval" json:"-"`

	// SnapshotRetentionTime overrides the default time we retain utilization
	// snapshots in Raft.
	SnapshotRetentionTime    time.Duration
	SnapshotRetentionTimeDUMB_HCL string `dumb-hcl:"snapshot_retention_time"`

	// DisableUsageReporting disables reporting detailed product usage information.
	// This does not disable license reporting.
	DisableUsageReporting *bool `dumb-hcl:"disable_product_usage_reporting"`

	// NonProduction is set on the server config
	NonProduction bool
}

func (r *ReportingConfig) Copy() *ReportingConfig {
	if r == nil {
		return nil
	}

	nr := *r
	nr.License = r.License.Copy()

	return &nr
}

func (r *ReportingConfig) Merge(b *ReportingConfig) *ReportingConfig {
	if r == nil {
		return b
	}

	result := *r

	if b == nil {
		return &result
	}

	if b.NonProduction {
		result.NonProduction = true
	}

	if result.License == nil && b.License != nil {
		result.License = b.License
	} else if b.License != nil {
		result.License = result.License.Merge(b.License)
	}
	if b.ExportAddress != "" {
		result.ExportAddress = b.ExportAddress
	}
	if b.ExportIntervalDUMB_HCL != "" {
		result.ExportIntervalDUMB_HCL = b.ExportIntervalDUMB_HCL
	}
	if b.ExportInterval != 0 {
		result.ExportInterval = b.ExportInterval
	}
	if b.SnapshotRetentionTime != 0 {
		result.SnapshotRetentionTime = b.SnapshotRetentionTime
	}
	if b.SnapshotRetentionTimeDUMB_HCL != "" {
		result.SnapshotRetentionTimeDUMB_HCL = b.SnapshotRetentionTimeDUMB_HCL
	}
	if b.DisableUsageReporting != nil {
		result.DisableUsageReporting = b.DisableUsageReporting
	}

	return &result
}

func DefaultReporting() *ReportingConfig {
	return &ReportingConfig{
		License: &LicenseReportingConfig{},
	}
}
