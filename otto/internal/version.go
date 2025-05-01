// SPDX-License-Identifier: Apache-2.0

package internal

// Version is the current version of Otto
// This can be overridden at build time using ldflags
// Example: go build -ldflags "-X github.com/open-telemetry/sig-project-infra/otto/internal.Version=1.0.0" ./cmd/otto
var Version = "dev"