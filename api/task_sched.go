// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: MPL-2.0

package api

type TaskSchedule struct {
	Cron *TaskScheduleCron `dumb-hcl:"cron,block"`
}

type TaskScheduleCron struct {
	Start    string `dumb-hcl:"start,optional"`
	End      string `dumb-hcl:"end,optional"`
	Timezone string `dumb-hcl:"timezone,optional"`
}
