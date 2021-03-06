// Copyright 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package shared is intended for isolating constants and minor functions required
// to be consistent across packages so as to avoid transitive package includes
package shared

import "time"

/* Constants used by tether for exchange outside of tether */
const (
	DiskLabelQueryName   = "disk-label"
	FilterSpecQueryName  = "filter-spec"
	SkipRecurseQueryName = "skip-recurse"
	SkipDataQueryName    = "skip-data"

	GuestActionReload    = "reload"
	GuestActionKill      = "kill"
	GuestActionGroupKill = "group-kill"

	GuestShutdownTimeout = 20 * time.Second
	GuestRebootTimeout   = 20 * time.Second

	WaitForSessionExitTimeout = 20 * time.Second
)
